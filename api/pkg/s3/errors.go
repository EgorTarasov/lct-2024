package s3

import "fmt"

// InitError represents an error that occurred during the initialization of the S3 client.
type InitError struct {
	err error // err is the underlying error that caused the initialization to fail.
}

// Error implements the error interface for InitError. It returns a string representation of the error.
func (ie InitError) Error() string {
	return fmt.Sprintf("Minio Err: %v", ie.err) // Format the error message to include the underlying error.
}
