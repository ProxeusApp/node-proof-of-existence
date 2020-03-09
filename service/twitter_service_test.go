package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpTwitterService_parseIDFromURL(t *testing.T) {
	twitterService := httpTwitterService{}

	type expectedResult struct {
		id    int
		error bool
	}

	type test struct {
		name     string
		url      string
		expected expectedResult
	}

	tests := []test{
		{name: "valid url", url: "https://twitter.com/meteoschweiz/status/1229676178715398146", expected: expectedResult{id: 1229676178715398146, error: false}},
		{name: "valid url", url: "https://twitter.com/MakerDAO/status/1235650873608425472", expected: expectedResult{id: 1235650873608425472, error: false}},
		{name: "broken link", url: "http://twitter.com/meteoschweiz/status/", expected: expectedResult{error: true}},
		{name: "empty string", url: "", expected: expectedResult{error: true}},
		{name: "id only", url: "1235650873608425472", expected: expectedResult{id: 1235650873608425472, error: false}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := twitterService.parseID(test.url)

			if test.expected.error {
				assert.NotNil(t, err, "should return error")
			} else {
				assert.Nil(t, err, "should not return error")
				assert.Equal(t, int64(test.expected.id), id, "should return the correct id")
			}
		})
	}

}
