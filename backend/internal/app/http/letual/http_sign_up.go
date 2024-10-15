package letual

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type signUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type signUpResponse struct {
	resp Response
}

func (s *Server) signUp(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "signUp"

		log := s.logger.With(
			slog.String("fn", fn),
		)

		var req signUpRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", "Error", err.Error())

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		err = s.userModule.Register(ctx, req.Name, req.Email, req.Password)
		if err != nil {
			log.Error("failed to register", err)

			w.WriteHeader(http.StatusBadRequest)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		render.JSON(w, r, signUpResponse{
			resp: OK(),
		})
	}
}
