CREATE TABLE IF NOT EXISTS shop (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  type VARCHAR(36) CHECK (type IN ('restaurant', 'grocery', 'pharmaceutical')) NOT NULL,
  status VARCHAR(8) CHECK (type IN ('open', 'closed')) NOT NULL,
  owner_id VARCHAR(36) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
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

