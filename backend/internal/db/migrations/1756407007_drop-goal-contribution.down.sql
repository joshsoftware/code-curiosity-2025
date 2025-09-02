CREATE TABLE "goal_contribution"(
    "id" SERIAL PRIMARY KEY,
    "goal_id" BIGINT NOT NULL,
    "contribution_score_id" BIGINT NOT NULL,
    "target_count" BIGINT NOT NULL,
    "is_custom" BOOLEAN NOT NULL,
    "set_by_user_id" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE
    "goal_contribution" ADD CONSTRAINT "goal_contribution_set_by_user_id_foreign" FOREIGN KEY("set_by_user_id") REFERENCES "users"("id") ON DELETE CASCADE;
ALTER TABLE
    "goal_contribution" ADD CONSTRAINT "goal_contribution_contribution_score_id_foreign" FOREIGN KEY("contribution_score_id") REFERENCES "contribution_score"("id");
ALTER TABLE
    "goal_contribution" ADD CONSTRAINT "goal_contribution_goal_id_foreign" FOREIGN KEY("goal_id") REFERENCES "goal"("id");