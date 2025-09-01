CREATE INDEX idx_users_current_balance 
ON users(current_balance DESC)
WHERE is_admin = false AND is_deleted = false;
