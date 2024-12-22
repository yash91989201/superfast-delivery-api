package queries

const (
	CREATE_AUTH = "INSERT INTO auth" +
		"(id, email, email_verified, phone, role)" +
		"VALUES (:id, :email, :email_verified, :phone, :role)"
	GET_AUTH_BY_ID            = "SELECT id, email, email_verified, phone, role FROM auth WHERE id=?"
	GET_AUTH_BY_EMAIL         = "SELECT * FROM auth WHERE email=?"
	GET_AUTH_BY_PHONE         = "SELECT * FROM auth WHERE phone=?"
	DELETE_AUTH               = "DELETE FROM auth WHERE id=?"
	CREATE_EMAIL_VERIFICATION = "INSERT INTO email_verification" +
		"(token, email, expires_at) VALUES(:token, :email, :expires_at)"
	CREATE_PHONE_VERIFICATION = "INSERT INTO phone_verification" +
		"(token, phone, expires_at) VALUES(:token, :phone, :expires_at)"
	GET_EMAIL_VERIFICATION    = "SELECT * FROM email_verification WHERE email=?"
	GET_PHONE_VERIFICATION    = "SELECT * FROM phone_verification WHERE phone=?"
	DELETE_EMAIL_VERIFICATION = "DELETE FROM email_verification WHERE email=?"
	DELETE_PHONE_VERIFICATION = "DELETE FROM email_verification WHERE phone=?"

	CREATE_PROFILE = "INSERT INTO profile" +
		"(id, name, image_url, dob, anniversary, gender, auth_id)" +
		"VALUES (:id, :name, :image_url, :dob, :anniversary, :gender, :auth_id)"

	GET_PROFILE_BY_ID      = "SELECT id, name, image_url, dob, anniversary, gender, auth_id, created_at, updated_at FROM profile WHERE id=?"
	GET_PROFILE_BY_AUTH_ID = "SELECT id, name, image_url, dob, anniversary, gender, auth_id, created_at, updated_at FROM profile WHERE auth_id=?"
	UPDATE_PROFILE         = "UPDATE profile SET name = :name, image_url = :image_url, dob = :dob, anniversary = :anniversary, gender = :gender WHERE id = :id"
	DELETE_PROFILE         = "DELETE FROM profile WHERE id = ?"
)
