package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func addContentHandler(c *fiber.Ctx) error {
	content := Content{Title: c.Query("title"), Url: c.Query("url")}
	err := insertContent(content)

	if err != nil {
		return err
	}
	return c.SendString("Success adding content")
}

func addUserHandler(c *fiber.Ctx) error {
	user := User{Username: c.Query("username"), NotionId: c.Query("notion_id")}
	err := insertUser(user)

	if err != nil {
		return err
	}
	return c.SendString("Success adding user")
}

func addEdgeHandler(c *fiber.Ctx) error {
	user_id, err_uid := strconv.Atoi(c.Query("user_id"))
	if err_uid != nil {
		log("Failed converting user_id to int from request: " + err_uid.Error())
	}

	content_id, err_cid := strconv.Atoi(c.Query("content_id"))
	if err_uid != nil {
		log("Failed converting content_id to int from request: " + err_cid.Error())
	}

	weight := 1
	err := insertEdge(Edge{UserId: user_id, ContentId: content_id, Weight: float64(weight)})
	if err != nil {
		return err
	}
	return c.SendString("Success adding edge")
}

func recommendContentHandler(c *fiber.Ctx) error {
	return c.SendString("Success")
}
