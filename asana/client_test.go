package asana

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewClient(Authorization{})
	if err != nil {
		t.Error(err.Error())
	}
}
