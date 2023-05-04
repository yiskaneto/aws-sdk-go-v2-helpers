package main

import (
	"flag"
	"log"
	"testing"
)

var amiNameTagValue string

func init() {
	flag.StringVar(&amiNameTagValue, "ami-name-tag-value", "", "AMI name tage value, example: 'RHEL 9.1-internal'")
}

// TestSearchAmiID tests SearchAmiID
func TestSearchAmiID(t *testing.T) {
	_, cfg, ctx := AWSLoadCreds() // full call is awsAccountID, cfg, ctx
	flag.Parse()
	if amiNameTagValue == "" {
		flag.CommandLine.Usage()
		log.Fatal("\nERROR: The --ami-name-tag-value must be filled with the correct value, example: --ami-name-tag-value \"RHEL 9.1-internal\"")
	}
	verticaAMI := SearchAmiID(ctx, cfg, amiNameTagValue)
	log.Println(verticaAMI)
}
