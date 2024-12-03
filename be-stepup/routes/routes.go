package routes

import (
	"be-stepup/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, orderController *controllers.OrderController) {
	// Grouping routes for products
	productGroup := app.Group("/api/products")
	productGroup.Get("/", controllers.GetAllProducts)             // Route untuk mengambil semua produk
	productGroup.Get("/:id", controllers.GetProductByID)          // Route untuk mengambil produk berdasarkan ID
	productGroup.Get("/code/:code", controllers.GetProductByCode) // Route untuk mengambil produk berdasarkan kode unik
	productGroup.Post("/", controllers.CreateProduct)             // Route untuk membuat produk baru
	productGroup.Put("/:id", controllers.UpdateProduct)           // Route untuk memperbarui produk berdasarkan ID
	productGroup.Delete("/:id", controllers.DeleteProduct)        // Route untuk menghapus produk berdasarkan ID

	// Rute untuk upload gambar
	app.Post("/api/upload", controllers.UploadImage) // Rute untuk upload gambar
	// Grouping routes for orders
	orderGroup := app.Group("/api/orders")
	orderGroup.Get("/", orderController.GetOrders)                   // Route untuk mengambil semua order
	orderGroup.Get("/:id", orderController.GetOrderByID)             // Route untuk mengambil order berdasarkan ID
	orderGroup.Post("/", orderController.CreateOrder)                // Route untuk membuat order baru
	orderGroup.Put("/:id/status", orderController.UpdateOrderStatus) // Route untuk memperbarui status order
	orderGroup.Delete("/:id", orderController.DeleteOrder)           // Route untuk menghapus order berdasarkan ID

}
