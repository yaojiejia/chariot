package db

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	p := NewPSQL("localhost", "5432", "alex", "jiayaojie0715", "newdb")
	err := p.Connect()
	if err != nil {
		t.Errorf("Error connecting to the database: %v", err)
	}

}
