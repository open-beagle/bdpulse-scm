package gitlab

import (
	"testing"
	"time"
)

func TestConvertBranchUpdated(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	b := &branch{
		Name: "dev",
		Commit: struct {
			ID            string    `json:"id"`
			AuthorDate    time.Time `json:"authored_date"`
			CommittedDate time.Time `json:"committed_date"`
			Created       time.Time `json:"created_at"`
		}{
			ID:            "123456",
			CommittedDate: now,
		},
	}
	ref := convertBranch(b)
	if ref.Name != "dev" {
		t.Errorf("expected dev, got %s", ref.Name)
	}
	if !ref.Updated.Equal(now) {
		t.Errorf("expected %v, got %v", now, ref.Updated)
	}
}
