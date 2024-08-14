package config

type Table struct {
	Name               string   // Table name
	Columns            []string // List of column names
	DbName             string   // Database name
	SecurityEncryption []string // Columns requiring end-to-end encryption
	PIIEncryption      []string // Columns requiring encryption at storage
}

var Tables = map[string]Table{
	// Authentication DB Tables
	"users": {
		Name:               "users",
		Columns:            []string{"user_id", "passwd", "last_updated"},
		DbName:             "auth_db",
		SecurityEncryption: []string{"user_id", "passwd"}, // Security-based encryption
		PIIEncryption:      []string{},                    // No PII encryption needed
	},
	"refresh": {
		Name:               "refresh",
		Columns:            []string{"refresh_token", "user_id", "user_agent", "last_login", "last_logout", "expires_at"},
		DbName:             "auth_db",
		SecurityEncryption: []string{"refresh_token", "user_id"}, // Security-based encryption
		PIIEncryption:      []string{},                           // No PII encryption needed
	},
	"llm": {
		Name:               "llm",
		Columns:            []string{"user_id", "llm_choice", "llm_version", "llm_token"},
		DbName:             "auth_db",
		SecurityEncryption: []string{"user_id", "llm_token"}, // Security-based encryption
		PIIEncryption:      []string{},                       // No PII encryption needed
	},
	"integrations": {
		Name:               "integrations",
		Columns:            []string{"user_id", "meta", "meta_token", "meta_expires_at", "google", "google_token", "google_expires_at", "twitter", "twitter_token", "twitter_expires_at"},
		DbName:             "auth_db",
		SecurityEncryption: []string{"user_id", "meta_token", "google_token", "twitter_token"}, // Security-based encryption
		PIIEncryption:      []string{},                                                         // No PII encryption needed
	},

	// User DB Tables
	"user_profile": {
		Name:               "user_profile",
		Columns:            []string{"user_id", "first_name", "last_name", "email", "phone", "age", "gender", "weight", "height"},
		DbName:             "user_db",
		SecurityEncryption: []string{"user_id"},                                                                        // Security-based encryption
		PIIEncryption:      []string{"first_name", "last_name", "email", "phone", "age", "gender", "weight", "height"}, // PII encryption
	},
	"health": {
		Name:               "health",
		Columns:            []string{"user_id", "health_conditions", "medications", "allergies", "dietary_restrictions"},
		DbName:             "user_db",
		SecurityEncryption: []string{"user_id"},                                                               // Security-based encryption
		PIIEncryption:      []string{"health_conditions", "medications", "allergies", "dietary_restrictions"}, // PII encryption
	},
	"conversations": {
		Name:               "conversations",
		Columns:            []string{"conversation_id", "user_id", "chat_history", "last_updated"},
		DbName:             "user_db",
		SecurityEncryption: []string{"user_id"},                         // Security-based encryption
		PIIEncryption:      []string{"conversation_id", "chat_history"}, // PII encryption
	},
	"preferences": {
		Name:               "preferences",
		Columns:            []string{"user_id", "favorite_recipes", "disliked_recipes", "favorite_items", "disliked_items", "favorite_categories", "disliked_categories"},
		DbName:             "user_db",
		SecurityEncryption: []string{"user_id"}, // Security-based encryption
		PIIEncryption:      []string{},          // No PII encryption needed
	},
	"inventory": {
		Name:               "inventory",
		Columns:            []string{"user_id", "items"},
		DbName:             "user_db",
		SecurityEncryption: []string{"user_id"}, // Security-based encryption
		PIIEncryption:      []string{},          // No PII encryption needed
	},

	// Butler DB Tables
	"recipes": {
		Name:               "recipes",
		Columns:            []string{"recipe_id", "name", "category", "cuisine", "ingredients", "instructions", "nutritional_info", "user_id", "recipe_html"},
		DbName:             "butler_db",
		SecurityEncryption: []string{"user_id"},   // Security-based encryption
		PIIEncryption:      []string{"recipe_id"}, // PII encryption
	},
}

var Schemas = map[string]string{
	// Authentication DB Schemas
	"users": `
	CREATE TABLE IF NOT EXISTS users (
		user_id VARBINARY(256) PRIMARY KEY,
		passwd VARBINARY(256),
		last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`,
	"refresh": `
	CREATE TABLE IF NOT EXISTS refresh (
		refresh_token VARBINARY(256) PRIMARY KEY,
		user_id VARBINARY(256),
		user_agent TEXT,
		last_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		last_logout TIMESTAMP,
		expires_at TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`,
	"llm": `
	CREATE TABLE IF NOT EXISTS llm (
		user_id VARBINARY(256),
		llm_choice ENUM('openai', 'gemini', 'llama3'),
		llm_version TEXT,
		llm_token VARBINARY(256),
		PRIMARY KEY (user_id),
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`,
	"integrations": `
	CREATE TABLE IF NOT EXISTS integrations (
		user_id VARBINARY(256) PRIMARY KEY,
		meta BOOLEAN DEFAULT FALSE,
		meta_token VARBINARY(256),
		meta_expires_at TIMESTAMP,
		google BOOLEAN DEFAULT FALSE,
		google_token VARBINARY(256),
		google_expires_at TIMESTAMP,
		twitter BOOLEAN DEFAULT FALSE,
		twitter_token VARBINARY(256),
		twitter_expires_at TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`,

	// User DB Schemas
	"user_profile": `
	CREATE TABLE IF NOT EXISTS user_profile (
		user_id VARBINARY(256) PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		email TEXT,
		phone TEXT,
		age INTEGER,
		gender TEXT,
		weight REAL,
		height REAL
	);`,
	"health": `
	CREATE TABLE IF NOT EXISTS health (
		user_id VARBINARY(256) PRIMARY KEY,
		health_conditions TEXT,
		medications TEXT,
		allergies TEXT,
		dietary_restrictions TEXT,
		FOREIGN KEY (user_id) REFERENCES user_profile(user_id) ON DELETE CASCADE
	);`,
	"conversations": `
	CREATE TABLE IF NOT EXISTS conversations (
		conversation_id VARBINARY(256) PRIMARY KEY,
		user_id VARBINARY(256),
		chat_history TEXT,
		last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES user_profile(user_id) ON DELETE CASCADE
	);`,
	"preferences": `
	CREATE TABLE IF NOT EXISTS preferences (
		user_id VARBINARY(256) PRIMARY KEY,
		favorite_recipes TEXT,
		disliked_recipes TEXT,
		favorite_items TEXT,
		disliked_items TEXT,
		favorite_categories TEXT,
		disliked_categories TEXT,
		FOREIGN KEY (user_id) REFERENCES user_profile(user_id) ON DELETE CASCADE
	);`,
	"inventory": `
	CREATE TABLE IF NOT EXISTS inventory (
		user_id VARBINARY(256) PRIMARY KEY,
		items JSON,
		FOREIGN KEY (user_id) REFERENCES user_profile(user_id) ON DELETE CASCADE
	);`,

	// Butler DB Schemas
	"recipes": `
	CREATE TABLE IF NOT EXISTS recipes (
		recipe_id VARBINARY(256) PRIMARY KEY,
		name TEXT,
		category TEXT,
		cuisine TEXT,
		ingredients TEXT,
		instructions TEXT,
		nutritional_info TEXT,
		user_id VARBINARY(256),
		recipe_html TEXT
	);`,
}
