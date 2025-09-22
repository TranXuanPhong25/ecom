
from typing import List

from langchain_core.messages import HumanMessage, SystemMessage, AIMessage

from app.models.chat import Message


def ai_sdk_to_langchain(messages: List[Message]):
    lc_messages = []
    for msg in messages:
        text_parts = [p.text for p in msg.parts if p.type == "text"]
        content = "\n".join(text_parts)  # gộp text parts
        if msg.role == "user":
            lc_messages.append(HumanMessage(content=content))
        elif msg.role == "assistant":
            lc_messages.append(AIMessage(content=content))
        elif msg.role == "system":
            lc_messages.append(SystemMessage(content=content))
    return lc_messages
