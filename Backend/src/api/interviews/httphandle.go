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
	utils.EnableCors(&w)
	var levelRequest interviewLevel
	if err := json.NewDecoder(r.Body).Decode(&levelRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
		if err := insertLevel(s.db, levelRequest.Name); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, "Inserted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
	case http.MethodPatch:

		break
	case http.MethodDelete:
		if err := deleteLevel(s.db, levelRequest.ID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, "Deleted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
    default:

    break;
	}
}

func (s *interviewHandler) HandlePositions(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var positionRequest interviewPosition
	if err := json.NewDecoder(r.Body).Decode(&positionRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		positions, err := getAllInterviewPositions(s.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		utils.WriteJson(w, positions, http.StatusOK)
		break
	case http.MethodPost:
		err := insertPosition(s.db, positionRequest.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, "Inserted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
	case http.MethodDelete:
		err := deletePositions(s.db, positionRequest.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, "Deleted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
	default:
		utils.WriteJson(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		break
	}

}

func (s *interviewHandler) HandleIndustries(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var industryRequest interviewIndustry
	if err := json.NewDecoder(r.Body).Decode(&industryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		industries, err := getAllInterviewIndustries(s.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
        if err := utils.WriteJson(w, industries, http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		break
	case http.MethodPost:
        if err := insertIndustry(s.db, industryRequest.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		if err := utils.WriteJson(w, "Inserted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		break
	case http.MethodDelete:
        if err := deleteIndustries(s.db, industryRequest.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		if err := utils.WriteJson(w, "Deleted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		break
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    break;
	}
}
