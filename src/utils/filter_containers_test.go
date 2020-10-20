package utils

import (
	"testing"

	"github.com/dhenkel92/pod-helper/src/config"
	v1 "k8s.io/api/core/v1"
)

func TestShouldReturnAllContainersIfNoConfig(t *testing.T) {
	conf := config.Config{Container: "", ContainerIndex: -1}
	containers := []v1.Container{{Name: "bar"}, {Name: "foo"}, {Name: "foobar"}}

	filtered, err := FilterContainers(&containers, &conf)
	if err != nil {
		t.Error(err)
	}
	if len(containers) != len(filtered) {
		t.Fail()
	}
}

func TestFilterByName(t *testing.T) {
	conf := config.Config{Container: "foo"}
	containers := []v1.Container{{Name: "bar"}, {Name: "foo"}}

	res, err := FilterContainers(&containers, &conf)
	if err != nil {
		t.Error(err)
	}

	if len(res) == 0 {
		t.Fail()
	}
}

func TestShouldReturnErrorWhenNoContainerMatches(t *testing.T) {
	conf := config.Config{Container: "foo"}
	containers := []v1.Container{{Name: "bar"}}

	_, err := FilterContainers(&containers, &conf)
	if err == nil {
		t.Fail()
	}
}

func TestShouldReturnContainerAtIndex(t *testing.T) {
	conf := config.Config{ContainerIndex: 2}
	containers := []v1.Container{{Name: "foo"}, {Name: "foobar"}, {Name: "bar"}}

	filtered, err := FilterContainers(&containers, &conf)
	if err != nil {
		t.Error(err)
	}
	if len(filtered) == 0 {
		t.Fail()
	}
	if containers[2].Name != filtered[0].Name {
		t.Fail()
	}
}

func TestShouldReturnErrorIfIndexOutOfBound(t *testing.T) {
	conf := config.Config{ContainerIndex: 20}
	containers := []v1.Container{{Name: "bar"}}

	_, err := FilterContainers(&containers, &conf)
	if err == nil {
		t.Fail()
	}
}
