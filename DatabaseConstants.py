butler_db = "butler.db"

get_conversation_by_id = """SELECT chat_history FROM conversations 
    WHERE conversation_id = ?"""

get_conversation_by_user_id = """SELECT chat_history FROM conversations 
    WHERE user_id = ?"""
