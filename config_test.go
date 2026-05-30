package i18n

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/text/language"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Dir != "./i18n" {
		t.Errorf("Dir = %q", cfg.Dir)
	}
	if cfg.DefaultLang != language.English.String() {
		t.Errorf("DefaultLang = %q", cfg.DefaultLang)
	}
}

func TestConfig_Validate(t *testing.T) {
	if err := (Config{Dir: t.TempDir()}).Validate(); err != nil {
		t.Errorf("Validate(dir) = %v", err)
	}
	if err := (Config{}).Validate(); err == nil {
		t.Error("Validate(empty) expected error")
	}
	filePath := filepath.Join(t.TempDir(), "file.txt")
	if err := os.WriteFile(filePath, []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := (Config{Dir: filePath}).Validate(); err == nil {
		t.Error("Validate(file) expected error")
	}
}
