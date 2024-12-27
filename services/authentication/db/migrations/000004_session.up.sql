CREATE TABLE IF NOT EXISTS `session` (
  `id` VARCHAR(36) PRIMARY KEY,
  `auth_id` VARCHAR(36) NOT NULL,
  `refresh_token` VARCHAR(512) NOT NULL,
  `is_revoked` BOOLEAN DEFAULT false NOT NULL, 
  `created_at` DATETIME DEFAULT NOW() NOT NULL, 
  `expires_at` DATETIME NOT NULL, 
  FOREIGN KEY (`auth_id`) REFERENCES `auth`(`id`) ON DELETE CASCADE
);
