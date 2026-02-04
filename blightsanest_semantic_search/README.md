This is the Python server necessary for semantic search.

This feature is built so you can perform advance search on the financial data you saved on the database.
- embedding port
- query port
- results port

Pass documents to be embedded to the server {[id of the row]: stringified data}
It's curucial that you stringify the data before sending it to the semantic search server.
Building the embeddings will take time, depending on the size of the data you send to this server.

Pass query to the server

Retrieve results

This is yet to be tested and integrated
