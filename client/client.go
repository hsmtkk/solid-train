package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:  "client address port message",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		portStr := args[1]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal(err)
		}
		message := args[2]
		run(address, port, message)
	},
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(address string, port int, message string) error {
	raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return fmt.Errorf("ResolveTCPAddr failed; %w", err)
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return fmt.Errorf("DialTCP failed; %w", err)
	}
	defer conn.Close()
	fmt.Printf("send: %s\n", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("Write failed; %w", err)
	}
	if err := conn.CloseWrite(); err != nil {
		return fmt.Errorf("CloseWrite failed; %w", err)
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return fmt.Errorf("Read failed; %w", err)
	}
	fmt.Printf("receive: %s\n", string(buf))
	return nil
}
