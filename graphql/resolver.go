package graphql

import (
	"context"

	"github.com/phob0s-pl/weatherllo/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) OpenWeatherMap() OpenWeatherMapResolver {
	return &openWeatherMapResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type openWeatherMapResolver struct{ *Resolver }

func (r *openWeatherMapResolver) Cities(ctx context.Context, obj *model.OpenWeatherMap, input []string) ([]*model.Weather, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	panic("not implemented")
}
func (r *queryResolver) Openweathermap(ctx context.Context, input *model.OpenWeatherMapInput) (*model.OpenWeatherMap, error) {
	panic("not implemented")
}
