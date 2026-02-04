from semantic_engine.semantic_search import SemanticSearch

# building connector
def embed_documents(documents):
    ss = SemanticSearch()
    ss.load_or_create_embeddings(documents)
