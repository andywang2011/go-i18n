# go-i18n

JSON message bundle loader for [nicksnyder/go-i18n](https://github.com/nicksnyder/go-i18n) (`github.com/andywang2011/go-i18n`).

```bash
go get github.com/andywang2011/go-i18n@latest
```

**Not** a fork of go-i18n — a thin wrapper that loads `*.json` from a directory and exposes `Engine.NewLocalizer(lang)`.

## Quick start

### 1. Message files

Place one or more JSON files under a directory (e.g. `./i18n/`):

```json
[
  {
    "id": "web.lib_no_route",
    "translation": "路由未找到"
  },
  {
    "id": "web.lib_rate_limit",
    "translation": "服务繁忙，请稍后再试"
  }
]
```

File names are arbitrary (`zh.json`, `en.json`, …); all `*.json` files in the directory are loaded.

### 2. Create engine

```go
import (
	"fmt"

	"github.com/andywang2011/go-i18n"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

engine, err := i18n.NewEngine(i18n.Config{
	Dir:         "./i18n",
	DefaultLang: "en", // fallback when request lang missing
})
if err != nil {
	return err
}
```

Or use defaults (`./i18n`, `en`):

```go
engine, err := i18n.NewEngine(i18n.DefaultConfig())
```

### 3. Per-request localizer

```go
func handleRequest(engine *i18n.Engine, lang string) (string, error) {
	loc := engine.NewLocalizer(lang) // e.g. "zh", "en", "ja"
	return loc.Localize(&goi18n.LocalizeConfig{
		MessageID: "web.lib_no_route",
	})
}
```

`NewLocalizer` tries `lang` first, then falls back to `DefaultLang`.

## Configuration

Set `Config` fields directly (no viper in this library):

| Field | Description |
|-------|-------------|
| `Dir` | Directory containing `*.json` message files (must exist). |
| `DefaultLang` | Fallback language tag (e.g. `en`, `zh`). |

In **demo-api**, map viper keys in `demo-api/wire/config.LoadI18nConfig()`:

```yaml
i18n:
  dir: ./i18n
  default_lang: en
```

Wire injects: `LoadI18nConfig()` → `NewI18nEngine(cfg)` → `*i18n.Engine`.

## Examples

Runnable example (creates temp message files, prints `zh` / `en` / fallback):

```bash
cd examples/basic
go run .
```

Expected output:

```text
[zh] 你好
[en] Hello
[fr] Hello
```

(`fr` falls back to `DefaultLang: en`.)

## API

| Method | Description |
|--------|-------------|
| `DefaultConfig()` | Returns `Dir: ./i18n`, `DefaultLang: en`. |
| `NewEngine(cfg)` | Walk `cfg.Dir`, load JSON bundles; returns error if dir invalid. |
| `Engine.NewLocalizer(lang)` | Per-request `*goi18n.Localizer` with fallback to default lang. |
| `Config.Validate()` | Ensures `Dir` exists and is a directory. |

## Design

- Config injection only — viper / nacos stay in the application layer.
- Stdlib walk + go-i18n bundle; no global singleton.
- Reuse one `*Engine` per process; create a `Localizer` per request.

## 中文摘要

从目录加载 JSON 多语言文件，封装 go-i18n Bundle。`Config{Dir, DefaultLang}` 由应用层注入；demo-api 通过 `wire/config.LoadI18nConfig()` 读 `i18n.dir` / `i18n.default_lang`。
