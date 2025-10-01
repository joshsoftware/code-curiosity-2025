ALTER TABLE users 
ADD COLUMN current_active_goal_id BIGINT DEFAULT NULL REFERENCES goal(id);
