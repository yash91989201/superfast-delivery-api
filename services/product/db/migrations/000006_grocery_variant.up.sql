CREATE TABLE IF NOT EXISTS `grocery_variant` (
  `id` VARCHAR(36) PRIMARY KEY,
  `variant_type` VARCHAR(128),
  `variant_value` VARCHAR(128),
  `product_id` VARCHAR(36) NOT NULL,
  CONSTRAINT fk_grocery_variant_product FOREIGN KEY (`product_id`) REFERENCES `product`(id) ON DELETE CASCADE
);
