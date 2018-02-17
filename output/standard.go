package output

import (
	"github.com/yoichiro/dialogflow-query-checker/check"
	"time"
	"fmt"
)

func Standard(holder *check.Holder, start time.Time, end time.Time) {
	if holder.AllFailureAssertResultCount() == 0 {
		fmt.Printf("Finished in %f seconds.\n", (end.Sub(start)).Seconds())
		fmt.Println("All tests passed.")
	} else {
		for _, testResult := range holder.AllTestResults() {
			if testResult.AllFailureAssertResultCount() > 0 {
				fmt.Printf("[%s]\n", testResult.Prefix)
				for _, assertResult := range testResult.AllFailureAssertResults() {
					fmt.Printf("  Failure: %s\n", assertResult.Message)
					fmt.Printf("    Expected: %s\n", assertResult.Expected)
					fmt.Printf("    Actual: %s\n", assertResult.Actual)
				}
			}
		}
		fmt.Printf("Finished in %f seconds.\n", (end.Sub(start)).Seconds())
	}
}