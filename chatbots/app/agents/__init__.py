from app.agents.prompts import RECOMMENDER_AGENT_PROMPT, SUPERVISOR_AGENT_PROMPT
from app.agents.configs import SupervisorAgent, RecommenderAgent

recommender_agent = RecommenderAgent(RECOMMENDER_AGENT_PROMPT, name="recommender")
supervisor_agent = SupervisorAgent(SUPERVISOR_AGENT_PROMPT, [recommender_agent.get()])