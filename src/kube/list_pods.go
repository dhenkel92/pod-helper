package kube

import (
	"github.com/dhenkel92/pod-exec/src/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func ListPods(client *kubernetes.Clientset, namespace string, labels []string) (*v1.PodList, error) {
	pods, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: strings.Join(labels, ",")})
	if err != nil {
		return nil, err
	}
	log.Trace.Printf("%d pods in ns %s\n", len(pods.Items), namespace)
	return pods, err
}
