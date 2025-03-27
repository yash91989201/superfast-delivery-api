package queries

const (
	CREATE_AUTH = "INSERT INTO auth" +
		"(id, email, email_verified, phone, auth_role)" +
		"VALUES (:id, :email, :email_verified, :phone, :auth_role)"
	GET_AUTH_BY_ID            = "SELECT id, email, email_verified, phone, auth_role FROM auth WHERE id = ?"
	GET_AUTH_BY_EMAIL         = "SELECT * FROM auth WHERE email = ?"
	GET_AUTH_BY_PHONE         = "SELECT * FROM auth WHERE phone = ?"
	DELETE_AUTH               = "DELETE FROM auth WHERE id=?"
	UPDATE_EMAIL              = "UPDATE auth SET email = ? WHERE id = ?"
	UPDATE_PHONE              = "UPDATE auth SET phone = ? WHERE id = ?"
	CREATE_EMAIL_VERIFICATION = "INSERT INTO email_verification" +
		"(token, email, expires_at) VALUES(:token, :email, :expires_at)"
	CREATE_PHONE_VERIFICATION = "INSERT INTO phone_verification" +
		"(token, phone, expires_at) VALUES(:token, :phone, :expires_at)"
	GET_EMAIL_VERIFICATION    = "SELECT * FROM email_verification WHERE email = ?"
	GET_PHONE_VERIFICATION    = "SELECT * FROM phone_verification WHERE phone = ?"
	DELETE_EMAIL_VERIFICATION = "DELETE FROM email_verification WHERE email = ?"
	DELETE_PHONE_VERIFICATION = "DELETE FROM email_verification WHERE phone = ?"

	CREATE_SESSION = "INSERT INTO session " +
		"(id, auth_id, refresh_token, is_revoked, expires_at)" +
		"VALUES (:id, :auth_id, :refresh_token, :is_revoked, :expires_at)"
	GET_SESSION            = "SELECT * FROM session WHERE refresh_token = ?"
	GET_SESSION_BY_ID      = "SELECT * FROM session WHERE id = ?"
	GET_SESSION_BY_AUTH_ID = "SELECT * FROM session WHERE auth_id = ?"
	REVOKE_SESSION         = "UPDATE session SET is_revoked = 1 WHERE id = ?"
	DELETE_SESSION         = "DELETE FROM session WHERE id = ?"
)
