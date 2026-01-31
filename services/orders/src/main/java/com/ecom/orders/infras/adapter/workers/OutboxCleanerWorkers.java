package com.ecom.orders.infras.adapter.workers;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import lombok.extern.slf4j.Slf4j;

@Component
@Slf4j
public class OutboxCleanerWorkers {

   @Autowired
   private JdbcTemplate jdbcTemplate;

   // Chạy hàng ngày lúc 2h sáng
   @Scheduled(cron = "0 0 2 * * *")
   public void cleanupProcessedEvents() {
      String sql = """
            DELETE FROM outbox
            WHERE processed_at IS NOT NULL
            AND processed_at < NOW() - INTERVAL '7 days'
            """;

      int deleted = jdbcTemplate.update(sql);
      log.info("Cleaned up {} processed outbox events", deleted);
   }
}