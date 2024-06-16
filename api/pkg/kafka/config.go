package kafka

// Config для создания кафки
type Config struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}
