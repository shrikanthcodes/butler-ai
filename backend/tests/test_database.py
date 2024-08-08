import unittest
from services.database_service import DatabaseService
from config.sql_config import SQLConfig
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
        cls.db_service.create_tables(drop_tables=True)

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
        self.db_service.add_conversation(user_id, 'test_chat_history')
        result = self.db_service.db.query(
            "SELECT chat_history FROM conversations WHERE user_id = ?;", (user_id,))
        self.assertEqual(result[0][0], 'test_chat_history')

    def test_fetch_conversation_chat_history(self):
        """
        Test that a conversation chat history is fetched by conversation ID.
        """
        user_id = self.db_service.add_user(name='test_user')
        self.db_service.add_conversation(user_id, 'test_chat_history')
        conversation_id = self.db_service.db.last_insert_id()
        chat_history = self.db_service.get_conversation_chat_history_by_id(
            conversation_id)
        self.assertEqual(chat_history, 'test_chat_history')

    def test_get_conversation_chat_history_by_user_id(self):
        """
        Test that all chat histories are fetched for a given user ID.
        """
        user_id = self.db_service.add_user(name='test_user')
        test_string = f"system::Hi!, How are you?;;user::GoodBye!;;System::Have a nice day!"
        test_update = f"system::Hi!, How are you?;;user::Nothing;;System::Have a bad day!"
        self.db_service.add_conversation(user_id, test_string)
        self.db_service.add_conversation(user_id, test_update)
        chat_histories = self.db_service.get_conversation_chat_history_by_user_id(
            user_id)
        self.assertEqual(len(chat_histories), 1)
        self.assertIn(
            [{'role': 'system', 'content': 'Hi!, How are you?'},
             {'role': 'user', 'content': 'GoodBye!'},
             {'role': 'System', 'content': 'Have a nice day!'},
             {'role': 'system', 'content': 'Hi!, How are you?'},
             {'role': 'user', 'content': 'Nothing'},
             {'role': 'System', 'content': 'Have a bad day!'}],
            chat_histories)

    def test_update_user(self):
        """
        Test that a user's name can be updated in the users table.
        """
        user_id = self.db_service.add_user(name='test_user')
        self.db_service.update_user(user_id, name="updated_user")
        result = self.db_service.db.query(
            "SELECT name FROM users WHERE user_id = ?;", (user_id,))
        self.assertEqual(result[0][0], 'updated_user')

    def test_update_item(self):
        """
        Test that an item's name can be updated in the items table.
        """
        item_id = self.db_service.add_items(
            name='Strawberries', category='Fruit', calories=150)
        self.db_service.update_item(item_id, name="Bananas")
        result = self.db_service.db.query(
            "SELECT name FROM items WHERE item_id = ?;", (item_id,))
        self.assertEqual(result[0][0], 'Bananas')

    def test_update_conversation(self):
        """
        Test that a conversation's chat_history can be updated in the conversations table.
        """
        user_id = self.db_service.add_user(name='test_user')
        test_string = f"system::Hi!, How are you?;;user::GoodBye!;;System::Have a nice day!"
        test_update = f"system::Hi!, How are you?;;user::Nothing;;System::Have a bad day!"
        conversation_id = self.db_service.add_conversation(
            user_id=user_id, chat_history=test_string)
        self.db_service.update_conversation(
            conversation_id, new_conversation=test_update)
        result = self.db_service.db.query(
            "SELECT chat_history FROM conversations WHERE conversation_id = ?;", (conversation_id,))
        self.assertEqual(result[0][0], test_update)

    def test_update_recipe(self):
        """
        Test that a recipe's name can be updated in the recipes table.
        """
        recipe_id = self.db_service.add_recipe(
            name="Dessert", ingredients="Chocolate, Ice Cream", instructions="Put Chocolate on top of the ice cream")
        self.db_service.update_recipe(recipe_id, name="Chocolate Ice Cream")
        result = self.db_service.db.query(
            "SELECT name FROM recipes WHERE recipe_id = ?;", (recipe_id,))
        self.assertEqual(result[0][0], "Chocolate Ice Cream")


if __name__ == '__main__':
    unittest.main()
