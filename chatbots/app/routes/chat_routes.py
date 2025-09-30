from http.client import responses
from typing import Union

from app.agents import supervisor_agent
from app.models.chat import ChatPayload
from fastapi import Request


from fastapi import APIRouter

from app.services.chat_services import process_query
from app.utils.chat_utils import ai_sdk_to_langchain

chat_router = APIRouter()


@chat_router.post("/api/chats")
def stream(payload: ChatPayload):
    messages = ai_sdk_to_langchain(payload.messages)
    return process_query(supervisor_agent.get(),  messages)

