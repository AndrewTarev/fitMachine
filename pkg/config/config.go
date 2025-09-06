package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

// New создает новый экземпляр конфигурации
func New() (IConfig, error) {
	v := viper.New()

	// Определяем окружение
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // значение по умолчанию
	}

	// Определяем имя конфигурационного файла
	var configName string
	switch env {
	case "local":
		configName = "local"
	case "prod":
		configName = "fitMachine/prod"
	default:
		return Config{}, fmt.Errorf("not supported config env: %s", env)
	}

	v.SetConfigName(configName)
	v.SetConfigType("yml")

	v.AddConfigPath("./configs")
	v.AutomaticEnv()
	v.SetEnvPrefix("")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return Config{}, fmt.Errorf("not supported config env: %s", configFileNotFoundError)
		}
	}

	return Config{Viper: v}, nil
}

// GetInt возвращает значение как int
func (c Config) GetInt(key string) int {
	return c.Viper.GetInt(key)
}

// GetUint8 возвращает значение как uint8
func (c Config) GetUint8(key string) uint8 {
	return uint8(c.Viper.GetUint(key))
}

// GetFloat64 возвращает значение как float64
func (c Config) GetFloat64(key string) float64 {
	return c.Viper.GetFloat64(key)
}

// GetString возвращает значение как string
func (c Config) GetString(key string) string {
	return c.Viper.GetString(key)
}

// GetBool возвращает значение как bool
func (c Config) GetBool(key string) bool {
	return c.Viper.GetBool(key)
}

// GetDuration возвращает значение как time.Duration
func (c Config) GetDuration(key string) time.Duration {
	return c.Viper.GetDuration(key)
}

// GetSliceOfInt возвращает значение как []int
func (c Config) GetSliceOfInt(key string) []int {
	return c.GetIntSlice(key)
}

// GetSliceOfString возвращает значение как []string
func (c Config) GetSliceOfString(key string) []string {
	return c.GetStringSlice(key)
}
