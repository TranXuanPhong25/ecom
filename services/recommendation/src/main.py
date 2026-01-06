from typing import Union

from fastapi import FastAPI
from routes import recommendation_router
app = FastAPI()

app.include_router(recommendation_router)
