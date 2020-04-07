package owm

import (
	"github.com/patrickmn/go-cache"
	"github.com/phob0s-pl/weatherllo/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	apiHost = "api.openweathermap.org"
	apiPath = "/data/2.5/weather"
)

type OpenWeatherMap struct {
	// cityCache contains cached result of cities
	// The cache data will expire in 10 minutes
	// ref: https://openweathermap.org/appid#work
	cityCache *cache.Cache

	// token is openweathermap.org access token
	token string

	client *http.Client
}

type CityInput struct {
	City  string
	Lang  *model.Lang
	Units *model.Units
}

func New(token string) *OpenWeatherMap {
	return &OpenWeatherMap{
		cityCache: cache.New(10*time.Minute, 10*time.Minute),
		token:     token,
		client: &http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (o *OpenWeatherMap) City(input *CityInput) (*ApiReply, error) {
	cachedWeather, ok := o.cachedCityWeather(input.City)
	if ok {
		log.WithField("city", input.City).Debug("returning cached result")
		return cachedWeather, nil
	}

	apiCityWeather, err := o.apiCallCity(input)
	if err != nil {
		return nil, err
	}

	o.cityCache.SetDefault(input.City, apiCityWeather)

	log.WithField("city", input.City).Debug("returning api result")
	return apiCityWeather, nil
}

func (o *OpenWeatherMap) cachedCityWeather(city string) (*ApiReply, bool) {
	result, exist := o.cityCache.Get(city)
	if exist {
		val, ok := result.(*ApiReply)
		if ok {
			return val, true
		}
	}

	return nil, false
}

func (o *OpenWeatherMap) apiCallCity(input *CityInput) (*ApiReply, error) {
	apiURL, err := o.buildUrlFromModel(input)
	if err != nil {
		return nil, err
	}

	response, err := o.client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, decodeError(response.Body)
	}

	return decodeReply(response.Body)
}
