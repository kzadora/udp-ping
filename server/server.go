package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	listenPort := getPort(os.Args)

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not resolve UDP address: %s\n", err.Error())
		os.Exit(2)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open UDP socket: %s\n", err.Error())
		os.Exit(3)
	}

	fmt.Fprintf(os.Stderr, "Listening on port %d...\n", listenPort)

	var reqBuf [512]byte
	bytesRead, clientAddr, err := conn.ReadFromUDP(reqBuf[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read UDP packet: %s\n", err.Error())
		os.Exit(4)
	}

	reply := fmt.Sprintf("Received '%s' from %s:%d", reqBuf[0:bytesRead], clientAddr.IP.String(), clientAddr.Port)
	fmt.Fprintf(os.Stderr, "%s\n", reply)

	_, err = conn.WriteToUDP([]byte(reply), clientAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not send a response: %s\n", err.Error())
		os.Exit(5)
	}

	fmt.Fprintf(os.Stderr, "Exiting...\n")
}

func getPort(args []string) int {
	if len(args) == 2 {
		port, err := strconv.Atoi(os.Args[1])
		if err == nil && port > 0 && port < 65535 {
			return port
		}
	}

	fmt.Fprintf(os.Stderr, "Usage: %s port-to-listen\n", args[0])
	os.Exit(1)
	return 0 // Make compiler happy
}
