package com.mahmud.orderservice.service;

public interface OrderApiService {
    /**
     * Updates the status of an order
     * @param orderId The ID of the order to update
     * @param userId The ID of the user making the request
     * @param status The new status (CANCELLED or COMPLETED)
     * @throws com.mahmud.orderservice.exception.ResourceNotFoundException if order not found or doesn't belong to user
     * @throws IllegalArgumentException if status is invalid
     */
    void updateOrderStatus(Long orderId, String userId, String status);
}
