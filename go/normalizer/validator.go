package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

const timeFmt = "2006-01-02 15:04:05.000"

func validateTime(s string) (*tspb.Timestamp, error) {
	// <TODO: Implement me!>
	// Parse a datetime string to a Timestamp proto message.
	// Return error fmt.Errorf("invalid timestamp found: %s", s) if the datetime string is not valid.
	// Hint #1: import "time".
	// Hint #2: make use of the timestamppb package.
	goTime, err := time.Parse(timeFmt, s)
	ts := tspb.New(goTime)
	if err != nil {
		fmt.Errorf("Invalid timestamp found: %s", s)
	}
	return ts, nil
}

func validateTimestamp(s string) (*tspb.Timestamp, error) {
	ts, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("unexpected timestamp found: %s", s)
	}
	// Other checks that could be implemented: check if timestamp is not in the future or too far in the past.
	return &tspb.Timestamp{Seconds: ts}, nil
}

func validateIP(s string) (string, error) {
	// <TODO: Implement me!>
	// Confirm the IP string contains a valid IP address.
	// Return back the valid IP.
	// Return error fmt.Errorf("invalid IP found: %s", s) if the IP is not valid.
	ipAddress := net.ParseIP(s)
	if ipAddress.String() == "" {
		fmt.Errorf("Invalid IP found: %s", s)
	}
	return ipAddress.String(), nil
}

func validatePort(s string) (int32, error) {
	port, err := strconv.ParseInt(s, 10, 32)
	if err != nil || port < 0 || port > 65535 {
		return 0, fmt.Errorf("unexpected port number found: %s", s)
	}
	return int32(port), nil
}

func validateQuery(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("Empty query found")
	}
	return s, nil
}

func validateReturnCode(s string) (nlpb.DNS_ReturnCode, error) {
	code, err := strconv.ParseInt(s, 10, 32)
	if err != nil || code > 9 {
		return 0, fmt.Errorf("Unexpected return code found: %s", s)
	}
	// Increment return code by 1 as the enum value 0 is reserved for default value (e.g. unspecified) only.
	return nlpb.DNS_ReturnCode(int32(code + 1)), nil
}

func validateInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Unexpected string found, expecting int64: %s", s)
	}
	return i, nil
}

func validatePlatform(s string) nlpb.Execution_Platform {
	p, ok := nlpb.Execution_Platform_value[strings.Trim(s, "\"")]
	if !ok {
		log.Printf("Invalid platform %s found, set to default", s)
		return nlpb.Execution_Platform(0)
	}
	return nlpb.Execution_Platform(p)
}
