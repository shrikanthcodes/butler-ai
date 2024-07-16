def convert_chat_to_string_print(chat):
    if isinstance(chat, list):
        formatted_messages = '\n'.join(
            [f"{message['role']}: {message['content']}" for message in chat])
    else:
        formatted_messages = chat
    return formatted_messages


def convert_chat_store_to_string_print(chat_str):
    if chat_str == "":
        return ""
    pretty_print = chat_str.replace(";;", "\n").replace("::", ": ")
    return "\n\n" + pretty_print


def convert_chat_to_string_store(chat):
    formatted_messages = ';;'.join(
        [f"{message['role']}::{message['content']}" for message in chat])
    return formatted_messages


def format_message_to_JSON(role, content):
    return {"role": role, "content": content}

# TODO: Error Handling for empty conversation_history string in convert_chat_to_list


def convert_chat_to_list(chat_string):
    if chat_string == "":
        return []
    chat_list = [message.split("::") for message in chat_string.split(";;")]
    return [format_message_to_JSON(message[0], message[1]) for message in chat_list]


def format_message_for_storage(old_chat, new_chat):
    if old_chat == "":
        chat = new_chat
    else:
        chat = old_chat + ";;" + new_chat
    return chat
