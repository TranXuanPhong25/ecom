from typing import List, Optional
from pydantic import BaseModel


class MessageParts(BaseModel):
    type: str
    text: Optional[str] = None


class Message(BaseModel):
    role: str
    id: str
    parts: List[MessageParts]


class ChatPayload(BaseModel):
    messages: List[Message]

    class Config:
        json_schema_extra = {
            "example": {
                "messages": [
                    {
                        "role": "user",
                        "id": "msg-1",
                        "parts": [
                            {"type": "text", "text": "Who are you?"}
                        ]
                    }
                ]
            }
        }
