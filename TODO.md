# TODO

## General

- [ ] Rate limiting tracking on the client
- [ ] Determine best type (probably not float64) for holding currency
- [ ] Handle errors gracefully
  - [ ] `json: Invalid access token`
  - [ ] `json: cannot unmarshal number 0E-8 into Go value of type int` on `LastFillQuantity` and such
  - [ ] Better error handling on `MarshalJSON` and `UnmarshalJSON`

## Structs


## User endpoints

|                | endpoint | test coverage |
|----------------|----------|---------------|
| user/profile   | ✓        | ✓             |
| user/balances  | ✓        |               |
| user/positions | ✓        |               |
| user/history   | ✓        |               |
| user/gainloss  | ✓        |               |
| user/orders    | ✓        |               |

- [x] ~~Fix `OrdersAccountEntry` to dynamically map to object~~
- [x] `Order` should support indexing if it is a slice


## Account endpoints

|                   | endpoint | test coverage |
|-------------------|----------|---------------|
| account/balances  |          |               |
| account/positions |          |               |
| account/history   |          |               |
| account/gainloss  |          |               |
| account/orders    |          |               |
| account/status    |          |               |

## Trading endpoints

## Market Data endpoints

## Fundamentals endpoints

## Watchlists endpoints

## Streaming endpoints
