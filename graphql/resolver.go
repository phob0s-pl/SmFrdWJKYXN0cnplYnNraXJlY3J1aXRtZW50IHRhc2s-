package graphql

import (
	"context"
	"github.com/phob0s-pl/weatherllo/owm"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/phob0s-pl/weatherllo/model"
)

type Resolver struct {
	client *owm.OpenWeatherMap
}

func NewResolver(client *owm.OpenWeatherMap) *Resolver {
	return &Resolver{client: client}
}

func (r *Resolver) OpenWeatherMap() OpenWeatherMapResolver {
	return &openWeatherMapResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type openWeatherMapResolver struct{ *Resolver }

func (r *openWeatherMapResolver) Cities(ctx context.Context, obj *model.OpenWeatherMap, input []string) (list []*model.Weather, err error) {
	log.Info("graphql/cities: new call")

	for i := range input {
		city, err := r.client.City(&owm.CityInput{
			City:  input[i],
			Lang:  obj.Language,
			Units: obj.Units,
		})
		if err != nil {
			return nil, err
		}
		list = append(list, owmWeatherToGql(city))
	}

	return list, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
func (r *queryResolver) Openweathermap(ctx context.Context, input *model.OpenWeatherMapInput) (*model.OpenWeatherMap, error) {
	openWeatherMap := &model.OpenWeatherMap{}

	if input != nil {
		openWeatherMap.Language = &input.Language
		openWeatherMap.Units = &input.Units
	}

	return openWeatherMap, nil
}

func owmWeatherToGql(w *owm.ApiReply) *model.Weather {
	weather := &model.Weather{
		City: w.Name,
		General: &model.GeneralWeather{
			TemparatureMin:    w.Main.TempMin,
			TemparatureMax:    w.Main.TempMax,
			TemparatureFeel:   w.Main.Temp,
			Pressure:          w.Main.Pressure,
			Humidity:          w.Main.Humidity,
			PressureSeaLevel:  w.Main.SeaLevel,
			PressureGrndLevel: w.Main.GrndLevel,
		},
		Clouds: &model.Clouds{
			Cloudiness: w.Clouds.All,
		},
		Rain: &model.Rain{
			Rain1h: w.Rain.OneH,
			Rain3h: w.Rain.ThreeH,
		},
		Snow: &model.Snow{
			Snow1h: w.Snow.OneH,
			Snow3h: w.Snow.ThreeH,
		},
		Wind: &model.Wind{
			Speed:  w.Wind.Speed,
			Degree: w.Wind.Deg,
		},
		Sunrise: &model.Sunrise{
			Sunrise: time.Unix(w.Sys.Sunrise, 0).Format("15:04:05"),
			Sunset:  time.Unix(w.Sys.Sunset, 0).Format("15:04:05"),
		},
	}

	for i := range w.Weather {
		weather.Description = append(weather.Description, &model.WeatherDescription{
			Main:        w.Weather[i].Main,
			Description: w.Weather[i].Description,
		})
	}

	return weather
}
