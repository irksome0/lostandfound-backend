package controllers

import (
	"log"
	"lostandfounditemmanagment/database"
	"lostandfounditemmanagment/models"
	"lostandfounditemmanagment/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	var data map[string]interface{}
	// var userData models.User
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal("Unable to parse body")
	}

	user := models.User{
		FullName:    strings.TrimSpace(data["full_name"].(string)),
		PhoneNumber: strings.TrimSpace(data["phoneNumber"].(string)),
	}
	user.SetPassword(data["password"].(string))

	database.DB.Create(&user)

	ctx.Status(200)
	ctx.JSON(fiber.Map{
		"message": "User has been successfully created!",
		"user":    user,
		"role":    "user",
	})

	return nil
}
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Unable to parse body!(Login)"))
	}
	var user models.User

	database.DB.Where("phone_number=?", data["phoneNumber"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User with such phone number does not exist!",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password!",
		})
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(500)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: "None",
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message":      "Successful login!",
		"access_token": token,
		"user":         user,
	})
}

func GetUser(c *fiber.Ctx) error {
	token := c.Cookies("jwt")
	var user models.User

	if token != "" {
		issuer, err := utils.ParseJwt(token)
		if err != nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Could not parse jwt",
			})
		}

		database.DB.Where("id=?", issuer).First(&user)
		if user.Id == 0 {
			c.Status(404)
			return c.JSON(fiber.Map{
				"message": "User was not found!",
			})
		}

		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Session has been confirmed",
			"user":    user,
		})
	}
	c.Status(400)
	return c.JSON(fiber.Map{
		"message": "Session was not confirmed",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Unable to parse body!(Login)"))
	}
	token := c.Cookies("jwt")
	var user models.User

	if token != "" {
		issuer, err := utils.ParseJwt(token)
		if err != nil {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Could not parse jwt",
			})
		}

		database.DB.Where("id=?", issuer).First(&user)
		if user.Id == 0 {
			c.Status(404)
			return c.JSON(fiber.Map{
				"message": "User was not found!",
			})
		}
		database.DB.Model(models.User{}).Where("id=?", issuer).Update("full_name", data["full_name"])
		database.DB.Model(models.User{}).Where("id=?", issuer).Update("phone_number", data["phone_number"])

		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Data has been updated",
		})
	}
	c.Status(400)
	return c.JSON(fiber.Map{
		"message": "Session was not confirmed",
	})
}
