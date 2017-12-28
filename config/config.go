package config

type Instance struct {
	Region       string `yaml:"region"`
	Instance     string `yaml:"instance"`
	AWSAccessKey string `yaml:"aws_access_key"`
	AWSSecretKey string `yaml:"aws_secret_key"`
}

type Config struct {
	Instances []Instance `yaml:"instances"`
}
