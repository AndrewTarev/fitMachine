package config

import "time"

// IConfig интерфейс для работы с конфигурацией
type IConfig interface {
	GetInt(key string) int
	GetFloat64(key string) float64
	GetString(key string) string
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetSliceOfInt(key string) []int
	GetSliceOfString(key string) []string
}
