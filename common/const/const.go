package constanta

const (

	// Encyrption
	KEYPAIR_BITS int = 2048

	// Role
	ROLE_ALL     int8 = 0
	ROLE_REGULAR int8 = 1
	ROLE_ADMIN   int8 = 5

	// Context
	CLAIMS_CONTEXT_KEY string = "claims"

	// Transaction
	TRANSACTION_TYPE_DEPOSIT  int8 = 1
	TRANSACTION_TYPE_WITHDRAW int8 = 2
)
