package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service, payload := extractArgs(os.Args)

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not resolve UDP address %s: %s\n", service, err.Error())
		os.Exit(2)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open UDP client socket: %s\n", err.Error())
		os.Exit(3)
	}

	payloadBytes := []byte(payload)
	bytesWritten, err := conn.Write(payloadBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not send UDP packet: %s\n", err.Error())
		os.Exit(4)
	}
	if bytesWritten != len(payloadBytes) {
		fmt.Fprintf(os.Stderr, "Attempted to send %d bytes but only sent %d\n", len(payloadBytes), bytesWritten)
		os.Exit(5)
	}

	var respBuf [512]byte
	bytesRead, err := conn.Read(respBuf[0:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read the response: %s\n", err.Error())
		os.Exit(6)
	}

	resultMsg := fmt.Sprintf("Server says: '%s'\n", respBuf[0:bytesRead])
	fmt.Fprintf(os.Stderr, resultMsg)

	fmt.Fprintf(os.Stderr, "Exiting...\n")
}

func extractArgs(args []string) (string, string) {
	if len(args) == 2 || len(args) == 3 {
		service := args[1]
		payload := "Hello, buddy!"

		if len(args) == 3 {
			payload = args[2]
		}

		return service, payload
	}

	fmt.Fprintf(os.Stderr, "Usage: %s target-host:port <payload>\n", args[0])
	os.Exit(1)
	return "", "" // Make compiler happy
}
