from services.databaseService import DatabaseService
from services.chatService import ChatService
from utils.formatting import ChatFormatter as fmt


class Routines:
    """
    A service class for defining run_time routines/operations.

    Attributes:
        db (SQLService): The database service instance.
    """

    def __init__(self, db_name):
        """Initialize a Routines instance.
        """
        self.db_name = db_name
        self.ds = DatabaseService(self.db_name)
        self.cs = ChatService()

    def new_user_new_chat(self, user_name):
        """
        Create a new user and a new chat conversation.

        Args:
            user_name (str): The name of the user.

        Returns:
            None: dumps the conversation into the conversations table.
        """
        self.ds.create_tables()
        user_id = self.ds.add_user(user_name)
        conversation_id = self.ds.add_conversation(user_id)
        self.cs.complete_chat(self.ds, conversation_id)
        return user_id, conversation_id

    def get_chat_and_print(self, conversation_id):
        """
        Get the chat conversation by ID and print it.

        Args:
            conversation_id (int): The ID of the conversation.
        """
        print("\nChat Transcript: for conversation_id: "+str(conversation_id)+"\n")
        chat_history = fmt.convert_chat_store_to_string_print(
            self.ds.get_conversation_chat_history_by_id(conversation_id))
        print(chat_history)

    def create_tables_initial(self):
        """Create the initial tables in the database."""
        self.ds.create_tables()
