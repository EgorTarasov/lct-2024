package s3

type Config struct {
	Host      string   `yaml:"host" env:"MINIO_HOST"`
	Port      int      `yaml:"port" env:"MINIO_PORT"`
	AccessKey string   `yaml:"access_key" env:"MINIO_ACCESS_KEY"`
	SecretKey string   `yaml:"secret_key" env:"MINIO_SECRET_KEY"`
	Ssl       bool     `yaml:"ssl" env:"MINIO_SSL"`
	Buckets   []Bucket `yaml:"buckets"`
}

type Bucket struct {
	Name   string `yaml:"name" `
	Region string `yaml:"region"`
	Lock   bool   `yaml:"lock"`
}
