package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var (
	cidrBlocku string = "10.68.0.0/16"
	vpcID      string = "vpc-invalidVPC"
)

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	ec2Client := ec2.NewFromConfig(cfg)

	// We can create an empty Input object
	filtro := &ec2.DescribeVpcsInput{}

	// Or we can create the DescribeVpcsInput with a filter
	// filtro := &ec2.DescribeVpcsInput{
	// 	Filters: []types.Filter{
	// 		{
	// 			Name:   aws.String("cidr"),
	// 			Values: []string{cidrBlocku},
	// 		},
	// 	},
	// }

	vpcDescribe, err := ec2Client.DescribeVpcs(ctx, filtro)
	if err != nil {
		log.Fatal(err.Error())
	}

	currentCidrBlocks := make([]string, 0)
	for _, v := range vpcDescribe.Vpcs {
		vpcInfo := fmt.Sprintf("%s associated to %s", *v.CidrBlock, *v.VpcId)
		currentCidrBlocks = append(currentCidrBlocks, vpcInfo)
	}

	// Optionally, we can directly access the fields of a VPC type:
	// log.Printf("VPC %s contains CIDR block %s", *vpcDescribe.Vpcs[0].VpcId, *vpcDescribe.Vpcs[0].CidrBlock)
	log.Println(CheckCidrBlock(vpcDescribe, currentCidrBlocks))
}

func CheckCidrBlock(filter *ec2.DescribeVpcsOutput, currentCidrBlocks []string) string {
	for _, v := range filter.Vpcs {
		if *v.CidrBlock == cidrBlocku {
			log.Printf("The provided CIDR block %s is already part of %s, you must specify another CIDR block", cidrBlocku, *v.VpcId)
			log.Printf("Current CIDR blocks:")
			for _, cblock := range currentCidrBlocks {
				log.Print(cblock)
			}
			log.Fatal("Take a look at the listed cidr blocks above and provide one that is NOT in the list")
		}
	}
	return "The provided CIDR block is available, continuing..."
}
