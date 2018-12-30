import sys

from thrift.protocol import TBinaryProtocol
from thrift.transport import TSocket, TTransport

from example.format_data import Client, Data

RPC_HOST = '127.0.0.1'
RPC_PORT = 8080

tsocket = TSocket.TSocket(RPC_HOST, RPC_PORT)
transport = TTransport.TBufferedTransport(tsocket)
protocol = TBinaryProtocol.TBinaryProtocol(transport)
client = Client(protocol)

data = Data(sys.argv[1])
transport.open()

print(type(client.artificial_idiot(data).text))
print(client.artificial_idiot(data).text)
