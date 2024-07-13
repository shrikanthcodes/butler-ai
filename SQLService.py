import sqlite3
from sqlite3 import Error
import os


class SQLService:
  def __init__(self, db_file, delete_db=False):
    if delete_db and os.path.exists(db_file):
      os.remove(db_file)
    self.conn = None
    try:
      self.conn = sqlite3.connect(db_file)
      self.conn.execute("PRAGMA foreign_keys = 1")
    except Error as e:
      print(e)

  def create_table(self, create_table_sql, drop_table_name=None):
    if drop_table_name:
      try:
        c = self.conn.cursor()
        c.execute("""DROP TABLE IF EXISTS %s""" % (drop_table_name))
      except Error as e:
        print(e)

    try:
      c = self.conn.cursor()
      c.execute(create_table_sql)
    except Error as e:
      print(e)

  def execute_sql_statement(self, sql_statement):
    cur = self.conn.cursor()
    cur.execute(sql_statement)
    rows = cur.fetchall()
    return rows

  def insert_rows(self, sql_statement, values):
    cur = self.conn.cursor()
    cur.execute(sql_statement, values)
    return cur.lastrowid
