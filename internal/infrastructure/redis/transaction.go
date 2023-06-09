package redis

import (
	"context"

	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/go-redis/redis/v8"
)

type Transaction struct {
	pipe redis.Pipeliner
	ctx  context.Context
}

func (d *Transaction) Commit() error {
	_, err := d.pipe.Exec(d.ctx)
	return err
}

func (d *Transaction) Rollback() error {
	return d.pipe.Discard()
}

func (d *Transaction) Context() context.Context {
	return d.ctx
}

func (d *Transaction) Transaction() interface{} {
	return d.pipe
}

func NewTransaction(pipe redis.Pipeliner, ctx context.Context) resolver.Transactioner {
	return &Transaction{pipe: pipe, ctx: ctx}
}
