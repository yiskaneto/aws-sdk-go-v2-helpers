package secretsmanager

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func TestDescribeSecret(t *testing.T) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal("error")
	}
	expected := true
	got, err := DescribeSecret(ctx, cfg, aws.String("us-west-1-oicd-cert"))
	if expected != got {
		t.Errorf("Expected: %v. Got: %v. Reported Error: %v", expected, got, err)
	}
}
