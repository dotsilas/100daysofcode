from fastapi import FastAPI
from pydantic import BaseModel, ValidationError
from datetime import date 
from enum import Enum

app = FastAPI()

# Standard field scalar types
class Person(BaseModel):
    first_name: str
    last_name: str
    age: int

# Standard field complex types
class Gender(str, Enum):
    MALE = "MALE"
    FEMALE = "FEMALE"
    NON_BYNARY = "NON_BYNARY"


class Address(BaseModel):
    country: str = "USA"
    # optional field
    street: str | None = None


class Person2(BaseModel):
    first_name: str
    last_name: str
    age: int
    gender: Gender
    birthdate: date
    interests: list[str]
    address: Address

try:
    Person2(
        first_name="John",
        last_name="Maxwell",
        age=25,
        gender=Gender.MALE,
        birthdate="1996-01-01",
        interests=["rock music", "painting"],
        address={"country": "Canadá", "street": "13"},
    )
except ValidationError as e:
    print(str(e))


@app.get("/")
def get_hello():
    return {"hello": "world"}
