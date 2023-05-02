package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

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
		cidrBlks: []string{"10.10.0.0/16", "10.11.0.0/16", "10.12.0.0/16", "172.31.0.0/16"},
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
		"10.36.0.0/16",
		"10.37.0.0/16",
		"10.38.0.0/16",
		"10.39.0.0/16",
		"10.40.0.0/16",
		"10.41.0.0/16",
		"10.42.0.0/16",
		"10.43.0.0/16",
		"10.44.0.0/16",
		"10.45.0.0/16",
		"10.46.0.0/16",
		"10.47.0.0/16",
		"10.48.0.0/16",
		"10.49.0.0/16",
		"10.50.0.0/16",
		"10.51.0.0/16",
		"10.52.0.0/16",
		"10.53.0.0/16",
		"10.54.0.0/16",
		"10.55.0.0/16",
		"10.56.0.0/16",
		"10.57.0.0/16",
		"10.58.0.0/16",
		"10.59.0.0/16",
		"10.60.0.0/16",
	}
	randomIndex := rand.Intn(len(cidrList))
	if !slices.Contains(currentCidr, cidrList[randomIndex]) {
		finalCidr = cidrList[randomIndex]
		log.Printf("SUCCESS: Assigning new available CIDR block: %s", cidrList[randomIndex])
	} else {
		log.Fatal("ERROR: No cidr availables")
	}
	return
}
