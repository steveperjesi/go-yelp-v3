# go-yelp-v3
Go service to communicate with the Yelp v3 API

Grab your Yelp v3 API key from 
https://www.yelp.com/developers/v3/manage_app

Add it to the `config.yaml`

or

Add to your environment

```export YELP_V3_API_KEY="<YOUR YELP V3 API KEY>"```

```go run examples/*.go```

```go test yelp/* -cover```

_ok  	command-line-arguments	6.946s	**coverage: 91.5%** of statements_
