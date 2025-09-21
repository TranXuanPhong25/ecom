import os
from langchain.chat_models import init_chat_model
from dotenv import load_dotenv
from langchain_tavily import TavilySearch
from langgraph.prebuilt import create_react_agent

try:
    load_dotenv()
except ImportError:
    pass

class AgentConfig:
    def __init__(self):
        os.environ["LANGSMITH_TRACING"] = "true"
        if "LANGSMITH_API_KEY" not in os.environ:
            print("LANGSMITH_API_KEY is missing")
        if "LANGSMITH_PROJECT" not in os.environ:
            print("LANGSMITH_PROJECT set to 'default'")
        if not os.environ.get("GOOGLE_API_KEY"):
            print("GOOGLE_API_KEY is missing")
        if not os.environ.get("TAVILY_API_KEY"):
            print("TAVILY_API_KEY is missing")

        model = init_chat_model("gemini-2.5-flash", model_provider="google_genai")
        search = TavilySearch(max_results=2)
        tools = [search]

        self.agent = create_react_agent(model, tools)
