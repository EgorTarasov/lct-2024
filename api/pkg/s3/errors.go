package s3

import "fmt"

type S3InitError struct {
	err error
}

// Error implements error.
func (me S3InitError) Error() string {
	return fmt.Sprintf("Minio Err: %v", me.err)
}
