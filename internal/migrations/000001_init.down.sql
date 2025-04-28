-- Drop foreign key constraints first
ALTER TABLE "summary" DROP CONSTRAINT IF EXISTS "summary_contribution_id_foreign";
ALTER TABLE "repositories" DROP CONSTRAINT IF EXISTS "repositories_user_id_foreign";
ALTER TABLE "goals" DROP CONSTRAINT IF EXISTS "goals_user_id_foreign";
ALTER TABLE "contributions" DROP CONSTRAINT IF EXISTS "contributions_user_id_foreign";
ALTER TABLE "badges" DROP CONSTRAINT IF EXISTS "badges_user_id_foreign";
ALTER TABLE "goals" DROP CONSTRAINT IF EXISTS "goals_goal_id_foreign";
ALTER TABLE "summary" DROP CONSTRAINT IF EXISTS "summary_user_id_foreign";
ALTER TABLE "repositories" DROP CONSTRAINT IF EXISTS "repositories_contributor_id_foreign";
ALTER TABLE "goal_criteria" DROP CONSTRAINT IF EXISTS "goal_criteria_contribution_id_foreign";
ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_user_id_foreign";
ALTER TABLE "contributions" DROP CONSTRAINT IF EXISTS "contributions_repository_id_foreign";
ALTER TABLE "score_configuration" DROP CONSTRAINT IF EXISTS "score_configuration_user_id_foreign";
ALTER TABLE "transactions" DROP CONSTRAINT IF EXISTS "transactions_contribution_id_foreign";

-- Drop tables
DROP TABLE IF EXISTS "score_configuration";
DROP TABLE IF EXISTS "goal_criteria";
DROP TABLE IF EXISTS "summary";
DROP TABLE IF EXISTS "transactions";
DROP TABLE IF EXISTS "badges";
DROP TABLE IF EXISTS "repositories";
DROP TABLE IF EXISTS "goals";
DROP TABLE IF EXISTS "contributions";
DROP TABLE IF EXISTS "leaderboard_hourly";
DROP TABLE IF EXISTS "users";
