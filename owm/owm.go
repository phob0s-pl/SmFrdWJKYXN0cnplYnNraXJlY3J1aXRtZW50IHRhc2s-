package owm

import (
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	apiUrl = ""
)

type OpenWeatherMap struct {
	// cityCache contains cached result of cities
	// The cache data will expire in 10 minutes
	// ref: https://openweathermap.org/appid#work
	cityCache *cache.Cache

	// token is openweathermap.org access token
	token string
}

func New(token string) *OpenWeatherMap {
	return &OpenWeatherMap{
		cityCache: cache.New(10*time.Minute, 10*time.Minute),
		token:     token,
	}
}

func (o *OpenWeatherMap) City(name string) ([]byte, error) {
	result, exist := o.cityCache.Get(name)
	if exist {
		val, ok := result.([]byte)
		if ok {
			return val, nil
		}
	}

	return nil, nil
}

func (o *OpenWeatherMap) apiCall(url string) ([]byte, error) {
	return nil, nil
}

func (o *OpenWeatherMap) lang() {

}
