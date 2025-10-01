CREATE TABLE "user_goal_progress" (
    "user_goal_target_id" BIGINT NOT NULL,
    "contribution_id" BIGINT NOT NULL,

    CONSTRAINT "user_goal_progress_pkey"
        PRIMARY KEY ("user_goal_target_id", "contribution_id")
);

-- Foreign keys
ALTER TABLE "user_goal_progress"
    ADD CONSTRAINT "user_goal_progress_user_goal_target_id_fkey"
    FOREIGN KEY ("user_goal_target_id") REFERENCES "user_goal_target"("id") ON DELETE CASCADE;

ALTER TABLE "user_goal_progress"
    ADD CONSTRAINT "user_goal_progress_contribution_id_fkey"
    FOREIGN KEY ("contribution_id") REFERENCES "contributions"("id") ON DELETE CASCADE;
