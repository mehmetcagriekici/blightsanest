from semantic_engine.semantic_search import SemantciSearch

# search connector
def embed_query_text(query):
    ss = SemanticSearch()
    embeddings = ss.generate_embedding(query)
