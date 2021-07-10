package config

type (
	Config struct {
		Database Database `yaml:"database"`
		I18n I18n `yaml:"i18n"`
	}

	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}

	I18n struct {
		BundlePath string `yaml:"bundle_path"`
	}
)
