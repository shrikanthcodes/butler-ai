import sqlite3
from sqlite3 import Error
import os


class SQLService:
    def __init__(self, db, delete_db=False):
        if delete_db and os.path.exists(db):
            os.remove(db)
        self.conn = None
        try:
            self.conn = sqlite3.connect(db)
            self.conn.execute("PRAGMA foreign_keys = 1")
        except Error as e:
            print(e)

    def create(self, table, schema, drop_table=False):
        if drop_table:
            try:
                c = self.conn.cursor()
                c.execute(f"DROP TABLE IF EXISTS {table}")
            except Error as e:
                print(e)

        try:
            c = self.conn.cursor()
            c.execute(f"CREATE TABLE {table} ({schema});")
        except Error as e:
            print(e)

    def query(self, sql, params=()):
        if not isinstance(params, tuple):
            params = (params,)
        cur = self.conn.cursor()
        cur.execute(sql, params)
        rows = cur.fetchall()
        return rows

    def insert(self, table, columns, values):
        columns_str = ', '.join(columns)
        placeholders = ', '.join(['?'] * len(values))
        sql = f"INSERT INTO {table} ({columns_str}) VALUES ({placeholders})"
        cur = self.conn.cursor()
        cur.execute(sql, values)
        self.conn.commit()
        return cur.lastrowid

    def update(self, table, set_column, set_value, condition_column, condition_value):
        sql = f"""UPDATE {table} SET {
            set_column} = ? WHERE {condition_column} = ?"""
        cur = self.conn.cursor()
        cur.execute(sql, (set_value, condition_value))
        self.conn.commit()
        return cur.lastrowid
