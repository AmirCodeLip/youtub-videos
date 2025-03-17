def singleton(cls):
    instances = {}

    def get_instance(*args, **kwargs):
        if cls not in instances:
            instances[cls] = cls(*args, **kwargs)
        return instances[cls]

    return get_instance


@singleton
class ServerConnection:
    def __init__(self, uri=None, ports=None, user_name=None, password=None):
        if uri is not None:
            self.uri = uri
        if ports is not None:
            self.ports = ports
        if user_name is not None:
            self.user_name = user_name
        if password is not None:
            self.password = password

    def connect(self) -> None:
        print(
            f"we are connect to {self.uri}:{self.ports} {self.user_name} {self.password}"
        )


server1 = ServerConnection("192.168.1.1", 80, "amir", "password1234")
# server1.connect()


server2 = ServerConnection()
server2.connect()
server2.user_name = "user2"

server1.connect()
