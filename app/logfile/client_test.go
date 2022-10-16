package logfile

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("../../data/qconsole.log")
	if err != nil {
		t.Error(err)
	}
}

func TestNewClient_fileNotExist(t *testing.T) {
	_, err := NewClient("../../data/unknown.log")
	if err == nil {
		t.Error(err)
	}
}

func TestRunClient(t *testing.T) {
	client, err := NewClient("../../data/qconsole.log")
	if err != nil {
		t.Error(err)
	}

	err = client.Run()
	if err != nil {
		t.Error(err)
	}
}

func TestRunClient_readFailed(t *testing.T) {
	client, err := NewClient("../../data")
	if err != nil {
		t.Error(err)
	}

	err = client.Run()
	if err == nil {
		t.Error(err)
	}
}
