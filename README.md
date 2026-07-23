# redirector

A URL shortener so minimal it forgets everything the moment you restart it.

## API

**Shorten a URL** — `POST /`

```
curl -X POST -d "https://example.com/very/long/path" http://localhost:8080/
```

You'll get back a UUID like `3f8a9b2c-e5f6-7890-abcd-ef1234567890`. That's your short link.

**Get redirected** — `GET /{uuid}`

```
curl -v http://localhost:8080/3f8a9b2c-e5f6-7890-abcd-ef1234567890
```

Returns a `301 Moved Permanently` to wherever you originally pointed it. Browsers work too — just paste the link and go.

## Run it

```sh
go run .             # quick and dirty
```

```sh
docker build -t redirector .
docker run -p 8080:8080 redirector    # for when you want to pretend it's production
```

```
services:
  redirector:
    image: ghcr.io/tomekbielaszewski/redirector:latest
    container_name: redirector
    restart: unless-stopped
```

The server binds to `:8080`. There's no config, no flags, no frills — just vibes.

## Deployment

Push a tag starting with `v` (e.g. `v1.2.3`) and GitHub Actions builds a Docker image, then publishes it to GHCR.

## Caveats

Everything lives in a map in memory. Restart the server, lose your links. This is a feature — it forces you to appreciate impermanence.
