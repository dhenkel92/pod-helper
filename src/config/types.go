package config

type Config struct {
	Kubeconfig string
	Namespaces []string
	Labels     []string

	ContainerIndex int64
	Container      string

	RunConfig  RunConfig
	LogsConfig LogsConfig
}

type RunConfig struct {
	Entrypoint []string
	Command    string
}

type LogsConfig struct {
	Tail int64
}
