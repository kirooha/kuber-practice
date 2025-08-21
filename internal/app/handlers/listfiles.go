package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kirooha/kuber-practice/internal/pkg/dbmodel"

	"github.com/gofiber/fiber/v2"
)

type ListHandler struct {
	queries *dbmodel.Queries
}

func NewListHandler(queries *dbmodel.Queries) *ListHandler {
	return &ListHandler{
		queries: queries,
	}
}

func (h *ListHandler) Handle(fiberCtx *fiber.Ctx) error {
	var (
		ctx       = fiberCtx.Context()
		msgPrefix = "app.handler.ListHandler.Handle"
	)

	files, err := h.queries.ListFiles(ctx)
	if err != nil {
		log.Printf("%s: h.queries.ListFiles error - %v", msgPrefix, err)
		return fiberCtx.SendStatus(http.StatusInternalServerError)
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name)
	}

	return json.NewEncoder(fiberCtx.Response().BodyWriter()).Encode(filenames)
}
