package secretsmanager

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/escanoru/aws-sdk-go-v2-helpers/tree/main/helper_errors"
)

func DescribeSecret(ctx context.Context, cfg aws.Config, secretID *string) {
	secretManagerClient := secretsmanager.NewFromConfig(cfg)
	filtro := &secretsmanager.DescribeSecretInput{
		SecretId: secretID,
	}
	secretOutput, err := secretManagerClient.DescribeSecret(ctx, filtro)
	helper_errors.CheckAWSError(err)
	log.Printf("AMI ID %s was successfully deregistered", *secretID)
}