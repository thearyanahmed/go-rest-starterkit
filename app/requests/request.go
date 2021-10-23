package requests

import "net/url"

type RequestError struct {
	ErrorBag url.Values
}

func (r RequestError) HasErrors() bool {
	return len(r.ErrorBag) > 0
}
