-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id VARCHAR(100) PRIMARY KEY,
    payer_vpa VARCHAR(255) NOT NULL,
    payee_vpa VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'INITIATED',
    bank_ref_id VARCHAR(100),
    failure_reason TEXT,
    remarks VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_psp_txn_payer_vpa ON transactions(payer_vpa);

-- +goose Down
DROP TABLE IF EXISTS transactions;