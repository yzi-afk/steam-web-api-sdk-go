package steam

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := New("test")
	if err != nil {
		t.Fatal(err)
	}

	if client.apiKey != "test" {
		t.Errorf("expected apiKey to be 'test', got %s", client.apiKey)
	}

}
