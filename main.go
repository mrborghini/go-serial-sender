package main

import (
	"fmt"
	"github.com/tarm/serial"
	"go-serial-sender/components"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: go-serial-sender <channel name> <serial device path>")
		return
	}

	channel := args[1]
	serialDevice := args[2]

	ws := components.NewWebsocket()
	go ws.Connect(channel)
	// Configure the serial port
	config := &serial.Config{
		Name:        serialDevice, // Change to your port (e.g., "/dev/ttyUSB0" for Linux)
		Baud:        9600,   // Set the correct baud rate
		ReadTimeout: time.Millisecond * 1,
		Size:        8, // Change the size to 8 bits
	}

	// Open the serial port
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	string_part := ""

	for {
		// Read data from the serial port
		buf := make([]byte, 1024) // Adjust buffer size as needed
		n, err := port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		buffer_string := string(buf[:n])

		if buffer_string == "" {
			continue
		}

		if !strings.HasSuffix(buffer_string, "\n") {
			string_part += buffer_string
			continue
		}

		temperature := strings.ReplaceAll(string_part+buffer_string, "\n", "")
		ws.Transmit(channel, temperature)
		string_part = ""
	}
}
