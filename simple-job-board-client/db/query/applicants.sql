-- name: CreateApplicant :one
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
) RETURNING id;