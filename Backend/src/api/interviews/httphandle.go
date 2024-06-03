package interviews

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Marcodave03/AceMyCareer/backend/src/api/utils"
)

type interviewHandler struct {
	db *sql.DB
}

func CreateInterviewHandler(db *sql.DB) *interviewHandler {
	return &interviewHandler{
		db: db,
	}
}

func (s *interviewHandler) HandleLevel(w http.ResponseWriter, r *http.Request) {
	var levelresponse interviewLevel
	if err := json.NewDecoder(r.Body).Decode(&levelresponse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(levelresponse)
	switch r.Method {
	case http.MethodGet:
		levels, err := getAllInterviewLevels(s.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, levels, http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
	case http.MethodPost:
		if err := insertLevel(s.db, levelresponse.Name); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		utils.WriteJson(w, "Inserted", http.StatusOK)

		break
	case http.MethodPatch:

		break
	case http.MethodDelete:
		if err := deleteLevel(s.db, levelresponse.ID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
	}
}
