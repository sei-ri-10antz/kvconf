package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/sei-ri/kvconf"
	"time"
)

const (
	DefaultDialTimeout    = 5 * time.Second
	DefaultRequestTimeout = 5 * time.Second
	DefaultEndpoint       = "localhost:2379"
)

type client struct {
	db             *clientv3.Client
	config         clientv3.Config
	requestTimeout time.Duration
}

func NewClient(ctx context.Context, opts ...Option) (*client, error) {
	c := &client{
		config: clientv3.Config{
			Endpoints:   []string{DefaultEndpoint},
			DialTimeout: DefaultDialTimeout,
		},
		requestTimeout: DefaultRequestTimeout,
	}

	for i := range opts {
		if err := opts[i](c); err != nil {
			return nil, err
		}
	}

	cli, err := clientv3.New(c.config)
	if err != nil {
		return nil, err
	}
	c.db = cli

	return c, nil
}

func (c *client) Close() error {
	return c.db.Close()
}

func (c *client) Get(ctx context.Context, key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, c.requestTimeout)
	defer cancel()

	resp, err := c.db.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if resp.Count > 0 {
		return resp.Kvs[0].Value, nil
	}

	return []byte{}, nil
}

func (c *client) Update(ctx context.Context, pair *kvconf.Pair) (*kvconf.Pair, error) {
	panic("implement me")
}

type Option func(c *client) error

func WithEndpoints(s ...string) Option {
	return func(c *client) error {
		if len(s) == 0 {
			return nil
		}
		c.config.Endpoints = s
		return nil
	}
}

func WithDialTimeout(d time.Duration) Option {
	return func(c *client) error {
		if d < 0 {
			return nil
		}
		c.config.DialTimeout = d
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
