package handlers

import (
	"net/http"
	"bytes"
	"encoding/json"
	"io"
)

type OrderServiceHandler struct {
	orderServiceURL string
	apiKey         string
}

func NewOrderServiceHandler(orderServiceURL, apiKey string) *OrderServiceHandler {
	return &OrderServiceHandler{
		orderServiceURL: orderServiceURL,
		apiKey:         apiKey,
	}
}

func (h *OrderServiceHandler) HandleOrderStatusUpdate(w http.ResponseWriter, r *http.Request) {
	// Get the order ID and action from the URL
	orderId := mux.Vars(r)["orderId"]
	action := mux.Vars(r)["action"] // will be either "cancel" or "complete"

	// Create a new request to the order service
	url := h.orderServiceURL + "/orders/" + orderId + "/" + action
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Add the API key header
	req.Header.Set("X-API-KEY", h.apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to order service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the status code and response body
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
