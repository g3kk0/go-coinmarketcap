# Go CoinMarketCap

Provides a Go library for working with the CoinMarketCap Professional API. [pro.coinmarketcap.com](https://pro.coinmarketcap.com)

## Installation

```sh
go get github.com/g3kk0/go-coinmarketcap
```

## Usage

Import package and create client.

```go
import coinmarketcap "github.com/g3kk0/go-coinmarketcap"

cmc := coinmarketcap.NewClient()
```

### Get Latest Quotes

```go
// by id
params := map[string]string{"id": "1"}
quote, err := cmc.QuotesLatest(params)
if err != nil {
    panic(err)
}

// by symbol
params := map[string]string{"symbol": "eth"}
quote, err := cmc.QuotesLatest(params)
if err != nil {
    panic(err)
}

// request multiple symbols (works for ids too)
params := map[string]string{"symbol": "bch,eth"}
quotes, err := cmc.QuotesLatest(params)
if err != nil {
    panic(err)
}

// specify quote currency (default: usd)
params := map[string]string{"symbol": "bch", "convert": "gbp"}
quote, err := cmc.QuotesLatest(params)
if err != nil {
    panic(err)
}
```
