package config

import (
	"testing"
)

func TestShouldEndConversation(t *testing.T) {
	if createExpectForTestShouldEndConversation("").ShouldEndConversation() {
		t.Fatal("ShouldEndConversation() should be false, if the endConversation is empty")
	}
	if createExpectForTestShouldEndConversation("false").ShouldEndConversation() {
		t.Fatal("ShouldEndConversation() should be false, if the endConversation is \"false\"")
	}
	if !createExpectForTestShouldEndConversation("true").ShouldEndConversation() {
		t.Fatal("ShouldEndConversation() should be true, if the endConversation is \"true\"")
	}
}

func createExpectForTestShouldEndConversation(endConversation string) *Expect {
	expect := Expect{
		Action: "",
		IntentName: "",
		Parameters: nil,
		Contexts: nil,
		Speech: "",
		Speeches: nil,
		EndConversation: endConversation,
	}
	return &expect
}