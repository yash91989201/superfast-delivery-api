package queries

const (
	CREATE_SHOP = "INSERT INTO shop" +
		"(id, name, shop_type, shop_status, owner_id)" +
		"VALUES (:id, :name, :shop_type, :shop_status, :owner_id)"

		// CREATE_SHOP_ADDRESS = "INSERT INTO shop_address" +
		// 	"(id, address1, address2, location, nearby_landmark, city, state, pincode, country, shop_id)" +
		// 	"VALUES (:id, :address1, :address2, ST_SetSRID(ST_MakePoint(:lng, :lat), 4326), :nearby_landmark, :city, :state, :pincode, :country, :shop_id)"

	CREATE_SHOP_ADDRESS = "INSERT INTO shop_address" +
		"(id, address1, address2, longitude, latitude, nearby_landmark, city, state, pincode, country, shop_id)" +
		"VALUES (:id, :address1, :address2, :longitude, :latitude, :nearby_landmark, :city, :state, :pincode, :country, :shop_id)"

	CREATE_SHOP_CONTACT = "INSERT INTO shop_contact" +
		"(id, name, phone_number, email, shop_id)" +
		"VALUES (:id, :name, :phone_number, :email, :shop_id)"

	CREATE_SHOP_IMAGE = "INSERT INTO shop_image" +
		"(id, image_url, description, shop_id)" +
		"VALUES (:id, :image_url, :description, :shop_id)"

	CREATE_SHOP_TIMING = "INSERT INTO shop_timing" +
		"(id, day, opens_at, closes_at, shop_id)" +
		"VALUES (:id, :day, :opens_at, :closes_at, :shop_id)"
)
