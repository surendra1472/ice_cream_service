package local_config

import (
	"context"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetTranslationMessage(ctx context.Context, messageId string) string {
	localizer := GetLocalizer(getLocaleFromContext(ctx))
	message := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageId})
	return message
}

func getLocaleFromContext(ctx context.Context) string {
	locale := ctx.Value("locale")
	if locale == "" || locale == nil {
		return "en"
	}
	return locale.(string)
}
