CREATE EXTENSION IF NOT EXISTS postgis;

-- Function that creates GEOMETRY point from given inputs
CREATE OR REPLACE FUNCTION create_geometry_point(longitude DOUBLE PRECISION, latitude DOUBLE PRECISION)
  RETURNS GEOMETRY
  LANGUAGE SQL 
  IMMUTABLE 
AS $$
  SELECT ST_SetSRID(ST_MakePoint(longitude, latitude), 4326)
$$;

CREATE TABLE IF NOT EXISTS shop_address (
  id VARCHAR(36) PRIMARY KEY,
  shop_id VARCHAR(36) NOT NULL, 
  longitude DOUBLE PRECISION NOT NULL,
  latitude DOUBLE PRECISION NOT NULL,
  location GEOMETRY GENERATED ALWAYS AS (create_geometry_point(longitude, latitude)) STORED,
  address VARCHAR(256) NOT NULL,
  nearby_landmark VARCHAR(128) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 

  CONSTRAINT fk_shop FOREIGN KEY (shop_id) REFERENCES shop(id) ON DELETE CASCADE
);

-- Trigger to update `updated_at` on row updates
CREATE OR REPLACE FUNCTION set_shop_address_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_shop_address_timestamp
BEFORE UPDATE ON shop_address
FOR EACH ROW
EXECUTE FUNCTION set_shop_address_updated_at();
