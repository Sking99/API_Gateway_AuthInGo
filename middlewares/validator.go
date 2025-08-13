package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var payload dto.UserLoginDTO

		if jsonErr := utils.ReadJsonRequest(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", jsonErr)
			return
		}
		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", validationErr)
			return
		}

		ctx := context.WithValue(r.Context(), "userLoginPayload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserRegisterRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var payload dto.UserRegisterDTO

		if jsonErr := utils.ReadJsonRequest(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", jsonErr)
			return
		}
		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", validationErr)
			return
		}

		ctx := context.WithValue(r.Context(), "useRegisterPayload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
