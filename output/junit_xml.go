package output

import (
	"github.com/yoichiro/dialogflow-query-checker/check"
	"time"
	"encoding/xml"
	"strconv"
	"fmt"
	"os"
)

func JunitXml(holder *check.Holder, path string, start time.Time, end time.Time) error {
	testsuite := TestSuite{
		Name: "dialogflow-query-checker",
		Tests: strconv.Itoa(holder.TestResults.Len()),
		Errors: "0",
		Failures: strconv.Itoa(holder.AllFailureAssertResultCount()),
		Time: fmt.Sprintf("%f", (end.Sub(start)).Seconds()),
	}
	for _, testResult := range holder.AllTestResults() {
		testcase := TestCase{
			Name: testResult.Prefix,
			Assertions: strconv.Itoa(testResult.AllAssertResultCount()),
			Time: fmt.Sprintf("%f", testResult.Time),
			Score: fmt.Sprintf("%f", testResult.Score),
		}
		for _, assertResult := range testResult.AllFailureAssertResults() {
			testcase.AddFailure(Failure{
				Message: assertResult.Message,
				Text: fmt.Sprintf("Expected: %s\nActual: %s", assertResult.Expected, assertResult.Actual),
			})
		}
		testsuite.AddTestCase(testcase)
	}
	writer, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = writer.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	if err != nil {
		return err
	}
	encoder := xml.NewEncoder(writer)
	encoder.Indent("", "  ")
	err = encoder.Encode(testsuite)
	if err != nil {
		return err
	}
	return nil
}

type TestSuite struct {
	XMLName xml.Name `xml:"testsuite"`
	Name string `xml:"name,attr"`
	Tests string `xml:"tests,attr"`
	Errors string `xml:"errors,attr"`
	Failures string `xml:"failures,attr"`
	Time string `xml:"time,attr"`
	TestCases []TestCase `xml:",omitempty"`
}

func (testsuite *TestSuite) AddTestCase(testcase TestCase) {
	r := append(testsuite.TestCases, testcase)
	testsuite.TestCases = r
}

type TestCase struct {
	XMLName xml.Name `xml:"testcase"`
	Name string `xml:"name,attr"`
	Assertions string `xml:"assertions,attr"`
	Time string `xml:"time,attr"`
	Score string `xml:"score,attr"`
	Failures []Failure `xml:",omitempty"`
}

func (testcase *TestCase) AddFailure(failure Failure) {
	r := append(testcase.Failures, failure)
	testcase.Failures = r
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Message string `xml:"message,attr"`
	Text string `xml:",cdata"`
}
