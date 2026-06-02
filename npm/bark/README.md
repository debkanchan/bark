# 🐕 bark

> An "embarrassment linter" — detects `BARK` comments in your code so you never push temporary code by accident.

This is the npm distribution of [`bark`](https://github.com/debkanchan/bark). It installs a small launcher plus a prebuilt native binary for your platform (selected automatically — no compiler needed).

## Install

```bash
npm install -g @debkanchan/bark
```

Then use it:

```bash
bark .                       # scan the current directory
bark ./src main.go           # scan specific paths
bark --format json .         # JSON output for CI
bark git-hook install        # install a pre-commit hook
```

## How it works

`@debkanchan/bark` declares one platform package per OS/arch (e.g. `@debkanchan/bark-darwin-arm64`) as an `optionalDependency`, gated by npm's `os`/`cpu` fields. npm installs only the package matching your machine, and the `bark` launcher execs the bundled binary.

Supported platforms: linux x64/arm64, macOS x64/arm64, Windows x64.

## Other install methods

- Homebrew: `brew install debkanchan/tap/bark`
- From source (needs a C compiler — uses CGO): `go install github.com/debkanchan/bark/cmd/bark@latest`

See the [project README](https://github.com/debkanchan/bark#readme) for full docs.

## License

AGPL-3.0-or-later
