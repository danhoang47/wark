package appcontext

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AppContext interface {
	GetDB() *sqlx.DB
	GetMemCached() *redis.Client
	GetSecret() string
}

type appContext struct {
	db     *sqlx.DB
	cache  *redis.Client
	secret string
}

func New(db *sqlx.DB, cache *redis.Client, secret string) *appContext {
	return &appContext{db, cache, secret}
}

func (appCtx *appContext) GetDB() *sqlx.DB { return appCtx.db }

func (appCtx *appContext) GetMemCached() *redis.Client { return appCtx.cache }

func (appCtx *appContext) GetSecret() string { return appCtx.secret }
