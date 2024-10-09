package handlers

// import (
// 	db "backend/pkg/services/db"
// 	http "net/http"

// 	gin "github.com/gin-gonic/gin"
// )

// func GetUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	user, err := db.GetUserByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Return the user information in the response
// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }

// func CreateUser(c *gin.Context) {
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")

// 	// Assuming the service layer handles DB interaction or business logic
// 	user, err := db.CreateUser(name, email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	// Return the created user information in the response
// 	c.JSON(http.StatusCreated, gin.H{"user": user})
// }

// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")

// 	// Assuming the service layer handles DB interaction or business logic
// 	user, err := db.UpdateUser(id, name, email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 		return
// 	}

// 	// Return the updated user information in the response
// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }

// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.DeleteUser(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
// 		return
// 	}

// 	// Return the deletion message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserConversations(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	conversations, err := db.GetUserConversations(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Conversations not found"})
// 		return
// 	}

// 	// Return the conversations associated with the user in the response
// 	c.JSON(http.StatusOK, gin.H{"conversations": conversations})
// }

// func GetUserProfile(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	profile, err := db.GetUserProfile(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
// 		return
// 	}

// 	// Return the user profile information in the response
// 	c.JSON(http.StatusOK, gin.H{"profile": profile})
// }

// func UpdateUserProfile(c *gin.Context) {
// 	id := c.Param("id")
// 	profile := c.PostForm("profile")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserProfile(id, profile)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserHealth(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	health, err := db.GetUserHealth(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User health information not found"})
// 		return
// 	}

// 	// Return the user health information in the response
// 	c.JSON(http.StatusOK, gin.H{"health": health})
// }

// func UpdateUserHealth(c *gin.Context) {
// 	id := c.Param("id")
// 	health := c.PostForm("health")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserHealth(id, health)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user health information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserDiet(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	diet, err := db.GetUserDiet(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User diet information not found"})
// 		return
// 	}

// 	// Return the user diet information in the response
// 	c.JSON(http.StatusOK, gin.H{"diet": diet})
// }

// func UpdateUserDiet(c *gin.Context) {
// 	id := c.Param("id")
// 	diet := c.PostForm("diet")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserDiet(id, diet)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user diet information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserInventory(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	inventory, err := db.GetUserInventory(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User inventory information not found"})
// 		return
// 	}

// 	// Return the user inventory information in the response
// 	c.JSON(http.StatusOK, gin.H{"inventory": inventory})
// }

// func UpdateUserInventory(c *gin.Context) {
// 	id := c.Param("id")
// 	inventory := c.PostForm("inventory")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserInventory(id, inventory)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user inventory information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserGoal(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	goal, err := db.GetUserGoal(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User goal information not found"})
// 		return
// 	}

// 	// Return the user goal information in the response
// 	c.JSON(http.StatusOK, gin.H{"goal": goal})
// }

// func UpdateUserGoal(c *gin.Context) {
// 	id := c.Param("id")
// 	goal := c.PostForm("goal")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserGoal(id, goal)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user goal information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserLLM(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	llm, err := db.GetUserLLM(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User LLM information not found"})
// 		return
// 	}

// 	// Return the user LLM information in the response
// 	c.JSON(http.StatusOK, gin.H{"llm": llm})
// }

// func UpdateUserLLM(c *gin.Context) {
// 	id := c.Param("id")
// 	llm := c.PostForm("llm")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserLLM(id, llm)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user LLM information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserScript(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	script, err := db.GetUserScript(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User script information not found"})
// 		return
// 	}

// 	// Return the user script information in the response
// 	c.JSON(http.StatusOK, gin.H{"script": script})
// }

// func UpdateUserScript(c *gin.Context) {
// 	id := c.Param("id")
// 	script := c.PostForm("script")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserScript(id, script)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user script information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserShopping(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	shopping, err := db.GetUserShopping(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User shopping information not found"})
// 		return
// 	}

// 	// Return the user shopping information in the response
// 	c.JSON(http.StatusOK, gin.H{"shopping": shopping})
// }

// func UpdateUserShopping(c *gin.Context) {
// 	id := c.Param("id")
// 	shopping := c.PostForm("shopping")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserShopping(id, shopping)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user shopping information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserMealChoices(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	mealChoices, err := db.GetUserMealChoices(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User meal choices information not found"})
// 		return
// 	}

// 	// Return the user meal choices information in the response
// 	c.JSON(http.StatusOK, gin.H{"mealChoices": mealChoices})
// }

// func UpdateUserMealChoices(c *gin.Context) {
// 	id := c.Param("id")
// 	mealChoices := c.PostForm("mealChoices")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserMealChoices(id, mealChoices)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user meal choices information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetAllChoices(c *gin.Context) {
// 	// Assuming the service layer handles DB interaction or business logic
// 	choices, err := db.GetAllChoices()
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Choices not found"})
// 		return
// 	}

// 	// Return all choices in the response
// 	c.JSON(http.StatusOK, gin.H{"choices": choices})
// }

// func UpdateChoices(c *gin.Context) {
// 	choices := c.PostFormArray("choices")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateChoices(choices)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update choices"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }

// func GetUserWriteUp(c *gin.Context) {
// 	id := c.Param("id")

// 	// Assuming the service layer handles DB interaction or business logic
// 	writeUp, err := db.GetUserWriteUp(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User write-up information not found"})
// 		return
// 	}

// 	// Return the user write-up information in the response
// 	c.JSON(http.StatusOK, gin.H{"writeUp": writeUp})
// }

// func UpdateUserWriteUp(c *gin.Context) {
// 	id := c.Param("id")
// 	writeUp := c.PostForm("writeUp")

// 	// Assuming the service layer handles DB interaction or business logic
// 	message, err := db.UpdateUserWriteUp(id, writeUp)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user write-up information"})
// 		return
// 	}

// 	// Return the update message in the response
// 	c.JSON(http.StatusOK, gin.H{"message": message})
// }
