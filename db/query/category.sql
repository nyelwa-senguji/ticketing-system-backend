-- name: GetCategory :one
SELECT * FROM category
WHERE id = ? LIMIT 1;

-- name: GetCategoryByName :one
SELECT * FROM category
WHERE category_name=? LIMIT 1;

-- name: ListCategories :many
SELECT * FROM category
ORDER BY category_name;

-- name: CreateCategory :execresult
INSERT INTO category (
  category_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
);

-- name: UpdateCategory :exec
UPDATE category
SET category_name=?, status=?, updated_at=?
WHERE id=?