package locale

import "embed"

type SettingService interface {
	// GetTgLang() (string, error)
}

func InitLocalizer(i18nFS embed.FS, settingService SettingService) error {
	// TODO
	return nil
}
