package account

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/azizyogo/bank-app/common/auth"
	constanta "github.com/azizyogo/bank-app/common/const"
	"github.com/azizyogo/bank-app/common/encryption"
	"github.com/azizyogo/bank-app/common/model"
	"github.com/azizyogo/bank-app/common/util/float"
	"github.com/azizyogo/bank-app/common/util/response"
	accDmn "github.com/azizyogo/bank-app/domain/account"
	"github.com/azizyogo/bank-app/domain/transaction"

	"errors"
)

func (account *AccountUsecase) GetBalance(ctx context.Context, userID int) (model.EncryptedRequestResponse, error) {

	resDB, err := account.acc.GetAccountByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Get user public and private key for encryption purpose
	userKey, err := account.u.GetUserByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Encrypt response
	encryptRes, err := response.EncryptResponse(userKey, GetBalanceResponse{
		UserID:  userID,
		Balance: resDB.Balance,
	})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	return model.EncryptedRequestResponse{
		Data: encryptRes,
	}, nil
}

func (account *AccountUsecase) GetEncryptionKey(ctx context.Context, userID int) (GetEncryptionKeyResponse, error) {

	res := GetEncryptionKeyResponse{}
	resDB, err := account.u.GetKeyByUserID(ctx, userID)
	if err != nil {
		return GetEncryptionKeyResponse{}, err
	}

	// Convert RSA key to string
	encryptionKey, err := encryption.EncodePublicKeyToPEM(account.cfg.RSAKeyPublic)
	if err != nil {
		return GetEncryptionKeyResponse{}, err
	}

	res = GetEncryptionKeyResponse{
		UserID:        userID,
		DecryptionKey: resDB.PrivateKey, // user private key
		EncryptionKey: encryptionKey,    // server public key
	}

	return res, nil
}

func (account *AccountUsecase) Deposit(ctx context.Context, encryptedReq model.EncryptedRequestResponse) (model.EncryptedRequestResponse, error) {

	// Get user id from jwt claims that stored in context before
	userID := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims).UserID

	// Decrypt data from request body
	decryptedReq, err := encryption.Decrypt(account.cfg.RSAKeyPrivate, encryptedReq.Data)
	if err != nil || decryptedReq == nil {
		return model.EncryptedRequestResponse{}, err
	}

	var req DepositRequest
	if err = json.Unmarshal(decryptedReq, &req); err != nil {
		fmt.Println(err)
		return model.EncryptedRequestResponse{}, err
	}

	// Get current balance
	userAcc, err := account.acc.GetAccountByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Start database transaction
	tx, err := account.u.GetDBTx(ctx)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Rollback if any errors were found
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Update account balance
	err = account.acc.UpdateAccountBalance(ctx, tx, accDmn.AccountEntity{
		UserID:  userID,
		Balance: userAcc.Balance + req.Amount,
	})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Update transaction history
	_, err = account.tr.InsertNewTransaction(ctx, tx, transaction.TransactionEntity{
		UserID: userID,
		Amount: req.Amount,
		Type:   constanta.TRANSACTION_TYPE_DEPOSIT,
	})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Get user public and private key for encryption purpose
	userKey, err := account.u.GetUserByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Encrypt response
	encryptRes, err := response.EncryptResponse(userKey, MessageResponse{Message: "Success Deposit"})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Database transaction commit changes
	err = tx.Commit()
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	return model.EncryptedRequestResponse{
		Data: encryptRes,
	}, nil
}

func (account *AccountUsecase) Withdraw(ctx context.Context, encryptedReq model.EncryptedRequestResponse) (model.EncryptedRequestResponse, error) {

	// Get user id from jwt claims that stored in context before
	userID := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims).UserID

	// Decrypt data from request body
	decryptedReq, err := encryption.Decrypt(account.cfg.RSAKeyPrivate, encryptedReq.Data)
	if err != nil || decryptedReq == nil {
		return model.EncryptedRequestResponse{}, err
	}

	var req WithdrawRequest
	if err = json.Unmarshal(decryptedReq, &req); err != nil {
		fmt.Println(err)
		return model.EncryptedRequestResponse{}, err
	}

	// Get current balance
	userAcc, err := account.acc.GetAccountByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	if userAcc.Balance <= 0 {
		return model.EncryptedRequestResponse{}, errors.New("insufficient balance")
	}

	// Start database transaction
	tx, err := account.u.GetDBTx(ctx)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// rollback if any errors were found
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Update account balance
	err = account.acc.UpdateAccountBalance(ctx, tx, accDmn.AccountEntity{
		UserID:  userID,
		Balance: userAcc.Balance - req.Amount,
	})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Update transaction history
	_, err = account.tr.InsertNewTransaction(ctx, tx, transaction.TransactionEntity{
		UserID: userID,
		Amount: req.Amount,
		Type:   constanta.TRANSACTION_TYPE_WITHDRAW,
	})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Get user public and private key for encryption purpose
	userKey, err := account.u.GetUserByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Encrypt response
	encryptRes, err := response.EncryptResponse(userKey, MessageResponse{Message: "Success Withdraw"})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Database transaction commit changes
	err = tx.Commit()
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	return model.EncryptedRequestResponse{
		Data: encryptRes,
	}, nil
}

func (account *AccountUsecase) GetCompoundInterest(ctx context.Context, encryptedReq model.EncryptedRequestResponse) (model.EncryptedRequestResponse, error) {

	// Get user id from jwt claims that stored in context before
	userID := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims).UserID

	// Decrypt data from request body
	decryptedReq, err := encryption.Decrypt(account.cfg.RSAKeyPrivate, encryptedReq.Data)
	if err != nil || decryptedReq == nil {
		return model.EncryptedRequestResponse{}, err
	}

	var req GetCompoundInterestRequest
	if err = json.Unmarshal(decryptedReq, &req); err != nil {
		fmt.Println(err)
		return model.EncryptedRequestResponse{}, err
	}

	accountInformation, err := account.acc.GetAccountByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	interestInformation, err := account.ir.Get(ctx)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// validate year

	p := accountInformation.Balance
	r := interestInformation.Rate
	n := interestInformation.Compound
	t := req.ExpectedYears - time.Now().Year()

	// Compount Interest Formula
	// A = P(1 + r/n)^n*t
	balanceWithInterest := p * math.Pow((1+(float64(r)/100)/float64(n)), float64(int(n)*t))

	// Get user public and private key for encryption purpose
	userKey, err := account.u.GetUserByUserID(ctx, userID)
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	// Encrypt response
	encryptRes, err := response.EncryptResponse(userKey, GetCompoundInterestResponse{
		UserID:            userID,
		CurrentBalance:    float.RoundToDecimalPlaces(accountInformation.Balance, 3),
		AccumulateBalance: float.RoundToDecimalPlaces(balanceWithInterest, 3),
		Interest:          float.RoundToDecimalPlaces(balanceWithInterest-accountInformation.Balance, 3),
		Rate:              interestInformation.Rate,
		CompoundInterval:  interestInformation.Compound,
		ExpectedYears:     req.ExpectedYears,
	})
	if err != nil {
		return model.EncryptedRequestResponse{}, err
	}

	return model.EncryptedRequestResponse{
		Data: encryptRes,
	}, nil
}
