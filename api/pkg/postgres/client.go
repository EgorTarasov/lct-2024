package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDb создает новое подключение к базе данных
func NewDb(ctx context.Context, cfg *Config) (*Database, error) {
	pool, err := pgxpool.New(ctx, createDsn(cfg))
	if err != nil {
		return nil, err
	}
	// проверка соединения
	// после создания pool и некорректными параметрами соединения (выключенной бд)
	// pgx.pool возвращает err = nil
	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}
	return newDatabase(pool), nil
}
