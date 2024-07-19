# TODO: Separation of Concerns by creating 3 different DBs for different data reqs
butler_db = "butler.db"
user_db = "user.db"
items_db = "items.db"


# def get_entity_by_id(entity_type, entity_id):
#     query = f"SELECT * FROM {entity_type}s WHERE id = ?"
#     return query, (entity_id,)
