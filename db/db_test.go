package db

import (
	"testing"
)

func TestNewInstance(t *testing.T) {
	i := NewInstance("demo", "12345")

	if i.host != "demo" {
		t.Error("Expected host to be 'demo'")
	}

	if i.port != "12345" {
		t.Error("Expected port to be '12345'")
	}
}

func TestCreateCtx(t *testing.T) {
	i := NewInstance("demo", "12345")
	i.createCtx()

	if i.ctx == nil {
		t.Error("Expected context to be set")
	}
}

func TestCreateClient(t *testing.T) {
	i := NewInstance("demo", "12345")
	i.createClient()

	if i.client == nil {
		t.Error("Expected client to be set")
	}
}

func TestCreateCollection(t *testing.T) {
	i := NewInstance("demo", "12345")
	i.createClient()
	i.createCollection("test", "events")

	if i.collection == nil {
		t.Error("Expected collection to be set")
	}
}

func TestConnect(t *testing.T) {
	i := NewInstance("demo", "12345")
	i.createClient()
	i.createCollection("test", "events")
	i.createCtx()

	err := i.client.Connect(i.ctx)
	if err != nil {
		t.Error("Expected no error")
	}
}

func TestClose(t *testing.T) {
	i := NewInstance("demo", "12345")
	i.createClient()
	i.createCollection("test", "events")
	i.createCtx()

	i.Close()

	if i.client != nil {
		t.Error("Expected client to be nil")
	}
}

func TestAddEntry(t *testing.T) {
	i := NewInstance("localhost", "27017")
	i.createClient()
	i.createCollection("test", "events")
	i.createCtx()

	err := i.client.Connect(i.ctx)
	if err != nil {
		t.Error("Expected no error")
	}

	_, err = i.AddEntry("senderX", "testAddEntry test running")
	if err != nil {
		t.Error("Expected no error")
	}
}

func TestSaveEntry(t *testing.T) {
	_, err := SaveEntry("senderX", "TestSaveEntry test running")
	if err != nil {
		t.Error("Expected no error")
	}
}
