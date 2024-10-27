package students

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/BibhabenduMukherjee/student-api/internal/types"
	"github.com/BibhabenduMukherjee/student-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			// json response
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}
		slog.Info("creating a student account")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
