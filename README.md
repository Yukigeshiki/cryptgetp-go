# cryptgetp-go
Crypt-Get-P - a just for fun CLI tool to fetch cryptocurrency prices written in Go.

### How to run:

Clone the repo, cd into it and run:

```bash
go build cryptgetp.go
```

The pricing data is fetched from coinapi.io, so you'll need to get a free API key from [here](https://www.coinapi.io/pricing?apikey).

Once you have an API key you can run: 

```bash
./cryptgetp fetch --crypto BTC --in USD --key <your-key-here>
```

You get 100 free calls per day with your API key. More information about the API can be found [here](https://www.coinapi.io).