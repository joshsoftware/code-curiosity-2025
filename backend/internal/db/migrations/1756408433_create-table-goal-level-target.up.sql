CREATE TABLE "goal_level_target" (
    "id" BIGSERIAL PRIMARY KEY,
    "goal_level_id" BIGINT NOT NULL,
    "contribution_score_id" BIGINT NOT NULL,
    "target" BIGINT NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Foreign keys
ALTER TABLE "goal_level_target"
    ADD CONSTRAINT "goal_level_target_goal_level_id_fkey"
    FOREIGN KEY ("goal_level_id") REFERENCES "goal_level"("id") ON DELETE CASCADE;

ALTER TABLE "goal_level_target"
    ADD CONSTRAINT "goal_level_target_contribution_score_id_fkey"
    FOREIGN KEY ("contribution_score_id") REFERENCES "contribution_score"("id");
