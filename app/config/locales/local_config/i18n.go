package local_config

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

var permittedLanguages = [...]string{"en"}

var localizers = make(map[string]*i18n.Localizer)

func localeFile(language string) string {
	switch language {
	case "en":
		return "en.json"
	default:
		return "en.json"
	}
}

func contains(a [1]string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func InitializeLocalizer() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, value := range permittedLanguages {
		_, err := bundle.LoadMessageFile("app/config/locales/" + localeFile(value))
		if err != nil {
			log.Print("Localizer: Error reading file ", localeFile(value))
		}
		localizers[value] = i18n.NewLocalizer(bundle, value)
	}
}

func GetLocalizer(lang string) *i18n.Localizer {
	if contains(permittedLanguages, lang) {
		return localizers[lang]
	}
	return localizers["en"]
}

