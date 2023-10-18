from enum import Enum
from fastapi import Body, FastAPI, File, Form, Path, UploadFile
from pydantic import BaseModel

app = FastAPI()

# basic end-point
@app.get("/")
def hello_world():
    return {"hello":"world!"}

# path parameters
@app.get("/user/{id}")
# parameters typing
def get_user(id: int):
    return {"user_id": f"{id}"}

# Limiting values for a parameter
class userType(str, Enum):
    STANDARD = "standard"
    ADMIN = "admin"

@app.get("/user/{type}/{id}")
# use a defined enumerate class that implements Enum class as parameter type
def get_user_with_type(type: userType, id: int):
    return {
            "usertype": f"{type}",
            "user_id": f"{id}"
            }

# Advanced validation with Path function from fastapi
@app.get("/users/{id}")
# use of Path as default parameter, elipsis syntax "..." means "not default value", therefore, "value is required"
# Path() allows:
# Numeric validation: gt, ge, lt, le
# String validation: min_length, max_length, pattern (regex) 
def get_user_with_path_validation(id: int = Path(...,ge=1)):
    return {"id": id}

# Query parameters == dymanic parameters after ? in URL
# if a parameter of functions isn't present in the path, automatically is query param
@app.get("/users")
def get_user_query(page: int = 1, size: int = 10):
    return {"page": page, "size": size}

# To define required query parameters simply leaves out default value
class UserFormats(str, Enum):
    FULL = "full"
    SHORT = "short"

@app.get("/users_query")
def get_user_query_required(format: UserFormats):
    return {"format": format}

# Request Body
# like query parameters but always use Body() function as default value
@app.post("/users_body")
def create_User(name: str = Body(...), age: int = Body(...)):
    return {"name": name, "age": age}

# Using pydantic models
# Every model inherits from BaseModel class
class User(BaseModel):
    name: str
    age: int

@app.post("users_models")
# use your Model (that inherits from BaseModel) as type of parameter
def create_user_pydantic_method(user: User):
    # if you return a model, fastapi automatically convert it to JSON to produce HTTP response
    return user

# use python-multipart (pip3 install python-multipart) to handle forms and files

# form data
@app.post("/users_form")
def create_User_form(name: str = Form(...), age: int = Form(...)):
    return {"name": name, "age": age}

# files, little files can be managed with File() function
@app.post("/files")
def upload_file(file: bytes = File(...)):
    return {"file_size": len(file)}

# For Big Files is more recommended use UploadFile class as type instead of bytes
@app.post("/big_files")
# Acepting multiple files typehinting parameter like a list of UploadFile
def upload_big_file(files: list[UploadFile] = File(...)):
    return [{"file_name": file.filename, "content_type": file.content_type} for file in files]

## CONTINUE...
# headers
# cookies
# request object
# path operation parameters
# response parameter
# raising http errors
# custom response
# structuring a project


