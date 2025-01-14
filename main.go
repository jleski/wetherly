package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/influxdata/go-syslog/v3/rfc5424"
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
	CyanColor   = "\033[1;36m"
	GreenColor  = "\033[1;32m"
	RedColor    = "\033[1;31m"
	YellowColor = "\033[1;33m"
	ResetColor  = "\033[0m"
)

type RFC5424Message struct {
	Priority       int
	Version        string
	Timestamp      time.Time
	Hostname       string
	AppName        string
	ProcID         string
	MsgID          string
	StructuredData string // New field for structured data
	Message        string
}

func parseRFC5424Message(msg string) (*rfc5424.SyslogMessage, error) {
	parser := rfc5424.NewParser()
	parsedMsg, err := parser.Parse([]byte(msg))
	if err != nil {
		return nil, fmt.Errorf("error parsing RFC5424 message: %w", err)
	}

	// Type assertion to convert syslog.Message to *rfc5424.SyslogMessage
	rfc5424Msg, ok := parsedMsg.(*rfc5424.SyslogMessage)
	if !ok {
		return nil, fmt.Errorf("parsed message is not of type *rfc5424.SyslogMessage")
	}

	return rfc5424Msg, nil
}

func printStartupInfo() {
	fmt.Print(CyanColor)
	fmt.Print(BANNER)
	fmt.Print(ResetColor)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	fmt.Print(GreenColor)
	fmt.Printf("🚀 Starting Wetherly Syslog Server...\n")
	fmt.Printf("📅 Time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Printf("💻 Hostname: %s\n", hostname)
	fmt.Printf("🔌 Protocol: TCP\n")
	fmt.Printf("🎯 Port: 6601\n")
	fmt.Printf("📦 Buffer Size: %d bytes\n", BUFFER_SIZE)
	fmt.Print(ResetColor)
	fmt.Println("==========================================")
}

func main() {
	printStartupInfo()

	listener, err := net.Listen("tcp", SYSLOG_PORT)
	if err != nil {
		log.Fatalf("%sError creating TCP listener: %v%s", RedColor, err, ResetColor)
	}
	defer listener.Close()

	fmt.Printf("%s✅ Server is ready to accept connections%s\n\n", GreenColor, ResetColor)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("%sError accepting connection: %v%s", RedColor, err, ResetColor)
			continue
		}

		fmt.Printf("%s📥 New connection from %s%s\n", GreenColor, conn.RemoteAddr(), ResetColor)
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
				log.Printf("%sError reading from connection: %v%s", RedColor, err, ResetColor)
			}
			fmt.Printf("%s📤 Connection closed from %s%s\n", YellowColor, conn.RemoteAddr(), ResetColor)
			return
		}

		message := string(buffer[:n])
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		if strings.HasPrefix(message, "<") {
			parsedMsg, err := parseRFC5424Message(message)
			if err != nil {
				fmt.Printf("%sError parsing RFC5424 message: %v%s\n", RedColor, err, ResetColor)
			} else {
				fmt.Printf("%s[%s] Parsed RFC5424 Message:%s\n%s%+v%s\n", GreenColor, timestamp, ResetColor, GreenColor, parsedMsg, ResetColor)
			}
		} else {
			fmt.Printf("%s[%s] Message from %v:%s\n%s%s%s\n", GreenColor, timestamp, conn.RemoteAddr(), ResetColor, GreenColor, message, ResetColor)
		}
	}
}
