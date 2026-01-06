from fastapi import APIRouter
recommendation_router = APIRouter()

@recommendation_router.get("/api/recommendations/daily-discoveries")
async def get_daily_discoveries():
    # Placeholder for actual recommendation logic
    recommendations = [
        {"id": 1, "title": "Discover the Mountains", "description": "Explore the majestic mountain ranges."},
        {"id": 2, "title": "Beach Getaway", "description": "Relax on the sunny beaches."},
        {"id": 3, "title": "City Lights", "description": "Experience the vibrant city life."}
    ]
    return {"daily_discoveries": recommendations}
