from DatabaseService import DatabaseService
import ChatService
from Formatting import convert_chat_to_string_print, convert_chat_to_list, convert_chat_to_string_print


if __name__ == "__main__":
    user_id = 2
    conversation_id = 2
    user_name = "admin"
    db = DatabaseService()
    db.create_tables()
    user_id = db.add_user(user_id, user_name)
    conversation_id = db.add_conversation(conversation_id, user_id, "")
    ChatService.complete_chat(db, conversation_id, [])
    print("\nCompleted Chat Transcript:\n")
    convert_chat_to_string_print(convert_chat_to_list(
        db.get_conversation(conversation_id)))
