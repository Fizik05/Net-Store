package letual

import (
	"context"
	"github.com/go-chi/render"
	"letual/internal/models"
	"log/slog"
	"net/http"
)

type getAllRequest struct{}

type getAllResponse struct {
	resp     Response
	Products []*models.Product `json:"products,omitempty"`
}

func (s *Server) getProducts(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "getProducts"

		log := s.logger.With(
			slog.String("fn", fn),
		)

		products, err := s.productModule.GetProducts(ctx)
		if err != nil {
			log.Error("failed to get products", err)

			w.WriteHeader(http.StatusInternalServerError)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		render.JSON(w, r, getAllResponse{
			resp:     OK(),
			Products: products,
		})
	}
}
