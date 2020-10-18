package kube

import (
	"github.com/dhenkel92/pod-helper/src/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func listPodsForNamespace(client *kubernetes.Clientset, namespace string, labels []string) ([]v1.Pod, error) {
	pods, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: strings.Join(labels, ",")})
	if err != nil {
		return nil, err
	}
	log.Trace.Printf("%d pods in ns %s\n", len(pods.Items), namespace)
	return pods.Items, err
}

func ListPods(client *kubernetes.Clientset, namespaces []string, labels []string) ([]v1.Pod, error) {
	var results []v1.Pod
	for _, ns := range namespaces {
		pods, _ := listPodsForNamespace(client, ns, labels)
		results = append(results, pods...)
	}
	return results, nil
}
