package invoice

import (
	"github.com/google/uuid"
	"github.com/rcgc/go-ecommerce/model"
)

type UseCase interface {
	Create(m *model.PurchaseOrder) error
	GetByUserID(userID uuid.UUID) (model.InvoicesReport, error)
	GetAll() (model.InvoicesReport, error)
}

type Storage interface {
	Create(m *model.Invoice, ms model.InvoiceDetails) error
}

type StorageInvoiceDetailReport interface {
	HeadByInvoiceID(ID uuid.UUID) (model.InvoiceReport, error)
	HeadsByUserID(userID uuid.UUID) (model.InvoicesReport, error)
	AllHead() (model.InvoicesReport, error)
	AllDetailsByInvoiceID(ID uuid.UUID) (model.InvoiceDetailsReports, error)
}
