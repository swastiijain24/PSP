-- name: CreateTransaction :one
INSERT INTO transactions (
    transaction_id,
    payer_vpa,
    payee_vpa,
    amount,
    remarks
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTransactionHistory :many
SELECT 
    transaction_id,
    payer_vpa,
    payee_vpa,
    amount,
    status,
    bank_ref_id,
    failure_reason,
    remarks,
    created_at,
    updated_at
FROM transactions
WHERE payer_vpa = $1 OR payee_vpa = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;