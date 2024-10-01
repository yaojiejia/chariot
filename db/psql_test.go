package db

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	p := NewPSQL("localhost", "5432", "alex", "", "newdb")
	err := p.Connect()
	if err != nil {
		t.Errorf("Error connecting to the database: %v", err)
	}

}
