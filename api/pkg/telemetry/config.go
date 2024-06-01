package telemetry

// Config для трассировки
type Config struct {
	OTLPEndpoint string `yaml:"endpoint"`
}
