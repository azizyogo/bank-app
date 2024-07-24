package user

const (
	QueryGetUser string = `
	SELECT 
		u.id AS "id",
		COALESCE(u.username, '') AS "username",
		COALESCE(u.password, '') AS "password",
		COALESCE(u.role, -1) AS "role",
		COALESCE(u.public_key, '') AS "public_key",
		COALESCE(u.private_key, '') AS "private_key",
		COALESCE(u.created_at, '0001-01-01T00:00:00Z'::timestamp) AS "created_at"
	FROM users AS u
	`
	QueryGetUserByUsername string = QueryGetUser + ` WHERE u.username = $1;`

	QueryGetUserByUserID string = QueryGetUser + ` WHERE u.id = $1;`

	QueryGetUserKey string = `
	SELECT 
		u.id AS "id",
		COALESCE(u.public_key, '') AS "public_key",
		COALESCE(u.private_key, '') AS "private_key"
	FROM users AS u
	`

	QueryGetUserKeyByUserID string = QueryGetUserKey + ` WHERE u.id = $1`

	QueryUpdateKeyByUserID string = `
	UPDATE users
	SET public_key = $1,
    	private_key = $2
	WHERE id = $3;
	`
)
