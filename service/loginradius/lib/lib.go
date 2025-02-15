package lib

import (
	"KandeHarsha/service/loginradius/schema"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method   string
	Path     string
	Query    map[string]string
	Header   map[string]string
	Payload  interface{}
	Response interface{}
	Status   int
}

func (r *Request) getPath() string {
	if len(r.Query) == 0 {
		return r.Path
	}
	qValues := url.Values{}
	for i, i2 := range r.Query {
		qValues.Set(i, i2)
	}
	return r.Path + "?" + qValues.Encode()
}

func (r *Request) getPayload() io.Reader {
	if r.Payload != nil {
		b, _ := json.Marshal(r.Payload)
		return bytes.NewBuffer(b)
	}
	return nil
}

func (r *Request) Do() *schema.ErrorResponse {
	req, err := http.NewRequest(r.Method, r.getPath(), r.getPayload())
	if err != nil {
		eErr := schema.GetSomethingWentWrongError()
		eErr.ErrorInfo = err.Error()
		return eErr
	}

	if r.Method != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}
	for s, s2 := range r.Header {
		req.Header.Set(s, s2)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		eErr := schema.GetSomethingWentWrongError()
		eErr.ErrorInfo = err.Error()
		return eErr
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		eErr := schema.GetSomethingWentWrongError()
		eErr.ErrorInfo = err.Error()
		return eErr
	}
	r.Status = res.StatusCode
	if res.StatusCode == 200 {
		err = json.Unmarshal(resBody, r.Response)
		if err != nil {
			eErr := schema.GetSomethingWentWrongError()

			eErr.ErrorInfo = err.Error()
			return eErr
		}
		return nil
	}
	var errResponse schema.ErrorResponse
	_ = json.Unmarshal(resBody, &errResponse)
	return &errResponse

}
