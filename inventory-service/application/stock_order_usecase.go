package application

import (
	"context"
)

func (u *stockUseCase) ReleaseStockForOrder(ctx context.Context, orderID int64) error {
	// Get stock items for the order
	stockItems, err := u.stockRepo.GetStockByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	// Release each stock item back to inventory
	for _, item := range stockItems {
		err = u.stockRepo.ReleaseStock(ctx, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}

	// Update stock items status to cancelled
	return u.stockRepo.UpdateStockStatus(ctx, orderID, "CANCELLED")
}

func (u *stockUseCase) FinalizeStockForOrder(ctx context.Context, orderID int64) error {
	// Get stock items for the order
	stockItems, err := u.stockRepo.GetStockByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	// Update stock items status to completed
	return u.stockRepo.UpdateStockStatus(ctx, orderID, "COMPLETED")
}
