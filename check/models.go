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

func (holder *Holder) AllSuccessAssertResults() []*AssertResult {
	r := make([]*AssertResult, 0)
	for _, testResult := range holder.AllTestResults() {
		for _, assertResult := range testResult.AllAssertResults() {
			if assertResult.Success {
				r = append(r, assertResult)
			}
		}
	}
	return r
}

func (holder *Holder) AllSuccessAssertResultCount() int {
	return len(holder.AllSuccessAssertResults())
}

func (holder *Holder) AllFailureAssertResults() []*AssertResult {
	r := make([]*AssertResult, 0)
	for _, testResult := range holder.AllTestResults() {
		for _, assertResult := range testResult.AllFailureAssertResults() {
			r = append(r, assertResult)
		}
	}
	return r
}

func (holder *Holder) AllFailureAssertResultCount() int {
	return len(holder.AllFailureAssertResults())
}

func (holder *Holder) AllFailureTestResultCount() int {
	r := 0
	for _, testResult := range holder.AllTestResults() {
		if testResult.AllFailureAssertResultCount() > 0 {
			r += 1
		}
	}
	return r
}

type TestResult struct {
	Prefix string
	Time float64
	AssertResults *list.List
}

func NewTestResult(prefix string, time float64, assertResults *list.List) *TestResult {
	r := &TestResult{
		Prefix: prefix,
		Time: time,
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

func (test *TestResult) AllAssertResultCount() int {
	return len(test.AllAssertResults())
}

func (test *TestResult) AllFailureAssertResults() []*AssertResult {
	r := make([]*AssertResult, 0, test.AssertResults.Len())
	for e := test.AssertResults.Front(); e != nil; e = e.Next() {
		if x, ok := e.Value.(*AssertResult); ok {
			if !x.Success {
				r = append(r, x)
			}
		}
	}
	return r
}

func (test *TestResult) AllFailureAssertResultCount() int {
	return len(test.AllFailureAssertResults())
}

type AssertResult struct {
	Name string
	Success bool
	Message string
	Expected string
	Actual string
}

func NewSuccessAssertResult(name string) *AssertResult {
	return &AssertResult{
		Name: name,
		Success: true,
	}
}

func NewFailureAssertResult(name string, message string, expected string, actual string) *AssertResult {
	return &AssertResult{
		Name: name,
		Success: false,
		Message: message,
		Expected: expected,
		Actual: actual,
	}
}
