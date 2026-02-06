from fastapi import FastAPI
from pydantic import BaseModel

from semantic_engine.semantic_search import SemanticSearch

class EmbeddingsRequest(BaseModel):
    documnents: List[str]

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
    return {"results": results}
