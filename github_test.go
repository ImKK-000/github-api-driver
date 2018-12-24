package github_test

import (
	. "github-api-driver"
	"strings"
	"testing"
)

const headerLink = `<http://localhost:8882/user/15193133/repos?page=2>; rel="prev", <http://localhost:8882/user/15193133/repos?page=4>; rel="next", <http://localhost:8882/user/15193133/repos?page=5>; rel="last", <http://localhost:8882/user/15193133/repos?page=1>; rel="first"`

func Test_CallAPI_Input_URL_http_Should_Be_Get_Header_Response_Error_Message(t *testing.T) {
	expectedErrorMessage := "Header link not found"
	mockServer := mockServer(nil, nil)
	defer mockServer.Close()

	responseGithub := CallAPI(mockServer.Client(), mockServer.URL)

	if expectedErrorMessage != responseGithub.Error.Error() {
		t.Errorf("expect '%s' but it got '%s'", expectedErrorMessage, responseGithub.Error.Error())
	}
}

func Test_CallAPI_Input_URL_http_Should_Be_Get_Response_Error_Message(t *testing.T) {
	expectedErrorMessage := "Get http:: http: no Host in request URL"
	headerMapping := map[string]string{"Link": headerLink}
	mockServer := mockServer(headerMapping, nil)
	defer mockServer.Close()

	responseGithub := CallAPI(mockServer.Client(), "http://")

	if expectedErrorMessage != responseGithub.Error.Error() {
		t.Errorf("expect '%s' but it got '%s'", expectedErrorMessage, responseGithub.Error.Error())
	}
}

func Test_CallAPI_Input_URL_Should_Be_HeaderLink(t *testing.T) {
	expectedHeaderLinkPrevious := "http://localhost:8882/user/15193133/repos?page=2"
	expectedHeaderLinkNext := "http://localhost:8882/user/15193133/repos?page=4"
	expectedHeaderLinkLast := "http://localhost:8882/user/15193133/repos?page=5"
	expectedHeaderLinkFirst := "http://localhost:8882/user/15193133/repos?page=1"
	headerMapping := map[string]string{"Link": headerLink}
	mockServer := mockServer(headerMapping, []byte("{}"))
	defer mockServer.Close()

	responseGithub := CallAPI(mockServer.Client(), mockServer.URL)
	if !strings.Contains(headerMapping["Link"], responseGithub.HeaderLink.Previous) {
		t.Errorf("expect\n'%s'\nbut it got\n'%s'", expectedHeaderLinkPrevious, responseGithub.HeaderLink.Previous)
	}
	if !strings.Contains(headerMapping["Link"], responseGithub.HeaderLink.Next) {
		t.Errorf("expect\n'%s'\nbut it got\n'%s'", expectedHeaderLinkNext, responseGithub.HeaderLink.Next)
	}
	if !strings.Contains(headerMapping["Link"], responseGithub.HeaderLink.Last) {
		t.Errorf("expect\n'%s'\nbut it got\n'%s'", expectedHeaderLinkLast, responseGithub.HeaderLink.Last)
	}
	if !strings.Contains(headerMapping["Link"], responseGithub.HeaderLink.First) {
		t.Errorf("expect\n'%s'\nbut it got\n'%s'", expectedHeaderLinkFirst, responseGithub.HeaderLink.First)
	}
}

func Test_GetLink_Input_List_Link_Should_Be_Struct_List_Link(t *testing.T) {
	expectedListLink := ListLink{
		Previous: "http://localhost:8882/user/15193133/repos?page=2",
		Next:     "http://localhost:8882/user/15193133/repos?page=4",
		Last:     "http://localhost:8882/user/15193133/repos?page=5",
		First:    "http://localhost:8882/user/15193133/repos?page=1",
	}
	inputHeaderLink := headerLink

	actualListLink := GetLink(inputHeaderLink)

	if expectedListLink != actualListLink {
		t.Errorf("expect %v but it got %v", expectedListLink, actualListLink)
	}
}

func Test_ConvertTimeToTimestamp_Input_Time_RFC3339_Format_Should_Be_Error(t *testing.T) {
	inputTimeString := "1970-01-01 00:00:00"

	_, actualError := ConvertTimeToTimestamp(inputTimeString)

	if actualError == nil {
		t.Errorf("expect %v but it got %v", nil, actualError)
	}
}

func Test_ConvertTimeToTimestamp_Input_Time_RFC3339_Format_Should_Be_Unix_Time(t *testing.T) {
	expectedUnixTime := int64(0)
	inputTimeString := "1970-01-01T00:00:00Z"

	actualUnixTime, _ := ConvertTimeToTimestamp(inputTimeString)

	if expectedUnixTime != actualUnixTime {
		t.Errorf("expect '%d' but it got '%d'", expectedUnixTime, actualUnixTime)
	}
}
