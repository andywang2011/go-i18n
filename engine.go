package i18n

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Engine wraps a go-i18n bundle.
type Engine struct {
	Bundle      *goi18n.Bundle
	defaultLang string
}

// NewEngine creates an Engine from Config.
func NewEngine(cfg Config) (*Engine, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("i18n config: %w", err)
	}
	tag := language.English
	if cfg.DefaultLang != "" {
		if t, err := language.Parse(cfg.DefaultLang); err == nil {
			tag = t
		}
	}
	bundle := goi18n.NewBundle(tag)
	defLang := tag.String()
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	err := filepath.WalkDir(cfg.Dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}
		_, err = bundle.LoadMessageFile(path)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("i18n load messages: %w", err)
	}

	return &Engine{
		Bundle:      bundle,
		defaultLang: defLang,
	}, nil
}

// NewLocalizer returns a per-request localizer.
func (e *Engine) NewLocalizer(lang string) *goi18n.Localizer {
	return goi18n.NewLocalizer(e.Bundle, lang, e.defaultLang)
}
