package lib

import (
	"testing"
)

func TestConfigs(t *testing.T) {
	var c Config
	c.GetConfig()

	if c.Host != "localhost" {
		t.Errorf("host names are different")
	}
}
