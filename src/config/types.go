package config

type Config struct {
	Kubeconfig string
	Namespace  string
	Labels     []string
	Command    string
}
