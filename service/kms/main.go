package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

const (
	awsSecretManKMS = "alias/aws/secretsmanager"
)

var (
	aliasName = "alias/"
)

func main() {
	aliasName = fmt.Sprintf(aliasName + flagInit())
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	client := kms.NewFromConfig(cfg)

	kmsKey, err := client.ListAliases(ctx, &kms.ListAliasesInput{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// listAllAliases(kmsKey)
	// searchSpecificAliasArN(awsSecretManKMS, kmsKey)

	// Custom search for alias provided on the --alias flag
	searchSpecificAliasArN(aliasName, kmsKey)
}

// listAllAliases fetches and prints out all the available aliases
func listAllAliases(kmsKey *kms.ListAliasesOutput) {
	for _, aliName := range kmsKey.Aliases {
		fmt.Println(*aliName.AliasName)
	}
}

// searchSpecifiAlias searches for a specific alias and returns the alias's ARN
func searchSpecificAliasArN(aliasName string, kmsKey *kms.ListAliasesOutput) (aliasARN string) {
	for _, aliName := range kmsKey.Aliases {
		if *aliName.AliasName == aliasName {
			aliasARN = *aliName.AliasArn
			fmt.Println(aliasARN)
			break
		}
	}
	return
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
