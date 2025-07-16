from pydantic import BaseModel


class InputModel(BaseModel):
    prompt: str
