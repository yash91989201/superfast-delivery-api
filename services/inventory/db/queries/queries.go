package queries

const (
	INSERT_ITEM_STOCK = "INSERT INTO item_stock" +
		"(id, item_id, quantity)" +
		"VALUES (:id, :item_id, :quantity)"
	INSERT_VARIANT_STOCK = "INSERT INTO variant_stock" +
		"(id, variant_id, quantity)" +
		"VALUES (:id, :variant_id, :quantity)"
	INSERT_ADDON_STOCK = "INSERT INTO addon_stock" +
		"(id, addon_id, quantity)" +
		"VALUES (:id, :addon_id, :quantity)"
)
