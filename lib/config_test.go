package lib

import (
	"testing"
)

func configs(t *testing.T) {
	var c Config
	c.GetConfig()

	if c.Host != "123" {
		t.Errorf("host names are different")
	}
}
