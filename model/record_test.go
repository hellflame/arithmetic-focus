package model

import (
	"testing"
)

func TestSaveRecord(t *testing.T) {
	defer removeRecord("999 + 1")
	for i := 0; i < 10; i++ {
		if _, e := SaveRecord("999 + 1", 1000); e != nil {
			t.Fatal(e.Error())
		}
	}
	if ReadRecordsSummary().Corrects != 10 {
		t.Error("failed to summarize")
	}
}
