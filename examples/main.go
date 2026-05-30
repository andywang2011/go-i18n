// Example: load JSON message files and localize by language.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	libi18n "github.com/andywang2011/go-i18n"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	dir := filepath.Join(os.TempDir(), "go-i18n-example")
	if err := writeMessages(dir); err != nil {
		log.Fatal(err)
	}

	engine, err := libi18n.NewEngine(libi18n.Config{
		Dir:         dir,
		DefaultLang: "en",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, lang := range []string{"zh", "en", "fr"} {
		loc := engine.NewLocalizer(lang)
		msg, err := loc.Localize(&goi18n.LocalizeConfig{MessageID: "greeting"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%s] %s\n", lang, msg)
	}
}

func writeMessages(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	files := map[string]string{
		"zh.json": `[{"id":"greeting","translation":"你好"}]`,
		"en.json": `[{"id":"greeting","translation":"Hello"}]`,
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0644); err != nil {
			return err
		}
	}
	return nil
}
