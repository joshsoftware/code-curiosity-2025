
CREATE TABLE "users"(
    "user_id" SERIAL,
    "github_id" BIGINT NOT NULL,
    "github_username" VARCHAR(255) NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) ,
    "current_balance" BIGINT DEFAULT 0,
    "is_blocked" BOOLEAN DEFAULT FALSE,
    "is_admin" BOOLEAN DEFAULT FALSE,
    "password" VARCHAR(255) DEFAULT '',
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("user_id");
CREATE TABLE "leaderboard_hourly"(
    "leaderboard_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "github_id" BIGINT NOT NULL,
    "avatar_url" VARCHAR(255) NOT NULL,
    "current_balance" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "refreshed_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "leaderboard_hourly" ADD PRIMARY KEY("leaderboard_id");
CREATE TABLE "contributions"(
    "contribution_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "repository_id" BIGINT NOT NULL,
    "contribution_type" VARCHAR(255) NOT NULL,
    "balance_change" BIGINT NOT NULL,
    "contributed_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "contributions" ADD PRIMARY KEY("contribution_id");
CREATE TABLE "goals"(
    "goal_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "level_name" VARCHAR(255) NOT NULL,
    "is_completed" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "goals" ADD PRIMARY KEY("goal_id");
CREATE TABLE "repositories"(
    "repository_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "repo_name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "languages" JSON NOT NULL,
    "contributor_id" BIGINT NOT NULL,
    "owner_name" VARCHAR(255) NOT NULL,
    "update_date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "repositories" ADD PRIMARY KEY("repository_id");
CREATE TABLE "badges"(
    "badge_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "badge_type" VARCHAR(255) NOT NULL,
    "earned_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "badges" ADD PRIMARY KEY("badge_id");
CREATE TABLE "transactions"(
    "transaction_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "contribution_id" BIGINT NOT NULL,
    "transaction_type" VARCHAR(255) NOT NULL,
    "transacted_balance" BIGINT NOT NULL,
    "transacted_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "transactions" ADD PRIMARY KEY("transaction_id");
CREATE TABLE "summary"(
    "summary_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "month_year" BIGINT NOT NULL,
    "net_balance" BIGINT NOT NULL,
    "badges_count" BIGINT NOT NULL,
    "rank" BIGINT NOT NULL,
    "contribution_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "summary" ADD PRIMARY KEY("summary_id");
CREATE TABLE "goal_criteria"(
    "goal_id" BIGINT NOT NULL,
    "contribution_id" BIGINT NOT NULL,
    "goal_type" VARCHAR(255) NOT NULL,
    "progress_count" BIGINT NOT NULL,
    "target_count" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "goal_criteria" ADD PRIMARY KEY("goal_id");
CREATE TABLE "score_configuration"(
    "configuration_id" BIGINT NOT NULL,
    "user_id" BIGINT NOT NULL,
    "field1_score" BIGINT NOT NULL,
    "field2_score" BIGINT NOT NULL,
    "field3_score" BIGINT NOT NULL,
    "field4_score" BIGINT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "score_configuration" ADD PRIMARY KEY("configuration_id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_contribution_id_foreign" FOREIGN KEY("contribution_id") REFERENCES "contributions"("contribution_id");
ALTER TABLE
    "score_configuration" ADD CONSTRAINT "score_configuration_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "contributions" ADD CONSTRAINT "contributions_repository_id_foreign" FOREIGN KEY("repository_id") REFERENCES "repositories"("repository_id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "goal_criteria" ADD CONSTRAINT "goal_criteria_contribution_id_foreign" FOREIGN KEY("contribution_id") REFERENCES "contributions"("contribution_id");
ALTER TABLE
    "repositories" ADD CONSTRAINT "repositories_contributor_id_foreign" FOREIGN KEY("contributor_id") REFERENCES "users"("user_id");
ALTER TABLE
    "summary" ADD CONSTRAINT "summary_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "goals" ADD CONSTRAINT "goals_goal_id_foreign" FOREIGN KEY("goal_id") REFERENCES "goal_criteria"("goal_id");
ALTER TABLE
    "badges" ADD CONSTRAINT "badges_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "contributions" ADD CONSTRAINT "contributions_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "goals" ADD CONSTRAINT "goals_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "repositories" ADD CONSTRAINT "repositories_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("user_id");
ALTER TABLE
    "summary" ADD CONSTRAINT "summary_contribution_id_foreign" FOREIGN KEY("contribution_id") REFERENCES "contributions"("contribution_id");