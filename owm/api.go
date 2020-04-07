package owm

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

type ApiReply struct {
	Coordinates struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
		SeaLevel  float64 `json:"sea_level"`
		GrndLevel float64 `json:"grnd_level"`
	} `json:"main"`
	Visibility float64 `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Rain struct {
		OneH   float64 `json:"1h"`
		ThreeH float64 `json:"3h"`
	} `json:"rain"`
	Snow struct {
		OneH   float64 `json:"1h"`
		ThreeH float64 `json:"3h"`
	} `json:"snow"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int64  `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type ApiReplyError struct {
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}

func decodeReply(r io.Reader) (*ApiReply, error) {
	ret := &ApiReply{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func decodeError(r io.Reader) error {
	ret := &ApiReplyError{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, ret); err != nil {
		return err
	}

	return errors.New(ret.Message)
}
