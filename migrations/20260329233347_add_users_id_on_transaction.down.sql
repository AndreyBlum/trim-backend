ALTER TABLE transactions
DROP CONSTRAINT IF EXISTS fk_transactions_user,
DROP COLUMN IF EXISTS user_id;