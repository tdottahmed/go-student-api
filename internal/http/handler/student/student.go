package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/tdottahmed/students-api/internal/types"
	"github.com/tdottahmed/students-api/internal/utils/response"
)

func New() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating new student")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenerateError(fmt.Errorf("empty request body")))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	})
}
