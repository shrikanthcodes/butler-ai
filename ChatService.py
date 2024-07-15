from openai import OpenAI
from Credentials import OPENAI_API_KEY
import Formatting
import LLMConstants as lc


def complete_chat(db, conversation_id):
    conversation_history = db.get_conversation_chat_history_by_id(
        conversation_id)
    client = OpenAI(api_key=OPENAI_API_KEY)
    current_conversation = []
    current_conversation.append(
        Formatting.format_message_to_JSON(lc.roles[1], lc.initial_message))

    while True:
        prompt_text = input("Type Next Message: ")
        current_conversation.append(
            {"role": lc.roles[0], "content": prompt_text})
        response = client.chat.completions.create(
            model=lc.model,
            messages=conversation_history + current_conversation,
            temperature=0.7,
            n=1,
            stop=None,
        )

        response_message = response.choices[0].message.content
        print("AI: " + response_message)
        current_conversation.append(
            Formatting.format_message_to_JSON(lc.roles[1], response_message))

        if goal_reached(prompt_text):
            break

    current_conversation_string = Formatting.convert_chat_to_string_store(
        current_conversation)
    db.update_conversation(
        conversation_id, current_conversation_string)


def goal_reached(message):
    if "Goodbye" in message:
        return True
    return False
