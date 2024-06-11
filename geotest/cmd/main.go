package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://Dino:Dino-misos2024@192.168.1.70:54000dev?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = pool.Ping(ctx); err != nil {
		panic(err)
	}
	P
	pool
}
