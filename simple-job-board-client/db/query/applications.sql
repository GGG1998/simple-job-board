-- name: ListJobApplications :many
-- SELECT sqlc.embed(jobs.*), sqlc.embed(applicants.*)
--     FROM applications
--         JOIN jobs ON jobs.id = applications.job_id
--         JOIN applicants ON applicants.id = applications.applicant_id
-- WHERE (applications.job_id = @job_id::int OR applications.applicant_id = @applicant_id::int) AND applications.active;

-- name: CreateJobApplication :one
INSERT INTO applications(
    job_id, 
    applicant_id,
    application_date,
    status
) VALUES($1, $2, $3, $4) RETURNING *;