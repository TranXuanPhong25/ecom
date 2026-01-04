package com.ecom.orders.configs;

import com.ecom.orders.monitors.CustomMetricsService;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;

import org.springframework.lang.NonNull;
import org.springframework.lang.Nullable;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

@Component
public class MetricsInterceptor implements HandlerInterceptor {

    private final CustomMetricsService metricsService;

    public MetricsInterceptor(CustomMetricsService metricsService) {
        this.metricsService = metricsService;
    }

    @Override
    public void afterCompletion(@NonNull HttpServletRequest request, @NonNull HttpServletResponse response,
                                @NonNull Object handler, @Nullable Exception ex) {
        metricsService.trackHttpRequest(
                request.getMethod(),
                request.getRequestURI(),
                response.getStatus()
        );
    }
}