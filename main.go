package main

import (
	yiskaAwsS3Helpers "github.com/yiskaneto/aws-sdk-go-v2-helpers/service/s3"
)

func main() {
	yiskaAwsS3Helpers.ListNBuckets(10)
}
