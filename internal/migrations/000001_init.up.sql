CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "github_id" BIGINT NOT NULL UNIQUE,
    "github_username" VARCHAR(255) NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255),
    "current_balance" BIGINT DEFAULT 0,
    "is_blocked" BOOLEAN DEFAULT FALSE,
    "is_admin" BOOLEAN DEFAULT FALSE,
    "password" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "leaderboard_hourly" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "github_id" BIGINT NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "current_balance" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "refreshed_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "repositories" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "repo_name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "languages" JSON NOT NULL,
    "contributor_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "owner_name" VARCHAR(255) NOT NULL,
    "update_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "contributions" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "repository_id" INTEGER NOT NULL REFERENCES "repositories"("id"),
    "contribution_type" VARCHAR(255) NOT NULL,
    "balance_change" BIGINT NOT NULL,
    "contributed_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "goals" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "level_name" VARCHAR(255) NOT NULL,
    "is_completed" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "badges" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "badge_type" VARCHAR(255) NOT NULL,
    "earned_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "transactions" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "contribution_id" INTEGER NOT NULL REFERENCES "contributions"("id"),
    "transaction_type" VARCHAR(255) NOT NULL,
    "transacted_balance" BIGINT NOT NULL,
    "transacted_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "summary" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "month_year" BIGINT NOT NULL,
    "net_balance" BIGINT NOT NULL,
    "badges_count" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "contribution_id" INTEGER NOT NULL REFERENCES "contributions"("id"),
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "goal_criteria" (
    "id" SERIAL PRIMARY KEY,
    "contribution_id" INTEGER NOT NULL REFERENCES "contributions"("id"),
    "goal_type" VARCHAR(255) NOT NULL,
    "progress_count" BIGINT NOT NULL,
    "target_count" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "score_configuration" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id"),
    "field1_score" BIGINT NOT NULL,
    "field2_score" BIGINT NOT NULL,
    "field3_score" BIGINT NOT NULL,
    "field4_score" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
