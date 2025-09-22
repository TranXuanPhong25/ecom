from fastapi.responses import StreamingResponse
from app.models.chat import Message
import json


async def send_completion_events(agent, messages ):
    # phát sự kiện start
    yield f'data: {json.dumps({"type": "start", "messageMetadata": {"createdAt": "2025-09-20T15:00:00Z"}})}\n\n'
    yield f'data: {json.dumps({"type": "start-step"})}\n\n'
    yield f'data: {json.dumps({"type": "text-start", "id": "0"})}\n\n'
    response = agent.invoke(messages)

    for message in response["messages"]:
        message.pretty_print()
    # stream từng token
    async for patch in agent.astream_log(messages):
        for op in patch.ops:
            if op["op"] == "add" and (
                    op["path"].endswith("/streamed_output/-")
                    or op["path"].endswith("/content/-")
            ):
                val = op["value"]
                if isinstance(val, str):
                    content = val
                elif hasattr(val, "content"):
                    content = val.content
                elif isinstance(val, dict):
                    content = val.get("content", "")
                else:
                    content = str(val)

                if content.strip():
                    payload = {
                        "type": "text-delta",
                        "id": "0",
                        "delta": content,
                    }
                    yield f'data: {json.dumps(payload)}\n\n'

    # kết thúc text
    yield f'data: {json.dumps({"type": "text-end", "id": "0"})}\n\n'
    yield f'data: {json.dumps({"type": "finish-step"})}\n\n'
    yield f'data: {json.dumps({"type": "finish", "messageMetadata": {"reasoningDurationInMs": 0}})}\n\n'
    yield 'data: [DONE]\n\n'

def process_query(agent, messages: Message):
    return StreamingResponse(send_completion_events(agent, {"messages": messages}), media_type="text/event-stream")