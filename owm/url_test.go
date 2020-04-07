package owm

import (
	"github.com/phob0s-pl/weatherllo/model"
	"testing"
)

func TestBuildUrl(t *testing.T) {
	en := model.LangEnglish
	metric := model.UnitsMetric

	testCases := []struct {
		expected string
		token    string
		input    *CityInput
	}{
		{
			token:    "token",
			expected: "https://api.openweathermap.org/data/2.5/weather?appid=token&lang=en&q=Katowice&units=metric",
			input: &CityInput{
				City:  "Katowice",
				Lang:  &en,
				Units: &metric,
			},
		},

		{
			token:    "token",
			expected: "https://api.openweathermap.org/data/2.5/weather?appid=token&lang=en&q=Katowice",
			input: &CityInput{
				City: "Katowice",
				Lang: &en,
			},
		},

		{
			token:    "token",
			expected: "https://api.openweathermap.org/data/2.5/weather?appid=token&q=Katowice&units=metric",
			input: &CityInput{
				City:  "Katowice",
				Units: &metric,
			},
		},

		{
			token:    "token",
			expected: "https://api.openweathermap.org/data/2.5/weather?appid=token&q=Katowice",
			input: &CityInput{
				City: "Katowice",
			},
		},
	}

	for i, testCase := range testCases {
		url, err := New(testCase.token).buildUrlFromModel(testCase.input)
		if err != nil {
			t.Errorf("[%d] buildUrl(): failed: %s", i, err)
		}

		if url != testCase.expected {
			t.Errorf("[%d] buildUrl(): returned %s, expected: %s", i, url, testCase.expected)
		}
	}
}
