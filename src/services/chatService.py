from openai import OpenAI
from constants.credentials import OPENAI_API_KEY
from utils.formatting import ChatFormatter as cf
from llm.functions import goal_reached
import constants.llmConstants as lc
from utils.errorHandling import ErrorHandler


class ChatService:
    """
    A class to handle the chat interactions between the user and AI.

    Methods:
        complete_chat(db, conversation_id): Manages the complete chat session.
    """

    def __init__(self):
        """
        Initializes the ChatService with OpenAI client.
        """

        try:
            self.client = OpenAI(api_key=OPENAI_API_KEY)
        except Exception as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error initializing OpenAI client: {e}")

    def complete_chat(self, db, conversation_id):
        """
        Manages the complete chat session.

        Args:
            db: The database object to interact with conversation history.
            conversation_id (int): The ID of the conversation to continue.

        Raises:
            ValueError: If there is an error in chat completion.
        """
        try:
            conversation_history = cf.convert_chat_to_list(
                db.get_conversation_chat_history_by_id(conversation_id))
            current_conversation = [cf.format_message_to_JSON(
                lc.roles[1], lc.initial_message)]
            print("AI: " + lc.initial_message)

            while True:
                prompt_text = input("Type Next Message: ")
                current_conversation.append(
                    cf.format_message_to_JSON(lc.roles[0], prompt_text))

                response = self.client.chat.completions.create(
                    model=lc.model,
                    messages=conversation_history + current_conversation,
                    temperature=0.7,
                    n=1,
                    stop=None,
                )

                response_message = response.choices[0].message.content
                print("AI: " + response_message)
                current_conversation.append(
                    cf.format_message_to_JSON(lc.roles[1], response_message))

                if goal_reached(prompt_text):
                    break

            current_conversation_string = cf.convert_chat_to_string_store(
                current_conversation)
            db.update_conversation(
                conversation_id, current_conversation_string)

        except Exception as e:
            ErrorHandler.log_and_raise(
                ValueError, f"Error during chat completion: {e}")
