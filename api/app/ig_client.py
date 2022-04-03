from datetime import datetime, timedelta
from http import HTTPStatus

import requests

from .config import Settings


IG_API_VERSION = "v12.0"
IG_API_OAUTH_BASEURL = "https://api.instagram.com/oauth"
IG_API_OAUTH_AUTHORIZE = f"{IG_API_OAUTH_BASEURL}/authorize"
IG_API_OAUTH_ACCESS_TOKEN = f"{IG_API_OAUTH_BASEURL}/access_token"
IG_API_GRAPH_BASEURL = f"https://graph.instagram.com/{IG_API_VERSION}"


def IG_API_GRAPH_MEDIA(
    user_id): return f'{IG_API_GRAPH_BASEURL}/{user_id}/media'


def get_autorize_url(client_id: str, settings: Settings) -> str:

    if settings.ig_api_client_id == "" or settings.ig_api_redirect_url == "":
        raise Exception("invalid params")

    base_url = f"{IG_API_OAUTH_AUTHORIZE}?"
    params = [
        f'client_id={settings.ig_api_client_id}',
        f'redirect_uri={settings.ig_api_redirect_url}',
        'scope=user_profile,user_media',
        'response_type=code',
        f'state={client_id}'
    ]

    return base_url + "&".join(params)


def get_access_token(code: str, settings: Settings) -> dict:
    response = requests.post(
        url=f"{IG_API_OAUTH_ACCESS_TOKEN}",
        data={
            "client_id":     settings.ig_api_client_id,
            "client_secret": settings.ig_api_client_secret,
            "grant_type":    settings.ig_api_grant_type,
            "redirect_uri":  settings.ig_api_redirect_url,
            "code":          code,
        })

    if response.status_code != HTTPStatus.OK:
        raise Exception("Unable to get access_token")

    return response.json()


def get_user_media(user_id: str, access_token: str) -> dict:

    today = datetime.today()
    default_since = today - timedelta(days=60)
    default_until = today

    params = [
        f'access_token={access_token}',
        "limit=10",
        f"since={default_since}",
        f"until={default_until}",
        "fields=caption,media_type,media_url,timestamp"
    ]
    url = f'{IG_API_GRAPH_MEDIA(user_id)}?{"&".join(params)}'

    response = requests.get(url)
    if response.status_code != HTTPStatus.OK:
        error = response.json()
        raise Exception("error", str(error))

    body = response.json()
    return body["data"]
