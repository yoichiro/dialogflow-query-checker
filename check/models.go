package check

import (
	"container/list"
)

type Holder struct {
	TestResults *list.List
}

func (holder *Holder) AllTestResults() []*TestResult {
	r := make([]*TestResult, 0, holder.TestResults.Len())
	for e := holder.TestResults.Front(); e != nil; e = e.Next() {
		if x, ok := e.Value.(*TestResult); ok {
			r = append(r, x)
		}
	}
	return r
}

func (holder *Holder) AllAssertResults() []*AssertResult {
	r := make([]*AssertResult, 0)
	for _, testResult := range holder.AllTestResults() {
		for _, assertResult := range testResult.AllAssertResults() {
			r = append(r, assertResult)
		}
	}
	return r
}

func (holder *Holder) AllAssertResultCount() int {
	return len(holder.AllAssertResults())
}

type TestResult struct {
	Prefix string
	AssertResults *list.List
}

func NewTestResult(prefix string, assertResults *list.List) *TestResult {
	r := &TestResult{
		Prefix: prefix,
		AssertResults: assertResults,
	}
	return r
}

func (test *TestResult) AllAssertResults() []*AssertResult {
	r := make([]*AssertResult, 0, test.AssertResults.Len())
	for e := test.AssertResults.Front(); e != nil; e = e.Next() {
		if x, ok := e.Value.(*AssertResult); ok {
			r = append(r, x)
		}
	}
	return r
}

type AssertResult struct {
	Message string
	Expected string
	Actual string
}

func NewAssertResult(message string, expected string, actual string) *AssertResult {
	return &AssertResult{
		Message: message,
		Expected: expected,
		Actual: actual,
	}
}