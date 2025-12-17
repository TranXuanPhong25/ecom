from typing import Any
from langchain_core.tools import StructuredTool, tool

from app.utils.httpx_clients import ProductSearchSvcClient

def search_products(query: str, limit: int) -> Any :
    """ Get limited list of products have name similar query"""
    client = ProductSearchSvcClient()
    res = client.search(query, limit)
    return res

async def asearch_products(query: str, limit: int) -> Any:
    """ Async get limited list of products have name similar query"""
    client = ProductSearchSvcClient()
    res = await client.search(query, limit)
    return res

products_searcher = StructuredTool.from_function(func=search_products, coroutine=asearch_products)

