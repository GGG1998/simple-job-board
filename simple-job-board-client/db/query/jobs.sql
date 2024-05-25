-- name: CreateJob :one
INSERT INTO jobs(
    name,
    description,
    level,
    salary_min,
    salary_max,
    employer_id,
    valid_date,
    active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetJob :one
SELECT * FROM jobs WHERE id = $1 LIMIT 1;

-- name: ListJobs :many
SELECT * FROM jobs WHERE active = true ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: ListJobsBy :many
SELECT * FROM jobs 
    WHERE active = @active::boolean
    AND (title = COALESCE(@title::text, title)) 
    AND (location = COALESCE(@location::text, location)) 
    AND (salary_min = COALESCE(@salary_min::int, salary_min)) 
    AND (salary_max = COALESCE(@salary_max::int, salary_max)) 
ORDER BY created_at LIMIT $1 OFFSET $2;
