-- name: GetTicket :one
SELECT * FROM tickets
WHERE id = ? LIMIT 1;

-- name: ListTickets :many
SELECT * FROM tickets
ORDER BY id;

-- name: SearchTickets :many
SELECT * FROM tickets
WHERE subject LIKE '%?%';

-- name: CreateTicket :execresult
INSERT INTO tickets (
    subject, description, status, updated_at, created_at, user_id, category_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);