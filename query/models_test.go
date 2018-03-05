package query

import (
	"testing"
	"encoding/json"
)

func TestExpectUserResponse(t *testing.T) {
	var response Response
	err := json.Unmarshal([]byte(createSampleJsonWithoutExpectUserResponse()), &response)
	if err != nil {
		t.Error(err)
	}
	if !response.Result.Fulfillment.IsExpectUserResponse() {
		t.Fatal("ExpectUserResponse should be true, if there is no expect_user_response")
	}
	err = json.Unmarshal([]byte(createSampleJsonWithExpectUserResponseTrue()), &response)
	if err != nil {
		t.Error(err)
	}
	if !response.Result.Fulfillment.IsExpectUserResponse() {
		t.Fatal("ExpectUserResponse should be true, if the expect_user_response is true")
	}
	err = json.Unmarshal([]byte(createSampleJsonWithExpectUserResponseFalse()), &response)
	if err != nil {
		t.Error(err)
	}
	if response.Result.Fulfillment.IsExpectUserResponse() {
		t.Fatal("ExpectUserResponse should be false, if the expect_user_response is false")
	}
}

func createSampleJsonWithoutExpectUserResponse() string {
	return `
{
  "result": {
    "fulfillment": {
      "data": {
        "google": {
        }
      }
    }
  },
  "status": {
  }
}
`
}

func createSampleJsonWithExpectUserResponseTrue() string {
	return `
{
  "result": {
    "fulfillment": {
      "data": {
        "google": {
          "expect_user_response": true
        }
      }
    }
  },
  "status": {
  }
}
`
}

func createSampleJsonWithExpectUserResponseFalse() string {
	return `
{
  "result": {
    "fulfillment": {
      "data": {
        "google": {
          "expect_user_response": false
        }
      }
    }
  },
  "status": {
  }
}
`
}
