from pydantic import BaseSettings
import os


class Settings(BaseSettings):
    ig_api_client_id: str = os.getenv("IG_API_CLIENT_ID")
    ig_api_client_secret: str = os.getenv("IG_API_CLIENT_SECRET")
    ig_api_grant_type: str = os.getenv("IG_API_GRANT_TYPE")
    ig_api_redirect_url: str = os.getenv("IG_API_REDIRECT_URL")
    ig_api_state: str = os.getenv("IG_API_STATE")
