# Product: PortConnection
class PortConnection:

    def __init__(self):
        self.port_type = None
        self.port_name = None
        self.baud_rate = None
        self.parity = None
        self.timeout = None
        self.buffer_size = None
        self.ip_address = None
        self.port_number = None

    def __str__(self):
        if self.port_type == "Serial":
            return (
                f"SerialPort: name={self.port_name}, baud_rate={self.baud_rate}, "
                f"parity={self.parity}, timeout={self.timeout}, buffer_size={self.buffer_size}"
            )
        elif self.port_type == "TCP/IP":
            return f"TCP/IP Port: ip={self.ip_address}, port={self.port_number}, timeout={self.timeout}"
        else:
            return "Unknown port type"


class PortConnectionBuilder:

    def __init__(self, port_type):
        self.connection = PortConnection()
        self.connection.port_type = port_type

    def set_port_name(self, port_name):
        self.connection.port_name = port_name
        return self

    def set_baud_rate(self, baud_rate):
        self.connection.baud_rate = baud_rate
        return self

    def set_parity(self, parity):
        self.connection.parity = parity
        return self

    def set_timeout(self, timeout):
        self.connection.timeout = timeout
        return self

    def set_buffer_size(self, buffer_size):
        self.connection.buffer_size = buffer_size
        return self

    def set_port_number(self, port_number):
        self.connection.port_number = port_number
        return self

    def set_ip_address(self, ip_address):
        self.connection.ip_address = ip_address
        return self

    def set_baud_rate(self, baud_rate):
        self.connection.baud_rate = baud_rate
        return self

    def connect(self):
        print(self.connection.__str__())


tcp = (
    PortConnectionBuilder("TCP/IP")
    .set_ip_address("127.0.0.1")
    .set_port_number("80")
    .set_parity("none")
    .set_timeout(1000)
)
tcp.connect()

serial = (
    PortConnectionBuilder("Serial")
    .set_port_name("COM1")
    .set_baud_rate("9600")
    .set_timeout(10)
)

serial.connect()