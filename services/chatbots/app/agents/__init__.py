from app.agents.prompts import RECOMMENDER_AGENT_PROMPT, SUPERVISOR_AGENT_PROMPT, PRODUCTS_SEARCHER_AGENT_PROMPT
from app.agents.configs import SupervisorAgent, RecommenderAgent, ProductsSearcherAgent

recommender_agent = RecommenderAgent(RECOMMENDER_AGENT_PROMPT, name="recommender")
products_searcher_agent = ProductsSearcherAgent(PRODUCTS_SEARCHER_AGENT_PROMPT, name="products_searcher")
supervisor_agent = SupervisorAgent(SUPERVISOR_AGENT_PROMPT, [recommender_agent.get(), products_searcher_agent.get()])
