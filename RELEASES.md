# Release Process

1. Update the version in `cmd/constants.go`
2. PR to `main` branch and merge if tests pass
3. Ensure GitHub Actions tests pass on `main` before creating a release
4. Create a new release on GitHub on the `main` branch with a `v` prefixed tag (e.g. `v0.1.0`)
5. GoReleaser will automatically build cross-platform binaries and attach them to the release
