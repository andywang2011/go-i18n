package i18n

import (
	"fmt"
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

// ConfigDefault is an alias for DefaultConfig (naming consistency across go-* libs).
func ConfigDefault() Config { return DefaultConfig() }

// Validate checks that Dir exists and is a directory, and that a non-empty
// DefaultLang is a well-formed BCP 47 tag.
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
	if c.DefaultLang != "" {
		if _, err := language.Parse(c.DefaultLang); err != nil {
			return fmt.Errorf("invalid default lang %q: %w", c.DefaultLang, err)
		}
	}
	return nil
}
