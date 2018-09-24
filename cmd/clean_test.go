package cmd

import "testing"

func testExtractDuration(t *testing.T) {
	parameter := "5d"
	var duration = ExtractDuration(parameter)
	t.Logf("running test: %s, %d", parameter, duration)
	if duration == 0 {
		t.Error("duration", duration)
	}
}
