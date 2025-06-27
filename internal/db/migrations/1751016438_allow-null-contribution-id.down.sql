UPDATE transactions
SET contribution_id = 0 
WHERE contribution_id IS NULL;

ALTER TABLE transactions
ALTER COLUMN contribution_id SET NOT NULL;