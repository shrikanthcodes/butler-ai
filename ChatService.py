from openai import OpenAI
from Credentials import OPENAI_API_KEY
from DatabaseService import DatabaseService
import Formatting
import LLMConstants as lc


def complete_chat(conversation=[]):
    current_conversation = []
    client = OpenAI(api_key=OPENAI_API_KEY)
    current_conversation.append(
        Formatting.format_message_to_JSON(lc.roles[1], lc.initial_message))

    while True:
        prompt_text = input("Type Next Message: ")
        current_conversation.append(
            {"role": lc.roles[0], "content": prompt_text})
        response = client.chat.completions.create(
            model=lc.model,
            messages=conversation.extend(current_conversation),
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

    return Formatting.convert_chat_to_string_print(current_conversation)


def goal_reached(message):
    if "Goodbye" in message:
        return True
    return False


print("\nCompleted Chat Transcript:\n" + complete_chat())
