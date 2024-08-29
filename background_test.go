package byo_context

import "testing"

func Test_Background(t *testing.T) {
	bg := Background()
	if bg.Err() != nil {
		t.Error("Background().Err() should return nil")
	}
	if _, ok := bg.Deadline(); ok {
		t.Error("Background().Deadline() should return false")
	}
	if bg.Done() != nil {
		t.Error("Background().Done() should return nil")
	}
	if bg.Value(nil) != nil {
		t.Error("Background().Value() should return nil")
	}
}
