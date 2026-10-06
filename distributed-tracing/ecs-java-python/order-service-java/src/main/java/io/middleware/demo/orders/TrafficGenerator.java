package io.middleware.demo.orders;

import java.util.List;
import java.util.concurrent.ThreadLocalRandom;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestClient;
import org.springframework.web.client.RestClientResponseException;

/** Calls this service's own HTTP endpoint so every request starts as an inbound server request. */
@Component
public class TrafficGenerator {
    private static final Logger log = LoggerFactory.getLogger(TrafficGenerator.class);
    private static final List<String> PRODUCTS = List.of("P100", "P200", "P300", "P400", "P500");

    private final RestClient self;
    private final boolean enabled;

    public TrafficGenerator(RestClient.Builder builder,
                            @Value("${server.port:8080}") int port,
                            @Value("${traffic.enabled:true}") boolean enabled) {
        this.self = builder.baseUrl("http://localhost:" + port).build();
        this.enabled = enabled;
    }

    @Scheduled(initialDelay = 15000, fixedDelayString = "${traffic.interval-ms:3000}")
    public void tick() {
        if (!enabled) return;
        String productId = PRODUCTS.get(ThreadLocalRandom.current().nextInt(PRODUCTS.size()));
        try {
            self.get().uri("/api/orders/quote/{id}", productId).retrieve().toBodilessEntity();
            log.info("quote ok for {}", productId);
        } catch (RestClientResponseException e) {
            log.warn("quote failed for {}: HTTP {}", productId, e.getStatusCode().value());
        } catch (Exception e) {
            log.warn("quote request error for {}: {}", productId, e.toString());
        }
    }
}
