def convert_chat_to_string_print(chat):
    return '\n'.join([f"{message['role']}: {message['content']}" for message in chat]).join(",")

def convert_chat_to_string_store(chat):
    return '\n'.join([f"{message['role']}::{message['content']}" for message in chat]).join(";;")

def format_message_to_JSON(role, content):
    return {"role": role, "content": content}

def convert_chat_to_list(chat_string):
    chat_list = [message.split("::") for message in chat_string.split(";;")]
    return [format_message_to_JSON(message[0], message[1]) for message in chat_list]