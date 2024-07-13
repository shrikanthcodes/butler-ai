from SQLService import SQLService
import DatabaseConstants as dc

class DatabaseService:
    def __init__(self):
        self.user_count = 0
        self.conversation_count = 0

    def add_new_conversation(self, user_id, conversation_id):
        sql_statement = dc.add_new_conversation_sql
        values = [(user_id, conversation_id, "")]
        SQLService.insert_rows(sql_statement, values)
        self.user_count += 1
        self.conversation_count += 1

    def get_chat_history(self, conversation_id=None, user_id=None):
        if conversation_id is None:
            conversation_id = self.conversation_count
        if user_id is None:
            user_id = self.user_count

        if conversation_id:
            result = SQLService.execute_sql_statement(dc.chat_history_by_conversation_id_sql,(conversation_id,))
        elif user_id:
            result = SQLService.execute_sql_statement(dc.chat_history_by_user_id_sql,(user_id,))
        else:
            return None

        chat_history = result.fetchone()[0] if result else None
        return chat_history

    def update_chat_history(self, conversation_id, new_message):
        current_chat_history = self.get_chat_history(conversation_id=conversation_id)
        if current_chat_history is None:
            current_chat_history = ""
        updated_chat_history = current_chat_history +";;"+new_message

        sql_statement = dc.update_chat_history_sql
        values = (updated_chat_history, conversation_id)
        SQLService.execute_sql_statement(sql_statement, values)