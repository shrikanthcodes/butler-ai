from SQLService import SQLService, Error
from Schemas import Schemas, Tables
from Formatting import format_message_for_storage, convert_chat_to_list
from errorHandling import ErrorHandler, logger


class DatabaseService:
    """
    A service class for managing database interactions.

    Attributes:
        db (SQLService): The database service instance.
    """

    def __init__(self):
        """
        Initialize the DatabaseService.

        Initializes the database service.
        """
        self.db = SQLService("butler.db", delete_db=False)

    def create_tables(self, drop_tables=False):
        """
        Create database tables.

        Args:
            drop_tables (bool): If True, drop existing tables before creating new ones.
        """
        for table in Tables.values():
            self.db.create(table.name, Schemas.__dict__[
                           table.name], drop_tables)
        logger.info("The following tables were created successfully:\n" +
                    "\n".join(Tables.keys()))

    def add_entity(self, entity_type, **kwargs):
        """
        Add an entity to the database.

        Args:
            entity_type (str): The type of entity (e.g., user, conversation, recipe, item).
            **kwargs: Additional attributes for the entity.

        Raises:
            ErrorHandler.EntityAlreadyExistsError: If the entity with the given ID already exists.

        Returns:
            int: The ID of the added entity.
        """
        columns = Tables[entity_type].columns[1:]  # Exclude the id column
        values = tuple(kwargs[col] for col in columns)
        entity_id = self.db.insert(f"{entity_type}s", columns, values)

        logger.info(f"Added {entity_type}_id: {
                    entity_id} with values: {kwargs}")
        return entity_id

    def add_user(self, name=""):
        """
        Add a user to the database.

        Args:
            name (str, optional): The name of the user. Defaults to "admin".

        Returns:
            int: The ID of the added user.
        """
        try:
            return self.add_entity("user", name=name)
        except ErrorHandler.EntityAlreadyExistsError as e:
            logger.error(e)
            print(e)

    def add_conversation(self, user_id=0, conversation_history=""):
        """
        Add a conversation to the database.

        Args:
            user_id (int, optional): The ID of the user. Defaults to 0.
            conversation_history (str, optional): The conversation history. Defaults to "".

        Returns:
            int: The ID of the added conversation.
        """
        try:
            return self.add_entity("conversation", user_id=user_id, conversation_history=conversation_history)
        except ErrorHandler.EntityAlreadyExistsError as e:
            logger.error(e)
            print(e)

    def add_recipe(self, name="", ingredients="", instructions=""):
        """
        Add a recipe to the database.

        Args:
            name (str): The name of the recipe.
            ingredients (str, optional): The ingredients of the recipe. Defaults to "".
            instructions (str, optional): The instructions for the recipe. Defaults to "".

        Returns:
            int: The ID of the added recipe.
        """
        try:
            return self.add_entity("recipe", name=name, ingredients=ingredients, instructions=instructions)
        except ErrorHandler.EntityAlreadyExistsError as e:
            logger.error(e)
            print(e)

    def add_items(self, name="", category="", calories=0):
        """
        Add an item to the database.

        Args:
            name (str): The name of the item.
            category (str, optional): The category of the item. Defaults to "".
            calories (int, optional): The calories of the item. Defaults to 0.

        Returns:
            int: The ID of the added item.
        """
        try:
            return self.add_entity("item", name=name, category=category, calories=calories)
        except ErrorHandler.EntityAlreadyExistsError as e:
            logger.error(e)
            print(e)

    def get_entity(self, entity_type, entity_id):
        """
        Fetch an entity from the database.

        Args:
            entity_type (str): The type of entity (e.g., user, conversation, recipe, item).
            entity_id (int): The ID of the entity.

        Raises:
            ErrorHandler.EntityNotFoundError: If the entity with the given ID is not found.

        Returns:
            dict: The entity data if found.
        """
        if entity_id is None:
            error_message = ErrorHandler.entity_not_found_error_message(
                entity_type, entity_id)
            logger.error(error_message)
            raise ErrorHandler.EntityNotFoundError(error_message)

        result = self.db.fetch_one(f"{entity_type}s", f"{
                                   entity_type}_id", entity_id)

        if not result:
            error_message = ErrorHandler.entity_not_found_error_message(
                entity_type, entity_id)
            logger.error(error_message)
            raise ErrorHandler.EntityNotFoundError(error_message)

        entity_data = result
        logger.info(f"Fetched {entity_type}_id: {
                    entity_id} with data: {entity_data}")
        return entity_data

    def get_conversation(self, conversation_id):
        """
        Fetch a conversation from the database.

        Args:
            conversation_id (int): The ID of the conversation.

        Returns:
            dict: The conversation data if found.
        """
        return self.get_entity("conversation", conversation_id)

    def get_user(self, user_id):
        """
        Fetch a user from the database.

        Args:
            user_id (int): The ID of the user.

        Returns:
            dict: The user data if found.
        """
        return self.get_entity("user", user_id)

    def get_recipe(self, recipe_id):
        """
        Fetch a recipe from the database.

        Args:
            recipe_id (int): The ID of the recipe.

        Returns:
            dict: The recipe data if found.
        """
        return self.get_entity("recipe", recipe_id)

    def get_item(self, item_id):
        """
        Fetch an item from the database.

        Args:
            item_id (int): The ID of the item.

        Returns:
            dict: The item data if found.
        """
        return self.get_entity("item", item_id)

    def update_entity(self, entity_type, entity_id, **kwargs):
        """
        Update an entity in the database.

        Args:
            entity_type (str): The type of entity (e.g., user, conversation, recipe, item).
            entity_id (int): The ID of the entity.
            kwargs (dict): The updated entity data.

        Raises:
            ErrorHandler.EntityNotFoundError: If the entity with the given ID is not found.

        Returns:
            dict: The updated entity data.
        """
        existing_entity = self.get_entity(entity_type, entity_id)
        if not existing_entity:
            error_message = ErrorHandler.entity_not_found_error_message(
                entity_type, entity_id)
            logger.error(error_message)
            raise ErrorHandler.EntityNotFoundError(error_message)

        columns = ", ".join(f"{key} = ?" for key in kwargs.keys())
        values = list(kwargs.values()) + [entity_id]
        query = f"UPDATE {entity_type}s SET {
            columns} WHERE {entity_type}_id = ?"
        self.db.execute(query, values)

        logger.info(f"Updated {entity_type}_id: {
                    entity_id} with values: {kwargs}")
        return self.get_entity(entity_type, entity_id)

    def update_conversation(self, conversation_id, new_conversation):
        """
        Update a conversation in the database.

        Args:
            conversation_id (int): The ID of the conversation.
            new_conversation (str): The new conversation data to add.

        Returns:
            dict: The updated conversation data.
        """
        chat_history = self.get_conversation_chat_history_by_id(
            conversation_id=conversation_id)
        updated_chat_history = format_message_for_storage(
            chat_history, new_conversation)
        return self.update_entity("conversation", conversation_id, chat_history=updated_chat_history)

    def update_user(self, user_id, **kwargs):
        """
        Update a user in the database.

        Args:
            user_id (int): The ID of the user.
            kwargs (dict): The updated user data.

        Returns:
            dict: The updated user data.
        """
        return self.update_entity("user", user_id, **kwargs)

    def update_recipe(self, recipe_id, **kwargs):
        """
        Update a recipe in the database.

        Args:
            recipe_id (int): The ID of the recipe.
            kwargs (dict): The updated recipe data.

        Returns:
            dict: The updated recipe data.
        """
        return self.update_entity("recipe", recipe_id, **kwargs)

    def update_item(self, item_id, **kwargs):
        """
        Update an item in the database.

        Args:
            item_id (int): The ID of the item.
            kwargs (dict): The updated item data.

        Returns:
            dict: The updated item data.
        """
        return self.update_entity("item", item_id, **kwargs)

    def get_conversation_chat_history_by_id(self, conversation_id):
        """
        Fetch the chat history of a conversation from the database by conversation ID.

        Args:
            conversation_id (int): The ID of the conversation.

        Returns:
            list: The chat history as a list if found.

        Raises:
            ErrorHandler.EntityNotFoundError: If the conversation with the given ID is not found.
        """
        try:
            query = "SELECT chat_history FROM conversations WHERE conversation_id = ?"
            result = self.db.query(query, (conversation_id,))
            if result:
                # TODO: Find out what result looks like
                chat_history = result[0][0]
                return convert_chat_to_list(chat_history)
            else:
                error_message = ErrorHandler.entity_not_found_error_message(
                    "conversation", conversation_id)
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"""Error fetching chat history
                for conversation_id {conversation_id}: {e}""")

    def get_conversation_chat_history_by_user_id(self, user_id):
        """
        Fetch the chat histories of all conversations from the database by user ID.

        Args:
            user_id (int): The ID of the user.

        Returns:
            list: A list of chat histories, each converted to list format.

        Raises:
            ErrorHandler.EntityNotFoundError: If no conversations are found for the given user ID.
        """
        try:
            query = "SELECT chat_history FROM conversations WHERE user_id = ?"
            results = self.db.query(query, (user_id,))
            if results:
                chat_histories = [convert_chat_to_list(
                    chat_history[0]) for chat_history in results]
                return chat_histories
            else:
                error_message = ErrorHandler.entity_not_found_error_message(
                    "user", user_id)
                ErrorHandler.log_and_raise(
                    ErrorHandler.EntityNotFoundError, error_message)
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error fetching chat histories for user_id {user_id}: {e}")
