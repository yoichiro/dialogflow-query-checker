package config

import (
	"testing"
	"time"
)

func TestDetermineDateMacro(t *testing.T) {
	def := Definition{
		DateMacroFormat: "2006-01-02",
		Tests: []Test{
			{
				Condition: Condition{
					Query: "start1${date.today}end1",
				},
				Expect: Expect{
					Speech: "start2${date.today}end2",
					Speeches: []string{
						"start3${date.today}end3",
					},
					Parameters: map[interface{}]interface{}{
						"key1": "start4${date.today}end4",
						"key2": "start5${date.today}end5",
					},
				},
			},
		},
	}
	determineDateMacro(&def)
	test := def.Tests[0]
	date := time.Now().Format("2006-01-02")
	if test.Condition.Query != "start1" + date + "end1" {
		t.Fatalf("test.Condition.Query is not expected value: %s", test.Condition.Query)
	}
	if test.Expect.Speech != "start2" + date + "end2" {
		t.Fatalf("test.Expect.Speech is not expected value: %s", test.Expect.Speech)
	}
	if test.Expect.Speeches[0] != "start3" + date + "end3" {
		t.Fatalf("test.Expect.Speeches[0] is not expected value: %s", test.Expect.Speeches[0])
	}
	if test.Expect.Parameters["key1"] != "start4" + date + "end4" {
		t.Fatalf("test.Expect.Parameters[\"key1\"] is not expected value: %s", test.Expect.Parameters["key1"])
	}
	if test.Expect.Parameters["key2"] != "start5" + date + "end5" {
		t.Fatalf("test.Expect.Parameters[\"key2\"] is not expected value: %s", test.Expect.Parameters["key2"])
	}
}