package account

const (
	QueryGetAccount string = `
	SELECT 
		acc.id AS "id",
		COALESCE(acc.user_id, 0) AS "user_id",
		COALESCE(acc.account_number, 0) AS "account_number",
		COALESCE(acc.balance, 0) AS "balance"
	FROM accounts as acc
	`

	QueryGetAccountByUserID string = QueryGetAccount + ` WHERE acc.user_id = $1;`

	QueryUpdateAccountBalance string = `
	UPDATE accounts
	SET balance = $1
	WHERE user_id = $2;
	`
)
