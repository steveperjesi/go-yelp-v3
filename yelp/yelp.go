package yelp

type Client struct {
	AuthOptions AuthOptions
	Debug       bool
}

func NewClient(authOptions AuthOptions) Client {
	return Client{
		AuthOptions: authOptions,
	}
}
