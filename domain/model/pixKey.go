package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixkey *PixKey) (*PixKey, error)
	FindKeyByType(key string, pixType string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	PixType   string   `json:"pix_type" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	Account   *Account `valid:"-"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixkey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixkey)

	if pixkey.PixType != "email" && pixkey.PixType != "cpf" {
		return errors.New("invalid pix type")
	}

	if pixkey.Status != "active" && pixkey.Status != "inactive" {
		return errors.New("Invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(account *Account, pix_type string, key string) (*PixKey, error) {
	pixKey := PixKey{
		PixType: pix_type,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
