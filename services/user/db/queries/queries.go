package queries

const (
	CREATE_PROFILE = `
		INSERT INTO profile 
		(id, name, image_url, dob, anniversary, gender, auth_id) 
		VALUES 
		(:id, :name, :image_url, :dob, :anniversary, :gender, :auth_id)
	`

	GET_PROFILE_BY_ID = `
		SELECT 
			id, name, image_url, dob, anniversary, gender, auth_id, 
			created_at, updated_at 
		FROM profile 
		WHERE id = ?
	`

	GET_PROFILE_BY_AUTH_ID = `
		SELECT 
			id, name, image_url, dob, anniversary, gender, auth_id, 
			created_at, updated_at 
		FROM profile 
		WHERE auth_id = ?
	`

	UPDATE_PROFILE = `
		UPDATE profile 
		SET 
			name = :name, 
			image_url = :image_url, 
			dob = :dob, 
			anniversary = :anniversary, 
			gender = :gender 
		WHERE id = :id
	`

	DELETE_PROFILE = `
		DELETE FROM profile 
		WHERE id = ?
	`

	CREATE_DELIVERY_ADDRESS = `
		INSERT INTO delivery_address (
			id, receiver_name, receiver_phone, address_alias, other_alias, 
			longitude, latitude, address, nearby_landmark, 
			delivery_instruction, auth_id, created_at, updated_at
		) VALUES (
			:id, :receiver_name, :receiver_phone, :address_alias, :other_alias, 
			:longitude, :latitude, :address, 
			:nearby_landmark, :delivery_instruction, :auth_id, :created_at, :updated_at
		)
	`

	DELETE_DELIVERY_ADDRESS = `
		DELETE FROM delivery_address 
		WHERE id = ?
	`

	GET_DELIVERY_ADDRESS_BY_ID = `
		SELECT 
			id, receiver_name, receiver_phone, address_alias, other_alias, 
			longitude, latitude, ST_AsText(location) AS location, address, 
			nearby_landmark, delivery_instruction, auth_id, created_at, updated_at 
		FROM delivery_address 
		WHERE id = ?
	`

	GET_DELIVERY_ADDRESSES_BY_AUTH_ID = `
		SELECT 
			id, receiver_name, receiver_phone, address_alias, other_alias, 
			longitude, latitude, ST_AsText(location) AS location, address, 
			nearby_landmark, delivery_instruction, auth_id, created_at, updated_at 
		FROM delivery_address 
		WHERE auth_id = ?
	`

	UPDATE_DELIVERY_ADDRESS = `
		UPDATE delivery_address 
		SET 
			receiver_name = :receiver_name, 
			receiver_phone = :receiver_phone, 
			address_alias = :address_alias, 
			other_alias = :other_alias, 
			longitude = :longitude, 
			latitude = :latitude, 
			address = :address, 
			nearby_landmark = :nearby_landmark, 
			delivery_instruction = :delivery_instruction, 
			updated_at = CURRENT_TIMESTAMP 
		WHERE id = :id
	`
)
