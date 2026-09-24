package lang

import "slices"

import "testing"

// TestAllowListFunctions ensures that only the allowed functions are present.
// It also ensures that assumptions made about how functions are registered (i.e. via makeBaseFunctionTable)
// are still valid. If the upstream function table changes, this test will fail.
func TestAllowListFunctions(t *testing.T) {
	// The subset of functions required by the AI Assistant
	allowedFunctions := []string{"jsondecode", "jsonencode", "keys", "length", "list", "trim", "trimprefix", "trimspace", "trimsuffix", "try"}

	// All the functions exposed by OpenTofu
	allFunctions := makeBaseFunctionTable("", nil)

	for _, allowed := range allowedFunctions {
		if _, ok := allFunctions[allowed]; !ok {
			t.Fatalf("expected function %q to be allowed, but it is not", allowed)
		}
	}

	for all := range allFunctions {
		found := slices.Contains(allowedFunctions, all)
		if !found {
			t.Fatalf("function %q is not allowed", all)
		}
	}

	if len(allFunctions) != len(allowedFunctions) {
		t.Fatalf("expected %d functions, got %d", len(allowedFunctions), len(allFunctions))
	}
}
