package queries

const (
	CREATE_AUTH = "INSERT INTO auth" +
		"(id, email, email_verified, phone, type)" +
		"VALUES (:id, :email, :email_verified, :phone, :type)"
	GET_AUTH_BY_ID            = "SELECT * FROM auth WHERE id=?"
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
)
