package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderStatusUpdateRequest struct {
	Status string `json:"status"`
}

func (h *StockHandler) HandleCancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.ParseInt(vars["orderId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Release the stock items back to inventory
	err = h.stockUseCase.ReleaseStockForOrder(r.Context(), orderID)
	if err != nil {
		http.Error(w, "Failed to release stock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *StockHandler) HandleCompleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.ParseInt(vars["orderId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Mark the stock items as completed/sold
	err = h.stockUseCase.FinalizeStockForOrder(r.Context(), orderID)
	if err != nil {
		http.Error(w, "Failed to finalize stock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
