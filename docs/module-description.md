# Module description

## What's in the module

- pgtype (`github.com/jackc/pgx/v5/pgtype`) conversion
- Postgres connection pool initialization
- Env variables loading
- Common error handling
- Logging (standard using `fmt` + web using `zap`)
- URL parsing and validation
- Validator initialization and usage

## How to use the module

It's open to public use, so you can use it by running:

```bash
go get github.com/BeaverHouse/go-common
```

## How to update the module

1. Change the code.
2. Push to the repository with tag.
3. Run `go get github.com/BeaverHouse/go-common@latest` to update the module.
