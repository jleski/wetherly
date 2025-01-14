package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	SYSLOG_PORT = ":6601"
	BUFFER_SIZE = 8192
	BANNER      = `
 __          __  _   _                _       
 \ \        / / | | | |              | |      
  \ \  /\  / /__| |_| |__   ___ _ __| |_   _ 
   \ \/  \/ / _ \ __| '_ \ / _ \ '__| | | | |
    \  /\  /  __/ |_| | | |  __/ |  | | |_| |
     \/  \/ \___|\__|_| |_|\___|_|  |_|\__, |
                                        __/ |
                                       |___/ 
    Syslog Server v1.0.0
    ==========================================
`
)

type RFC5424Message struct {
	Priority  int
	Version   string
	Timestamp time.Time
	Hostname  string
	AppName   string
	ProcID    string
	MsgID     string
	Message   string
}

func parseRFC5424Message(msg string) (*RFC5424Message, error) {
	parts := strings.Fields(msg)
	if len(parts) < 8 {
		return nil, fmt.Errorf("invalid RFC5424 message format")
	}

	//priority := parts[0][1:3] // Extract priority from <PRI>
	version := parts[1]
	timestamp, err := time.Parse("2006-01-02T15:04:05Z", parts[2])
	if err != nil {
		return nil, err
	}
	hostname := parts[3]
	appName := parts[4]
	procID := parts[5]
	msgID := parts[6]
	message := strings.Join(parts[7:], " ") // Join remaining parts as message

	return &RFC5424Message{
		Priority:  0,
		Version:   version,
		Timestamp: timestamp,
		Hostname:  hostname,
		AppName:   appName,
		ProcID:    procID,
		MsgID:     msgID,
		Message:   message,
	}, nil
}

func printStartupInfo() {
	fmt.Print("\033[1;36m") // Cyan color
	fmt.Print(BANNER)
	fmt.Print("\033[0m") // Reset color

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	fmt.Printf("\033[1;32m") // Green color
	fmt.Printf("🚀 Starting Wetherly Syslog Server...\n")
	fmt.Printf("📅 Time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Printf("💻 Hostname: %s\n", hostname)
	fmt.Printf("🔌 Protocol: TCP\n")
	fmt.Printf("🎯 Port: 6601\n")
	fmt.Printf("📦 Buffer Size: %d bytes\n", BUFFER_SIZE)
	fmt.Printf("\033[0m") // Reset color
	fmt.Println("==========================================")
}

func main() {
	printStartupInfo()

	// Create TCP listener
	listener, err := net.Listen("tcp", SYSLOG_PORT)
	if err != nil {
		log.Fatalf("\033[1;31mError creating TCP listener: %v\033[0m", err)
	}
	defer listener.Close()

	fmt.Printf("\033[1;32m✅ Server is ready to accept connections\033[0m\n\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("\033[1;31mError accepting connection: %v\033[0m", err)
			continue
		}

		fmt.Printf("\033[1;34m📥 New connection from %s\033[0m\n", conn.RemoteAddr())
		// Handle each connection in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, BUFFER_SIZE)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Printf("\033[1;31mError reading from connection: %v\033[0m", err)
			}
			fmt.Printf("\033[1;33m📤 Connection closed from %s\033[0m\n", conn.RemoteAddr())
			return
		}

		message := string(buffer[:n])
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		if strings.HasPrefix(message, "<") {
			parsedMsg, err := parseRFC5424Message(message)
			if err != nil {
				fmt.Printf("\033[1;31mError parsing RFC5424 message: %v\033[0m\n", err)
			} else {
				fmt.Printf("\033[1;32m[%s] Parsed RFC5424 Message:\033[0m\n\033[1;37m%+v\033[0m\n", timestamp, parsedMsg)
			}
		} else {
			fmt.Printf("\033[1;32m[%s] Message from %v:\033[0m\n\033[1;37m%s\033[0m\n",
				timestamp, conn.RemoteAddr(), message)
		}
	}
}
