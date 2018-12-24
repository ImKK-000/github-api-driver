package github_test

import (
	. "github-api-driver"
	"testing"
)

const headerLink = `<http://localhost:8882/user/15193133/repos?page=2>; rel="prev", <http://localhost:8882/user/15193133/repos?page=4>; rel="next", <http://localhost:8882/user/15193133/repos?page=5>; rel="last", <http://localhost:8882/user/15193133/repos?page=1>; rel="first"`

func Test_CallAPI_Input_URL_Should_Be_Response_And_Header_Link(t *testing.T) {
	expectedResponseBody := []byte("{}")
	expectedHeader := map[string]string{"Link": headerLink}
	mockServer := mockServer(expectedHeader, expectedResponseBody)
	defer mockServer.Close()

	actualResponseBody, actualHeaderLink := CallAPI(mockServer.Client(), mockServer.URL)

	if string(expectedResponseBody) != string(actualResponseBody) {
		t.Errorf("expect\n'%s'\nbut it got\n'%s'", string(expectedResponseBody), string(actualResponseBody))
	}
	if expectedHeader["Link"] != actualHeaderLink {
		t.Errorf("expect\n'%s'\nbut it got\n'%s'", expectedHeader["Link"], actualHeaderLink)
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
