package entities

import (
	"time"
)

type User struct {
	ID                   string `xorm:"not null pk VARCHAR(36) id"`
	Name                 string `xorm:"not null TEXT default '' name"`
	Email                string `xorm:"-"`
	EmailHash            string `xorm:"not null TEXT email_hash"`
	EncryptedEmail       []byte `xorm:"not null BYTEA encrypted_email"`
	EmailDecryptionNonce []byte `xorm:"not null BYTEA email_decryption_nonce"`

	EmailAcceptedAt *time.Time `xorm:"TIMESTAMPZ email_accepted_at"`
	LastLoginAt     time.Time  `xorm:"TIMESTAMPZ last_login_at"`

	CreatedAt time.Time  `xorm:"created TIMESTAMPZ created_at"`
	UpdatedAt time.Time  `xorm:"updated TIMESTAMPZ updated_at"`
	DeletedAt *time.Time `xorm:"TIMESTAMPZ deleted_at"`
}

func (User) TableName() string {
	return "users"
}