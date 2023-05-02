package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"golang.org/x/exp/slices"
)

var (
	cidrBlocku string = "172.31.0.0/16"
	vpcID      string = "vpc-invalidVPC"
)

type existingCidrBlks struct {
	cidrBlks []string
}

func (cidr *existingCidrBlks) appendCidrBlk(cidrItem string) (updatedCidrBlks []existingCidrBlks) {
	cidr.cidrBlks = append(cidr.cidrBlks, cidrItem)
	return
}

func main() {
	defaultCidrBlks := existingCidrBlks{
		cidrBlks: []string{"10.10.0.0/16", "10.11.0.0/16", "10.12.0.0/16", "172.31.0.0/16", "10.36.0.0/16", "10.37.0.0/16"},
	}
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
	listCurrentCidrBlocks := CheckAllCidrBlocks(vpcDescribe)
	for _, v := range listCurrentCidrBlocks {
		defaultCidrBlks.appendCidrBlk(v)
	}
	GetNewCidrBlock(defaultCidrBlks.cidrBlks)
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

func CheckAllCidrBlocks(filter *ec2.DescribeVpcsOutput) (allCidrs []string) {
	allCidrs = make([]string, 0)
	for _, v := range filter.Vpcs {
		allCidrs = append(allCidrs, *v.CidrBlock)
	}
	return
}

// GetNewCidrBlock returns an unsed CIDR block from a given lists
func GetNewCidrBlock(currentCidr []string) (finalCidr string) {
	cidrList := []string{
		"10.100.0.0/16",
		"10.101.0.0/16",
		"10.102.0.0/16",
		"10.103.0.0/16",
		"10.104.0.0/16",
		"10.105.0.0/16",
		"10.106.0.0/16",
		"10.107.0.0/16",
		"10.108.0.0/16",
		"10.109.0.0/16",
		"10.110.0.0/16",
		"10.111.0.0/16",
		"10.112.0.0/16",
		"10.113.0.0/16",
		"10.114.0.0/16",
		"10.115.0.0/16",
		"10.116.0.0/16",
		"10.117.0.0/16",
		"10.118.0.0/16",
		"10.119.0.0/16",
		"10.120.0.0/16",
		"10.121.0.0/16",
		"10.122.0.0/16",
		"10.123.0.0/16",
		"10.124.0.0/16",
		"10.125.0.0/16",
	}

	for _, CB := range cidrList {
		if slices.Contains(currentCidr, CB) {
			continue
		} else if !slices.Contains(currentCidr, CB) {
			finalCidr = CB
			log.Printf("SUCCESS: Assigning new available CIDR block: %s", finalCidr)
			break
		}
	}
	if finalCidr == "" {
		log.Fatal("ERROR: No available cidr, consider expanding the CIDR block on GetNewCidrBlock")
	}

	// The approach below is based on a random selection, but is not the best option:
	// randomIndex := rand.Intn(len(cidrList))
	// if !slices.Contains(currentCidr, cidrList[randomIndex]) {
	// 	finalCidr = cidrList[randomIndex]
	// 	log.Printf("SUCCESS: Assigning new available CIDR block: %s", cidrList[randomIndex])
	// } else {
	// 	log.Fatal("ERROR: No cidr availables")
	// }
	return
}
