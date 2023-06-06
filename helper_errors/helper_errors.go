package helper_errors

import (
	"errors"
	"fmt"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
)

// CheckAWSError check if the passed error is nil, if not then we check the type of error, it returns an error
func CheckAWSError(err error) error {
	if err != nil {
		var ae smithy.APIError
		var re *awshttp.ResponseError
		if errors.As(err, &ae) {
			return fmt.Errorf("Failure Code: %s, Message: %s, Fault is on: %s", ae.ErrorCode(), ae.ErrorMessage(), ae.ErrorFault().String())
		}
		if errors.As(err, &re) {
			return fmt.Errorf("requestID: %s", re.ServiceRequestID())
		}
	}
	return nil
}
