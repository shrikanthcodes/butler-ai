class Table:
    def __init__(self, name, columns):
        self.name = name
        self.columns = columns


Tables = {
    "users": Table("users", ["user_id", "name"]),
    "items": Table("items", ["item_id", "name", "category", "calories"]),
    "recipes": Table("recipes", ["recipe_id", "name", "ingredients", "instructions"]),
    "conversations": Table("conversations", ["conversation_id", "user_id", "chat_history"])
}


class Schemas:
    conversations = """
    conversation_id INTEGER PRIMARY KEY,
    user_id INTEGER,
    chat_history TEXT,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
    """

    users = """
    user_id INTEGER PRIMARY KEY,
    name TEXT
    """

    items = """
    item_id INTEGER PRIMARY KEY,
    name TEXT,
    category TEXT,
    calories INTEGER
    """

    recipes = """
    recipe_id INTEGER PRIMARY KEY,
    name TEXT,
    ingredients TEXT,
    instructions TEXT
    """
