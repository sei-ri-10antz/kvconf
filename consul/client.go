package consul

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sei-ri/kvconf"
)

const (
	DefaultRequestTimeout = 5 * time.Second
	DefaultEndpoint       = "localhost:8500"
	DefaultSchema         = "http"
)

type client struct {
	endpoint       string
	requestTimeout time.Duration
	schema         string
	httpClient     *http.Client
}

func NewClient(ctx context.Context, opts ...Option) (*client, error) {
	c := &client{
		endpoint:       DefaultEndpoint,
		requestTimeout: DefaultRequestTimeout,
		schema:         DefaultSchema,
		httpClient:     &http.Client{Timeout: DefaultRequestTimeout},
	}

	for i := range opts {
		if err := opts[i](c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *client) Get(ctx context.Context, key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, c.requestTimeout)
	defer cancel()

	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: c.schema,
			Host:   c.endpoint,
			Path:   fmt.Sprintf("/v1/kv/%s", strings.TrimPrefix(key, "/")),
		},
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("%s not found", key)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var pairs kvconf.Pairs
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&pairs); err != nil {
		return nil, err
	}

	if len(pairs) > 0 {
		v, err := base64.StdEncoding.DecodeString(string(pairs[0].Value))
		if err != nil {
			fmt.Errorf("failed to base64 decode %v", err)
			return []byte{}, nil
		}
		return v, nil
	}

	return []byte{}, nil
}

type Option func(c *client) error

func WithEndpoint(s string) Option {
	return func(c *client) error {
		if len(s) == 0 {
			return nil
		}
		c.endpoint = s
		return nil
	}
}

func WithRequestTimeout(d time.Duration) Option {
	return func(c *client) error {
		if d < 0 {
			return nil
		}
		c.requestTimeout = d
		return nil
	}
}

func WithSchema(s string) Option {
	return func(c *client) error {
		if len(s) == 0 {
			return nil
		}
		c.schema = s
		return nil
	}
}
