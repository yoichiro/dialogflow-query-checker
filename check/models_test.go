package check

import (
	"testing"
	"container/list"
	"reflect"
)

func TestAllTestResults(t *testing.T) {
	holder := createTestData()
	actual := holder.AllTestResults()
	if reflect.TypeOf(actual).String() != "[]*check.TestResult" {
		t.Fatalf("actual is %s, not []*check.TestResult", reflect.TypeOf(actual))
	}
	if len(actual) != 2 {
		t.Fatalf("The length of actual is %d, not 2", len(actual))
	}
	if actual[0].Prefix != "prefix1" {
		t.Fatal("actual[0] is not expected TestResult")
	}
	if actual[1].Prefix != "prefix2" {
		t.Fatal("actual[1] is not expected TestResult")
	}
}

func TestAllAssertResults(t *testing.T) {
	holder := createTestData()
	actual := holder.AllAssertResults()
	if reflect.TypeOf(actual).String() != "[]*check.AssertResult" {
		t.Fatalf("actual is %s, not []*check.AssertResult", reflect.TypeOf(actual))
	}
	if len(actual) != 2 {
		t.Fatalf("The length of actual is %d, not 2", len(actual))
	}
	if actual[0].Message != "message1" {
		t.Fatal("actual[0] is not expected AssertResult")
	}
	if actual[1].Message != "message2" {
		t.Fatal("actual[1] is not expected TestResult")
	}
}

func TestAllAssertResultCount(t *testing.T) {
	holder := createTestData()
	actual := holder.AllAssertResultCount()
	if actual != 2 {
		t.Fatalf("actual is %d, not 2", actual)
	}
}

func createTestData() *Holder {
	testResults1 := list.New()
	assertResults1 := list.New()

	assertResults1.PushBack(NewFailureAssertResult("name1", "message1", "expected1", "actual1"))

	testResult1 := NewTestResult("prefix1", assertResults1)
	testResults1.PushBack(testResult1)

	assertResults2 := list.New()

	assertResults2.PushBack(NewFailureAssertResult("name2", "message2", "expected2", "actual2"))

	testResult2 := NewTestResult("prefix2", assertResults2)
	testResults1.PushBack(testResult2)

	holder := &Holder{
		TestResults: testResults1,
	}

	return holder
}