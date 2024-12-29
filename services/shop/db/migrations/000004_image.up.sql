CREATE TABLE IF NOT EXISTS image(
  id VARCHAR(36) PRIMARY KEY,
  url VARCHAR(1024) NOT NULL,
  description VARCHAR(256),
  shop_id VARCHAR(36) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_shop FOREIGN KEY (shop_id) REFERENCES shop(id) ON DELETE CASCADE
);

-- using trigger to automatically update updated_at field
CREATE OR REPLACE FUNCTION update_shop_updated_at() 
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_shop_updated_at_trigger
BEFORE UPDATE ON shop
FOR EACH ROW
EXECUTE FUNCTION update_shop_updated_at();
