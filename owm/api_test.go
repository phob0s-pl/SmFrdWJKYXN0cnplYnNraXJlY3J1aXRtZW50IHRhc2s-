package owm

import (
	"bytes"
	"testing"
)

const city1 = `
{"coord":{"lon":19.03,"lat":50.26},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"base":"stations","main":{"temp":285.67,"feels_like":282.12,"temp_min":284.26,"temp_max":287.04,"pressure":1030,"humidity":24},"visibility":10000,"wind":{"speed":1},"clouds":{"all":0},"dt":1586284523,"sys":{"type":1,"id":1705,"country":"PL","sunrise":1586232392,"sunset":1586280294},"timezone":7200,"id":3096472,"name":"Katowice","cod":200}
`

const city2 = `
{"coord":{"lon":34.03,"lat":38.37},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10n"}],"base":"stations","main":{"temp":282.68,"feels_like":274.12,"temp_min":282.68,"temp_max":282.68,"pressure":1015,"humidity":41,"sea_level":1015,"grnd_level":903},"wind":{"speed":8.81,"deg":42},"rain":{"3h":0.31},"clouds":{"all":99},"dt":1586285439,"sys":{"country":"TR","sunrise":1586229553,"sunset":1586275933},"timezone":10800,"id":324496,"name":"Aksaray","cod":200}
`

const city3 = `
{"coord":{"lon":-147.72,"lat":64.84},"weather":[{"id":600,"main":"Snow","description":"light snow","icon":"13d"},{"id":701,"main":"Mist","description":"mist","icon":"50d"}],"base":"stations","main":{"temp":262.65,"feels_like":258.37,"temp_min":261.48,"temp_max":263.71,"pressure":1020,"humidity":85},"visibility":6437,"wind":{"speed":1.5,"deg":50},"clouds":{"all":90},"dt":1586285705,"sys":{"type":1,"id":7684,"country":"US","sunrise":1586270541,"sunset":1586322168},"timezone":-28800,"id":5861897,"name":"Fairbanks","cod":200}
`

func TestDecodeReply(t *testing.T) {
	cases := []string{city1, city2, city3}

	for i, city := range cases {
		if _, err := decodeReply(bytes.NewBufferString(city)); err != nil {
			t.Errorf("[%d] decodeReply() failed: %s", i, err)
		}
	}
}
