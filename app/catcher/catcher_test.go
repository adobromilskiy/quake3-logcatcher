package catcher

import (
	"context"
	"testing"
)

func TestNewCatcher(t *testing.T) {
	_, err := New(context.Background(), "mongodb://localhost:27017", "quake3")
	if err != nil {
		t.Error(err)
	}
}

func TestNewCatcher_failed(t *testing.T) {
	_, err := New(context.Background(), "mongo://localhost:27017", "quake3")
	if err == nil {
		t.Error("Must be error")
	}
}

func TestIsNewGame(t *testing.T) {
	ct := Catcher{}
	if !ct.isNewGame("Game_Start: FFA in arena 0") {
		t.Error("Must be true")
	}

	if ct.isNewGame("Game_Start:  in arena 0") {
		t.Error("Must be false")
	}
}

func TestIsEndGame(t *testing.T) {
	ct := Catcher{}
	if !ct.isEndGame("Exit: Fraglimit hit.") {
		t.Error("Must be true")
	}

	if ct.isEndGame("exit:") {
		t.Error("Must be false")
	}
}

func TestGetKill(t *testing.T) {
	ct := Catcher{}
	kill := ct.getKill("Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0")
	if kill.Killer != "JavaScripter" {
		t.Error("Must be JavaScripter")
	}
	if kill.Victim != "twist" {
		t.Error("Must be twist")
	}
}

func TestGetKill_notFound(t *testing.T) {
	ct := Catcher{}
	kill := ct.getKill("Item: 1 weapon_rocketlauncher")
	if kill != nil {
		t.Error("Must be nil")
	}
}

func TestDo(t *testing.T) {
	ct, err := New(context.Background(), "mongodb://localhost:27017", "quake3")
	if err != nil {
		t.Error(err)
	}

	logs := []byte(`Game_Start: FFA in arena 0
		Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0
		Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0
		Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0
		Game_Start: FFA in arena 0
		Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0
		Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0
		Kill: 1 2 6: JavaScripter killed twist by MOD_ROCKET 5 in arena 0
		Exit: Timelimit hit.`)

	if err := ct.Do(logs); err != nil {
		t.Error(err)
	}
}
