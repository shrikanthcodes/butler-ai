package handler

// import (
// 	db "github.com/shrikanthcodes/butler-ai/backend/pkg/services/db"
// 	api "net/api"

// 	server "github.com/server-gonic/server"
// )

// func GetUserByID(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	user, err := db.GetUserByID(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User not found"})
// 		return
// 	}

// 	// Return the user information in the response
// 	c.JSON(api.StatusOK, server.H{"user": user})
// }

// func CreateUser(c *server.Context) {
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")

// 	// Assuming the service layer handles DB interaction or business logic
// 	user, err := db.CreateUser(name, email)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to create user"})
// 		return
// 	}

// 	// Return the created user information in the response
// 	c.JSON(api.StatusCreated, server.H{"user": user})
// }

// func UpdateUser(c *server.Context) {
// 	id := c.Param("id")
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")

// 	// Assuming the service layer handles DB interaction or business logic
// 	user, err := db.UpdateUser(id, name, email)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user"})
// 		return
// 	}

// 	// Return the updated user information in the response
// 	c.JSON(api.StatusOK, server.H{"user": user})
// }

// func DeleteUser(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.DeleteUser(id)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to delete user"})
// 		return
// 	}

// 	// Return the deletion message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserConversations(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	conversations, err := db.GetUserConversations(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "Conversations not found"})
// 		return
// 	}

// 	// Return the conversations associated with the user in the response
// 	c.JSON(api.StatusOK, server.H{"conversations": conversations})
// }

// func GetUserProfile(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	profile, err := db.GetUserProfile(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User profile not found"})
// 		return
// 	}

// 	// Return the user profile information in the response
// 	c.JSON(api.StatusOK, server.H{"profile": profile})
// }

// func UpdateUserProfile(c *server.Context) {
// 	id := c.Param("id")
// 	profile := c.PostForm("profile")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserProfile(id, profile)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user profile"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserHealth(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	health, err := db.GetUserHealth(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User health information not found"})
// 		return
// 	}

// 	// Return the user health information in the response
// 	c.JSON(api.StatusOK, server.H{"health": health})
// }

// func UpdateUserHealth(c *server.Context) {
// 	id := c.Param("id")
// 	health := c.PostForm("health")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserHealth(id, health)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user health information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserDiet(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	diet, err := db.GetUserDiet(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User diet information not found"})
// 		return
// 	}

// 	// Return the user diet information in the response
// 	c.JSON(api.StatusOK, server.H{"diet": diet})
// }

// func UpdateUserDiet(c *server.Context) {
// 	id := c.Param("id")
// 	diet := c.PostForm("diet")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserDiet(id, diet)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user diet information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserInventory(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	inventory, err := db.GetUserInventory(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User inventory information not found"})
// 		return
// 	}

// 	// Return the user inventory information in the response
// 	c.JSON(api.StatusOK, server.H{"inventory": inventory})
// }

// func UpdateUserInventory(c *server.Context) {
// 	id := c.Param("id")
// 	inventory := c.PostForm("inventory")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserInventory(id, inventory)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user inventory information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserGoal(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	goal, err := db.GetUserGoal(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User goal information not found"})
// 		return
// 	}

// 	// Return the user goal information in the response
// 	c.JSON(api.StatusOK, server.H{"goal": goal})
// }

// func UpdateUserGoal(c *server.Context) {
// 	id := c.Param("id")
// 	goal := c.PostForm("goal")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserGoal(id, goal)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user goal information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserLLM(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	llm, err := db.GetUserLLM(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User LLM information not found"})
// 		return
// 	}

// 	// Return the user LLM information in the response
// 	c.JSON(api.StatusOK, server.H{"llm": llm})
// }

// func UpdateUserLLM(c *server.Context) {
// 	id := c.Param("id")
// 	llm := c.PostForm("llm")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserLLM(id, llm)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user LLM information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserScript(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	script, err := db.GetUserScript(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User script information not found"})
// 		return
// 	}

// 	// Return the user script information in the response
// 	c.JSON(api.StatusOK, server.H{"script": script})
// }

// func UpdateUserScript(c *server.Context) {
// 	id := c.Param("id")
// 	script := c.PostForm("script")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserScript(id, script)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user script information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserShopping(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	shopping, err := db.GetUserShopping(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User shopping information not found"})
// 		return
// 	}

// 	// Return the user shopping information in the response
// 	c.JSON(api.StatusOK, server.H{"shopping": shopping})
// }

// func UpdateUserShopping(c *server.Context) {
// 	id := c.Param("id")
// 	shopping := c.PostForm("shopping")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserShopping(id, shopping)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user shopping information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserMealChoices(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	mealChoices, err := db.GetUserMealChoices(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User meal choices information not found"})
// 		return
// 	}

// 	// Return the user meal choices information in the response
// 	c.JSON(api.StatusOK, server.H{"mealChoices": mealChoices})
// }

// func UpdateUserMealChoices(c *server.Context) {
// 	id := c.Param("id")
// 	mealChoices := c.PostForm("mealChoices")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserMealChoices(id, mealChoices)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user meal choices information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetAllChoices(c *server.Context) {
// 	// Assuming the service layer handles DB interaction or business logic
// 	choices, err := db.GetAllChoices()
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "Choices not found"})
// 		return
// 	}

// 	// Return all choices in the response
// 	c.JSON(api.StatusOK, server.H{"choices": choices})
// }

// func UpdateChoices(c *server.Context) {
// 	choices := c.PostFormArray("choices")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateChoices(choices)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update choices"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }

// func GetUserWriteUp(c *server.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	writeUp, err := db.GetUserWriteUp(id)
// 	if err != nil {
// 		c.JSON(api.StatusNotFound, server.H{"error": "User write-up information not found"})
// 		return
// 	}

// 	// Return the user write-up information in the response
// 	c.JSON(api.StatusOK, server.H{"writeUp": writeUp})
// }

// func UpdateUserWriteUp(c *server.Context) {
// 	id := c.Param("id")
// 	writeUp := c.PostForm("writeUp")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserWriteUp(id, writeUp)
// 	if err != nil {
// 		c.JSON(api.StatusInternalServerError, server.H{"error": "Failed to update user write-up information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(api.StatusOK, server.H{"message": message})
// }
