package job

import (
	http2 "net/http"
	"simple-job-board/internal/http"
	"simple-job-board/pkg/data"
)

// Create Job Body Params
type createJobParams struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Level       []string `json:"level" validate:"required"`
	SalaryMin   int      `json:"salary_min"`
	SalaryMax   int      `json:"salary_max"`
	EmployerID  int      `json:"employer_id" validate:"required"`
}

func CreateJobBody(r *http2.Request) (createJobParams, error) {
	var p createJobParams

	err := http.ReadJSON(r.Body, &p)
	if err != nil {
		return p, err
	}

	err = data.Valid(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}
