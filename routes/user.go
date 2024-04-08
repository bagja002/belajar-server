package routes

import (
	"fmt"
	"meeting3/database"
	"meeting3/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

//membuat file data user

func AllData(c *fiber.Ctx) error {

	var users []models.User

	database.DB.Find(&users)

	fmt.Println(users)

	//Buat Data kalo Untuk memb uat data excel baru atau yang lain

	// Untuk Generate File
	// Misal Darai Login di tas sudah menghasilkan File yaitu file.xlsxx

	return c.Download("../files/file.xlsx")
}

func Register(c *fiber.Ctx) error {

	//Membuat Inputan JSON
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//Encrip password
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	NoTelpon, _ := strconv.Atoi(data["no_telpon"])

	user := models.User{
		NamaLengkap: data["nama_lengkap"],
		NoTelpon:    NoTelpon,
		Email:       data["email"],
		Password:    string(password),
	}

	//Pengecekan data email yang sudah di pakai
	var email models.User

	//Aku Akana mengambil data berdasarkan email di database

	database.DB.Where("email = ?", user.Email).Find(&email)

	if email.Email == user.Email {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": "Email Telah Digunakan",
		})
	}
	//Untuk Menambahkan Database yang baru
	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"Pesan": "Berhasil Membuat Database User",
		"data":  user,
	})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	//Mengambil data berdasarkan Email
	database.DB.Where("email = ? ", data["email"]).Find(&user)

	if user.IdUser == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": " Tidak Ada di dalam Database",
		})
	}

	//Compare Hass Password

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": "Password Tidak Sesuai",
		})
	} else {
		//Kita Akan Buatkan Sebuah JWT Token Barel
		claims := jwt.MapClaims{

			//Ini Berisi tentang Informasi dari Token
			"id_user": user.IdUser,
			"role":    3123,
			"exp":     time.Now().Add(time.Hour * 48).Unix(),
			"status":  "Mencoba JWT",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("Rahasia"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"Pesan": "Gagal Membuat JWT",
			})
		}
		return c.JSON(fiber.Map{
			"token": t,
		})
	}

}

func UpdateDataUser(c *fiber.Ctx)error{

	id_user, _ := c.Locals("id_user").(int)
	role, _ := c.Locals("role").(float64)
	//names, _ := c.Locals("name").(string)
	//satminkal, _ := c.Locals("satminkal").(string)

	if role != 3123 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"pesan": "Role Bukan Super Admin ",
		})
	}
	if id_user == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Admin tidak terdaftar ",
		})
	}
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//untuk update data 

	var user models.User

	//ambil data untuk model user 
	database.DB.Where("id_user = ? ", id_user).Find(&user)

	//inisiasi apa saja yang akan di update 

	update:= models.User{
		NamaLengkap: data["nama_lengkap"],
	}

	//update data baru 
	database.DB.Model(&user).Updates(&update)



	return c.JSON(fiber.Map{
		"pesan":"Sukses",
		"data":user,
	})
}

// Kita Buat MiddleWare Untuk Memakai JWT Token yang di buat

// Update Data User

// Delete Data User

//Kita Akses Server VPS

//Konviguasi Nginx Dasar
