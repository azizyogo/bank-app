package sandbox

import (
	"context"

	"github.com/azizyogo/bank-app/common/auth"
	constanta "github.com/azizyogo/bank-app/common/const"
	"github.com/azizyogo/bank-app/common/encryption"
)

func (s SandboxUsecase) SandboxEncrypt(userID int, msg []byte) (string, error) {

	servicePublic := s.cfg.RSAKeyPublic

	encrypted, err := encryption.Encrypt(servicePublic, msg)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (s SandboxUsecase) SandboxDecrypt(ctx context.Context, msg string) (res string, err error) {

	claims := ctx.Value(constanta.CLAIMS_CONTEXT_KEY).(*auth.Claims)

	userKey, err := s.u.GetKeyByUserID(ctx, claims.UserID)
	if err != nil {
		return
	}

	privateKey, err := encryption.DecodePEMToPrivateKey(userKey.PrivateKey)
	if err != nil {
		return
	}

	decrypted, err := encryption.Decrypt(privateKey, msg)
	if err != nil {
		return
	}

	return string(decrypted), nil
}
