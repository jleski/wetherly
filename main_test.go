package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"
)

func structuredDataToString(sd map[string]map[string]string) string {
	var sb strings.Builder
	for id, params := range sd {
		sb.WriteString("[")
		sb.WriteString(id)

		// Sort keys for consistent output
		keys := make([]string, 0, len(params))
		for key := range params {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			sb.WriteString(fmt.Sprintf(" %s=\"%s\"", key, params[key]))
		}
		sb.WriteString("]")
	}
	return sb.String()
}

func TestParseRFC5424Message_Basic(t *testing.T) {
	msg := "<13>1 2023-10-10T14:48:00Z myhost myapp 1234 ID47 - Test message"
	expectedTimestamp, _ := time.Parse(time.RFC3339, "2023-10-10T14:48:00Z")

	parsedMsg, err := parseRFC5424Message(msg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedMsg.Version != 1 {
		t.Errorf("Expected version '1', got %d", parsedMsg.Version)
	}

	if !parsedMsg.Timestamp.Equal(expectedTimestamp) {
		t.Errorf("Expected timestamp %v, got %v", expectedTimestamp, parsedMsg.Timestamp)
	}

	if *parsedMsg.Hostname != "myhost" {
		t.Errorf("Expected hostname 'myhost', got %s", *parsedMsg.Hostname)
	}

	if *parsedMsg.Appname != "myapp" {
		t.Errorf("Expected app name 'myapp', got %s", *parsedMsg.Appname)
	}

	if *parsedMsg.ProcID != "1234" {
		t.Errorf("Expected proc ID '1234', got %s", *parsedMsg.ProcID)
	}

	if *parsedMsg.MsgID != "ID47" {
		t.Errorf("Expected msg ID 'ID47', got %s", *parsedMsg.MsgID)
	}

	if parsedMsg.StructuredData != nil {
		t.Errorf("Expected no structured data, got '%v'", parsedMsg.StructuredData)
	}

	if *parsedMsg.Message != "Test message" {
		t.Errorf("Expected message 'Test message', got %s", *parsedMsg.Message)
	}
}

func TestParseRFC5424Message_WithStructuredData(t *testing.T) {
	msg := "<13>1 2023-10-10T14:48:00Z myhost myapp 1234 ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] Test message"
	expectedStructuredData := "[exampleSDID@32473 eventID=\"1011\" eventSource=\"Application\" iut=\"3\"]"

	parsedMsg, err := parseRFC5424Message(msg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedMsg.StructuredData == nil || len(*parsedMsg.StructuredData) == 0 {
		t.Fatalf("Expected structured data, got none")
	}

	// Convert structured data to string for comparison
	sdString := structuredDataToString(*parsedMsg.StructuredData)
	if sdString != expectedStructuredData {
		t.Errorf("Expected structured data '%s', got '%s'", expectedStructuredData, sdString)
	}
}

func TestParseRFC5424Message_MissingOptionalFields(t *testing.T) {
	msg := "<13>1 2023-10-10T14:48:00Z myhost myapp - - - Test message"

	parsedMsg, err := parseRFC5424Message(msg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedMsg.ProcID != nil {
		t.Errorf("Expected proc ID '<nil>', got %v", parsedMsg.ProcID)
	}

	if parsedMsg.MsgID != nil {
		t.Errorf("Expected msg ID '<nil>', got %v", parsedMsg.MsgID)
	}
}

func TestParseRFC5424Message_EmptyStructuredData(t *testing.T) {
	msg := "<13>1 2023-10-10T14:48:00Z myhost myapp 1234 ID47 - Test message"

	parsedMsg, err := parseRFC5424Message(msg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedMsg.StructuredData != nil {
		t.Errorf("Expected empty structured data, got '%v'", parsedMsg.StructuredData)
	}
}

func TestParseRFC5424Message_InvalidFormat(t *testing.T) {
	msg := "<13>1 2023-10-10T14:48:00Z myhost myapp 1234 ID47 Test message"

	_, err := parseRFC5424Message(msg)
	if err == nil {
		t.Fatalf("Expected error for invalid format, got none")
	}
}
