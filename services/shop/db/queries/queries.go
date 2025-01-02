package queries

import (
	"fmt"

	"github.com/yash91989201/superfast-delivery-api/common/types"
)

const (
	allowedShopAddressColumns = " id, address1, address2, longitude, latitude, nearby_landmark, city, state, pincode, country, shop_id, created_at "
)

const (
	CREATE_SHOP = "INSERT INTO shop" +
		"(id, name, shop_type, shop_status, owner_auth_id)" +
		"VALUES (:id, :name, :shop_type, :shop_status, :owner_auth_id)"

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

	GET_SHOP = "SELECT * FROM shop WHERE id = $1"

	GET_SHOP_BY_OWNER_ID = "SELECT * FROM shop WHERE owner_id = $1"

	// TODO: implement paginated queries
	GET_PAGINATED_SHOPS = `
  SELECT * 
  FROM shop
  WHERE deleted_at IS NULL
  `

	GET_SHOP_ADDRESS            = "SELECT" + allowedShopAddressColumns + "FROM shop_address WHERE id = $1"
	GET_SHOP_ADDRESS_BY_SHOP_ID = "SELECT" + allowedShopAddressColumns + "FROM shop_address WHERE shop_id = $1"
	GET_SHOP_CONTACT            = "SELECT * FROM shop_contact WHERE id = $1"
	GET_SHOP_CONTACT_BY_SHOP_ID = "SELECT * FROM shop_contact WHERE shop_id = $1"
	GET_SHOP_TIMINGS            = "SELECT * FROM shop_timing WHERE shop_id = $1"
	GET_SHOP_IMAGES             = "SELECT * FROM shop_image WHERE shop_id = $1"
	GET_SHOP_TIMING             = "SELECT * FROM shop_timing WHERE id = $1"
	GET_SHOP_IMAGE              = "SELECT * FROM shop_image WHERE id = $1"
)

func GetListShopQueryAndArgs(filters *types.ListShopFilters) (string, []interface{}) {
	query := "SELECT * FROM shop WHERE 1=1"
	var args []interface{}
	argIndex := 1

	if filters.Name != nil {
		if *filters.Name != "" {
			query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
			args = append(args, "%"+*filters.Name+"%")
			argIndex++
		}
	}

	if filters.ShopType != nil {
		query += fmt.Sprintf(" AND shop_type = $%d", argIndex)
		args = append(args, *filters.ShopType)
		argIndex++
	}

	if filters.ShopStatus != nil {
		query += fmt.Sprintf(" AND shop_status = $%d", argIndex)
		args = append(args, *filters.ShopStatus)
		argIndex++
	}

	if filters.OrderBy != nil {
		direction := *filters.OrderBy

		if direction != "ASC" && direction != "DESC" {
			direction = "ASC"
		}

		query += fmt.Sprintf(" ORDER BY created_at %s", direction)
	} else {
		query += " ORDER BY created_at DESC"
	}

	if filters.Limit != nil {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, *filters.Limit)
		argIndex++
	}

	if filters.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, *filters.Offset)
	}

	return query, args
}
