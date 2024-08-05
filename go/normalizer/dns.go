package main

import (
	"fmt"
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
)

func (dn *DNSNormalizer) getInput() string {
	return dn.input
}

func (dn *DNSNormalizer) getBinaryOutput() string {
	return dn.binaryOutput
}

func (dn *DNSNormalizer) getJsonOutput() string {
	return dn.jsonOutput
}

func (dn *DNSNormalizer) normalize(line string) *nlpb.NormalizedLog {
	fields := strings.Split(line, ",")

	// Validate fields.
	if len(fields) != 8 {
		log.Printf("invalid number of fields found; expect 8, found %d: %s\n", len(fields), line)
		return nil
	}
	// <TODO: Implement me!>
	// Implement the validate function in validator.go file.
	// Parse and return `datetime` field with validateTime().
	timeField := fields[0]
	validatedTime, err := validateTime(timeField)
	if err != nil {
		log.Printf("%v, Failed to validate time field: %s\n", err, line)
		return nil
	}
	fmt.Printf("Validated time is %v\n", validatedTime)

	// <TODO: Implement me!>
	// Parse and return `src_ip` field with validateIP().
	srcIP := fields[2]
	validatedSrcIP, err := validateIP(srcIP)
	if err != nil {
		log.Printf("%v, Failed to validate IP address: %s\n", err, line)
		return nil
	}
	fmt.Printf("Validated Source IP address is %v\n", validatedSrcIP)

	// <TODO: Implement me!>
	// Parse and return `resolver_ip` field with validateIP().
	resolverIP := fields[3]
	validatedResolverIP, err := validateIP(resolverIP)
	if err != nil {
		log.Printf("%v, Failed to validate IP address: %s\n", err, line)
		return nil
	}
	fmt.Printf("Validated Resolver IP address is %v\n", validatedResolverIP)
	// <TODO: Implement me!>
	// Parse and return `query` field with validateQuery().
	query := fields[4]
	validatedQuery, err := validateQuery(query)
	if err != nil {
		log.Printf("%v, Failed to validate DNS query: %s\n", err, line)
		return nil
	}
	fmt.Printf("Validated DNS Query is %v\n", validatedQuery)
	/*
		// <TODO: Implement me!>
		// Parse and return `return_code` field with validateReturnCode().
		if err != nil {
			log.Printf("%v, skipping: %s\n", err, line)
			return nil
		}

		// <TODO: Implement me!>
		// Return a populated NormalizedLog proto message.
	*/
	// return &nlpb.NormalizedLog{}
	return nil
}
