package secretsmanager

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/escanoru/aws-sdk-go-v2-helpers/tree/main/helper_errors"
)

func TestDescribeSecret(t) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	DescribeSecret()
}
