# To build the client:

(The client was built in golang, so some tags had to be removed)
The removed tags were market_data tags in openapi.json file for /time_series and /time_series_cross because they alread had time_series tag,
and so duplicated code would be generated inside time_series.go for market_data and time_series and same inside time_series_cross.go

So before you could do:
```go
apiClient.MarketDataAPI.GetTimeSeries() // OR
apiClient.TimeSeriesAPI.GetTimeSeries()
```
And get the same result, but now, you can not do that with MarketDataAPI

In pwsh:

```pwsh
docker run --rm `
  -v <path_to_where_you_wanna_put_it>:/local `
  openapitools/openapi-generator-cli:latest generate `
  -i /local/<name_of_the_file_provided>.json `
  -g go `
  -o /local/client
```

Or you can start the pwsh in same directory:
```pwsh
docker run --rm `
  -v "${PWD}:/local" `
  openapitools/openapi-generator-cli:latest generate `
  -i /local/openapi.json `
  -g go `
  -o /local/client
```

Also required modules had to be installed in client library:

In client directory:
```go
go get github.com/stretchr/testify@latest
go mod tidy
```



