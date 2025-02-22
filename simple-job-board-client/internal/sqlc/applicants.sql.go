// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: applicants.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createApplicant = `-- name: CreateApplicant :one
INSERT INTO applicants (
    full_name, 
    email, 
    cv_link, 
    location, 
    availability, 
    monthly_salary_expectations,
    linkedin_profile, 
    github_profile, 
    telegram_profile, 
    languages
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING id
`

type CreateApplicantParams struct {
	FullName                  string         `json:"full_name"`
	Email                     string         `json:"email"`
	CvLink                    string         `json:"cv_link"`
	Location                  sql.NullString `json:"location"`
	Availability              string         `json:"availability"`
	MonthlySalaryExpectations int32          `json:"monthly_salary_expectations"`
	LinkedinProfile           sql.NullString `json:"linkedin_profile"`
	GithubProfile             sql.NullString `json:"github_profile"`
	TelegramProfile           sql.NullString `json:"telegram_profile"`
	Languages                 []string       `json:"languages"`
}

func (q *Queries) CreateApplicant(ctx context.Context, arg CreateApplicantParams) (int32, error) {
	row := q.queryRow(ctx, q.createApplicantStmt, createApplicant,
		arg.FullName,
		arg.Email,
		arg.CvLink,
		arg.Location,
		arg.Availability,
		arg.MonthlySalaryExpectations,
		arg.LinkedinProfile,
		arg.GithubProfile,
		arg.TelegramProfile,
		pq.Array(arg.Languages),
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
