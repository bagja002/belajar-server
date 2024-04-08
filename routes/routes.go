package routes

import(
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx)error{

	//Menggunakan Query Parameter 
	nama:= c.Query("nama")
	umur := c.Query("umur")

	//Parameter
	return c.JSON(fiber.Map{
		"pesan":"hello selamat datang "+nama,
		"umur_saya": umur,
	})
}

func HelloParams(c *fiber.Ctx)error{

	//Menambahkan Parameters

	id:= c.Params("id")

	return c.JSON(fiber.Map{
		"Pesan":"Berhasil",
		"id":id,
	})
}
