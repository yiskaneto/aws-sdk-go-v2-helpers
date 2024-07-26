package main

import (
	"fmt"

	yiska_aws_helpers "github.com/yiskaneto/aws-sdk-go-v2-helpers/service/s3"
)

func main() {
	yiska_aws_helpers.S3_list_n_buckets()
	fmt.Println("Test")
}
