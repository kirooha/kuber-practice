package main

import (
	"context"
	"log"

	"github.com/kirooha/kuber-practice/internal/app/handlers"
	"github.com/kirooha/kuber-practice/internal/pkg/dbmodel"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres dbname=kuber_practice sslmode=disable")
	if err != nil {
		log.Fatalf("pgx.Connect error - %v", err)
	}
	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("conn.Ping error - %v", err)
	}

	queries := dbmodel.New(conn)

	app := fiber.New()

	app.Get("/files", handlers.NewListHandler(queries).Handle)
	app.Post("/file", handlers.NewSaveHandler(queries).Handle)

	log.Fatal(app.Listen(":8080"))
}
