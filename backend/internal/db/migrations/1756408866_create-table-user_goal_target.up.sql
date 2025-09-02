CREATE TABLE "user_goal_target" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_goal_id" BIGINT NOT NULL,
    "contribution_score_id" BIGINT NOT NULL,
    "target" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT "user_goal_target_unique_user_goal_score"
        UNIQUE ("user_goal_id", "contribution_score_id")
);

-- Foreign keys
ALTER TABLE "user_goal_target"
    ADD CONSTRAINT "user_goal_target_user_goal_id_fkey"
    FOREIGN KEY ("user_goal_id") REFERENCES "user_goal"("id") ON DELETE CASCADE;

ALTER TABLE "user_goal_target"
    ADD CONSTRAINT "user_goal_target_contribution_score_id_fkey"
    FOREIGN KEY ("contribution_score_id") REFERENCES "contribution_score"("id");
