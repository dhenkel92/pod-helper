package config

type Config struct {
	Kubeconfig string
	Namespace  string
	Labels     []string
	RunConfig  RunConfig
	LogsConfig LogsConfig
}

type RunConfig struct {
	Command []string
}

type LogsConfig struct {
	ContainerIndex int64
}
