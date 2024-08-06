package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

const indicatorPath = "../../data/indicators/bad_domain.csv"

func (bdd *BadDomainDetection) ruleName() string {
	return bdd.name
}

func fmtRegex(ind []string) string {
	return fmt.Sprintf(".*(%s)$", strings.Join(ind, "|"))
}

func (bdd *BadDomainDetection) setFilterRegex() error {
	// Get the list of domain indicators.
	var ind []string
	f, err := os.Open(indicatorPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", indicatorPath, err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		ind = append(ind, strings.Split(s.Text(), ",")[0])
	}
	if err = s.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %v", err)
	}
	// Compile a regex expression for matching indicators of compromise

	str := fmtRegex(ind)
	bdd.rr, err = regexp.Compile(str)
	if err != nil {
		return fmt.Errorf("Failed compiling regex %s: %v", str, err)
	}
	return nil
}

func (bdd *BadDomainDetection) run() ([]*spb.Signal, error) {
	// Set regex for filtering.
	if err := bdd.setFilterRegex(); err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", bdd.logs)

	// <TODO: Implement me!>
	// Find any logs that contain indicators of compromise from indicatorPath:
	//   1. Filter logs to what is relevant, then
	//   2. [Optional] Aggregate logs based on source IP address.
	//   3. Return the set of interesting logs as a list of spb.Signal

	// Expected output:
	// Option #1: If the aggregation step is skipped, the list of spb.Signal returned should have `event` field set to `bad_domain_filtered`.
	// Option #2: If both filtering and aggregation are performed, the list of spb.Signal returned should have `event` field set to `bad_domain`.

	// Hint #1: Make use of bdd.rr and the `regexp` package.
	// Hint #2: Aggregation is easier using a map data structure.
	// Hint #3: Check the fields you need to populate by inspecting the spb.BadDomain protobuf message.

	// Match any domains to see if there are any in bad_domains, Query field is the domain field in our normalized DNS log
	matchedDomains := make([]*nlpb.DNS, 0)
	for _, log := range bdd.logs.dns {
		if bdd.rr.MatchString(log.Query) || bdd.rr.MatchString(log.Answer) {
			matchedDomains = append(matchedDomains, log)
		}
	}

	// fmt.Printf("Matched domains: %+v\n", matchedDomains)

	// Make map of matched signals by Source IP because source could have multiple matches
	signals := make([]*spb.Signal, 0)
	aggregatedDNSSignals := make(map[string][]*nlpb.DNS)
	for _, matched := range matchedDomains {
		aggregatedDNSSignals[matched.SourceIp] = append(aggregatedDNSSignals[matched.SourceIp], matched)
	}

	// Get timestamps for earliest and latest of the "bad" signals
	for sourceIP, dnsLogs := range aggregatedDNSSignals {
		earliest, latest := dnsLogs[0].Timestamp.AsTime(), dnsLogs[0].Timestamp.AsTime()
		for _, time := range dnsLogs[1:] {
			if time.Timestamp.AsTime().Before(earliest) {
				earliest = time.Timestamp.AsTime()
				continue
			}
			if time.Timestamp.AsTime().After(latest) {
				latest = time.Timestamp.AsTime()
			}
		}

		badDomainRegexResults := bdd.rr.FindStringSubmatch(dnsLogs[0].Query)
		if len(badDomainRegexResults) == 0 {
			badDomainRegexResults = bdd.rr.FindStringSubmatch(dnsLogs[0].Answer)
		}

		signals = append(signals, &spb.Signal{
			Event: &spb.Signal_BadDomain{
				BadDomain: &spb.BadDomain{
					TimestampStart: tspb.New(earliest),
					TimestampEnd:   tspb.New(latest),
					BadDomain:      badDomainRegexResults[1],
					SourceIp:       sourceIP,
					Hostname:       "",
					DnsLog:         dnsLogs,
				},
			},
		})
	}

	log.Printf("Signals: %+v\n", signals)

	return signals, nil
}
