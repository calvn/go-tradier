# go-tradier
Golang library for interacting with the Tradier API

***Note:*** *This library is still under development, please use with caution!*

## Authentication
go-tradier does not directly handle authentication. However, uses the `http.Client`, so authentication can be done by passing an `http.Client` that can handle authentication. We encourage using the [oauth2](https://github.com/golang/oauth2) library to achieve proper authentication.

```go
import "golang.org/x/oauth2"

func main() {
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "... your access token ..."},
  )
  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := tradier.NewClient(tc)

  // list all repositories for the authenticated user
  profile, _, err := client.User.Profile()
}
```

## License

This library is licensed under the MIT License as provided in [here](LICENSE.md).

*Made with <3 in Go. Heavily borrowed from Google's go-github library*
