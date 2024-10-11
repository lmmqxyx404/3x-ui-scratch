package locale

import (
	"embed"
	"io/fs"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

var (
	i18nBundle   *i18n.Bundle
	LocalizerBot *i18n.Localizer
)

type SettingService interface {
	GetTgLang() (string, error)
}

func InitLocalizer(i18nFS embed.FS, settingService SettingService) error {
	// set default bundle to english
	i18nBundle = i18n.NewBundle(language.MustParse("en-US"))
	i18nBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// parse files
	if err := parseTranslationFiles(i18nFS, i18nBundle); err != nil {
		return err
	}

	// setup bot locale
	if err := initTGBotLocalizer(settingService); err != nil {
		return err
	}

	return nil
}

func parseTranslationFiles(i18nFS embed.FS, i18nBundle *i18n.Bundle) error {
	err := fs.WalkDir(i18nFS, "translation",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			data, err := i18nFS.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = i18nBundle.ParseMessageFileBytes(data, path)
			return err
		})
	if err != nil {
		return err
	}

	return nil
}

func initTGBotLocalizer(settingService SettingService) error {
	botLang, err := settingService.GetTgLang()
	if err != nil {
		return err
	}

	LocalizerBot = i18n.NewLocalizer(i18nBundle, botLang)
	return nil
}
