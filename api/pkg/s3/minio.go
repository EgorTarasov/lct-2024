package s3

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// S3 represents an S3 client.
type S3 struct {
	cfg  *Config
	conn *minio.Client
}

// MustNew creates a new S3 client. It will return an error if the client cannot be created.
func MustNew(cfg *Config) (*S3, error) {
	uri := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	minioClient, err := minio.New(uri, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.Ssl,
	})

	if err != nil {
		return nil, InitError{err: err}
	}

	return &S3{
		cfg:  cfg,
		conn: minioClient,
	}, nil
}

// CreateBucket creates a new bucket with the provided name and options.
func (s3 S3) CreateBucket(ctx context.Context, bucket *Bucket) error {
	return s3.conn.MakeBucket(ctx, bucket.Name, minio.MakeBucketOptions{
		Region: bucket.Region, ObjectLocking: bucket.Lock,
	})
}

// CreateBuckets creates all the buckets specified in the S3 client's configuration.
func (s3 S3) CreateBuckets(ctx context.Context) error {
	for _, bucket := range s3.cfg.Buckets {
		if exists, _ := s3.BucketExists(ctx, bucket.Name); !exists {
			err := s3.conn.MakeBucket(ctx, bucket.Name, minio.MakeBucketOptions{
				Region: bucket.Region, ObjectLocking: bucket.Lock,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ListBuckets lists all the buckets in the S3 client.
func (s3 S3) ListBuckets(ctx context.Context) ([]minio.BucketInfo, error) {
	return s3.conn.ListBuckets(ctx)
}

// BucketExists checks if a bucket with the provided name exists.
func (s3 S3) BucketExists(ctx context.Context, name string) (bool, error) {
	return s3.conn.BucketExists(ctx, name)
}

// RemoveBucket removes a bucket with the provided name.
func (s3 S3) RemoveBucket(ctx context.Context, name string) error {
	return s3.conn.RemoveBucket(ctx, name)
}

// ListObjects lists all the objects in a bucket.
func (s3 S3) ListObjects(ctx context.Context, bucket string, opts minio.ListObjectsOptions) <-chan minio.ObjectInfo {
	return s3.conn.ListObjects(ctx, bucket, opts)
}

// GetObject retrieves an object from a bucket.
func (s3 S3) GetObject(ctx context.Context, bucket, obj string, opts minio.GetObjectOptions) (*minio.Object, error) {
	return s3.conn.GetObject(ctx, bucket, obj, opts)
}

// PresignedGetObject generates a presigned URL to retrieve an object.
func (s3 S3) PresignedGetObject(ctx context.Context, bucket, obj string, exp time.Duration, params url.Values) (*url.URL, error) {
	return s3.conn.PresignedGetObject(ctx, bucket, obj, exp, params)
}

// PresignedPutObject generates a presigned URL to put an object.
func (s3 S3) PresignedPutObject(ctx context.Context, bucketName, objectName string, expires time.Duration) (*url.URL, error) {
	return s3.conn.PresignedPutObject(ctx, bucketName, objectName, expires)
}

// PutObject uploads an object to a bucket.
func (s3 S3) PutObject(ctx context.Context, bucket, obj string, reader io.Reader, size int64, opts minio.PutObjectOptions) (minio.UploadInfo, error) {
	return s3.conn.PutObject(ctx, bucket, obj, reader, size, opts)
}

// RemoveObject removes an object from a bucket.
func (s3 S3) RemoveObject(ctx context.Context, bucket, obj string, opts minio.RemoveObjectOptions) error {
	return s3.conn.RemoveObject(ctx, bucket, obj, opts)
}
