package main

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
)

func TestSearchSpecificAliasArN(t *testing.T) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("error")
	}
	expected := true
	got, err := searchSpecificAliasArN(ctx, cfg, "aws/ebs")
	if expected != got {
		t.Errorf("Expected: %v. Got: %v. Reported Error: %v", expected, got, err)
	}
}
