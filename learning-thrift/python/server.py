from thrift.protocol import TBinaryProtocol
from thrift.server import TServer
from thrift.transport import TSocket, TTransport

from example import format_data, ttypes

RPC_HOST = '127.0.0.1'
RPC_PORT = 8080

class FormatDataHandler():
    def artificial_idiot(self, data):
        print(type(data.text))
        return ttypes.Data(data.text.strip('么吗？?'))


if __name__ == '__main__':
    handler = FormatDataHandler()

    processor = format_data.Processor(handler)
    transport = TSocket.TServerSocket(RPC_HOST, RPC_PORT)
    tfactory = TTransport.TBufferedTransportFactory()
    pfactory = TBinaryProtocol.TBinaryProtocolFactory()

    rpcServer = TServer.TSimpleServer(processor,transport, tfactory, pfactory)

    print('Starting the rpc server at', RPC_HOST,':', RPC_PORT)
    rpcServer.serve()
