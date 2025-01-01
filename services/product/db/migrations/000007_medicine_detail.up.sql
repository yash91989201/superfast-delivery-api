CREATE TABLE `medicine_detail` (
  `id` VARCHAR(36) PRIMARY KEY,
  `prescription_required` BOOLEAN DEFAULT FALSE,
  `dosage` VARCHAR(128),
  `expiry_date` DATE,
  `product_id` VARCHAR(36) NOT NULL,
  CONSTRAINT fk_medicine_detail_product FOREIGN KEY (`product_id`) REFERENCES `product`(id) ON DELETE CASCADE
);

