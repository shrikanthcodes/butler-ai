class Table:
    """
    A class representing a database table.

    Attributes:
        name (str): The name of the table.
        columns (list): A list of column names in the table.
    """

    def __init__(self, name, columns):
        """
        Initialize a Table instance.

        Args:
            name (str): The name of the table.
            columns (list): A list of column names in the table.
        """
        self.name = name
        self.columns = columns


# A dictionary of Table instances representing the tables in the database.
"""DB
1) Authentication (encrypted)
    a) user table : user_id (Primary, not autoincrement, unique but hashed), passwd (hashed), last_updated
    b) refresh table : refresh token (unique, hashed), user_id, user_agent, last_login, last_logout, expires_at
    c) llm table: user_id, llm_choice, llm_version, llm_token
    d) integrations table: user_id, meta (bool), meta_token, meta_expires_at, google (bool), google_token, google_expires_at, twitter (bool), twitter_token, twitter_expires_at

enum: openai (gpt-3.5-turbo, gpt-4o, gpt-4), gemini (gemini-1.5-pro, gemini-1.0), llama3

2) User
    a) user table : user_id (Same ID as Authentication.User table), first_name, last_name, email, phone, age, gender, weight, height
    b) health table : user_id, health_conditions, medications, allergies, dietary_restrictions
    c) conversation table : conversation_id (Primary (hashed)), user_id, chat_history, last_updated
    d) preference table : user_id, favorite_recipes, disliked_recipes, favorite_items, disliked_items, favorite_categories, disliked_categories
    e) inventory table : user_id, items (JSON format: {"item_id": "count", "count_unit"})

3) Items
    from USDA

4) Butler
    a) recipes table : recipe_id (Primary (hashed)), name, category, cuisine, ingredients, instructions, nutritional_info, user_id (optional), recipe_html

    """


Tables = {
    "users": Table("users", ["user_id", "name"]),
    "items": Table("items", ["item_id", "name", "category", "calories"]),
    "recipes": Table("recipes", ["recipe_id", "name", "ingredients", "instructions"]),
    "conversations": Table("conversations", ["conversation_id", "user_id", "chat_history"])
}


class Schemas:
    """
    A class containing the schema definitions for the database tables.

    Each schema is a string defining the SQL CREATE TABLE statement for the corresponding table.

    conversations table columns: conversation_id, user_id, chat_history.
    users table columns: user_id, name.
    items table columns: item_id, name, category, calories.
    recipes table columns: recipe_id, name, ingredients, instructions.
    """

    conversations = """
    conversation_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    chat_history TEXT,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
    """

    users = """
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT
    """

    items = """
    item_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    category TEXT,
    calories INTEGER
    """

    recipes = """
    recipe_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    ingredients TEXT,
    instructions TEXT
    """
