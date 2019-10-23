package config
import "time"

type Config struct {
	Period time.Duration `config:"period"`
	Path   string        `config:"path"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	Path:   ".",
}
