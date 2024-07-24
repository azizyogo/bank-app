package transaction

const (
	QueryGetTransaction string = ``

	QueryInsertNewTransaction string = `
	INSERT INTO transactions (user_id, amount, "type")
	VALUES ($1, $2, $3);
	`
)
