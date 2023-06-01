package secretsmanager

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/escanoru/aws-sdk-go-v2-helpers/helper_errors"
)

func DescribeSecret(ctx context.Context, cfg aws.Config, secretID *string) {
	secretManagerClient := secretsmanager.NewFromConfig(cfg)
	filtro := &secretsmanager.DescribeSecretInput{
		SecretId: secretID,
	}
	secretOutput, err := secretManagerClient.DescribeSecret(ctx, filtro)
	helper_errors.CheckAWSError(err)
	log.Printf("The %s secret is accessible and it has the folowwing description: %v", *secretID, *secretOutput.Description)
}
