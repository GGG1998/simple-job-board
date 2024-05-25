-- -- name: CreateJobCategory :one
-- INSERT INTO job_categories(job_id, category_id) VALUES($1, $2) RETURNING id;

-- -- name: ListJobCategory :many
-- SELECT sqlc.embed(jobs.*) as jobs, array_agg(sqlc.embed(categories.*)) as categories
--     FROM job_categories jc
--         JOIN jobs ON jobs.id = jc.job_id
--         JOIN categories ON categories.id = jc.categories_id 
-- WHERE jc.job_id = $1 GROUP BY jc.job_id;