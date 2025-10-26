# ai-python/app/db.py
import os
import asyncpg

# Default to the internal Docker network connection
DB_URL = os.getenv("DB_URL", "postgres://vector:vectorpass@db:5432/vector_insight")

_pool = None

async def get_pool():
    global _pool
    if _pool is None:
        _pool = await asyncpg.create_pool(DB_URL)
    return _pool

async def fetch(query: str, *args):
    pool = await get_pool()
    async with pool.acquire() as conn:
        return await conn.fetch(query, *args)

async def execute(query: str, *args):
    pool = await get_pool()
    async with pool.acquire() as conn:
        return await conn.execute(query, *args)
