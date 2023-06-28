package testutil

import (
	"os"
	"testing"
)

func MarkAsIntegrationTest(t *testing.T) {
	if v := os.Getenv("INTEGRATION"); v != "true" {
		t.SkipNow()
	}
}
