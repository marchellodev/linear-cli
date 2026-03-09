package cli

import "testing"

func TestIssuesUpdateCmd_HasRemoveLabelsFlag(t *testing.T) {
	cmd := newIssuesUpdateCmd()
	flag := cmd.Flags().Lookup("remove-labels")
	if flag == nil {
		t.Fatal("issues update command should have --remove-labels flag")
	}
	if flag.Shorthand != "" {
		t.Fatalf("expected no shorthand for --remove-labels, got %q", flag.Shorthand)
	}
}
