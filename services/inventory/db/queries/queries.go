package queries

const (
	INSERT_ITEM_STOCK = "INSERT INTO item_stock" +
		"(id, item_id, quantity, restock_qty)" +
		"VALUES (:id, :item_id, :quantity, :restock_qty)"
	INSERT_VARIANT_STOCK = "INSERT INTO variant_stock" +
		"(id, variant_id, quantity restock_qty)" +
		"VALUES (:id, :variant_id, :quantity, :restock_qty)"
	INSERT_ADDON_STOCK = "INSERT INTO addon_stock" +
		"(id, addon_id, quantity, restock_qty)" +
		"VALUES (:id, :addon_id, :quantity, :restock_qty)"
)
