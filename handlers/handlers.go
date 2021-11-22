package handlers

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/redisApi/models"
	"github.com/zeimedee/redisApi/services"
)

func Login(c *fiber.Ctx) error {
	ctx := context.Background()

	student := new(models.Student)
	err := c.BodyParser(student)
	services.Check(err)

	json, err := json.Marshal(student)
	services.Check(err)

	setStudent := services.RD.Client.Set(ctx, student.ID, json, 0).Err()
	services.Check(setStudent)
	return c.Status(200).JSON("SENT")
}

func Student(c *fiber.Ctx) error {
	ctx := context.Background()

	student := new(models.Student)
	err := c.BodyParser(student)
	services.Check(err)

	getStudent, err := services.RD.Client.Get(ctx, string(student.ID)).Result()
	services.Check(err)

	json.Unmarshal([]byte(getStudent), &student)
	return c.Status(200).JSON(student)
}
