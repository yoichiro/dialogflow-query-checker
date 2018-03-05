package query

// Request models

type RequestBody struct {
	Contexts []string `json:"contexts,omitempty"`
	Language string `json:"lang"`
	OriginalRequest OriginalRequest `json:"originalRequest"`
	Query string `json:"query,omitempty"`
	Event Event `json:"event,omitempty"`
	SessionId string `json:"sessionId"`
}

type Event struct {
	Name string `json:"name,omitempty"`
}

type OriginalRequest struct {
	Source string `json:"source"`
	Data Data `json:"data"`
}

type Data struct {
	User User `json:"user"`
}

type User struct {
	Locale string `json:"locale"`
	AccessToken string `json:"accessToken,omitempty"`
}

// Response models

type Response struct {
	Result Result `json:"result"`
	Status Status `json:"status"`
}

type Result struct {
	Action string `json:"action"`
	Parameters map[string]string `json:"parameters"`
	Metadata Metadata `json:"metadata"`
	Contexts []Context `json:"contexts"`
	Fulfillment Fulfillment `json:"fulfillment"`
}

type Status struct {
	Code int `json:"code"`
	ErrorDetails string `json:"errorDetails"`
	ErrorType string `json:"errorType"`
}

type Context struct {
	Name string `json:"name"`
}

type Metadata struct {
	IntentName string `json:"intentName"`
}

type Fulfillment struct {
	Speech string `json:"speech"`
	Data struct {
		Google struct {
			ExpectUserResponse *bool `json:"expect_user_response,omitempty"`
		} `json:"google"`
	} `json:"data"`
}
