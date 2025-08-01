package com.ecom.products.configs;

import com.ecom.products.monitors.CustomMetricsService;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;

@Component
public class MetricsInterceptor implements HandlerInterceptor {

    private final CustomMetricsService metricsService;

    public MetricsInterceptor(CustomMetricsService metricsService) {
        this.metricsService = metricsService;
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response,
                                Object handler, Exception ex) {
        metricsService.trackHttpRequest(
                request.getMethod(),
                request.getRequestURI(),
                response.getStatus()
        );
    }
}