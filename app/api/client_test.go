package api

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("/var/run/docker.sock", "q3srv", time.Second*5)
	if err != nil {
		t.Error(err)
	}
}

func TestNewClient_fileNotExist(t *testing.T) {
	_, err := NewClient("/var/run/unknown", "q3srv", time.Second*5)
	if err == nil {
		t.Error("Must be error")
	}
}

func TestGetContainer(t *testing.T) {
	client, err := NewClient("/var/run/docker.sock", "q3srv", time.Second*5)
	if err != nil {
		t.Error(err)
	}

	_, err = client.getcontainer()
	if err != nil {
		t.Error(err)
	}
}

func TestGetContainer_failed(t *testing.T) {
	client, err := NewClient("/var/run/docker.sock", "unknown", time.Second*5)
	if err != nil {
		t.Error(err)
	}

	container, err := client.getcontainer()
	if err != nil {
		t.Error(err)
	}

	if container.ID != "" {
		t.Error("Must be empty")
	}
}

func TestGetLogs(t *testing.T) {
	client, err := NewClient("/var/run/docker.sock", "q3srv", time.Second*5)
	if err != nil {
		t.Error(err)
	}

	container, err := client.getcontainer()
	if err != nil {
		t.Error(err)
	}

	_, err = client.getlogs(container.ID)
	if err != nil {
		t.Error(err)
	}
}
