package helper_errors

import (
	"errors"
	"log"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
)

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
