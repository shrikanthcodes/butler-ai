from DatabaseService import DatabaseService
import ChatService
from Formatting import convert_chat_to_string_print

if __name__ == "__main__":
    user_id = 1
    conversation_id = 1
    user_name = "admin"
    db = DatabaseService()
    # db.create_tables()
    # user_id = db.add_user(user_name)
    # conversation_id = db.add_conversation(user_id)
    ChatService.complete_chat(db, conversation_id, [])
    print("\nCompleted Chat Transcript:\n")
    convert_chat_to_string_print(db.get_conversation(conversation_id))
