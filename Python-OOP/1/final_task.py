from copy import deepcopy
from typing import List


class Server:
    count_servers = 0

    def __init__(self):
        self.ip = Server.count_servers
        self.buffer: List[Data] = []
        Server.count_servers += 1

    def send_data(self, data):
        if self in Router.servers:
            Router.buffer.append(data)

    def get_data(self):
        copy_buf = deepcopy(self.buffer)
        self.buffer = []

        return copy_buf

    def get_ip(self):
        return self.ip


class Data:
    def __init__(self, data, ip):
        self.data = data
        self.send_ip = ip


class Router:
    buffer: List[Data] = []
    servers: List[Server] = []

    def link(self, server: Server):
        self.servers.append(server)

    def unlink(self, server: Server):
        self.servers.remove(server)

    def send_data(self):
        for elem in self.buffer:
            for server in self.servers:
                if elem.send_ip == server.ip:
                    server.buffer.append(elem)
        self.buffer = []