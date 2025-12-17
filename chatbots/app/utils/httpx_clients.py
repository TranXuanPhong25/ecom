import os

import httpx

class ProductSearchSvcClient:
    _instance = None

    def __new__(cls, *args, **kwargs):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance

    def __init__(self):
        if not hasattr(self, "client"):
            base_url = os.getenv("PRODUCTS_SEARCH_SVC_URL")
            if not base_url:
                print("PRODUCTS_SEARCH_SVC_URL is missing, use default value")
                base_url = "http://products-search-svc.services"
            self.base_url =base_url
            self.client = httpx.AsyncClient(timeout=5.0, http2=True)

    async def search(self, query: str, limit: int = 10):
        # res = await self.client.get(
        #     f"{self.base_url}/search",
        #     params={"q": query, "limit": limit}
        # )

        # return res.json().get("results", [])
        return [{
            "name":"sunflower"
        }]
