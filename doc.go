// Package i18n loads JSON message bundles via go-i18n with config injection (no viper).
//
// Design constraints:
//   - Callers inject Config; the library does not read viper or environment variables.
//   - No package-level singletons; construct an Engine per process or scope.
package i18n
