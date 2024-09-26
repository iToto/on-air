package client

// Option represents the client options
type Option func(*HTTPClient)

// WithCustomClient sets a custom http client
func WithCustomClient(client Doer) Option {
	return func(c *HTTPClient) {
		c.client = client
	}
}

// WithMiddleware adds a RequestLogger middleware to the httpclient
func WithMiddleware(mid Middleware) Option {
	return func(c *HTTPClient) {
		if mid != nil {
			c.middlewares = append(c.middlewares, mid)
		}
	}
}
