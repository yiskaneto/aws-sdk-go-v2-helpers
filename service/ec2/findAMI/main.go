package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const amiNameTag = "Soma Value" // This need to be changed to a flag approach

func main() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	ec2Client := ec2.NewFromConfig(cfg)

	// Create empty Input object
	filtro := &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{amiNameTag},
			},
		},
	}

	getAMIs, err := ec2Client.DescribeImages(ctx, filtro)
	if err != nil {
		log.Fatal(err.Error())
	}

	amiList := make(map[string]string, 0)
	// if &getAMIs.Images == nil {
	// 	log.Fatal("\n\n\nNo Images were found for image with tag Name: %s\n\n\n", amiNameTag)
	// }
	for _, v := range getAMIs.Images {
		amiList[*v.CreationDate] = *v.ImageId
	}
	fmt.Println(fetchLatestAmi(amiList))

}

func fetchLatestAmi(amiList map[string]string) string {
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
