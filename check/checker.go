package check

import (
	"github.com/yoichiro/dialogflow-query-checker/config"
	"github.com/yoichiro/dialogflow-query-checker/query"
	"container/list"
	"fmt"
	"regexp"
	"strings"
	"time"
	"strconv"
)

func Execute(def *config.Definition) (*Holder, error) {
	results := list.New()
	for _, test := range def.Tests {
		if def.Environment.Debug {
			fmt.Printf("[Start] %s\n", test.CreatePrefix())
		}

		start := time.Now()

		assertResults := list.New()

		actual, err := query.Execute(&test, def)
		if err != nil {
			return nil, err
		}

		displayResult(assertResults, assertStatus(&actual.Status))
		displayResult(assertResults, assertStringEquals("action", test.Expect.Action, actual.Result.Action))
		displayResult(assertResults, assertStringEquals("intentName", test.Expect.IntentName, actual.Result.Metadata.IntentName))
		actualContexts := make([]string, len(actual.Result.Contexts))
		for i, context := range actual.Result.Contexts {
			actualContexts[i] = context.Name
		}
		if test.Expect.Contexts != nil {
			displayResult(assertResults, assertArrayContains("contexts", test.Expect.Contexts, actualContexts))
		}
		for name, expected := range test.Expect.Parameters {
			displayResult(assertResults, assertStringEquals(name, expected, actual.Result.Parameters[name]))
		}
		if test.Expect.Speeches != nil {
			displayResult(assertResults, assertByMultipleRegexps( "speech", test.Expect.Speeches, actual.Result.Fulfillment.Speech))
		} else {
			re := regexp.MustCompile(test.Expect.Speech)
			displayResult(assertResults, assertByRegexp("speech", re, actual.Result.Fulfillment.Speech))
		}
		expectedEndConversation := test.Expect.EndConversation == "true"
		displayResult(assertResults, assertBoolEquals("endConversation", expectedEndConversation, !actual.Result.Fulfillment.Data.Google.ExpectUserResponse))

		end := time.Now()
		results.PushBack(NewTestResult(test.CreatePrefix(), (end.Sub(start)).Seconds(), assertResults))

		if def.Environment.Debug {
			fmt.Printf("\n[End] %s\n", test.CreatePrefix())
		}
	}
	return &Holder{
		TestResults: results,
	}, nil
}

func displayResult(results *list.List, assertResult *AssertResult) {
	if assertResult.Success {
		fmt.Print(".")
	} else {
		fmt.Print("F")
	}
	results.PushBack(assertResult)
}

func assertStatus(status *query.Status) *AssertResult {
	if status.Code != 200 {
		return NewFailureAssertResult("status", fmt.Sprintf("status is %d, not 200. (%s: %s)", status.Code, status.ErrorType, status.ErrorDetails), strconv.Itoa(200), strconv.Itoa(status.Code))
	} else {
		return NewSuccessAssertResult("status")
	}
}

func assertBoolEquals(name string, expected bool, actual bool) *AssertResult {
	if expected != actual {
		return NewFailureAssertResult(name, fmt.Sprintf("%s is not same as expected value.", name), strconv.FormatBool(expected), strconv.FormatBool(actual))
	} else {
		return NewSuccessAssertResult(name)
	}
}

func assertStringEquals(name string, expected string, actual string) *AssertResult {
	if expected != actual {
		return NewFailureAssertResult(name, fmt.Sprintf("%s is not same as expected value.", name), expected, actual)
	} else {
		return NewSuccessAssertResult(name)
	}
}

func assertArrayContains(name string, expected []string, actual []string) *AssertResult {
	for _, e := range expected {
		if !contains(actual, e) {
			return NewFailureAssertResult(name, fmt.Sprintf("%s does not contain %s", name, e), "Contained", "Not contained")
		}
	}
	return NewSuccessAssertResult(name)
}

func assertByRegexp(name string, expected *regexp.Regexp, actual string) *AssertResult {
	if !expected.Match([]byte(actual)) {
		return NewFailureAssertResult(name, fmt.Sprintf("%s is not matched to expected regular expression.", name), expected.String(), actual)
	} else {
		return NewSuccessAssertResult(name)
	}
}

func assertByMultipleRegexps(name string, regexps []string, actual string) *AssertResult {
	for _, exp := range regexps {
		re := regexp.MustCompile(exp)
		if re.Match([]byte(actual)) {
			return NewSuccessAssertResult(name)
		}
	}
	f := func(x string) string {
		return fmt.Sprintf("\"%s\"", x)
	}
	return NewFailureAssertResult(name, fmt.Sprintf("%s is not matched to expected regular expression.", name), strings.Join(mapString(regexps, f), ", "), actual)
}

func mapString(x []string, f func(string) string) []string {
	r := make([]string, len(x))
	for i, e := range x {
		r[i] = f(e)
	}
	return r
}

func contains(array []string, s string) bool {
	for _, e := range array {
		if s == e {
			return true
		}
	}
	return false
}
