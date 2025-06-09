# Release process

The process of issuing a new release is as follows:

1. Pick a commit to release.
2. Create and push a new tag with the format `vX.Y.Z`, where `X`, `Y`, and `Z` are the major, minor, and patch version numbers respectively.
3. Wait for the CI to build the release artifacts.

That's it! The CI will automatically create a new release on GitHub with the tag name
and the release notes generated from the commit messages,
all thanks to the `Goreleaser` tool.
