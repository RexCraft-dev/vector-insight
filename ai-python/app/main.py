# ai-python/app/main.py
from fastapi import FastAPI
from app.db import fetch

app = FastAPI()

@app.get("/health")
def health_check():
    return {"status": "ok"}

@app.get("/plan")
def plan():
    return {"strategy": "Play a conservative tee shot, aim center fairway."}

@app.get("/analyze_round")
async def analyze_round(user_id: int, round_id: int):
    shots = await fetch("""
        SELECT club, distance, result
        FROM shots
        WHERE round_id = $1
        ORDER BY shot_number;
    """, round_id)

    if not shots:
        return {"message": "No shots found for this round."}

    avg_distance = sum(s["distance"] for s in shots if s["distance"]) / len(shots)
    results = [s["result"] for s in shots if s["result"]]
    result_summary = {r: results.count(r) for r in set(results)}

    return {
        "user_id": user_id,
        "round_id": round_id,
        "avg_distance": round(avg_distance, 1),
        "result_summary": result_summary,
        "total_shots": len(shots)
    }
