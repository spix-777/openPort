// Author: spix-777
// Date: 2023-08-25
// Version: 1.0.0
// Description: This program will list all open ports and the process name and PID that is using the port.
// This program is written in Go and uses the nmap library to scan for open ports.
// The lsof command is used to get the process name and PID.
// The program will then match the PID with the process name and port number.
// The output will be PID, process name and port number.
// This is a short cut to nmap -sV -p- <IP> and lsof -i :<port>

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Ullaakut/nmap"
)

// portToString converts a nmap.Port to a string
func portToString(port nmap.Port) string {
	serviceProduct := port.Service.Product

	// Construct the formatted string
	portInfo := fmt.Sprintf(
		"-%d-%s",
		port.ID, serviceProduct,
	)

	return portInfo
}

func nm() []string {
	var portInfo []string
	target := "127.0.0.1"

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts("0-65535"),
		nmap.WithServiceInfo(),
	)
	if err != nil {
		log.Fatalf("Failed to initialize scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("Scan failed: %v", err)
	}

	for _, host := range result.Hosts {
		for _, port := range host.Ports {
			if port.State.State == "open" {
				portInfo = append(portInfo, portToString(port))
			}
		}

	}
	return portInfo
}

func lsof(port string) []string {
	var pid []string
	cmd := exec.Command("lsof", "-i", ":"+port)
	var output bytes.Buffer
	cmd.Stdout = &output

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	lines := strings.Split(output.String(), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			pid = append(pid, fields[1])
		}
	}
	return pid
}

func removeStringsContainingPID(slice []string) []string {
	var result []string

	for _, s := range slice {
		if !strings.Contains(s, "PID") {
			result = append(result, s)
		}
	}

	return result
}

func removeDuplicates(slice []string) []string {
	encountered := map[string]bool{} // Map to track encountered strings
	result := []string{}             // New slice without duplicates

	for _, s := range slice {
		if !encountered[s] {
			encountered[s] = true
			result = append(result, s)
		}
	}

	return result
}

func banner() {
	banner := `
                          ___           _   
  ___  _ __   ___ _ __   / _ \___  _ __| |_ 
 / _ \| '_ \ / _ \ '_ \ / /_)/ _ \| '__| __|
| (_) | |_) |  __/ | | / ___/ (_) | |  | |_ 
 \___/| .__/ \___|_| |_\/    \___/|_|   \__|
      |_|                                   
                                   @spix-777  
`
	fmt.Println(banner)
}

func main() {
	// banner is Always a good idea!
	banner()

	// nm() returns a slice of strings containing port number and service name in nmap format
	portInfo := nm()

	// portSlice is port number and serviceProductSlice is service name
	var portSlice []string
	var serviceProductSlice []string
	for _, port := range portInfo {
		buffer := strings.Split(port, "-")
		portSlice = append(portSlice, buffer[1])
		serviceProductSlice = append(serviceProductSlice, buffer[2])
	}

	// pidSlice is a slice of strings containing PID
	var pid []string
	for _, port := range portSlice {
		buffer := lsof(port)
		pid = append(pid, buffer...)
	}

	pidSlice := removeStringsContainingPID(pid)
	pidSlice = removeDuplicates(pidSlice)

	// pidNamePort is a slice of strings containing PID, service name and port number
	var pidNamePort []string

	// This loop will match PID with service name and port number
	for _, pid := range pidSlice {
		for i, port := range portSlice {
			buffer := lsof(port)
			for _, b := range buffer {
				if pid == b {
					bufferStr := pid + "|" + serviceProductSlice[i] + "|" + port
					pidNamePort = append(pidNamePort, bufferStr)
				}
			}
		}
	}

	buffer := removeDuplicates(pidNamePort)
	fmt.Println("PID:	Process Name:                  Port:")
	fmt.Println("--------------------------------------------")
	for _, b := range buffer {
		a := strings.Split(b, "|")
		fmt.Printf("%s	%s	       %s\n", a[0], a[1], a[2])
	}
}
