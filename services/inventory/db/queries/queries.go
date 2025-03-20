package queries

const (
	INSERT_ITEM_STOCK    = "INSERT INTO item_stock (id, item_id, quantity) VALUES (:id, :item_id, :quantity)"
	GET_ITEM_STOCK_BY_ID = "SELECT * FROM item_stock WHERE id = ?"
	UPDATE_ITEM_STOCK    = "UPDATE item_stock SET quantity = :quantity, restock_qty = :restock_qty WHERE id = :id"
	DELETE_ITEM_STOCK    = "DELETE FROM item_stock WHERE id = ?"

	INSERT_VARIANT_STOCK    = "INSERT INTO variant_stock (id, variant_id, quantity) VALUES (:id, :variant_id, :quantity)"
	GET_VARIANT_STOCK_BY_ID = "SELECT * FROM variant_stock WHERE id = ?"
	UPDATE_VARIANT_STOCK    = "UPDATE variant_stock SET quantity = :quantity, restock_qty = :restock_qty WHERE id = :id"
	DELETE_VARIANT_STOCK    = "DELETE FROM variant_stock WHERE id = ?"

	INSERT_ADDON_STOCK    = "INSERT INTO addon_stock (id, addon_id, quantity) VALUES (:id, :addon_id, :quantity)"
	GET_ADDON_STOCK_BY_ID = "SELECT * FROM addon_stock WHERE id = "
	UPDATE_ADDON_STOCK    = "UPDATE addon_stock SET quantity = :quantity, restock_qty = :restock_qty WHERE id = :id"
	DELETE_ADDON_STOCK    = "DELETE FROM addon_stock WHERE id = ?"
)
