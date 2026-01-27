# version

Current version: **v1.24.4**

Based on upstream: [drone/go-scm](https://github.com/drone/go-scm)

## Changelog

### v1.24.4 (2026-01-27)

- Fixed OAuth2 token refresh trace context loss by passing context to Refresh method

### v1.24.3 (2026-01-24)

- Migrated module path from `github.com/drone/go-scm` to `github.com/open-beagle/go-scm`
- Removed all unit test files
- Removed copyright headers
- Updated README.md with comprehensive documentation
- Cleaned up project structure

## Upstream Sync

```bash
# Add upstream remote
git remote add upstream git@github.com:drone/go-scm.git

# Fetch upstream changes
git fetch upstream

# Merge upstream version
git merge v1.24.0
```

## Release

```bash
# build test
go build ./...

# Create a release tag
git tag v1.24.4

# Push tag (use -f to force update)
git push -f origin v1.24.4

# Delete local tag
git tag -d v1.24.4

# Delete remote tag
git push origin :refs/tags/v1.24.4
```

## Version History

- **v1.24.4**: Fixed OAuth2 trace context propagation
- **v1.24.3**: Current version with open-beagle migration
- **v1.24.0-beagle**: Previous version based on upstream v1.24.0
