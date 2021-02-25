package repository

const (
	createUserQuery   = `INSERT INTO users(username, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
	findUserByIDQuery = `SELECT user_id, username, password, email, created_at, updated_at FROM users WHERE user_id = $1`
	updateUserQuery   = `UPDATE users SET username = $1, email = $2, updated_at = $3 WHERE user_id = $4`
	destroyUserQuery  = `DELETE FROM users WHERE user_id = $1`
)
