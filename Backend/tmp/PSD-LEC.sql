CREATE TABLE "users" (
  "username" varchar(30) PRIMARY KEY,
  "firstname" varchar(30),
  "lastname" varchar(30),
  "email" varchar(30),
  "profile_picture_url" varchar(50),
  "score_technical" int,
  "score_leadership" int,
  "score_teamwork" int,
  "score_organization" int
);

CREATE TABLE "interview_histories" (
  "id" serial PRIMARY KEY,
  "user" varchar(30),
  "score" int NOT NULL,
  "interview_id" serial,
  "liked_interview" bool DEFAULT false
);

CREATE TABLE "interviews" (
  "id" serial PRIMARY KEY,
  "Title" varchar(30),
  "image_url" varchar(50),
  "industry_name" varchar(30),
  "position_name" varchar(30),
  "level_name" varchar(30),
  "experience" text,
  "likes" int DEFAULT 0
);

CREATE TABLE "interview_tags" (
  "id" serial PRIMARY KEY,
  "interview_id" serial,
  "tag_id" serial
);

CREATE TABLE "tags" (
  "id" serial PRIMARY KEY,
  "name" varchar(30)
);

CREATE TABLE "interview_requirements" (
  "id" serial PRIMARY KEY,
  "requirement_name" varchar(30),
  "interview_id" serial
);

CREATE TABLE "interview_levels" (
  "id" serial PRIMARY KEY,
  "name" varchar(30) NOT NULL
);

CREATE TABLE "interview_industries" (
  "name" varchar(30)
);

CREATE TABLE "interview_positions" (
  "name" varchar(30)
);

CREATE TABLE "interview_technical_requirements" (
  "name" varchar(30)
);

ALTER TABLE "interview_histories" ADD FOREIGN KEY ("user") REFERENCES "users" ("username");

ALTER TABLE "interview_histories" ADD FOREIGN KEY ("interview_id") REFERENCES "interviews" ("id");

ALTER TABLE "interviews" ADD FOREIGN KEY ("industry_name") REFERENCES "interview_industries" ("name");

ALTER TABLE "interviews" ADD FOREIGN KEY ("position_name") REFERENCES "interview_positions" ("name");

ALTER TABLE "interviews" ADD FOREIGN KEY ("level_name") REFERENCES "interview_levels" ("id");

ALTER TABLE "interview_tags" ADD FOREIGN KEY ("interview_id") REFERENCES "interviews" ("id");

ALTER TABLE "interview_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "interview_requirements" ADD FOREIGN KEY ("requirement_name") REFERENCES "interview_technical_requirements" ("name");

ALTER TABLE "interview_requirements" ADD FOREIGN KEY ("interview_id") REFERENCES "interviews" ("id");
