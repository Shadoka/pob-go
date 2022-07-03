package data_import

import "testing"

func TestImportPassives(t *testing.T) {
	character_data, err := ImportCharacterPassives("TheShadoka", "CrashingShadoka", "pc")
	if err != nil {
		t.Fatalf("unexpected error occurred: %v", err)
	}

	if character_data == nil {
		t.Fatalf("result from import was unexpectedly nil")
	}

	if len(character_data.Items) != 12 {
		t.Fatalf("expected 12 items in imported dataset, got %v", len(character_data.Items))
	}
}
