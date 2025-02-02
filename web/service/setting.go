package service

import (
	_ "embed"
	"strconv"
	"strings"
	"time"
	"x-ui-scratch/database"
	"x-ui-scratch/database/model"
	"x-ui-scratch/logger"
	"x-ui-scratch/util/common"
	"x-ui-scratch/util/random"
)

type SettingService struct{}

var defaultValueMap = map[string]string{
	"timeLocation": "Asia/Tehran",

	"webDomain": "",

	"webBasePath": "/",

	"tgLang": "en-US",

	"secret": random.Seq(32),

	"webListen": "",
	"webPort":   "2054",

	"secretEnable":       "false",
	"xrayTemplateConfig": xrayTemplateConfig,
}

//go:embed config.json
var xrayTemplateConfig string

func (s *SettingService) GetTimeLocation() (*time.Location, error) {
	l, err := s.getString("timeLocation")
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocation(l)
	if err != nil {
		defaultLocation := defaultValueMap["timeLocation"]
		// logger.Errorf("location <%v> not exist, using default location: %v", l, defaultLocation)
		return time.LoadLocation(defaultLocation)
	}
	return location, nil
}

func (s *SettingService) getString(key string) (string, error) {
	logger.Info(key)
	setting, err := s.getSetting(key)
	if database.IsNotFound(err) {
		value, ok := defaultValueMap[key]
		if !ok {
			return "", common.NewErrorf("key <%v> not in defaultValueMap", key)
		}
		return value, nil
	} else if err != nil {
		return "", err
	}
	return setting.Value, nil
}

func (s *SettingService) getSetting(key string) (*model.Setting, error) {
	db := database.GetDB()
	setting := &model.Setting{}
	err := db.Model(model.Setting{}).Where("key = ?", key).First(setting).Error
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (s *SettingService) GetWebDomain() (string, error) {
	return s.getString("webDomain")
}

func (s *SettingService) GetSecret() ([]byte, error) {
	secret, err := s.getString("secret")
	if secret == defaultValueMap["secret"] {
		err := s.saveSetting("secret", secret)
		if err != nil {
			// TODO: change it to Warning
			logger.Info("save secret failed:", err)
		}
	}
	return []byte(secret), err
}

func (s *SettingService) saveSetting(key string, value string) error {
	setting, err := s.getSetting(key)
	db := database.GetDB()
	if database.IsNotFound(err) {
		return db.Create(&model.Setting{
			Key:   key,
			Value: value,
		}).Error
	} else if err != nil {
		return err
	}
	setting.Key = key
	setting.Value = value
	return db.Save(setting).Error
}

func (s *SettingService) GetBasePath() (string, error) {
	basePath, err := s.getString("webBasePath")
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(basePath, "/") {
		basePath = "/" + basePath
	}
	if !strings.HasSuffix(basePath, "/") {
		basePath += "/"
	}
	return basePath, nil
}

func (s *SettingService) GetTgLang() (string, error) {
	return s.getString("tgLang")
}

func (s *SettingService) GetListen() (string, error) {
	return s.getString("webListen")
}

func (s *SettingService) GetPort() (int, error) {
	return s.getInt("webPort")
}

func (s *SettingService) getInt(key string) (int, error) {
	str, err := s.getString(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}

func (s *SettingService) GetSessionMaxAge() (int, error) {
	return s.getInt("sessionMaxAge")
}

func (s *SettingService) GetSecretStatus() (bool, error) {
	return s.getBool("secretEnable")
}

func (s *SettingService) getBool(key string) (bool, error) {
	logger.Info("haha")
	str, err := s.getString(key)
	if err != nil {
		logger.Info("haha2")
		return false, err
	}
	return strconv.ParseBool(str)
}

func (s *SettingService) GetXrayConfigTemplate() (string, error) {
	return s.getString("xrayTemplateConfig")
}

func (s *SettingService) GetDefaultSettings(host string) (interface{}, error) {
	type settingFunc func() (interface{}, error)
	settings := map[string]settingFunc{
		/* "expireDiff":    func() (interface{}, error) { return s.GetExpireDiff() },
		"trafficDiff":   func() (interface{}, error) { return s.GetTrafficDiff() },
		"pageSize":      func() (interface{}, error) { return s.GetPageSize() },
		"defaultCert":   func() (interface{}, error) { return s.GetCertFile() },
		"defaultKey":    func() (interface{}, error) { return s.GetKeyFile() },
		"tgBotEnable":   func() (interface{}, error) { return s.GetTgbotEnabled() },
		"subEnable":     func() (interface{}, error) { return s.GetSubEnable() },
		"subURI":        func() (interface{}, error) { return s.GetSubURI() },
		"subJsonURI":    func() (interface{}, error) { return s.GetSubJsonURI() },
		"remarkModel":   func() (interface{}, error) { return s.GetRemarkModel() },
		"datepicker":    func() (interface{}, error) { return s.GetDatepicker() },
		"ipLimitEnable": func() (interface{}, error) { return s.GetIpLimitEnable() }, */
	}

	result := make(map[string]interface{})

	for key, fn := range settings {
		value, err := fn()
		if err != nil {
			return "", err
		}
		result[key] = value
	}

	// todo: to complement the code.
	if result["subEnable"].(bool) && (result["subURI"].(string) == "" || result["subJsonURI"].(string) == "") {
		// subURI := ""
		/* subPort, _ := s.GetSubPort()
		subPath, _ := s.GetSubPath()
		subJsonPath, _ := s.GetSubJsonPath()
		subDomain, _ := s.GetSubDomain()
		subKeyFile, _ := s.GetSubKeyFile()
		subCertFile, _ := s.GetSubCertFile()
		subTLS := false
		if subKeyFile != "" && subCertFile != "" {
			subTLS = true
		}
		if subDomain == "" {
			subDomain = strings.Split(host, ":")[0]
		}
		if subTLS {
			subURI = "https://"
		} else {
			subURI = "http://"
		}
		if (subPort == 443 && subTLS) || (subPort == 80 && !subTLS) {
			subURI += subDomain
		} else {
			subURI += fmt.Sprintf("%s:%d", subDomain, subPort)
		}
		if result["subURI"].(string) == "" {
			result["subURI"] = subURI + subPath
		}
		if result["subJsonURI"].(string) == "" {
			result["subJsonURI"] = subURI + subJsonPath
		} */
	}

	return result, nil
}
