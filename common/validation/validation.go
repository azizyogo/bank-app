package validation

import (
	"errors"
	"regexp"
)

func ValidatePrincipalAmount(amount float64) error {
	if amount < 0 {
		return errors.New("principal amount must be non-negative")
	}
	return nil
}

func ValidateInterestRate(rate float64) error {
	if rate < 0 || rate > 100 {
		return errors.New("interest rate must be between 0 and 100")
	}
	return nil
}

func ValidateTransactionAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("transaction amount must be non-negative")
	}
	return nil
}

func ValidateTimePeriod(period int) error {
	if period < 0 {
		return errors.New("time period must be non-negative")
	}
	return nil
}

func ValidateInput(input string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !re.MatchString(input) {
		return errors.New("invalid input format")
	}
	return nil
}

func ValidateAccountID(accountID int) error {
	if accountID <= 0 {
		return errors.New("invalid account ID")
	}
	return nil
}
