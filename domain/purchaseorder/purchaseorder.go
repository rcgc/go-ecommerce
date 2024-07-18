package purchaseorder

import (
	"github.com/rcgc/go-ecommerce/model"

	"github.com/google/uuid"
)

type UseCase interface {
	Create(m *model.PurchaseOrder) error

	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}

type Storage interface {
	Create(m *model.PurchaseOrder) error

	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}
