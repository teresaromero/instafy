from http import HTTPStatus
from fastapi import Depends, FastAPI, Header, Request
from fastapi.responses import RedirectResponse, JSONResponse, HTMLResponse
from sse_starlette.sse import EventSourceResponse
from app.auth_storage import authStorage
from app.config import Settings, get_settings

from app.ig_client import get_access_token, get_autorize_url, get_user_media
from app.streams import login_event_generator


app = FastAPI()


@app.get('/stream-login')
async def runStatus(
        client_id: str,
        request: Request
):
    event_generator = login_event_generator(request, client_id)
    return EventSourceResponse(event_generator)


@app.get("/login")
async def get_login(client_id: str = "", settings: Settings = Depends(get_settings)) -> dict:

    try:
        url = get_autorize_url(client_id, settings)
        return RedirectResponse(url)
    except Exception as err:
        return JSONResponse(err, HTTPStatus.BAD_REQUEST)


@app.get("/login/callback")
async def post_login_callback(code: str, state: str, settings: Settings = Depends(get_settings)) -> dict:

    if state == "" or code == "":
        return JSONResponse({"error": "invalid parameters code or state"}, HTTPStatus.BAD_REQUEST)

    try:
        res = get_access_token(code, settings)
    except Exception as err:
        return err

    authStorage.saveAuth(state, res["access_token"], res["user_id"])

    html_content = """
    <html>
        <head>
            <title>Instafy API</title>
        </head>
        <body>
            <p>Success!! you can continue at the cli</p>
        </body>
    </html>
    """
    return HTMLResponse(html_content)


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
            HTTPStatus.BAD_REQUEST
        )

    try:
        res = get_user_media(x_user_id, x_access_token)
        return JSONResponse(res, HTTPStatus.OK)
    except Exception as err:
        return JSONResponse(err, HTTPStatus.INTERNAL_SERVER_ERROR)
