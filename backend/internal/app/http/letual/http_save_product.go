package letual

import (
	"context"
	"log/slog"
	"net/http"

	"letual/internal/models"

	"github.com/go-chi/render"
)

type saveRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price" validate:"required,min=0"`
	ImageUrl    string `json:"image-url" validate:"required"`
}

type saveResponse struct {
	resp Response
}

func (s *Server) saveProduct(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "saveProduct"

		log := s.logger.With(
			slog.String("fn", fn),
		)

		var req saveRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", "Error", err.Error())

			w.WriteHeader(http.StatusBadRequest)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		err = s.productModule.SaveProduct(ctx, &models.Product{
			Title:       req.Title,
			Description: req.Description,
			Price:       req.Price,
			ImageUrl:    req.ImageUrl,
		})
		if err != nil {
			log.Error("failed to save product", err)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		render.JSON(w, r, saveResponse{
			resp: OK(),
		})
	}
}
