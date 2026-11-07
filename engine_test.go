package i18n

import (
	"os"
	"path/filepath"
	"testing"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestNewEngine(t *testing.T) {
	dir := t.TempDir()
	content := `[{"id":"test.key","translation":"test value"}]`
	if err := os.WriteFile(filepath.Join(dir, "zh.json"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := Config{Dir: dir, DefaultLang: "en"}
	engine, err := NewEngine(cfg)
	if err != nil {
		t.Fatalf("NewEngine err = %v", err)
	}
	if engine.Bundle == nil {
		t.Fatal("Bundle is nil")
	}
	loc := engine.NewLocalizer("zh")
	if loc == nil {
		t.Fatal("NewLocalizer returned nil")
	}
	msg, err := loc.Localize(&goi18n.LocalizeConfig{MessageID: "test.key"})
	if err != nil || msg != "test value" {
		t.Errorf("Localize = %q, err = %v", msg, err)
	}
}

func TestNewEngine_InvalidDir(t *testing.T) {
	cfg := Config{Dir: filepath.Join(t.TempDir(), "nonexistent")}
	_, err := NewEngine(cfg)
	if err == nil {
		t.Error("expected error for invalid dir")
	}
}

func TestNewLocalizer_DefaultDir(t *testing.T) {
	engine, err := NewEngine(DefaultConfig())
	if err != nil {
		t.Skipf("default i18n dir not present: %v", err)
	}
	loc := engine.NewLocalizer("zh")
	if loc == nil {
		t.Fatal("NewLocalizer returned nil")
	}
}
