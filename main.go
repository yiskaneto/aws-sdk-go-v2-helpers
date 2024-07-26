package main

import (
	"fmt"

	aws_helpers "github.com/escanoru/aws-sdk-go-v2-helper/service/s3"
)

func main() {
	s3.s3_list_n_buckets(10)
	fmt.Println("Test")
}

