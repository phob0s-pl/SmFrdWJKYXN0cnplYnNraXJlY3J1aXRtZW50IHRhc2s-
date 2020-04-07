package owm

import (
	"github.com/phob0s-pl/weatherllo/model"
	"net/url"
)

const (
	unitMetric   = "metric"
	unitImperial = "imperial"
)

const (
	urlValKeyToken = "appid"
	urlValKeyCity  = "q"
	urlValKeyUnits = "units"
	urlValKeyLang  = "lang"
)

func (o *OpenWeatherMap) buildUrlFromModel(input *CityInput) (string, error) {
	values, err := o.buildValues(input)
	if err != nil {
		return "", err
	}

	apiURL := url.URL{
		Scheme:     "https",
		Host:       apiHost,
		Path:       apiPath,
		ForceQuery: true,
		RawQuery:   values.Encode(),
	}
	return apiURL.String(), nil
}

func (o *OpenWeatherMap) buildValues(input *CityInput) (url.Values, error) {
	if o.token == "" {
		return nil, ErrMissingToken
	}

	if input.City == "" {
		return nil, ErrNoCity
	}

	values := url.Values{}
	values.Add(urlValKeyCity, input.City)
	values.Add(urlValKeyToken, o.token)

	if input.Lang != nil {
		lang, err := modelLangToOwn(*input.Lang)
		if err != nil {
			return nil, err
		}
		values.Add(urlValKeyLang, lang)
	}

	if input.Units != nil {
		unit, err := modelUnitToOwm(*input.Units)
		if err != nil {
			return nil, err
		}
		values.Add(urlValKeyUnits, unit)
	}

	return values, nil
}

func modelUnitToOwm(unit model.Units) (string, error) {
	switch unit {
	case model.UnitsMetric:
		return unitMetric, nil
	case model.UnitsImperial:
		return unitImperial, nil
	default:
		return "", ErrUnsupportedUnit(unit)
	}
}

func modelLangToOwn(lang model.Lang) (string, error) {
	switch lang {
	case model.LangAfrikaans:
		return "af", nil
	case model.LangArabic:
		return "ar", nil
	case model.LangAzerbaijani:
		return "az", nil
	case model.LangBulgarian:
		return "bg", nil
	case model.LangCatalan:
		return "ca", nil
	case model.LangCzech:
		return "cz", nil
	case model.LangDanish:
		return "da", nil
	case model.LangGerman:
		return "de", nil
	case model.LangGreek:
		return "el", nil
	case model.LangEnglish:
		return "en", nil
	case model.LangBasque:
		return "eu", nil
	case model.LangPersian:
		return "fa", nil
	case model.LangFinnish:
		return "fi", nil
	case model.LangFrench:
		return "fr", nil
	case model.LangGalician:
		return "gl", nil
	case model.LangHebrew:
		return "he", nil
	case model.LangHindi:
		return "hi", nil
	case model.LangCroatian:
		return "hr", nil
	case model.LangHungarian:
		return "hu", nil
	case model.LangIndonesian:
		return "id", nil
	case model.LangItalian:
		return "it", nil
	case model.LangJapanese:
		return "ja", nil
	case model.LangKorean:
		return "kr", nil
	case model.LangLatvian:
		return "la", nil
	case model.LangLithuanian:
		return "lt", nil
	case model.LangMacedonian:
		return "mk", nil
	case model.LangNorwegian:
		return "no", nil
	case model.LangDutch:
		return "nl", nil
	case model.LangPolish:
		return "pl", nil
	case model.LangPortuguese:
		return "pt", nil
	case model.LangBrasilPortugues:
		return "pt_br", nil
	case model.LangRomanian:
		return "ro", nil
	case model.LangRussian:
		return "ru", nil
	case model.LangSwedish:
		return "se", nil
	case model.LangSlovak:
		return "sk", nil
	case model.LangSlovenian:
		return "sl", nil
	case model.LangSpanish:
		return "es", nil
	case model.LangSerbian:
		return "sr", nil
	case model.LangThai:
		return "th", nil
	case model.LangTurkish:
		return "tr", nil
	case model.LangUk:
		return "uk", nil
	case model.LangVietnamese:
		return "vi", nil
	case model.LangChineseSimplified:
		return "zh_cn", nil
	case model.LangChineseTraditional:
		return "zh_tw", nil
	case model.LangZulu:
		return "zu", nil
	default:
		return "", ErrUnsupportedLanguage(lang)
	}
}
