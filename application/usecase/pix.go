package usecase

import (
	"errors"

	"github.com/allanCordeiro/fc-codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, pixType string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, pixType, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new pix key at the moment")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, pixType string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByType(key, pixType)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
