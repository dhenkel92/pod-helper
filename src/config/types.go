package config

type Config struct {
	Kubeconfig string
	Namespace  string
	Labels     []string
	RunConfig  RunConfig
	LogsConfig LogsConfig
}

type RunConfig struct {
	Entrypoint []string
	Command    string
}

type LogsConfig struct {
	ContainerIndex int64
	Tail           int64
}
