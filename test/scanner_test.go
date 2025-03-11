package test

import (
	"CI-CD-Security-Scanner/internal"
	"testing"
)

func TestRunScan(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RunScan panicked: %v", r)
		}
	}()
	internal.RunScan("../examples/secure-gitlab.yml")
}
