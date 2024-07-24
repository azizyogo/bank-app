package interestrate

const (
	QueryGetInterestRate string = `
	SELECT
		ir.id AS "id",
		COALESCE(ir.rate, 0) AS "rate",
		COALESCE(ir.compound, 0) AS "compound",
		COALESCE(ir.effective_date, '0001-01-01T00:00:00Z'::timestamp) AS "effective_date"
	FROM interest_rates as ir
	`

	QueryUpdateInterestRate string = `
	UPDATE interest_rates
	SET rate = $1,
    	compound = $2,
    	effective_date = $3
	WHERE id = 1;
	`
)
