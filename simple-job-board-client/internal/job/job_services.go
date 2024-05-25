package job

import (
	"context"
	"database/sql"
	db "simple-job-board/internal/sqlc"
	"simple-job-board/pkg/data"
	"simple-job-board/pkg/domain"
	"time"
)

type service struct {
	repo db.Queries
}

func NewJobService(repo db.Queries) domain.JobService {
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context, id int) (domain.Job, error) {
	offerDb, err := s.repo.GetJob(ctx, int32(id))
	if err != nil {
		return domain.Job{}, err
	}
	offer, err := data.RemapStructJSON[db.Job, domain.Job](&offerDb)
	if err != nil {
		return domain.Job{}, err
	}
	return *offer, nil
}

func (s *service) Create(ctx context.Context, job *domain.Job) (domain.Job, error) {
	var lvls []db.LevelJob
	for _, l := range job.Level {
		ns := &db.NullLevelJob{}
		ns.Scan(l)
		lvls = append(lvls, ns.LevelJob)
	}

	params := db.CreateJobParams{
		Name:        job.Name,
		Description: job.Description,
		Level:       lvls,
		SalaryMin:   sql.NullString{String: string(job.SalaryMax)},
		SalaryMax:   sql.NullString{String: string(job.SalaryMax)},
		EmployerID:  sql.NullInt32{Int32: int32(job.EmployerID)},
		Active:      sql.NullBool{Bool: job.Active},
		ValidDate:   time.Now(),
	}
	offerDb, err := s.repo.CreateJob(ctx, params)
	if err != nil {
		return domain.Job{}, err
	}

	offer, err := data.RemapStructJSON[db.Job, domain.Job](&offerDb)
	if err != nil {
		return domain.Job{}, err
	}

	return *offer, nil
}

func (s *service) Update(ctx context.Context, job *domain.Job) (domain.Job, error) {
	panic("implement me")
}

func (s *service) ListBy(ctx context.Context, limit, offset int, active bool, title, location string, salaryMin, salaryMax int) ([]domain.Job, error) {
	params := db.ListJobsByParams{
		Limit:     int32(limit),
		Offset:    int32(offset),
		Active:    active,
		Title:     title,
		Location:  location,
		SalaryMin: int32(salaryMin),
		SalaryMax: int32(salaryMax),
	}
	offersDb, err := s.repo.ListJobsBy(ctx, params)
	if err != nil {
		return nil, err
	}

	offers, err := data.RemapArrayStructJSON[db.Job, domain.Job](&offersDb)
	return *offers, nil
}
