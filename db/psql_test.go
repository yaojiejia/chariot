package db

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	p := NewPSQL("localhost", "5432", "alex", "jiayaojie0715", "newdb")
	_, err := p.Connect()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
