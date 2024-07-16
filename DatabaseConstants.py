butler_db = "butler.db"


def get_entity_by_id(entity_type, entity_id):
    query = f"SELECT * FROM {entity_type}s WHERE id = ?"
    return query, (entity_id,)
