package s3

import (
	"context"
	"log"
	"math"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ListNBuckets uses the AWS SDK for Go V2 to create an Amazon Simple Storage Service
// (Amazon S3) client and list up to 10 buckets in your account.
// This example uses the default settings specified in your shared credentials
// and config files.
func ListNBuckets(nBuckets int) string {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("Couldn't load default configuration. Have you set up your AWS account?")
		log.Println(err)
		return failure
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	count := nBuckets
	log.Printf("Let's list up to %v buckets for your account.\n", count)
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		return failure
	}

	if math.Signbit(float64(nBuckets)) {
		return failure
	}
	if len(result.Buckets) == 0 {
		outcome := "You don't have any buckets!"
		return outcome
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count] {
			log.Printf("\t%v\n", *bucket.Name)
		}
		return success
	}

}
