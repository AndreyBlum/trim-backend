CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    type TEXT CHECK (type IN ('income','expense')),
    amount NUMERIC(12,2) NOT NULL,
    category TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);