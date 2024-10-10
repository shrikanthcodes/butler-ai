package service

// Given user's email, and password, check if the email exists in users table, if it doesnt,
// then return an error for no user found. If the user is found, then check if the password
// matches the hashed password in the database. If it does, return the userID, else return a different error.
