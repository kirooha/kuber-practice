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
	apiKey  string
}

func NewListHandler(queries *dbmodel.Queries, apiKey string) *ListHandler {
	return &ListHandler{
		queries: queries,
		apiKey:  apiKey,
	}
}

func (h *ListHandler) Handle(fiberCtx *fiber.Ctx) error {
	var (
		ctx          = fiberCtx.Context()
		msgPrefix    = "app.handler.ListHandler.Handle"
		headerApiKey = fiberCtx.Get("Authorization")
	)

	if h.apiKey != headerApiKey {
		return fiberCtx.SendStatus(http.StatusForbidden)
	}

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
