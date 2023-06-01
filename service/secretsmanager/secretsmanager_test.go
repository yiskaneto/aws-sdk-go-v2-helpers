package secretsmanager

import (
	"context"
	"log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/escanoru/aws-sdk-go-v2-helpers/tree/main/helper_errors"
)

func TestDescribeSecret(t, ctx context.Context, cfg aws.Config, secretID *string) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	secretManagerClient := secretsmanager.NewFromConfig(cfg)
	filtro := &secretsmanager.DescribeSecretInput{
		SecretId: secretID,
	}
	secretOutput, err := secretManagerClient.DescribeSecret(ctx, filtro)
	helper_errors.CheckAWSError(err)
	log.Printf("AMI ID %s was successfully deregistered", *secretID)
}
