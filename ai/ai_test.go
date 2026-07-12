package ai

import (
	"context"
	"testing"
)

func TestMockProvider(t *testing.T) {
	provider := &MockProvider{}
	ctx := context.Background()

	reply, err := provider.GenerateResponse(ctx, "Hello", nil)
	if err != nil {
		t.Fatalf("Mock provider returned error: %v", err)
	}

	if reply == "" {
		t.Error("Mock provider returned empty reply")
	}

	score, err := provider.QualifyLead(ctx, "Standard inquiry about website.")
	if err != nil {
		t.Fatalf("Mock qualify lead returned error: %v", err)
	}

	if score.Score == 0 {
		t.Error("Mock qualify lead returned 0 score")
	}
}
