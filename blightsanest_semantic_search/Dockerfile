FROM python:3.12-slim

WORKDIR /app

# 1) Copy project metadata and install your project (and deps) from pyproject.toml
COPY pyproject.toml ./

# CPU-only torch build first, so sentence-transformers' torch dependency
# resolves against it instead of pulling the ~7GB CUDA wheels (no GPU in k8s).
RUN pip install --no-cache-dir torch --index-url https://download.pytorch.org/whl/cpu
RUN pip install --no-cache-dir -e .

# 2) Copy the actual source code
COPY . .

# 3) Run FastAPI (entry: main.py -> module main)
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]