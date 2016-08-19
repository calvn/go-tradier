# TODO

## General
- [ ] Rate limiting tracking on the client
- [ ] Determine best type (probably not float64) for holding currency
- [ ] Handle errors gracefully (e.g. "Invalid access token")

## User
|                | endpoint | test coverage |
|----------------|----------|---------------|
| user/profile   |          |               |
| user/balances  |          |               |
| user/positions |          |               |
| user/history   |          |               |
| user/gainloss  |          |               |
| user/orders    |          |               |


## Account

## Trading

## Market Data

## Fundamentals

## Watchlists

## Streaming


## Notes

Response returned can be formed differently depending on underlying requested object size.

Example is based off `https://api.tradier.com/v1/accounts/6YA05708/orders` endpoint

- When the result is empty, it returns object with `"null"` as its value
```json
{
  "orders": "null"
}
```

- When the result is single value, it returns single object
```json
{
  "orders": {
    "order": {
      "id": 179686,
      "type": "market",
      "symbol": "QQQ",
      "side": "sell",
      "quantity": 25,
      "status": "pending",
      "duration": "day",
      "avg_fill_price": 0,
      "exec_quantity": 0,
      "last_fill_price": 0,
      "last_fill_quantity": 0,
      "remaining_quantity": 25,
      "create_date": "2016-08-18T04:37:06.511Z",
      "transaction_date": "2016-08-18T04:37:06.527Z",
      "class": "equity"
    }
  }
}
```

- When the result is multiple values, it returns an array of objects
```json
{
  "orders": {
    "order": [
      {
        "id": 179686,
        "type": "market",
        "symbol": "QQQ",
        "side": "sell",
        "quantity": 25,
        "status": "pending",
        "duration": "day",
        "avg_fill_price": 0,
        "exec_quantity": 0,
        "last_fill_price": 0,
        "last_fill_quantity": 0,
        "remaining_quantity": 25,
        "create_date": "2016-08-18T04:37:06.511Z",
        "transaction_date": "2016-08-18T04:37:06.527Z",
        "class": "equity"
      },
      {
        "id": 179687,
        "type": "market",
        "symbol": "SPY",
        "side": "sell",
        "quantity": 225,
        "status": "pending",
        "duration": "day",
        "avg_fill_price": 0,
        "exec_quantity": 0,
        "last_fill_price": 0,
        "last_fill_quantity": 0,
        "remaining_quantity": 225,
        "create_date": "2016-08-18T04:39:46.585Z",
        "transaction_date": "2016-08-18T04:39:46.600Z",
        "class": "equity"
      }
    ]
  }
}
```
