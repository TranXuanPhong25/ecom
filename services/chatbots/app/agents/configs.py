import os
from langchain.chat_models import init_chat_model
from dotenv import load_dotenv
from langchain_tavily import TavilySearch
from langgraph.prebuilt import create_react_agent
from langgraph_supervisor import create_supervisor

from app.agents.tools.products_search import products_searcher

try:
    load_dotenv()
except ImportError:
    pass


class Agent:
    agent = None
    def __init__(self):
        os.environ["LANGSMITH_TRACING"] = "true"
        if "LANGSMITH_API_KEY" not in os.environ:
            print("LANGSMITH_API_KEY is missing")
        if "LANGSMITH_PROJECT" not in os.environ:
            print("LANGSMITH_PROJECT set to 'default'")
        if not os.environ.get("GOOGLE_API_KEY"):
            print("GOOGLE_API_KEY is missing")
        self.model = init_chat_model("gemini-2.5-flash-lite", model_provider="google_genai")
    def get(self):
        return self.agent

class RecommenderAgent(Agent):
    def __init__(self, prompt, name):
        super().__init__()
        if not os.environ.get("TAVILY_API_KEY"):
            print("TAVILY_API_KEY is missing")
        search = TavilySearch(max_results=2)
        tools = [search]
        self.name = name
        self.agent = create_react_agent(self.model, tools, prompt=prompt, name=self.name)

class ProductsSearcherAgent(Agent):
    def __init__(self, prompt, name):
        super().__init__()
        tools = [products_searcher]
        self.name = name
        self.agent = create_react_agent(self.model, tools, prompt=prompt, name=self.name)



class SupervisorAgent(Agent):
    def __init__(self, prompt, agents):
        super().__init__()
        self.agent = create_supervisor(
            model=self.model,
            agents=agents,
            prompt=prompt,
            add_handoff_back_messages=True,
            output_mode="full_history",
        ).compile()

