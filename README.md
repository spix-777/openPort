Open Port Finder

Author: spix-777
Date: 2023-08-25
Version: 1.0.0

Description

This program identifies all open ports and maps them with their associated process name and PID (Process ID). It's essentially a more condensed tool combining the features of nmap for port scanning and lsof for identifying processes. Upon execution, it provides an output detailing the PID, process name, and port number.

Dependencies

* Go: Ensure you have the Go programming language installed.
* nmap: This program utilizes the nmap Go library (github.com/Ullaakut/nmap). Ensure you have nmap tool installed on your system and it's accessible from the command line.


Installation

1. Clone the repository or download the source code.
2. Navigate to the directory containing the code.
3. Ensure nmap is installed on your system.
4. Install the necessary Go packages:

go get github.com/Ullaakut/nmap


Usage

1. Navigate to the directory containing the code.
2. Run the program:

go run <filename>.go

3. View the output, which will present a list of PIDs, their associated process names, and the port numbers.


Functions Breakdown

1. portToString: Converts the nmap port structure into a string format.
2. nm: Uses nmap to scan all available ports on the localhost and returns those that are open.
3. lsof: Takes a port number as an argument and returns a slice of PIDs using the port.
4. removeStringsContainingPID: Filters out strings containing "PID".
5. removeDuplicates: Removes any duplicated strings from a slice.
6. banner: Displays an aesthetic banner on the console.
7. main: The driver function that integrates all other functions to display the desired output.


Notes

1. The program only scans the localhost (127.0.0.1). To scan different hosts, modify the target variable in the nm function.
2. The program is hardcoded to scan all ports (0-65535). Adjust the port range in the nm function as needed.
3. Ensure you have the necessary permissions to run port scans on your machine or network.
License

MIT License

