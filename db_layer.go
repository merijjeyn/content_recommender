package main

import (
	"fmt"
)

func insertContent(content Content) error {
	_, err := db.Exec("INSERT INTO contents (title, url) VALUES ($1, $2)", content.Title, content.Url)
	if err != nil {
		log("Failed adding content into database: " + err.Error())
		return fmt.Errorf("Failed adding content into database: %w", err)
	}

	log("Succesfully added content: " + content.Title + " " + content.Url)
	return nil
}

func insertUser(user User) error {
	_, err := db.Exec("INSERT INTOo users (notion_id, username) VALUES ($1, $2)", user.NotionId, user.Username)
	if err != nil {
		log("Failed adding user into database: " + err.Error())
		return fmt.Errorf("Failed adding user into database: %w", err)
	}

	log("Succesfully added user to database: " + user.NotionId + " " + user.Username)
	return nil
}

func insertEdge(edge Edge) error {
	_, err := db.Exec("INSERT INTO edges (user_id, content_id, weight) VALUES ($1, $2, $3)", edge.UserId, edge.ContentId, edge.Weight)
	if err != nil {
		log("Failed adding edge into database: " + err.Error())
		return fmt.Errorf("Failed adding edge into database: %w", err)
	}

	log("Succesfully added edge to database: " + string(edge.UserId) + " " + string(edge.ContentId))
	return nil
}
