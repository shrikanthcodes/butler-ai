from SQLService import SQLService
import DatabaseConstants as dc
from Schemas import Schemas, Tables
from Formatting import convert_chat_to_list


class DatabaseService:
    def __init__(self):
        # TODO: Logic for user_count, user_id, conversation_count, conversation_id must be improved
        self.counts = {"user": 0,
                       "conversation": 0,
                       "recipe": 0,
                       "item": 0}
        self.db = SQLService(dc.butler_db, delete_db=False)
        self.chat_history = []

    def create_tables(self):
        for table in Tables.values():
            self.db.create(table.name, Schemas.__dict__[
                table.name], drop_table=True)
        # Logging
        print("The following tables were created successfully: " +
              "\n".join(Tables.keys()))

    # TODO: Error handling for already existing user
    def add_user(self, id=None, name="admin"):
        if id == None:
            user_id = self.counts["user"]
            self.counts["user"] += 1
            print("User ID: None and ", user_id)
        else:
            user_id = id
            print("User ID: Not None and ", user_id)
        values = (user_id, name)
        self.db.insert(Tables["users"].name,
                       Tables["users"].columns, values)
        # Logging
        print("Added user_id: {} for user_name: {}".format(
            user_id, name))
        return user_id

    def add_conversation(self, id=None, user_id=0, conversation_history=""):
        if id == None:
            conversation_id = self.counts["conversation"]
            self.counts["conversation"] += 1
        else:
            conversation_id = id
        values = (conversation_id, user_id, conversation_history)
        self.db.insert(Tables["conversations"].name,
                       Tables["conversations"].columns, values)
        # Logging
        print("Added conversation_id: {} for user id: {}".format(
            conversation_id, user_id))
        return conversation_id

    def add_recipe(self, name, ingredients="", instructions=""):
        recipe_id = self.counts["recipe"]
        values = (recipe_id, name, ingredients, instructions)
        self.db.insert(Tables["recipes"].name,
                       Tables["recipes"].columns, values)
        self.counts["recipe"] += 1
        # Logging
        print("Added recipe_id: {} for recipe_name: {}".format(
            recipe_id, name))
        return recipe_id

    def add_items(self, name, category="", calories=0):
        item_id = self.counts["item"]
        values = (item_id, name, category, calories)
        self.db.insert(Tables["items"].name,
                       Tables["items"].columns, values)
        self.counts["item"] += 1
        # Logging
        print("Added item_id: {} for item_name: {}".format(
            item_id, name))
        return item_id

    def get_conversation(self, conversation_id=None, user_id=None):
        if conversation_id is None and user_id is None:
            return "No conversation found. Need a valid conversation_id or user_id."

        if conversation_id:
            result = self.db.query(
                dc.get_conversation_by_id, (conversation_id,))
        elif user_id:
            result = self.db.query(dc.get_conversation_by_user_id, (user_id,))

        if result:
            self.chat_history = [row[0]
                                 for row in result] if user_id else result[0][0]
        else:
            return "No Conversation Found"

        return self.chat_history

    def update_conversation(self, conversation_id, new_conversation):
        chat_history = self.get_conversation(
            conversation_id=conversation_id)
        if chat_history is "":
            updated_chat_history = new_conversation
        else:
            updated_chat_history = chat_history + ";;" + new_conversation
        self.db.update(Tables["conversations"].name, "chat_history",
                       updated_chat_history, "conversation_id", conversation_id)
        # Logging
        print("Updated conversation_id: {}".format(conversation_id))
        return convert_chat_to_list(updated_chat_history)

    # TODO: Functions for get/update recipes, items and users must be added
