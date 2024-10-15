package letual

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type signInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type signInResponse struct {
	resp  Response
	token string
}

func (s *Server) signIn(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "signIn"

		log := s.logger.With(
			slog.String("fn", fn),
		)

		var req signInRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", "Error", err.Error())

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		token, err := s.userModule.Login(ctx, req.Email, req.Password)
		if err != nil {
			log.Error("failed to sign in", err)

			w.WriteHeader(http.StatusBadRequest)

			render.JSON(w, r, Error("failed to decode request"))

			return
		}

		render.JSON(w, r, signInResponse{
			resp:  OK(),
			token: token,
		})
	}
}
