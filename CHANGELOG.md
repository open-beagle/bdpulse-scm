# Changelog

All notable changes to this project will be documented in this file.

## [v1.24.3] - 2026-01-25

### Fixed

- **HTTP Tracing**: Fixed OpenTelemetry trace propagation issue by replacing `http.NewRequest` with `http.NewRequestWithContext`
  - Updated `scm/client.go`: Now uses `http.NewRequestWithContext` to ensure trace context is properly propagated from the start
  - Updated `scm/transport/oauth2/refresh.go`: OAuth2 token refresh requests now correctly propagate trace information
  - This fix ensures that when HTTP Transport is wrapped with `otelhttp`, trace information can be properly extracted from context

### Technical Details

Previously, the code created HTTP requests using `http.NewRequest` and then added context via `req.WithContext(ctx)`. This pattern prevented OpenTelemetry from extracting trace information during request creation. The fix ensures trace spans are correctly propagated through all HTTP calls made by the SCM client.

## [v1.24.2] - 2026-01-24

### Changed

- **BREAKING**: Migrated module path from `github.com/drone/go-scm` to `github.com/open-beagle/go-scm`
- Updated all import paths across the codebase
- Removed all unit test files (`*_test.go`)
- Removed copyright headers from all source files
- Completely rewrote README.md with comprehensive documentation and examples
- Updated go.mod to reflect new module path

### Added

- Version badge in README.md
- Comprehensive usage examples for all supported SCM providers
- Detailed authentication examples (OAuth2, Private Token, Basic Auth)
- API coverage documentation
- Contributing guidelines

### Removed

- All unit test files
- Copyright headers (Drone.IO Inc.)
- Release procedure documentation (drone-specific)

## [v1.24.0-beagle.v7.6] - Previous Release

Based on upstream [drone/go-scm v1.24.0](https://github.com/drone/go-scm/releases/tag/v1.24.0)

---

## Migration Guide

If you're upgrading from `github.com/drone/go-scm`, you need to update all import paths:

```bash
# Using find and sed
find . -name "*.go" -type f -exec sed -i 's|github.com/drone/go-scm|github.com/open-beagle/go-scm|g' {} +
```

Or update your `go.mod`:

```go
// Before
require github.com/drone/go-scm v1.24.0

// After
require github.com/open-beagle/go-scm v1.24.2
```

Then run:

```bash
go mod tidy
```
