package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:  "server port",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		portStr := args[0]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal(err)
		}
		run(port)
	},
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(port int) error {
	laddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return fmt.Errorf("ResolveTCPAddr failed; %w", err)
	}
	lis, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return fmt.Errorf("ListenTCP failed; %w", err)
	}
	defer lis.Close()

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			return fmt.Errorf("AcceptTCP failed; %w", err)
		}
		defer conn.Close()
		handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	conn.CloseRead()
	recv := string(buf)
	send := strings.ToUpper(recv)
	fmt.Printf("receive: %s\n", recv)
	fmt.Printf("send: %s\n", send)
	_, err = conn.Write([]byte(send))
	if err != nil {
		log.Fatal(err)
	}
}
