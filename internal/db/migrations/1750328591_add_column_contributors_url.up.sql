ALTER TABLE repositories ADD COLUMN contributors_url VARCHAR(255);

UPDATE repositories SET contributors_url = '' WHERE contributors_url IS NULL;

ALTER TABLE repositories ALTER COLUMN contributors_url SET NOT NULL;
