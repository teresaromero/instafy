from http import HTTPStatus
from fastapi import Depends, FastAPI
from fastapi.responses import RedirectResponse, JSONResponse
from functools import lru_cache
import requests

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


@app.get("/login/callback")
async def post_login_callback(code: str, state: str, settings: Settings = Depends(get_settings)) -> dict:

    if state != settings.ig_api_state:
        return JSONResponse({"error": "invalid origin"}, HTTPStatus.UNAUTHORIZED)

    base_url = "https://api.instagram.com/oauth/access_token"
    data = {
        "client_id":     settings.ig_api_client_id,
        "client_secret": settings.ig_api_client_secret,
        "grant_type":    settings.ig_api_grant_type,
        "redirect_uri":  settings.ig_api_redirect_url,
        "code":          code,
    }

    response = requests.post(url=base_url, data=data)
    body = response.json()
    if response.status_code == HTTPStatus.OK:
        return JSONResponse(body)
    return JSONResponse({"err": body}, HTTPStatus.INTERNAL_SERVER_ERROR)
