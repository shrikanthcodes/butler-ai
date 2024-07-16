import os
import sqlite3
from sqlite3 import Error
from errorHandling import ErrorHandler, logger


class SQLService:
    """
    A service class for managing SQLite database interactions.

    Attributes:
        conn (sqlite3.Connection): The connection object to the SQLite database.
    """

    def __init__(self, db, delete_db=False):
        """
        Initialize the SQLService.

        Args:
            db (str): The path to the SQLite database file.
            delete_db (bool): If True, delete the existing database file before creating a new connection.
        """
        if delete_db and os.path.exists(db):
            os.remove(db)
        self.conn = None
        try:
            self.conn = sqlite3.connect(db)
            self.conn.execute("PRAGMA foreign_keys = 1")
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error connecting to database: {e}")

    def create(self, table, schema, drop_table=False):
        """
        Create a table in the database.

        Args:
            table (str): The name of the table.
            schema (str): The schema of the table.
            drop_table (bool): If True, drop the table if it already exists.
        """
        try:
            with self.conn:
                if drop_table:
                    self.conn.execute(f"DROP TABLE IF EXISTS {table}")
                self.conn.execute(f"CREATE TABLE IF NOT EXISTS {
                                  table} ({schema});")
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error creating table '{table}': {e}")

    def query(self, sql, params=()):
        """
        Execute a query and fetch all results.

        Args:
            sql (str): The SQL query to execute.
            params (tuple): The parameters to bind to the query.

        Returns:
            list: A list of rows returned by the query.
        """
        try:
            with self.conn:
                cur = self.conn.cursor()
                cur.execute(sql, params)
                rows = cur.fetchall()
                return rows
        except Error as e:
            ErrorHandler.log_and_raise(ErrorHandler.DatabaseError, f"Error executing query: {
                                       sql} with params: {params} - {e}")

    def fetch_one(self, table, column, value):
        """
        Fetch a single row from a table based on a condition.

        Args:
            table (str): The name of the table.
            column (str): The column name to match.
            value (any): The value to match in the column.

        Returns:
            tuple: The row that matches the condition, or None if no match is found.
        """
        sql = f"SELECT * FROM {table} WHERE {column} = ?"
        try:
            with self.conn:
                cur = self.conn.cursor()
                cur.execute(sql, (value,))
                row = cur.fetchone()
                return row
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error fetching row from table '{table}': {e}")

    def insert(self, table, columns, values):
        """
        Insert a row into a table.

        Args:
            table (str): The name of the table.
            columns (list): The list of column names to insert values into.
            values (tuple): The values to insert.

        Returns:
            int: The ID of the last inserted row.
        """
        columns_str = ', '.join(columns)
        placeholders = ', '.join(['?'] * len(values))
        sql = f"INSERT INTO {table} ({columns_str}) VALUES ({placeholders})"
        try:
            with self.conn:
                cur = self.conn.cursor()
                cur.execute(sql, values)
                self.conn.commit()
                return cur.lastrowid
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error inserting row into table '{table}': {e}")

    def update(self, table, set_column, set_value, condition_column, condition_value):
        """
        Update a row in a table.

        Args:
            table (str): The name of the table.
            set_column (str): The column to update.
            set_value (any): The new value for the column.
            condition_column (str): The column to use in the WHERE clause.
            condition_value (any): The value to match in the WHERE clause.

        Returns:
            int: The number of rows affected.
        """
        sql = f"UPDATE {table} SET {
            set_column} = ? WHERE {condition_column} = ?"
        try:
            with self.conn:
                cur = self.conn.cursor()
                cur.execute(sql, (set_value, condition_value))
                self.conn.commit()
                return cur.rowcount
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error updating row in table '{table}': {e}")

    def last_insert_id(self):
        """
        Get the ID of the last inserted row.

        Returns:
            int: The ID of the last inserted row.
        """
        try:
            with self.conn:
                return self.conn.execute('SELECT last_insert_rowid()').fetchone()[0]
        except Error as e:
            ErrorHandler.log_and_raise(
                ErrorHandler.DatabaseError, f"Error getting last insert id: {e}")
