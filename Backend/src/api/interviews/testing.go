package interviews

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Marcodave03/AceMyCareer/backend/src/api/utils"
)

func (s *interviewHandler) HandleIndustriesWithoutRepo(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var industryRequest interviewIndustry
	if err := json.NewDecoder(r.Body).Decode(&industryRequest); err != nil && err.Error() != "EOF" {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		// GET {{{
		query := `
        SELECT * FROM interviews.interview_industries;
        `
		rows, err := s.db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer rows.Close()

		var industries []interviewIndustry
		for rows.Next() {
			var curIndustrie interviewIndustry
			err := rows.Scan(&curIndustrie.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			industries = append(industries, curIndustrie)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, industries, http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} // }}}
		break
	case http.MethodPost:
		// POST {{{
		query := ` INSERT INTO interviews.interview_industries (name) VALUES ($1);`
		if _, err := s.db.Exec(query, industryRequest.Name); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, "Inserted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// }}}
		break
	case http.MethodDelete:
        // Delete {{{
		query := `DELETE FROM interviews.interview_industries WHERE name = $1`
		affected, err := s.db.Exec(query, industryRequest.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		total, err := affected.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if total == 0 {
			http.Error(w, "Not Found", http.StatusBadRequest)
			return
		}
		if err := utils.WriteJson(w, "Deleted", http.StatusOK); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}// }}}
		break
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
