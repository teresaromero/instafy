from http import HTTPStatus
from datetime import datetime, timedelta
from fastapi import Depends, FastAPI, Header
from fastapi.responses import RedirectResponse, JSONResponse, HTMLResponse
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
        f'state={settings.ig_api_state}'
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


@app.get("/")
async def root() -> str:
    html_content = """
    <html>
        <head>
            <title>Instafy API</title>
        </head>
        <body>
            <h1>🎉</h1>
            <a href="/login">Login with Instagram</a>
        </body>
    </html>
    """
    return HTMLResponse(html_content)


@app.get("/media")
async def get_media(
    x_access_token: str = Header(default=""),
    x_user_id: str = Header(default="")
) -> dict:

    if x_access_token == "" or x_user_id == "":
        return JSONResponse(
            {"error": "x-access-token or x-user-id headers missing"},
            HTTPStatus.UNPROCESSABLE_ENTITY
        )

    today = datetime.today()
    default_since = today - timedelta(days=60)
    default_until = today

    api_version = "v12.0"

    base_url = f'https://graph.instagram.com/{api_version}'
    endpoint = f'{x_user_id}/media'
    params = [
        f'access_token={x_access_token}',
        "limit=10",
        f"since={default_since}",
        f"until={default_until}",
        "fields=caption,media_type,media_url,timestamp"
    ]
    url = f'{base_url}/{endpoint}?{"&".join(params)}'

    response = requests.get(url)
    if response.status_code != HTTPStatus.OK:
        return JSONResponse(response.json(), response.status_code)

    body = response.json()
    data = body["data"]

    return JSONResponse(data)
