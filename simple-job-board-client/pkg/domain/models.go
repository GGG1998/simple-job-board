package domain

import (
	"time"
)

type Application struct {
	ID              int       `json:"id"`
	JobID           int       `json:"job_id"`
	ApplicantID     int       `json:"applicant_id"`
	ApplicationDate time.Time `json:"application_date"`
	Status          bool      `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Employer struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	SocialMediaLinks []byte    `json:"social_media_links"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Job struct {
	Active      bool      `json:"active"`
	ID          int       `json:"id"`
	SalaryMin   int       `json:"salary_min"`
	SalaryMax   int       `json:"salary_max"`
	EmployerID  int       `json:"employer_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Level       []string  `json:"level"`
	ValidDate   time.Time `json:"valid_date"`
	CreatedAt   time.Time `json:"created_at"`
}

type Applicant struct {
	ID                        int32     `json:"id"`
	FullName                  string    `json:"full_name"`
	Email                     string    `json:"email"`
	CvLink                    string    `json:"cv_link"`
	Location                  string    `json:"location"`
	Availability              string    `json:"availability"`
	MonthlySalaryExpectations int32     `json:"monthly_salary_expectations"`
	LinkedinProfile           string    `json:"linkedin_profile"`
	GithubProfile             string    `json:"github_profile"`
	TelegramProfile           string    `json:"telegram_profile"`
	Languages                 []string  `json:"languages"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}
