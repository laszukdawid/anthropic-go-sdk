# (unofficial) Anthropic Go SDK

Repo contains unofficial Golang SDK for Anthropic (https://docs.anthropic.com/en/api/getting-started).

The client is based on OpenAPI specification from https://github.com/laszukdawid/anthropic-openapi-spec. Self-curated specs were requires since there are no official ones.

## Reproduce

To reproduce the client execute:

```shell
task spec:download  # populates /specs
task spec:clean     # gets rid of incorrect paths
task run:generate   # generates client
```

## Example

The client isn't the easiest to work with, sorry. In case you want a quick example on what to do, check out [client/client_test.go](client/client_test.go).

To execute the test locally, when you have this repo built, execute

```shell
API_KEY="your-api-key-from-anthropic" go test ./...
```
