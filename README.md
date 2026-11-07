# go-i18n

Load JSON message bundles for [nicksnyder/go-i18n](https://github.com/nicksnyder/go-i18n) with config injection.

## Install

```bash
go get github.com/andywang2011/go-i18n@v0.1.0
```

## Quick start

```go
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andywang2011/go-i18n"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

dir, err := os.MkdirTemp("", "go-i18n-example-*")
if err != nil {
	fmt.Println("tmpdir error")
	return
}
defer os.RemoveAll(dir)

content := `[{"id":"greeting","translation":"Hello"}]`
if err := os.WriteFile(filepath.Join(dir, "en.json"), []byte(content), 0644); err != nil {
	fmt.Println("write error")
	return
}

engine, err := i18n.NewEngine(i18n.Config{Dir: dir, DefaultLang: "en"})
if err != nil {
	fmt.Println("engine error")
	return
}

msg, err := engine.NewLocalizer("en").Localize(&goi18n.LocalizeConfig{MessageID: "greeting"})
if err != nil {
	fmt.Println("localize error")
	return
}
fmt.Println(msg)
```

## Examples

```bash
cd go-i18n/examples && go run .
```

## License

MIT — see [LICENSE](LICENSE).
