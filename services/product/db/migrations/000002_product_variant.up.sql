CREATE TABLE `product_variant` (
  `id` VARCHAR(36) PRIMARY KEY,
  `variant_name` VARCHAR(128),
  `relative_price` DECIMAL(10,2) DEFAULT 0,
  -- is variant pricing related to the product +/-
  `relative_pricing` BOOLEAN DEFAULT TRUE, 
  `price` DECIMAL(10,2) DEFAULT 0,
  `product_id` VARCHAR(36) NOT NULL,
  CONSTRAINT fk_product_variant_product FOREIGN KEY (`product_id`) REFERENCES `product`(id) ON DELETE CASCADE
);
