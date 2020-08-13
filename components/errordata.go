package components

// ResponseErrorData all 4xx response bodies and maybe some 5xx should unmarshal to this
type ResponseErrorData struct {
	DocumentationUrl string `json:"documentation_url,omitempty"`
	Message          string `json:"message,omitempty"`
	Errors           []struct {
		Code     string `json:"code,omitempty"`
		Field    string `json:"field,omitempty"`
		Message  string `json:"message,omitempty"`
		Resource string `json:"resource,omitempty"`
	} `json:"errors,omitempty"`
}
