package eping

import (
	"encoding/json"
	"os"
	"testing"
)

func TestOfficialSearchFixture(t *testing.T) {
	b, err := os.ReadFile("testdata/notifications_search.json")
	if err != nil {
		t.Fatal(err)
	}
	var result SearchResult
	if err := json.Unmarshal(b, &result); err != nil {
		t.Fatal(err)
	}
	if len(result.Items) == 0 || result.Items[0].ID == nil || result.TotalCount == nil {
		t.Fatal("missing typed fields")
	}
}
