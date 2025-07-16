from fastapi import FastAPI

from schemas import InputModel
from generator import generate_image


app = FastAPI(title="Commercial studio photots generator", docs_url=None)


@app.post("/generate")
def generate(input_data: InputModel):
    generate_image(input_data.prompt)
