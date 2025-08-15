package models

import "context"

type StockRepository interface {
	ReserveStock(ctx context.Context, productID string, quantity int) error
	GetStock(ctx context.Context, productID string) (int, error)
	GetStockByOrderID(ctx context.Context, orderID int64) ([]StockItem, error)
	ReleaseStock(ctx context.Context, productID string, quantity int) error
	UpdateStockStatus(ctx context.Context, orderID int64, status string) error
}

type StockItem struct {
	ProductID string
	Quantity  int
	OrderID   int64
	Status    string
}
