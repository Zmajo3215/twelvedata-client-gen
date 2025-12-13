# To build the client:

In pwsh:

```pwsh
docker run --rm `
  -v <path to where you wanna put it>:/local `
  openapitools/openapi-generator-cli:latest generate `
  -i /local/<name of the file provided>.json `
  -g go `
  -o /local/client
```

Or you can start the pwsh in same directory:
```pwsh
docker run --rm `
  -v "${PWD}:/local" `
  openapitools/openapi-generator-cli:latest generate `
  -i /local/<name of the file provided>.json `
  -g go `
  -o /local/client
```

