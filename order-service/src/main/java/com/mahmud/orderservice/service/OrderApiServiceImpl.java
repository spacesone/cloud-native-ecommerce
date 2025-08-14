package com.mahmud.orderservice.service;

import com.mahmud.orderservice.entity.Order;
import com.mahmud.orderservice.exception.ResourceNotFoundException;
import com.mahmud.orderservice.repository.OrderRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

@Service
public class OrderApiServiceImpl implements OrderApiService {

    private static final Set<String> ALLOWED_STATUS_UPDATES = new HashSet<>(Arrays.asList("CANCELLED", "COMPLETED"));
    private final OrderRepository orderRepository;

    public OrderApiServiceImpl(OrderRepository orderRepository) {
        this.orderRepository = orderRepository;
    }

    @Override
    @Transactional
    public void updateOrderStatus(Long orderId, String userId, String status) {
        // Validate the status
        if (!ALLOWED_STATUS_UPDATES.contains(status)) {
            throw new IllegalArgumentException("Invalid status update: " + status);
        }

        // Get and validate the order
        Order order = orderRepository.findById(orderId)
                .orElseThrow(() -> new ResourceNotFoundException("Order not found with id: " + orderId));

        // Check if the order belongs to the user
        if (!order.getUserId().equals(userId)) {
            throw new ResourceNotFoundException("Order does not belong to user: " + userId);
        }

        // Validate current status
        String currentStatus = order.getStatus();
        if ("CANCELLED".equals(currentStatus) || "COMPLETED".equals(currentStatus)) {
            throw new IllegalStateException("Cannot update order with status: " + currentStatus);
        }

        // Only allow cancellation of PENDING orders
        if ("CANCELLED".equals(status) && !"PENDING".equals(currentStatus)) {
            throw new IllegalStateException("Only PENDING orders can be cancelled");
        }

        // Only allow completion of PAID orders
        if ("COMPLETED".equals(status) && !"PAID".equals(currentStatus)) {
            throw new IllegalStateException("Only PAID orders can be marked as completed");
        }

        // Update the status
        order.setStatus(status);
        orderRepository.save(order);
    }
}
