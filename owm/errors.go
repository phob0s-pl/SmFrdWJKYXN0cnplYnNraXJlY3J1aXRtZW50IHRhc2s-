package owm

import (
	"errors"
	"fmt"
	"github.com/phob0s-pl/weatherllo/model"
)

func ErrUnsupportedUnit(unit model.Units) error {
	return fmt.Errorf("unsupported unit: %s", unit)
}

func ErrUnsupportedLanguage(lang model.Lang) error {
	return fmt.Errorf("unsupported language: %s", lang)
}

var ErrMissingToken = errors.New("missing token")
var ErrNoCity = errors.New("no city")
