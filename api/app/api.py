from fastapi import Depends, FastAPI
from fastapi.responses import RedirectResponse, JSONResponse
from functools import lru_cache

from .config import Settings


@lru_cache()
def get_settings():
    return Settings()


app = FastAPI()


@app.get("/login")
async def get_login(settings: Settings = Depends(get_settings)) -> dict:

    base_url = "https://api.instagram.com/oauth/authorize?"
    params = [
        f'client_id={settings.ig_api_client_id}',
        f'redirect_uri={settings.ig_api_redirect_url}',
        'scope=user_profile,user_media',
        'response_type=code',
        'state=hola'
    ]

    url = base_url + "&".join(params)

    return RedirectResponse(url)


@app.post("/login/callback")
async def post_login_callback(settings: Settings = Depends(get_settings)) -> dict:

    return JSONResponse({"error": "WIP"})
