# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import enrichment_pb2 as enrichment__pb2

GRPC_GENERATED_VERSION = '1.65.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.66.0'
SCHEDULED_RELEASE_DATE = 'August 6, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in enrichment_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class EnrichmentStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.IPToHost = channel.unary_unary(
                '/enrichment.Enrichment/IPToHost',
                request_serializer=enrichment__pb2.IP.SerializeToString,
                response_deserializer=enrichment__pb2.Host.FromString,
                _registered_method=True)
        self.HostToIP = channel.unary_unary(
                '/enrichment.Enrichment/HostToIP',
                request_serializer=enrichment__pb2.Host.SerializeToString,
                response_deserializer=enrichment__pb2.IP.FromString,
                _registered_method=True)


class EnrichmentServicer(object):
    """Missing associated documentation comment in .proto file."""

    def IPToHost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HostToIP(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EnrichmentServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'IPToHost': grpc.unary_unary_rpc_method_handler(
                    servicer.IPToHost,
                    request_deserializer=enrichment__pb2.IP.FromString,
                    response_serializer=enrichment__pb2.Host.SerializeToString,
            ),
            'HostToIP': grpc.unary_unary_rpc_method_handler(
                    servicer.HostToIP,
                    request_deserializer=enrichment__pb2.Host.FromString,
                    response_serializer=enrichment__pb2.IP.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'enrichment.Enrichment', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('enrichment.Enrichment', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Enrichment(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def IPToHost(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/enrichment.Enrichment/IPToHost',
            enrichment__pb2.IP.SerializeToString,
            enrichment__pb2.Host.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def HostToIP(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/enrichment.Enrichment/HostToIP',
            enrichment__pb2.Host.SerializeToString,
            enrichment__pb2.IP.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
