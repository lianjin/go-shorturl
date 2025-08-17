CREATE TABLE `short_urls` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `short_code` varchar(255) NOT NULL DEFAULT '',
  `origin_url` varchar(2048) NOT NULL DEFAULT '' COMMENT 'URL',
  `hash_code` varchar(255) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_hash` (`hash_code`),
  UNIQUE KEY `uniq_short_code` (`short_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;