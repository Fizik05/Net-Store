package letual

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strconv"

	"letual/internal/models"

	"github.com/go-chi/render"
)

type getRequest struct {
}

type getResponse struct {
	resp    Response
	Product *models.Product `json:"product,omitempty"`
}

func (s *Server) getProduct(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "getProduct"

		log := s.logger.With(
			slog.String("fn", fn),
		)

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to decode the url", "Error", err.Error())

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		product, err := s.productModule.GetProduct(ctx, id)
		if err != nil {
			log.Error("failed to get product", err)

			w.WriteHeader(http.StatusInternalServerError)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		render.JSON(w, r, getResponse{
			resp:    OK(),
			Product: product,
		})
	}
}
