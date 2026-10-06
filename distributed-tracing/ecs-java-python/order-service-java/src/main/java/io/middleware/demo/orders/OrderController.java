package io.middleware.demo.orders;

import java.util.LinkedHashMap;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestClient;
import org.springframework.web.client.RestClientResponseException;

@RestController
public class OrderController {
    private static final Logger log = LoggerFactory.getLogger(OrderController.class);

    private final RestClient inventory;

    public OrderController(RestClient.Builder builder, @Value("${inventory.url}") String inventoryUrl) {
        this.inventory = builder.baseUrl(inventoryUrl).build();
    }

    @GetMapping("/health")
    public Map<String, String> health() {
        return Map.of("status", "ok");
    }

    @GetMapping("/api/orders/quote/{productId}")
    public ResponseEntity<Map<String, Object>> quote(@PathVariable String productId) {
        try {
            Map<?, ?> product = inventory.get().uri("/api/products/{id}", productId).retrieve().body(Map.class);
            Map<?, ?> rating = inventory.get().uri("/api/products/{id}/rating", productId).retrieve().body(Map.class);

            Map<String, Object> quote = new LinkedHashMap<>();
            quote.put("productId", productId);
            quote.put("name", product.get("name"));
            quote.put("price", product.get("price"));
            quote.put("inStock", ((Number) product.get("stock")).intValue() > 0);
            quote.put("averageRating", rating.get("averageRating"));
            return ResponseEntity.ok(quote);
        } catch (RestClientResponseException e) {
            log.error("inventory-service call failed for product {}: HTTP {}", productId, e.getStatusCode().value());
            return ResponseEntity.status(HttpStatus.BAD_GATEWAY)
                    .body(Map.of("productId", productId, "error", "inventory-service returned " + e.getStatusCode().value()));
        }
    }
}
