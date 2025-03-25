CREATE INDEX idx_shop_address_location 
ON shop_address 
USING GIST (location);
