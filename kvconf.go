package kvconf

import (
	"context"
)

type Client interface {
	Get(context.Context, string) ([]byte, error)
}

type Pair struct {
	Key   string
	Value string
}

type Pairs []Pair