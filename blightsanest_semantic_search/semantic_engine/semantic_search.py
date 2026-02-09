from semantic_engine.cosine_similarity import cosine_similarity

from sentence_transformers import SentenceTransformer
import numpy as np
import os

class SemanticSearch:
    def __init__(self, model_name="all-MiniLM-L6-v2"):
        self.model = SentenceTransformer(model_name)
        self.embeddings = None
        self.documents = None
        self.document_map = {}

    # convert text into embeddings
    def generate_embedding(self, text):
        # raise a value error if text is empty
        if text.strip() == "":
            raise ValueError("input text must not be empty.")

        # generate embeddings for the input text
        embeddings = self.model.encode([text])
        return embeddings[0]

    # create embeddings for the documents
    def build_embeddings(self, documents):
        # list to store document data to be embedded
        data_list = []
        # list of dictionaries {id: string, data: string}
        self.documents = documents
        for doc in self.documents:
            doc_id = doc["id"]
            doc_data = doc["data"]
            self.document_map[doc_id] = doc_data
            data_list.append(doc_data)
        # embed the documents
        self.embeddings = self.model.encode(data_list, show_progress_bar=True)

        # save the embeddings for the next time
        os.makedirs("cache", exist_ok=True)
        np.save("cache/db_embeddings.npy", self.embeddings)
                
        return self.embeddings

    # store embeddings locally
    def load_or_create_embeddings(self, documents):
        # populate the self.documents and self.document_map
        self.documents = documents
        for doc in self.documents:
            self.document_map[doc["id"]] = doc["data"]

        # check if embeddings already exists locally
        if os.path.exists("cache/db_embeddings.npy"):
            self.embeddings = np.load("cache/db_embeddings.npy")
            if len(self.embeddings) == len(self.documents):
                return self.embeddings
        return self.build_embeddings(documents)

    # semantic search method
    def search(self, query, limit):
        # check if the embeddings loaded
        if self.embeddings is None:
            raise ValueError("Please call <load_or_create_embeddings> first.")

        # generate an embedding for the query
        query_embedding = self.generate_embedding(query)

        # calculate cosine similarity between the query embedding and each document embeddings.
        similarities = []
        for i in range(len(self.embeddings)):
            cs_score = cosine_similarity(query_embedding, self.embeddings[i])
            similarities.append((cs_score, self.documents[i]))

        # sort the similarities
        similarities = sorted(similarities, key=lambda kv: kv[0], reverse=True)

        return similarities[:limit]
