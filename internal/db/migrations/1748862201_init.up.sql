CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "github_id" BIGINT NOT NULL UNIQUE,
    "github_username" VARCHAR(255) NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NULL,
    "current_active_goal_id" BIGINT NULL,
    "current_balance" BIGINT DEFAULT 0,
    "is_blocked" BOOLEAN DEFAULT FALSE,
    "is_admin" BOOLEAN DEFAULT FALSE,
    "password" VARCHAR(255) DEFAULT '',
    "is_deleted" BOOLEAN DEFAULT FALSE,
    "deleted_at" BIGINT,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

CREATE TABLE "leaderboard_hourly"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "github_id" BIGINT NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "current_balance" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "refreshed_at" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL
);

CREATE TABLE "contributions"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "repository_id" BIGINT NOT NULL,
    "contribution_score_id" BIGINT NOT NULL,
    "contribution_type" VARCHAR(255) NOT NULL,
    "balance_change" BIGINT NOT NULL,
    "contributed_at" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

CREATE TABLE "repositories"(
    "id" SERIAL PRIMARY KEY,
    "github_repo_id" BIGINT NOT NULL,
    "repo_name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "languages_url" VARCHAR(255) NOT NULL,
    "repo_url" VARCHAR(255) NOT NULL,
    "owner_name" VARCHAR(255) NOT NULL,
    "update_date" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

CREATE TABLE "badges"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "badge_type" VARCHAR(255) NOT NULL,
    "earned_at" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL
);

CREATE TABLE "transactions"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "contribution_id" BIGINT NOT NULL,
    "is_redeemed" BOOLEAN NOT NULL,
    "is_gained" BOOLEAN NOT NULL,
    "transacted_balance" BIGINT NOT NULL,
    "transacted_at" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL
);

CREATE TABLE "summary"(
    "id" SERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "month_year" BIGINT NOT NULL,
    "net_balance" BIGINT NOT NULL,
    "badges_count" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "contribution_id" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

CREATE TABLE "contribution_score"(
    "id" SERIAL PRIMARY KEY,
    "admin_id" BIGINT NOT NULL,
    "contribution_type" VARCHAR(255) NOT NULL,
    "score" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

CREATE TABLE "goal"(
    "id" SERIAL PRIMARY KEY,
    "level" VARCHAR(255) NOT NULL,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

CREATE TABLE "goal_contribution"(
    "id" SERIAL PRIMARY KEY,
    "goal_id" BIGINT NOT NULL,
    "contribution_score_id" BIGINT NOT NULL,
    "target_count" BIGINT NOT NULL,
    "is_custom" BOOLEAN NOT NULL,
    "set_by_user_id" BIGINT NOT NULL,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);

ALTER TABLE
    "goal_contribution" ADD CONSTRAINT "goal_contribution_set_by_user_id_foreign" FOREIGN KEY("set_by_user_id") REFERENCES "users"("id");
ALTER TABLE
    "goal_contribution" ADD CONSTRAINT "goal_contribution_contribution_score_id_foreign" FOREIGN KEY("contribution_score_id") REFERENCES "contribution_score"("id");
ALTER TABLE
    "contribution_score" ADD CONSTRAINT "contribution_score_admin_id_foreign" FOREIGN KEY("admin_id") REFERENCES "users"("id");
ALTER TABLE
    "summary" ADD CONSTRAINT "summary_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "contributions" ADD CONSTRAINT "contributions_contribution_score_id_foreign" FOREIGN KEY("contribution_score_id") REFERENCES "contribution_score"("id");
ALTER TABLE
    "badges" ADD CONSTRAINT "badges_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "goal_contribution" ADD CONSTRAINT "goal_contribution_goal_id_foreign" FOREIGN KEY("goal_id") REFERENCES "goal"("id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_contribution_id_foreign" FOREIGN KEY("contribution_id") REFERENCES "contributions"("id");
ALTER TABLE
    "contributions" ADD CONSTRAINT "contributions_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "contributions" ADD CONSTRAINT "contributions_repository_id_foreign" FOREIGN KEY("repository_id") REFERENCES "repositories"("id");
ALTER TABLE
    "leaderboard_hourly" ADD CONSTRAINT "leaderboard_hourly_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "summary" ADD CONSTRAINT "summary_contribution_id_foreign" FOREIGN KEY("contribution_id") REFERENCES "contributions"("id");