package main

import (
	"fmt"
	yiska_aws_helpers "github.com/yiskaneto/aws-sdk-go-v2-helpers/service/s3"
)

func main() {
	yiska_aws_helpers.s3_list_n_buckets(10)
	fmt.Println("Test")
}

