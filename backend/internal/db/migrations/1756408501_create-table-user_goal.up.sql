CREATE TABLE "user_goal" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "goal_level_id" BIGINT,
    "status" VARCHAR NOT NULL,
    "month_started_at" TIMESTAMPTZ NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Foreign keys
ALTER TABLE "user_goal"
    ADD CONSTRAINT "user_goal_user_id_fkey"
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE;

ALTER TABLE "user_goal"
    ADD CONSTRAINT "user_goal_goal_level_id_fkey"
    FOREIGN KEY ("goal_level_id") REFERENCES "goal_level"("id");
