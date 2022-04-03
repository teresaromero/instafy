
import json

import asyncio
from fastapi import Request

from app.auth_storage import authStorage


status_stream_delay = 5  # second
status_stream_retry_timeout = 30000  # milisecond


async def login_event_generator(request: Request, client_id: str):
    while True:
        if await request.is_disconnected():
            print('Request disconnected')
            break

        if authStorage.hasClientID(client_id):
            a = authStorage.getAuth(client_id)
            authStorage.deleteAuth(client_id)
            print('Request completed. Disconnecting now')

            yield dict(data=json.dumps(a.__dict__))
            break

        await asyncio.sleep(status_stream_delay)
