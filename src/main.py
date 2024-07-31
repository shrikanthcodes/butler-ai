from utils.runtime_routines import Routines

if __name__ == "__main__":
    # Create a Routines instance
    run = Routines("butler.db")

    # Create the initial tables in the database
    run.create_tables_initial()

    # Create a new user and chat conversation
    user_id, conversation_id = run.new_user_new_chat("Hithesh")

    # Get and print the chat conversation
    run.get_chat_and_print(conversation_id)
