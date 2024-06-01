package s3

// Config represents the configuration for the S3 client.
type Config struct {
	Host      string   `yaml:"host" env:"MINIO_HOST"`             // Host is the hostname of the S3 server.
	Port      int      `yaml:"port" env:"MINIO_PORT"`             // Port is the port number on which the S3 server is running.
	AccessKey string   `yaml:"access_key" env:"MINIO_ACCESS_KEY"` // AccessKey is the access key for S3 server authentication.
	SecretKey string   `yaml:"secret_key" env:"MINIO_SECRET_KEY"` // SecretKey is the secret key for S3 server authentication.
	Ssl       bool     `yaml:"ssl" env:"MINIO_SSL"`               // Ssl indicates whether to use SSL for the connection.
	Buckets   []Bucket `yaml:"buckets"`                           // Buckets is a list of buckets to be managed by the S3 client.
}

// Bucket represents a bucket in the S3 server.
type Bucket struct {
	Name   string `yaml:"name"`   // Name is the name of the bucket.
	Region string `yaml:"region"` // Region is the region in which the bucket is located.
	Lock   bool   `yaml:"lock"`   // Lock indicates whether object locking should be enabled for the bucket.
}
