package com.mahmud.orderservice.controller;

import com.mahmud.orderservice.service.OrderApiService;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.oauth2.jwt.Jwt;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/orders")
public class OrderApiController {

    private final OrderApiService orderApiService;

    public OrderApiController(OrderApiService orderApiService) {
        this.orderApiService = orderApiService;
    }

    @PutMapping("/{id}/cancel")
    public ResponseEntity<Void> cancelOrder(@PathVariable Long id) {
        orderApiService.updateOrderStatus(id, "SYSTEM", "CANCELLED");
        return ResponseEntity.ok().build();
    }

    @PutMapping("/{id}/complete")
    public ResponseEntity<Void> completeOrder(@PathVariable Long id) {
        orderApiService.updateOrderStatus(id, "SYSTEM", "COMPLETED");
        return ResponseEntity.ok().build();
    }
}
