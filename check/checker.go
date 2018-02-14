package check

import (
	"github.com/yoichiro/dialogflow-query-checker/config"
	"github.com/yoichiro/dialogflow-query-checker/query"
	"container/list"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Execute(def *config.Definition) (*list.List, error) {
	results := list.New()
	for _, test := range def.Tests {
		actual, err := query.Execute(&test, def.ClientAccessToken)
		if err != nil {
			return nil, err
		}
		displayResult(assertIntEquals(results, &test, "status.code", 200, actual.Status.Code))
		displayResult(assertStringEquals(results, &test,"action", test.Expect.Action, actual.Result.Action))
		displayResult(assertStringEquals(results, &test,"intentName", test.Expect.IntentName, actual.Result.Metadata.IntentName))
		actualContexts := make([]string, len(actual.Result.Contexts))
		for i, context := range actual.Result.Contexts {
			actualContexts[i] = context.Name
		}
		if test.Expect.Contexts != nil {
			displayResult(assertArrayEquals(results, &test,"contexts", test.Expect.Contexts, actualContexts))
		}
		displayResult(assertStringEquals(results, &test,"date", evaluateDateMacro(test.Expect.Parameters.Date, "2006-01-02"), actual.Result.Parameters.Date))
		displayResult(assertStringEquals(results, &test,"prefecture", test.Expect.Parameters.Prefecture, actual.Result.Parameters.Prefecture))
		displayResult(assertStringEquals(results, &test,"keyword", test.Expect.Parameters.Keyword, actual.Result.Parameters.Keyword))
		displayResult(assertStringEquals(results, &test,"event", test.Expect.Parameters.Event, actual.Result.Parameters.Event))
		re := regexp.MustCompile(evaluateDateMacro(test.Expect.Speech, "1月2日"))
		displayResult(assertByRegexp(results, &test,"speech", re, actual.Result.Fulfillment.Speech))
	}
	return results, nil
}

func displayResult(result bool) {
	if result {
		fmt.Print("\x1b[32m.\x1b[0m")
	} else {
		fmt.Print("\x1b[31mF\x1b[0m")
	}
}

func assertIntEquals(results *list.List, test *config.Test, name string, expected int, actual int) bool {
	if expected != actual {
		results.PushBack(
			fmt.Sprintf("%s %s is not same. expected:%d actual:%d",
				createPrefix(test), name, expected, actual))
		return false
	}
	return true
}

func assertStringEquals(results *list.List, test *config.Test, name string, expected string, actual string) bool {
	if expected != actual {
		results.PushBack(
			fmt.Sprintf("%s %s is not same. expected:%s actual:%s",
				createPrefix(test), name, expected, actual))
		return false
	}
	return true
}

func assertArrayEquals(results *list.List, test *config.Test, name string, expected []string, actual []string) bool {
	if len(expected) != len(actual) {
		results.PushBack(
			fmt.Sprintf("%s The length of %s is not same. expected:%d actual:%d",
				createPrefix(test), name, len(expected), len(actual)))
		return false
	}
	for _, e := range expected {
		if !contains(actual, e) {
			results.PushBack(fmt.Sprintf("%s does not contain %s", name, e))
			return false
		}
	}
	return true
}

func assertByRegexp(results *list.List, test *config.Test, name string, expected *regexp.Regexp, actual string) bool {
	if !expected.Match([]byte(actual)) {
		results.PushBack(
			fmt.Sprintf("%s %s does not match. expected:%s actual:%s",
				createPrefix(test), name, expected, actual))
		return false
	}
	return true
}

func createPrefix(test *config.Test) string {
	if test.Condition.Contexts != nil {
		return "[" + strings.Join(test.Condition.Contexts, ",") + " " + test.Condition.Query + "]"
	} else {
		return "[" + test.Condition.Query + "]"
	}
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
