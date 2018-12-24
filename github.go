package github

import (
	"regexp"
	"strings"
	"time"
)

type ListLink struct {
	Previous, Next, Last, First string
}

func ConvertTimeToTimestamp(timeString string) (int64, error) {
	dateTime, err := time.Parse(time.RFC3339, timeString)
	return dateTime.Unix(), err
}

func GetLink(headerLink string) ListLink {
	linkMapping := map[string]string{}
	headerLink = regexp.MustCompile("<|>| |\"|rel=").ReplaceAllString(headerLink, "")
	tempSplitComma := strings.Split(headerLink, ",")
	for _, temp := range tempSplitComma {
		// [0]: link url, [1]: link type {next, prev, first, last}
		tempSplitSemicolon := strings.Split(temp, ";")
		linkMapping[tempSplitSemicolon[1]] = tempSplitSemicolon[0]
	}
	return ListLink{
		Previous: linkMapping["prev"],
		Next:     linkMapping["next"],
		Last:     linkMapping["last"],
		First:    linkMapping["first"],
	}
}
