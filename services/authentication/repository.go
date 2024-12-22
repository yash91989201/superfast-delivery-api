package authentication

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yash91989201/superfast-delivery-api/common/types"
	"github.com/yash91989201/superfast-delivery-api/services/authentication/db/queries"
)

type Repository interface {
	Close() error
	CreateAuth(ctx context.Context, a *types.Auth) error
	GetAuthById(ctx context.Context, id string) (*types.Auth, error)
	GetAuthByEmail(ctx context.Context, email string) (*types.Auth, error)
	GetAuthByPhone(ctx context.Context, phone string) (*types.Auth, error)
	DeleteAuth(ctx context.Context, id string) error

	CreateEmailVerification(ctx context.Context, v *types.EmailVerification) error
	CreatePhoneVerification(ctx context.Context, v *types.PhoneVerification) error
	GetEmailVerification(ctx context.Context, email string) (*types.EmailVerification, error)
	GetPhoneVerification(ctx context.Context, phone string) (*types.PhoneVerification, error)
	DeleteEmailVerification(ctx context.Context, email string) error
	DeletePhoneVerification(ctx context.Context, phone string) error
}

type mysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(dbUrl string) (Repository, error) {
	db, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %w", err)

	}

	return &mysqlRepository{
		db: db,
	}, nil
}

func (r *mysqlRepository) Close() error {
	return r.db.Close()
}

func (r *mysqlRepository) CreateAuth(ctx context.Context, a *types.Auth) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.CREATE_AUTH, &a)
	if err != nil {
		return nil
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to insert product, 0 rows affected: %w", err)
	}

	return nil
}
func (r *mysqlRepository) GetAuthById(ctx context.Context, id string) (*types.Auth, error) {
	var auth *types.Auth
	if err := r.db.GetContext(ctx, queries.GET_AUTH_BY_ID, id); err != nil {
		return nil, err
	}

	return auth, nil
}
func (r *mysqlRepository) GetAuthByEmail(ctx context.Context, email string) (*types.Auth, error) {
	auth := &types.Auth{}
	if err := r.db.GetContext(ctx, &auth, queries.GET_AUTH_BY_EMAIL, email); err != nil {
		fmt.Print(err)
		return nil, err
	}

	return auth, nil
}

func (r *mysqlRepository) GetAuthByPhone(ctx context.Context, phone string) (*types.Auth, error) {
	var auth *types.Auth
	if err := r.db.GetContext(ctx, queries.GET_AUTH_BY_PHONE, phone); err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *mysqlRepository) DeleteAuth(ctx context.Context, id string) error {
	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_AUTH, id)
	if err != nil {
		return err
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to delete auth with id %s, 0 rows affected : %w", id, err)
	}

	return nil
}

func (r *mysqlRepository) CreateEmailVerification(ctx context.Context, v *types.EmailVerification) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.CREATE_EMAIL_VERIFICATION, &v)
	if err != nil {
		return fmt.Errorf("Error creating email verification: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create email verification, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) CreatePhoneVerification(ctx context.Context, v *types.PhoneVerification) error {
	queryRes, err := r.db.NamedExecContext(ctx, queries.CREATE_PHONE_VERIFICATION, &v)
	if err != nil {
		return fmt.Errorf("Error creating phone verification: %w", err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to create email verification, 0 rows affected: %w", err)
	}

	return nil
}

func (r *mysqlRepository) GetEmailVerification(ctx context.Context, email string) (*types.EmailVerification, error) {
	ev := &types.EmailVerification{}
	if err := r.db.GetContext(ctx, ev, queries.GET_EMAIL_VERIFICATION, email); err != nil {
		return nil, fmt.Errorf("Error getting email verification")
	}

	return ev, nil
}

func (r *mysqlRepository) GetPhoneVerification(ctx context.Context, phone string) (*types.PhoneVerification, error) {
	var pv *types.PhoneVerification
	if err := r.db.QueryRowContext(ctx, queries.GET_PHONE_VERIFICATION, &pv); err != nil {
		return nil, fmt.Errorf("Error getting email verification")
	}

	return pv, nil
}

func (r *mysqlRepository) DeleteEmailVerification(ctx context.Context, email string) error {

	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_EMAIL_VERIFICATION, email)
	if err != nil {
		return fmt.Errorf("Failed to delete email verification with email %s: %w", email, err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to delete email verification with email %s, 0 rows affected : %w", email, err)
	}

	return nil
}

func (r *mysqlRepository) DeletePhoneVerification(ctx context.Context, phone string) error {

	queryRes, err := r.db.ExecContext(ctx, queries.DELETE_PHONE_VERIFICATION, phone)
	if err != nil {
		return fmt.Errorf("Failed to delete phone verification with phone %s: %w", phone, err)
	}

	if rowsAffected, err := queryRes.RowsAffected(); rowsAffected == 0 || err != nil {
		return fmt.Errorf("Failed to delete phone verification with phone %s, 0 rows affected : %w", phone, err)
	}

	return nil
}
