package i18n

import (
	"os"

	"golang.org/x/text/language"
)

// Config holds bundle directory and default language.
type Config struct {
	// Dir is the filesystem path containing locale JSON/YAML files.
	Dir string
	// DefaultLang is the BCP 47 tag used when no locale matches (e.g. "en").
	DefaultLang string
}

// DefaultConfig returns sensible defaults (./i18n, en).
func DefaultConfig() Config {
	return Config{
		Dir:         "./i18n",
		DefaultLang: language.English.String(),
	}
}

// Validate checks that Dir exists and is a directory.
func (c Config) Validate() error {
	if c.Dir == "" {
		return os.ErrInvalid
	}
	info, err := os.Stat(c.Dir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return os.ErrInvalid
	}
	return nil
}
