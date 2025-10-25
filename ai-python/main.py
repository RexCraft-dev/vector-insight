from fastapi import FastAPI

app = FastAPI()

@app.get("/health")
def health_check():
    return {"status": "ok"}

@app.get("/plan")
def get_plan():
    return {
        "strategy": "Play a conservative tee shot, aim center fairway, approach from 130y with 8-iron."
    }