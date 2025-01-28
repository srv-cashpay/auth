package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.19.53:4370") // IP dan port perangkat
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "COMMAND_HERE\n")                  // Kirim perintah ke perangkat
	message, _ := bufio.NewReader(conn).ReadString('\n') // Baca respons
	fmt.Println("Message from device:", message)
}
