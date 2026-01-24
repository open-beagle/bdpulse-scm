# Changelog

All notable changes to this project will be documented in this file.

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
