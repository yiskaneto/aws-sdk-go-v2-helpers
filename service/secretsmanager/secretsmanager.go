package secretsmanager

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/yiskaneto/aws-sdk-go-v2-helpers/helper_errors"
)

// DescribeSecret attempts to find and get the description of the passed secret
func DescribeSecret(ctx context.Context, cfg aws.Config, secretID *string) (bool, error) {
	secretManagerClient := secretsmanager.NewFromConfig(cfg)
	filtro := &secretsmanager.DescribeSecretInput{
		SecretId: secretID,
	}
	secretOutput, err := secretManagerClient.DescribeSecret(ctx, filtro)
	if err != nil {
		return false, helper_errors.CheckAWSError(err)
	}
	log.Printf("The %s secret is accessible and it has the folowwing description: %v", *secretID, *secretOutput.Description)
	return true, nil
}
