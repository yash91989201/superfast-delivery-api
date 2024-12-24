package queries

const (
	CREATE_PROFILE = "INSERT INTO profile" +
		"(id, name, image_url, dob, anniversary, gender, auth_id)" +
		"VALUES (:id, :name, :image_url, :dob, :anniversary, :gender, :auth_id)"

	GET_PROFILE_BY_ID      = "SELECT id, name, image_url, dob, anniversary, gender, auth_id, created_at, updated_at FROM profile WHERE id = ?"
	GET_PROFILE_BY_AUTH_ID = "SELECT id, name, image_url, dob, anniversary, gender, auth_id, created_at, updated_at FROM profile WHERE auth_id = ?"
	UPDATE_PROFILE         = "UPDATE profile SET name = :name, image_url = :image_url, dob = :dob, anniversary = :anniversary, gender = :gender WHERE id = :id"
	DELETE_PROFILE         = "DELETE FROM profile WHERE id = ?"
)
