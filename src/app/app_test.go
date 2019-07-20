package main

import "testing"

func TestServer(t *testing.T) {
	textHtml := greeting("Code.education Rocks!")
	if textHtml != "<b>Code.education Rocks!</b>" {
		t.Errorf("Text html was incorrect, got: %s, want: %s.", textHtml, "<b>Code.education Rocks!</b>")
	}
}
