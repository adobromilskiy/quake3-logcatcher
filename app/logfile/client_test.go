package logfile

import (
	"context"
	"testing"

	"github.com/adobromilskiy/quake3-logcatcher/app/catcher"
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

	ct, _ := catcher.New(context.Background(), "mongodb://localhost:27017", "quake3")
	err = client.Run(ct)
	if err != nil {
		t.Error(err)
	}
}

func TestRunClient_readFailed(t *testing.T) {
	client, err := NewClient("../../data")
	if err != nil {
		t.Error(err)
	}

	ct, _ := catcher.New(context.Background(), "mongodb://localhost:27017", "quake3")
	err = client.Run(ct)
	if err == nil {
		t.Error(err)
	}
}
