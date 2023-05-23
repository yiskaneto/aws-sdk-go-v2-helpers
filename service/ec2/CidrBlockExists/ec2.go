package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go"
	"golang.org/x/exp/slices"
)

var (
	cidrBlocku      string   = "172.31.0.0/16"
	DefaultCidrBlks []string = []string{"172.31.0.0/16", "172.32.0.0/16"}
	vpcID           string   = "vpc-invalidVPC"
)

type existingCidrBlks struct {
	cidrBlks []string
}

func (cidr *existingCidrBlks) appendCidrBlk(cidrItem string) (updatedCidrBlks []existingCidrBlks) {
	cidr.cidrBlks = append(cidr.cidrBlks, cidrItem)
	return
}

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	CheckAWSError(err)

	// // Initialize ec2 client
	// ec2Client := ec2.NewFromConfig(cfg)

	log.Println(vpcExistTagNameValue(ctx, cfg, "allowOpsTeam", "true"))

	// // Optionallu, we can create an empty Input object
	// filtro := &ec2.DescribeVpcsInput{}

	// Or we can create the DescribeVpcsInput with a filter
	// filtro := &ec2.DescribeVpcsInput{
	// 	Filters: []types.Filter{
	// 		{
	// 			Name:   aws.String("cidr"),
	// 			Values: []string{cidrBlocku},
	// 		},
	// 	},
	// }

	// // Now we attempt to fetch the vpc information in the current region
	// vpcDescribe, err := ec2Client.DescribeVpcs(ctx, filtro)
	// CheckAWSError(err)

	// // Get the current subnets and their vpc associations.
	// currentCidrBlocks := make([]string, 0)
	// for _, v := range vpcDescribe.Vpcs {
	// 	vpcInfo := fmt.Sprintf("%s associated to %s", *v.CidrBlock, *v.VpcId)
	// 	currentCidrBlocks = append(currentCidrBlocks, vpcInfo)
	// }
	// GetNewCidrBlock(vpcDescribe)
}

// vpcExistWithTagNameValue checks wether a vpc exists with the provided tag value as well as with the tag value.
func vpcExistWithTagNameValue(ctx context.Context, cfg aws.Config, vpcTag, tagValue string) bool {
	ec2Client := ec2.NewFromConfig(cfg)
	tagAssigment := fmt.Sprintf("tag:%s", vpcTag)
	// Create the DescribeVpcsInput with a filter
	filtro := &ec2.DescribeVpcsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String(tagAssigment),
				Values: []string{tagValue},
			},
		},
	}
	getVpcTagName, err := ec2Client.DescribeVpcs(ctx, filtro)
	CheckAWSError(err)
	if len(getVpcTagName.Vpcs) >= 1 {
		log.Println("There is already a vpc with the passed tag name and tag value, skipping")
		return false
	} else {
		log.Println("No vpc found with the provided tag not tag value")
		return true
	}
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

//// Generate a new CIDR block, 3 functions

// genNewCidrBlk calls GenerateNewCidrBlock to get an unused cidr block
func genNewCidrBlk(ctx *context.Context, cfg *aws.Config) string {
	ec2Client := ec2.NewFromConfig(*cfg)
	vpcDescribe, err := ec2Client.DescribeVpcs(*ctx, &ec2.DescribeVpcsInput{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return GetNewCidrBlock(vpcDescribe)
}

// CheckAllCidrBlocks return all the current cidr blocks in the current account
func CheckAllCidrBlocks(filter *ec2.DescribeVpcsOutput) (allCidrs []string) {
	allCidrs = make([]string, 0)
	for _, v := range filter.Vpcs {
		allCidrs = append(allCidrs, *v.CidrBlock)
	}
	return
}

// GetNewCidrBlock returns an unsed CIDR block from a given lists
func GetNewCidrBlock(filter *ec2.DescribeVpcsOutput) (finalCidr string) {
	listCurrentCidrBlocks := CheckAllCidrBlocks(filter)
	for _, v := range listCurrentCidrBlocks {
		DefaultCidrBlks = append(DefaultCidrBlks, v)
	}
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
		if slices.Contains(DefaultCidrBlks, CB) {
			continue
		} else if !slices.Contains(DefaultCidrBlks, CB) {
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

// GetLatestAmi returns the amiID of the passed name tag, if there are duplicates it returns the newest one.
// Example on how to use it: myAMI := SearchAmiID(ctx, cfg, "coolest-tag")
func SearchAmiID(ctx *context.Context, cfg *aws.Config, amiNameTag string) string {
	ec2Client := ec2.NewFromConfig(*cfg)
	filtro := &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{amiNameTag},
			},
		},
	}

	getAMIs, err := ec2Client.DescribeImages(*ctx, filtro)
	if err != nil {
		log.Fatal(err.Error())
	}

	amiList := make(map[string]string, 0)
	if len(getAMIs.Images) == 0 {
		log.Fatalf("\n\n\nERROR: No images were found containing the tag Name: %s\n\n\n", amiNameTag)
	}
	for _, v := range getAMIs.Images {
		amiList[*v.CreationDate] = *v.ImageId
	}

	cDate := make([]string, 0)

	for date, _ := range amiList {
		cDate = append(cDate, date)
	}
	maxDate := cDate[0]
	for _, sDate := range cDate {
		if sDate > maxDate {
			maxDate = sDate
		}
	}
	return amiList[maxDate]
}

// RemoveAmiId attetmps to deregister the passed ami
// example on how to use it with the help of the SearchAmiID()
// myAMI := SearchAmiID(ctx, cfg, "coolest-tag")
// RemoveAmiId(*ctx, *cfg, &myAMI)
func RemoveAmiId(ctx context.Context, cfg aws.Config, amiID *string) {
	ec2Client := ec2.NewFromConfig(cfg)
	filtro := &ec2.DeregisterImageInput{
		ImageId: amiID,
	}
	_, err := ec2Client.DeregisterImage(ctx, filtro)
	CheckAWSError(err)
	log.Printf("AMI ID %s was successfully deregistered", *amiID)
}

// AWSLoadCreds loads AWS conf file from ~/.aws or environment variables.
func AWSLoadCreds() (string, *aws.Config, *context.Context) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	if aws.ToString(&cfg.Region) == "" {
		log.Fatal("ERROR: Cannot read the AWS credentials on ~/.aws  nor from the environmental variables,verify that you have authenticated")
	}

	svc := sts.NewFromConfig(cfg)
	testout, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	CheckAWSError(err)
	return *testout.Account, &cfg, &ctx
}

func CheckAWSError(err error) {
	if err != nil {
		var ae smithy.APIError
		var re *awshttp.ResponseError
		if errors.As(err, &ae) {
			log.Printf("Failure Code: %s, Message: %s, Fault is on: %s", ae.ErrorCode(), ae.ErrorMessage(), ae.ErrorFault().String())
			if errors.As(err, &re) {
				log.Fatalf("requestID: %s", re.ServiceRequestID())
			}
		}
		return
	}
}
