<p align="center">
  <a href="https://github.com/BeaverHouse/go-common">
    <img src="logo.png" alt="Logo" height="100"> 
  </a>

  <p align="center">
    Common Go library for personal use, in clean, standard code style
  </p>

  <p align="center">
    <a href="https://golang.org/">
      <img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=go&logoColor=white" alt="Go">
    </a>
    <a href="./LICENSE">
      <img src="https://img.shields.io/github/license/BeaverHouse/go-common" alt="License">
    </a>
  </p>
</p>

<!-- Content -->

<br>

## Description

Common Go library for personal use, in clean, standard code style.

<br>

## What's in the module

- `conv` — `pgtype` (`github.com/jackc/pgx/v5/pgtype`) ⇄ Go pointer conversion
- `database/postgres` — Postgres connection pool initialization
- `env` — environment variable loading
- `errorhandle` — protocol-agnostic error kinds and adapter mapping
- `logger` — structured logging (`SimpleLogger` for local, JSON `ZapLogger` for deployed)
- `urlutil` — URL parsing and validation
- `validation` — validator initialization and usage

<br>

## Installation

```bash
go get github.com/BeaverHouse/go-common
```

<br>

## Releasing

1. Change the code and merge to `main`.
2. Run the `release` workflow (Actions tab). It tags `v1.0.<YYYYMMDD>` by default,
   or any SemVer tag you pass as input.
3. Run `go get github.com/BeaverHouse/go-common@latest` to pull the new version.

<br>

## Contributing

See the [CONTRIBUTING.md](./CONTRIBUTING.md).

<br>
<br>

## Attribution

Logo icon is downloaded from [Go official homepage](https://go.dev/blog/go-brand/Go-Logo/PNG/).
