package com.ecom.orders.monitors;

import io.micrometer.core.instrument.Counter;
import io.micrometer.core.instrument.MeterRegistry;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class CustomMetricsService {

    private final MeterRegistry meterRegistry;

    public void trackHttpRequest(String method, String path, int status) {
        Counter.builder("http_request_total")
                .description("Total number of HTTP requests")
                .tags("method", method)
                .tags("path", path)
                .tags("status", String.valueOf(status))
                .register(meterRegistry)
                .increment();
    }
}