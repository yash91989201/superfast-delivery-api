CREATE TABLE IF NOT EXISTS `variant_stock`(
  `id` VARCHAR(36) PRIMARY KEY,
  `variant_id` VARCHAR(48) NOT NULL,
  `quantity` INT NOT NULL DEFAULT 0,
  `restock_qty` INT NOT NULL DEFAULT 5, 
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `unique_variant_id` (`variant_id`)
);
