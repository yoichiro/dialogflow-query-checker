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
		assertResults := list.New()

		actual, err := query.Execute(&test, def.ClientAccessToken, def.DefaultLanguage)
		if err != nil {
			return nil, err
		}

		displayResult(assertIntEquals(assertResults, "status.code", 200, actual.Status.Code))
		displayResult(assertStringEquals(assertResults, "action", test.Expect.Action, actual.Result.Action))
		displayResult(assertStringEquals(assertResults, "intentName", test.Expect.IntentName, actual.Result.Metadata.IntentName))
		actualContexts := make([]string, len(actual.Result.Contexts))
		for i, context := range actual.Result.Contexts {
			actualContexts[i] = context.Name
		}
		if test.Expect.Contexts != nil {
			displayResult(assertArrayEquals(assertResults, "contexts", test.Expect.Contexts, actualContexts))
		}
		displayResult(assertStringEquals(assertResults, "date", evaluateDateMacro(test.Expect.Parameters.Date, "2006-01-02"), actual.Result.Parameters.Date))
		displayResult(assertStringEquals(assertResults, "prefecture", test.Expect.Parameters.Prefecture, actual.Result.Parameters.Prefecture))
		displayResult(assertStringEquals(assertResults, "keyword", test.Expect.Parameters.Keyword, actual.Result.Parameters.Keyword))
		displayResult(assertStringEquals(assertResults, "event", test.Expect.Parameters.Event, actual.Result.Parameters.Event))
		if test.Expect.Speeches != nil {
			displayResult(assertByMultipleRegexps(assertResults,  "speech", test.Expect.Speeches, actual.Result.Fulfillment.Speech))
		} else {
			re := regexp.MustCompile(evaluateDateMacro(test.Expect.Speech, "1月2日"))
			displayResult(assertByRegexp(assertResults, "speech", re, actual.Result.Fulfillment.Speech))
		}

		if assertResults.Len() > 0 {
			results.PushBack(NewTestResult(test.CreatePrefix(), assertResults))
		}
	}
	return &Holder{
		TestResults: results,
	}, nil
}

func displayResult(result bool) {
	if result {
		fmt.Print(".")
	} else {
		fmt.Print("F")
	}
}

func assertIntEquals(results *list.List, name string, expected int, actual int) bool {
	if expected != actual {
		results.PushBack(
			NewAssertResult(fmt.Sprintf("%s is not same as expected value.", name), strconv.Itoa(expected), strconv.Itoa(actual)))
		return false
	}
	return true
}

func assertStringEquals(results *list.List, name string, expected string, actual string) bool {
	if expected != actual {
		results.PushBack(NewAssertResult(fmt.Sprintf("%s is not same as expected value.", name), expected, actual))
		return false
	}
	return true
}

func assertArrayEquals(results *list.List, name string, expected []string, actual []string) bool {
	if len(expected) != len(actual) {
		results.PushBack(
			NewAssertResult(fmt.Sprintf("The length of %s is not same as expected length.", name), strconv.Itoa(len(expected)), strconv.Itoa(len(actual))))
		return false
	}
	for _, e := range expected {
		if !contains(actual, e) {
			results.PushBack(
				NewAssertResult(fmt.Sprintf("%s does not contain %s", name, e), "Contained", "Not contained"))
			return false
		}
	}
	return true
}

func assertByRegexp(results *list.List, name string, expected *regexp.Regexp, actual string) bool {
	if !expected.Match([]byte(actual)) {
		results.PushBack(NewAssertResult(fmt.Sprintf("%s is not matched to expected regular expression.", name), expected.String(), actual))
		return false
	}
	return true
}

func assertByMultipleRegexps(results *list.List, name string, regexps []string, actual string) bool {
	for _, exp := range regexps {
		re := regexp.MustCompile(evaluateDateMacro(exp, "1月2日"))
		if re.Match([]byte(actual)) {
			return true
		}
	}
	f := func(x string) string {
		return fmt.Sprintf("\"%s\"", x)
	}
	results.PushBack(NewAssertResult(fmt.Sprintf("%s is not matched to expected regular expression.", name), strings.Join(mapString(regexps, f), ", "), actual))
	return false
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

func evaluateDateMacro(s string, layout string) string {
	t := time.Now()
	today := t.Format(layout)
	t = t.AddDate(0, 0, 1)
	tomorrow := t.Format(layout)
	result := strings.Replace(s, "${date.tomorrow}", tomorrow, -1)
	result = strings.Replace(result, "${date.today}", today, -1)
	return result
}
