package domain

import (
	"context"
	"time"
)

// Base
type Service[T any] interface {
	Get(ctx context.Context, id int) (T, error)
	Create(ctx context.Context, t *T) (T, error)
	Update(ctx context.Context, t *T) (T, error)
}

// JobApplications
type JobApplicationService interface {
	ListBy(ctx context.Context, jobID, applicantID int, applicationDate *time.Time, active string) ([]Application, error)
}

// Employers
type EmployerService interface {
	List(ctx context.Context, limit, offset int) ([]*Employer, error)
}

// Jobs
type JobService interface {
	Service[Job]
	ListBy(ctx context.Context, limit, offset int, active bool, title, location string, salaryMin, salaryMax int) ([]Job, error)
}
