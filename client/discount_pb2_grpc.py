# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import discount_pb2 as discount__pb2


class DiscountServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Get = channel.unary_unary(
        '/discount.DiscountService/Get',
        request_serializer=discount__pb2.DiscountResquest.SerializeToString,
        response_deserializer=discount__pb2.Discount.FromString,
        )


class DiscountServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Get(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_DiscountServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Get': grpc.unary_unary_rpc_method_handler(
          servicer.Get,
          request_deserializer=discount__pb2.DiscountResquest.FromString,
          response_serializer=discount__pb2.Discount.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'discount.DiscountService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
