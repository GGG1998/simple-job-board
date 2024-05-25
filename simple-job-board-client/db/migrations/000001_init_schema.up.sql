CREATE TYPE "level_job" AS ENUM (
  'junior',
  'medium',
  'senior'
);

CREATE TABLE "jobs" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "level" level_job[] NOT NULL,
  "location" varchar NOT NULL,
  "salary_min" varchar,
  "salary_max" varchar,
  "employer_id" int,
  "valid_date" timestamptz NOT NULL,
  "active" bool,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "employers" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "social_media_links" json,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "applicants" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "cv_link" varchar NOT NULL,
  "location" varchar,
  "availability" varchar NOT NULL,
  "monthly_salary_expectations" int NOT NULL,
  "linkedin_profile" varchar,
  "github_profile" varchar,
  "telegram_profile" varchar,
  "languages" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "applications" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "job_id" int,
  "applicant_id" int,
  "application_date" timestamptz NOT NULL DEFAULT (now()),
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "tags" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "categories" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "job_categories" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "job_id" int,
  "category_id" int
);

CREATE TABLE "applicant_tags" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "applicant_id" int,
  "tag_id" int
);

CREATE INDEX ON "applications" ("job_id");

CREATE INDEX ON "applications" ("applicant_id");

CREATE INDEX ON "applications" ("job_id", "applicant_id");

ALTER TABLE "jobs" ADD FOREIGN KEY ("employer_id") REFERENCES "employers" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("applicant_id") REFERENCES "applicants" ("id");

ALTER TABLE "job_categories" ADD FOREIGN KEY ("job_id") REFERENCES "jobs" ("id");

ALTER TABLE "job_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "applicant_tags" ADD FOREIGN KEY ("applicant_id") REFERENCES "applicants" ("id");

ALTER TABLE "applicant_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");
