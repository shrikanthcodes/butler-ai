create_table_sql = """CREATE TABLE conversations(
    conversation_id INT PRIMARY KEY,
    user_id INT FOREIGN KEY,
    chat_history TEXT);"""

db_file = "recipes.db"

chat_history_by_conversation_id_sql = """SELECT chat_history FROM conversations 
                                    WHERE conversation_id = ?"""

chat_history_by_user_id_sql = """SELECT chat_history FROM conversations WHERE 
                            user_id = ?"""

update_chat_history_sql = """UPDATE conversations SET chat_history = ? WHERE 
                        conversation_id = ?"""

add_new_conversation_sql = """INSERT INTO conversations (user_id, conversation_id,
                         chat_history) VALUES (?, ?, ?)"""
