
## Recommended Project Structure
- project
   - schemas ----------> pydantic models
      - __init__.py
      - post.py
      - user.py
   - routers ----------> routers
      - __init__.py
      - posts.py
      - usets.py
   - __init__.py
   - app.py
   - db.py


## How routers look like
```py
from fastapi import APIRouter, HTTPException, status

from project.db import db
from project.schemas.user import User, UserCreate

# use APIRouter() instead of FastAPI()
router = APIRouter()

# path decorator
@router.get("/")
# path operation function
async def all() -> list[User]:
   return list(db.users.values())
```

## How to import router in main app
```py
from fastapi import FastAPI

from project.routers.posts import router as posts_router
from project.routers.users import router as users_router

app = FastAPI()

app.include_router(posts_router, prefix="/posts", tags=["posts"])
app.include_router(users_router, prefix="/users", tags=["users"])
```

