package controllers

import (
	"log"
	"lostandfounditemmanagment/database"
	"lostandfounditemmanagment/models"
	"lostandfounditemmanagment/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateReport(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Could not parse body"))
	}
	var user models.User
	token := c.Cookies("jwt")
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

		if data["item"] != nil {
			itemID, err := CreateItem(data["item"].(map[string]interface{}))
			if err != nil {
				c.Status(404)
				return c.JSON(fiber.Map{
					"message": "Could not create an item",
				})
			}
			newReport := models.Report{
				UserID:     user.Id,
				ItemID:     itemID,
				ReportDate: time.Now(),
				Status:     data["status"].(string),
			}
			database.DB.Create(&newReport)
			c.Status(200)
			return c.JSON(fiber.Map{
				"message": "Report has been successfully created!",
			})
		}
	}
	c.Status(404)
	return c.JSON(fiber.Map{
		"message": "Error has occured",
		"item":    data["item"],
	})
}

func CreateItem(data map[string]interface{}) (uint, error) {
	itemTime, err := utils.ConvertTime(data["time"].(string))
	if err != nil {
		return 0, err
	}
	newItem := models.Item{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		DateFound:   itemTime,
		WhereFound:  data["where_found"].(string),
		Status:      data["status"].(string),
	}
	database.DB.Model(&models.Item{}).Create(&newItem)

	var item models.Item
	database.DB.Where("description=?", data["description"].(string)).First(&item)
	return item.Id, nil
}

func GetUserReports(c *fiber.Ctx) error {
	token := c.Cookies("jwt")
	if token != "" {
		var user models.User

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
		var reports []models.Report

		database.DB.Where("user_id=?", user.Id).Where("status=?", "Lost").Preload("Item").Find(&reports)
		c.Status(200)
		return c.JSON(fiber.Map{
			"reports": reports,
		})
	}
	c.Status(404)
	return c.JSON(fiber.Map{
		"message": "Please sign in",
	})
}

func DeleteReport(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Could not parse body"))
	}
	var user models.User
	token := c.Cookies("jwt")
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
		database.DB.Model(models.Report{}).Where("id=?", data["report_id"]).Update("status", "Closed")
		c.Status(200)
		return c.JSON(fiber.Map{
			"message": "Report has been closed",
		})
	}
	c.Status(404)
	return c.JSON(fiber.Map{
		"message": "Please sign in",
	})
}

func GetActiveReports(c *fiber.Ctx) error {
	token := c.Cookies("jwt")
	if token != "" {
		var user models.User

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
		var reports []models.Report

		database.DB.Where("status=?", "Lost").Preload("Item").Preload("User").Find(&reports)
		c.Status(200)
		return c.JSON(fiber.Map{
			"reports": reports,
		})
	}
	c.Status(404)
	return c.JSON(fiber.Map{
		"message": "Please sign in",
	})
}
