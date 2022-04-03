class Auth:
    def __init__(self, access_token: str, user_id: str) -> None:
        self.access_token = access_token
        self.user_id = user_id


class AuthStorage:
    def __init__(self) -> None:
        self.data = dict()

    def saveAuth(self, client_id: str, access_token: str, user_id: str):
        a = Auth(access_token, user_id)
        self.data[client_id] = a

    def getAuth(self, client_id: str) -> Auth:
        if client_id in self.data:
            return self.data[client_id]
        return None

    def deleteAuth(self, client_id: str):
        if client_id in self.data:
            del self.data[client_id]

    def hasClientID(self, client_id: str) -> bool:
        return client_id in self.data


# temp storage for auths
authStorage = AuthStorage()
