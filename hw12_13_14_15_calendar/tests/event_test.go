package scripts

import (
	"bytes"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v16"
	"io"
	"net/http"
	"strings"
)

type EventTest struct {
	r                  *http.Response
	responseStatusCode int
	responseBody       io.ReadCloser
}

func (test *EventTest) iSendRequestTo(httpMethod, addr string) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodDelete:
		req, err := http.NewRequest(http.MethodDelete, addr, nil)
		if err != nil {
			return err
		}
		client := &http.Client{}
		r, err = client.Do(req)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}
	if err != nil {
		return
	}

	test.responseStatusCode = r.StatusCode
	test.responseBody = r.Body
	test.r = r
	return
}

func (test *EventTest) iSendRequestToWithData(httpMethod, addr, contentType string, data *messages.PickleDocString) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodPost:
		replacer := strings.NewReplacer("\n", "", "\t", "")
		cleanJson := replacer.Replace(data.Content)
		r, err = http.Post(addr, contentType, bytes.NewReader([]byte(cleanJson)))
	case http.MethodPut:
		replacer := strings.NewReplacer("\n", "", "\t", "")
		cleanJson := replacer.Replace(data.Content)
		req, err := http.NewRequest(http.MethodPut, addr, bytes.NewReader([]byte(cleanJson)))
		if err != nil {
			return err
		}
		client := &http.Client{}
		r, err = client.Do(req)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody = r.Body
	test.r = r
	return
}

func (test *EventTest) theResponseCodeShouldBe(code int) (err error) {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func InitializeScenario(s *godog.ScenarioContext) {
	test := new(EventTest)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^I send "([^"]*)" request to "([^"]*)" with "([^"]*)" data:$`, test.iSendRequestToWithData)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
}
