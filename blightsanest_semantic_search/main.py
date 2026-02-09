from fastapi import FastAPI
from pydantic import BaseModel
from typing import List

from semantic_engine.semantic_search import SemanticSearch

class EmbeddingDoc(BaseModel):
    id: str
    data: str

class EmbeddingsRequest(BaseModel):
    documents: List[EmbeddingDoc]

class SearchRequest(BaseModel):
    query: str
    limit: int

# init semantic search class
ss = SemanticSearch()

# create app
app = FastAPI()

# load or create embeddings
@app.post("/embeddings")
def createEmbeddings(req: EmbeddingsRequest):
    ss.load_or_create_embeddings(req.documents)
    return {"count": len(req.documents)}

# perform search
@app.post("/search")
def search(req: SearchRequest):
    results = ss.search(req.query, req.limit)
    # turn tuples into JSON-serializable objects
    out = [{"score": score, "document": doc} for score, doc in results]
    return {"results": out}
