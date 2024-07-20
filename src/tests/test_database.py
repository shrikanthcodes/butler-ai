import unittest
from services.databaseService import DatabaseService
from config.sqlConfig import SQLConfig
from config.schemas import Tables
import os


class TestDatabaseService(unittest.TestCase):
    """
    Test suite for the DatabaseService class.
    """
    @classmethod
    def setUpClass(cls):
        """
        Set up the test environment by creating a temporary database and initializing the DatabaseService.
        """
        cls.test_db_path = 'test_butler.db'
        cls.db_service = DatabaseService(cls.test_db_path)
        cls.db_service.db = SQLConfig(cls.test_db_path, delete_db=True)

    @classmethod
    def tearDownClass(cls):
        """
        Tear down the test environment by removing the temporary database.
        """
        if os.path.exists(cls.test_db_path):
            os.remove(cls.test_db_path)

    def test_initialization(self):
        """
        Test that the DatabaseService initializes correctly with a SQLConfig instance.
        """
        self.assertIsInstance(self.db_service.db, SQLConfig)

    def test_create_tables(self):
        """
        Test that the create_tables method creates the required tables.
        """
        self.db_service.create_tables()
        for table in Tables.values():
            result = self.db_service.db.query(
                f"SELECT name FROM sqlite_master WHERE type='table' AND name='{table.name}';")
            self.assertEqual(len(result), 1)
            self.assertEqual(result[0][0], table.name)

    def test_add_user(self):
        """
        Test that a user is added to the users table.
        """
        user_id = self.db_service.add_user(name='test_user')
        result = self.db_service.db.query(
            "SELECT name FROM users WHERE user_id = ?;", (user_id,))
        self.assertEqual(result[0][0], 'test_user')

    def test_add_item(self):
        """
        Test that an item is added to the items table.
        """
        item_id = self.db_service.add_items(
            name='Tomato', category='Fruit', calories=250)
        result = self.db_service.db.query(
            "SELECT name FROM items WHERE item_id = ?;", (item_id,))
        self.assertEqual(result[0][0], 'Tomato')

    def test_add_recipe(self):
        """
        Test that a recipe is added to the recipes table.
        """
        recipe_id = self.db_service.add_recipe(
            name="Dessert", ingredients="Milk", instructions="Add Something Sweet")
        result = self.db_service.db.query(
            "SELECT name FROM recipes WHERE recipe_id = ?;", (recipe_id,))
        self.assertEqual(result[0][0], 'Dessert')

    def test_store_conversation(self):
        """
        Test that a conversation is stored in the conversations table.
        """
        user_id = self.db_service.add_user(name='test_user')
        self.db_service.store_conversation(user_id, 'test_chat_history')
        result = self.db_service.db.query(
            "SELECT chat_history FROM conversations WHERE user_id = ?;", (user_id,))
        self.assertEqual(result[0][0], 'test_chat_history')

    def test_fetch_conversation_chat_history(self):
        """
        Test that a conversation chat history is fetched by conversation ID.
        """
        user_id = self.db_service.add_user(name='test_user')
        self.db_service.store_conversation(user_id, 'test_chat_history')
        conversation_id = self.db_service.db.last_insert_id()
        chat_history = self.db_service.fetch_conversation_chat_history(
            conversation_id)
        self.assertEqual(chat_history, 'test_chat_history')

    def test_get_conversation_chat_history_by_user_id(self):
        """
        Test that all chat histories are fetched for a given user ID.
        """
        user_id = self.db_service.add_user(name='test_user')
        self.db_service.store_conversation(user_id, 'test_chat_history_1')
        self.db_service.store_conversation(user_id, 'test_chat_history_2')
        chat_histories = self.db_service.get_conversation_chat_history_by_user_id(
            user_id)
        self.assertEqual(len(chat_histories), 2)
        self.assertIn(['test_chat_history_1'], chat_histories)
        self.assertIn(['test_chat_history_2'], chat_histories)

    def test_update_user(self):
        """
        Test that a user's name can be updated in the users table.
        """
        user_id = self.db_service.add_user(name='test_user')
        self.db_service.db.update_user(user_id, "Steve")
        result = self.db_service.db.query(
            "SELECT name FROM users WHERE user_id = ?;", (user_id,))
        self.assertEqual(result[0][0], 'updated_user')

    def test_update_item(self):
        """
        Test that an item's name can be updated in the items table.
        """
        item_id = self.db_service.add_items(
            name='Strawberries', category='Fruit', calories=150)
        self.db_service.db.update_item(item_id, "Bananas")
        result = self.db_service.db.query(
            "SELECT name FROM items WHERE item_id = ?;", (item_id,))
        self.assertEqual(result[0][0], 'Bananas')

    def test_update_conversation(self):
        """
        Test that a conversation's chat_history can be updated in the conversations table.
        """
        conversation_id = self.db_service.add_conversation(
            chat_history="Hey, How you doin'?")
        self.db_service.db.update_item(conversation_id, "I am doin' good")
        result = self.db_service.db.query(
            "SELECT chat_history FROM conversations WHERE conversation_id = ?;", (conversation_id,))
        self.assertEqual(result[0][0], "I am doin' good")

    def test_update_recipe(self):
        """
        Test that a recipe's name can be updated in the recipes table
        """
        recipe_id = self.db_service.add_recipe(
            name="Dessert", ingredients="Chocolate, Ice Cream", instructions="Put Chocolate on top of the ice cream")
        self.db_service.db.update_item(recipe_id, "Chocolate Ice Cream")
        result = self.db_service.db.query(
            "SELECT name FROM recipes WHERE recipe_id = ?;", (recipe_id,))
        self.assertEqual(result[0][0], "Chocolate Ice Cream")


if __name__ == '__main__':
    unittest.main()
