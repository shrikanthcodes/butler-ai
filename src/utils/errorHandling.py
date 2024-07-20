import logging

# Configure logging
logging.basicConfig(level=logging.WARNING)
logger = logging.getLogger(__name__)


class ErrorHandler:
    """Handles errors and provides logging and error messages."""

    class EntityAlreadyExistsError(Exception):
        """Exception raised when an entity with the given ID already exists."""
        pass

    class DatabaseError(Exception):
        """Exception raised for database errors."""
        pass

    @staticmethod
    def entity_exists_error_message(entity_type, entity_id):
        """
        Generate an error message for an existing entity.

        Args:
            entity_type (str): The type of entity (e.g., user, conversation).
            entity_id (int): The ID of the entity.

        Returns:
            str: An error message indicating the entity already exists.
        """
        return f"{entity_type.capitalize()} with ID {entity_id} already exists."

    @staticmethod
    def entity_not_found_error_message(entity_type, entity_id):
        """
        Generate an error message for a non-existent entity.

        Args:
            entity_type (str): The type of entity (e.g., user, conversation).
            entity_id (int): The ID of the entity.

        Returns:
            str: An error message indicating the entity was not found.
        """
        return f"{entity_type.capitalize()} with ID {entity_id} not found."

    @staticmethod
    def log_and_raise(error, message):
        """
        Raises a custom exception.

        Args:
            error (Exception): The exception to raise.
            message (str): The error message to log and raise.
        """
        raise error(message)
