package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"default:now()"`

	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`

	ProductID uuid.UUID
	Product   Product `gorm:"foreignKey:ProductID"`

	VoucherID *uuid.UUID
	Voucher   *Voucher `gorm:"foreignKey:VoucherID"`

	Total int
}
