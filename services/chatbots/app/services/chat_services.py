import asyncio

from fastapi.responses import StreamingResponse
from langchain_core.messages import convert_to_messages
from langchain_core.messages.ai import AIMessage
from langchain_core.runnables.base import Runnable

from app.models.chat import Message
import json


async def send_completion_events(agent, messages):
    # phát sự kiện start
    yield f'data: {json.dumps({"type": "start", "messageMetadata": {"createdAt": "2025-09-20T15:00:00Z"}})}\n\n'
    yield f'data: {json.dumps({"type": "start-step"})}\n\n'
    yield f'data: {json.dumps({"type": "text-start", "id": "0"})}\n\n'
    # stream từng token
    for chunk in agent.stream(
       messages
    ):
        pretty_print_messages(chunk)

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
                    print(content)
                    yield f'data: {json.dumps(payload)}\n\n'

    # kết thúc text
    yield f'data: {json.dumps({"type": "text-end", "id": "0"})}\n\n'
    yield f'data: {json.dumps({"type": "finish-step"})}\n\n'
    yield f'data: {json.dumps({"type": "finish", "messageMetadata": {"reasoningDurationInMs": 0}})}\n\n'
    yield 'data: [DONE]\n\n'

def pretty_print_message(message, indent=False):
    pretty_message = message.pretty_repr(html=True)
    if not indent:
        print(pretty_message)
        return

    indented = "\n".join("\t" + c for c in pretty_message.split("\n"))
    print(indented)


def pretty_print_messages(update, last_message=False):
    is_subgraph = False
    if isinstance(update, tuple):
        ns, update = update
        # skip parent graph updates in the printouts
        if len(ns) == 0:
            return

        graph_id = ns[-1].split(":")[0]
        print(f"Update from subgraph {graph_id}:")
        print("\n")
        is_subgraph = True

    for node_name, node_update in update.items():
        update_label = f"Update from node {node_name}:"
        if is_subgraph:
            update_label = "\t" + update_label

        print(update_label)
        print("\n")

        messages = convert_to_messages(node_update["messages"])
        if last_message:
            messages = messages[-1:]

        for m in messages:
            pretty_print_message(m, indent=is_subgraph)
        print("\n")
async def sse_stream(graph: Runnable, messages):
    # phát sự kiện start
    yield f'data: {json.dumps({"type": "start"})}\n\n'
    yield f'data: {json.dumps({"type": "start-step"})}\n\n'
    yield f'data: {json.dumps({"type": "text-start", "id": "0"})}\n\n'

    async for chunk in graph.astream(messages, stream_mode="updates"):
        for actor, update in chunk.items():
            msgs = update.get("messages", [])
            for msg in msgs:
                if isinstance(msg, (AIMessage)) and actor == "supervisor":
                    pretty_print_messages(chunk)
                    payload ={
                        "type": "text-delta",
                        "id": "0",
                        "delta": msg.content
                    }
                    yield f'data: {json.dumps(payload)}\n\n'
            # event = {
            #     "type": "step",
            #     "actor": actor,
            #     "keys": list(update.keys())
            # }
            # yield f"data: {json.dumps(event)}\n\n"

    # kết thúc text
    yield f'data: {json.dumps({"type": "text-end", "id": "0"})}\n\n'
    yield f'data: {json.dumps({"type": "finish-step"})}\n\n'
    yield f'data: {json.dumps({"type": "finish"})}\n\n'
    yield 'data: [DONE]\n\n'


def process_query(agent, messages: Message):
    return StreamingResponse(send_completion_events(agent, {"messages": messages}), media_type="text/event-stream")
