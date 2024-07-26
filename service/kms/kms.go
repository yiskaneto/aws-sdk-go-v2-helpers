package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/yiskaneto/aws-sdk-go-v2-helpers/helper_errors"
	"golang.org/x/exp/slices"
)

func main() {
	// aliasName := fmt.Sprintf(aliasName + flagInit())
	// ctx := context.TODO()
	// cfg, err := config.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// client := kms.NewFromConfig(cfg)

	// kmsKey, err := client.ListAliases(ctx, &kms.ListAliasesInput{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// listAllAliases(kmsKey)
	// searchSpecificAliasArN(awsSecretManKMS, kmsKey)

	// Custom search for alias provided on the --alias flag
	// searchSpecificAliasArN(aliasName, kmsKey)
}

// listAllAliases fetches and prints out all the available aliases
func listAllAliases(kmsKey *kms.ListAliasesOutput) {
	for _, aliName := range kmsKey.Aliases {
		fmt.Println(*aliName.AliasName)
	}
}

// searchSpecifiAlias searches for a specific alias and returns the alias's ARN
func searchSpecificAliasArN(ctx context.Context, cfg aws.Config, aliasName string) (bool, error) {
	aliasName = fmt.Sprintf("alias/%v", aliasName)
	kmsClient := kms.NewFromConfig(cfg)
	kmsKey, err := kmsClient.ListAliases(ctx, &kms.ListAliasesInput{})
	if err != nil {
		return false, helper_errors.CheckAWSError(err)
	}
	aliasARN := ""
	aliases := make([]string, 0)
	for _, aliName := range kmsKey.Aliases {
		aliases = append(aliases, *aliName.AliasName)
		if *aliName.AliasName == aliasName {
			aliasARN = *aliName.AliasArn
			break
		}
	}
	if !slices.Contains(aliases, aliasName) {
		return false, errors.New("the provided alias doesn't exist")
	}
	log.Printf("The %v alias was found, its full ARN is: %v", aliasName, aliasARN)
	return true, nil
}

// flagInit initializes a set of flags to be used on the main program
func flagInit() string {

	aliasName := flag.String("alias", "", "Name of the alias")
	flag.Parse()
	if *aliasName == "" {
		log.Fatal("You must pass the --alias flag with the intended alias name")
	}
	return *aliasName
}
