from utils.errorHandling import ErrorHandler


class ChatFormatter:
    """
    A class to handle formatting of chat data for storage and display.

    Methods:
        convert_chat_to_string_print(chat): Converts chat to a formatted string for printing.
        convert_chat_store_to_string_print(chat_str): Converts stored chat string to a printable format.
        convert_chat_to_string_store(chat): Converts chat to a formatted string for storage.
        format_message_to_JSON(role, content): Formats a message to JSON.
        convert_chat_to_list(chat_string): Converts stored chat string to a list of JSON messages.
        format_message_for_storage(old_chat, new_chat): Formats messages for storage.
    """

    @staticmethod
    def convert_chat_to_string_print(chat):
        """
        Converts a chat list to a formatted string for printing.

        Args:
            chat (list or str): The chat data.

        Returns:
            str: Formatted chat string for printing.

        Raises:
            ValueError: If the chat data is not a list or a string.
        """
        if isinstance(chat, list):
            formatted_messages = '\n'.join(
                [f"{message['role']}: {message['content']}" for message in chat])
        elif isinstance(chat, str):
            formatted_messages = chat
        else:
            ErrorHandler.log_and_raise(
                ValueError, "Invalid chat data type. Must be a list or a string.")
        return formatted_messages

    @staticmethod
    def convert_chat_store_to_string_print(chat_str):
        """
        Converts a stored chat string to a printable format.

        Args:
            chat_str (str): The stored chat string.

        Returns:
            str: Formatted printable chat string.
        """
        if not chat_str:
            return ""
        pretty_print = chat_str.replace(";;", "\n").replace("::", ": ")
        return "\n\n" + pretty_print

    @staticmethod
    def convert_chat_to_string_store(chat):
        """
        Converts a chat list to a formatted string for storage.

        Args:
            chat (list): The chat data.

        Returns:
            str: Formatted chat string for storage.

        Raises:
            ValueError: If the chat data is not a list.
        """
        if not isinstance(chat, list):
            ErrorHandler.log_and_raise(
                ValueError, "Invalid chat data type. Must be a list.")
        formatted_messages = ';;'.join(
            [f"{message['role']}::{message['content']}" for message in chat])
        return formatted_messages

    @staticmethod
    def format_message_to_JSON(role, content):
        """
        Formats a message to a JSON-compatible dictionary.

        Args:
            role (str): The role of the sender.
            content (str): The message content.

        Returns:
            dict: A dictionary representing the message.
        """
        return {"role": role, "content": content}

    @staticmethod
    def convert_chat_to_list(chat_string):
        """
        Converts a stored chat string to a list of JSON messages.

        Args:
            chat_string (str): The stored chat string.

        Returns:
            list: A list of JSON messages.

        Raises:
            ValueError: If the chat_string is not a valid string.
        """
        if not isinstance(chat_string, str):
            ErrorHandler.log_and_raise(
                ValueError, "Invalid chat string type. Must be a string.")
        if not chat_string:
            return []
        chat_list = [message.split("::")
                     for message in chat_string.split(";;")]
        return [ChatFormatter.format_message_to_JSON(message[0], message[1]) for message in chat_list]

    @staticmethod
    def format_message_for_storage(old_chat, new_chat):
        """
        Formats messages for storage by appending new messages to old chat.

        Args:
            old_chat (str): The existing chat string.
            new_chat (str): The new chat string to append.

        Returns:
            str: The combined chat string.
        """
        if not old_chat:
            chat = new_chat
        else:
            chat = old_chat + ";;" + new_chat
        return chat
