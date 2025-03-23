package config;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.68.1)",
    comments = "Source: config.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class ConfigServiceGrpc {

  private ConfigServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "config.ConfigService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      config.Config.GetWorkflowResponse> getGetWorkflowFileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetWorkflowFile",
      requestType = com.google.protobuf.Empty.class,
      responseType = config.Config.GetWorkflowResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      config.Config.GetWorkflowResponse> getGetWorkflowFileMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, config.Config.GetWorkflowResponse> getGetWorkflowFileMethod;
    if ((getGetWorkflowFileMethod = ConfigServiceGrpc.getGetWorkflowFileMethod) == null) {
      synchronized (ConfigServiceGrpc.class) {
        if ((getGetWorkflowFileMethod = ConfigServiceGrpc.getGetWorkflowFileMethod) == null) {
          ConfigServiceGrpc.getGetWorkflowFileMethod = getGetWorkflowFileMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, config.Config.GetWorkflowResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetWorkflowFile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  config.Config.GetWorkflowResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ConfigServiceMethodDescriptorSupplier("GetWorkflowFile"))
              .build();
        }
      }
    }
    return getGetWorkflowFileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<config.Config.SetWorkflowRequest,
      config.Config.SetWorkflowResponse> getSetWorkflowFileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetWorkflowFile",
      requestType = config.Config.SetWorkflowRequest.class,
      responseType = config.Config.SetWorkflowResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<config.Config.SetWorkflowRequest,
      config.Config.SetWorkflowResponse> getSetWorkflowFileMethod() {
    io.grpc.MethodDescriptor<config.Config.SetWorkflowRequest, config.Config.SetWorkflowResponse> getSetWorkflowFileMethod;
    if ((getSetWorkflowFileMethod = ConfigServiceGrpc.getSetWorkflowFileMethod) == null) {
      synchronized (ConfigServiceGrpc.class) {
        if ((getSetWorkflowFileMethod = ConfigServiceGrpc.getSetWorkflowFileMethod) == null) {
          ConfigServiceGrpc.getSetWorkflowFileMethod = getSetWorkflowFileMethod =
              io.grpc.MethodDescriptor.<config.Config.SetWorkflowRequest, config.Config.SetWorkflowResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetWorkflowFile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  config.Config.SetWorkflowRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  config.Config.SetWorkflowResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ConfigServiceMethodDescriptorSupplier("SetWorkflowFile"))
              .build();
        }
      }
    }
    return getSetWorkflowFileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      config.Config.GetServicesResponse> getGetServicesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetServices",
      requestType = com.google.protobuf.Empty.class,
      responseType = config.Config.GetServicesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      config.Config.GetServicesResponse> getGetServicesMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, config.Config.GetServicesResponse> getGetServicesMethod;
    if ((getGetServicesMethod = ConfigServiceGrpc.getGetServicesMethod) == null) {
      synchronized (ConfigServiceGrpc.class) {
        if ((getGetServicesMethod = ConfigServiceGrpc.getGetServicesMethod) == null) {
          ConfigServiceGrpc.getGetServicesMethod = getGetServicesMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, config.Config.GetServicesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetServices"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  config.Config.GetServicesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ConfigServiceMethodDescriptorSupplier("GetServices"))
              .build();
        }
      }
    }
    return getGetServicesMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ConfigServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConfigServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConfigServiceStub>() {
        @java.lang.Override
        public ConfigServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConfigServiceStub(channel, callOptions);
        }
      };
    return ConfigServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ConfigServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConfigServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConfigServiceBlockingStub>() {
        @java.lang.Override
        public ConfigServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConfigServiceBlockingStub(channel, callOptions);
        }
      };
    return ConfigServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ConfigServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConfigServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConfigServiceFutureStub>() {
        @java.lang.Override
        public ConfigServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConfigServiceFutureStub(channel, callOptions);
        }
      };
    return ConfigServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void getWorkflowFile(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<config.Config.GetWorkflowResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetWorkflowFileMethod(), responseObserver);
    }

    /**
     */
    default void setWorkflowFile(config.Config.SetWorkflowRequest request,
        io.grpc.stub.StreamObserver<config.Config.SetWorkflowResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetWorkflowFileMethod(), responseObserver);
    }

    /**
     */
    default void getServices(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<config.Config.GetServicesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetServicesMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service ConfigService.
   */
  public static abstract class ConfigServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return ConfigServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service ConfigService.
   */
  public static final class ConfigServiceStub
      extends io.grpc.stub.AbstractAsyncStub<ConfigServiceStub> {
    private ConfigServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConfigServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConfigServiceStub(channel, callOptions);
    }

    /**
     */
    public void getWorkflowFile(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<config.Config.GetWorkflowResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetWorkflowFileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setWorkflowFile(config.Config.SetWorkflowRequest request,
        io.grpc.stub.StreamObserver<config.Config.SetWorkflowResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetWorkflowFileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getServices(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<config.Config.GetServicesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetServicesMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service ConfigService.
   */
  public static final class ConfigServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<ConfigServiceBlockingStub> {
    private ConfigServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConfigServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConfigServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public config.Config.GetWorkflowResponse getWorkflowFile(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetWorkflowFileMethod(), getCallOptions(), request);
    }

    /**
     */
    public config.Config.SetWorkflowResponse setWorkflowFile(config.Config.SetWorkflowRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetWorkflowFileMethod(), getCallOptions(), request);
    }

    /**
     */
    public config.Config.GetServicesResponse getServices(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetServicesMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service ConfigService.
   */
  public static final class ConfigServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<ConfigServiceFutureStub> {
    private ConfigServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConfigServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConfigServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<config.Config.GetWorkflowResponse> getWorkflowFile(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetWorkflowFileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<config.Config.SetWorkflowResponse> setWorkflowFile(
        config.Config.SetWorkflowRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetWorkflowFileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<config.Config.GetServicesResponse> getServices(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetServicesMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_WORKFLOW_FILE = 0;
  private static final int METHODID_SET_WORKFLOW_FILE = 1;
  private static final int METHODID_GET_SERVICES = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_WORKFLOW_FILE:
          serviceImpl.getWorkflowFile((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<config.Config.GetWorkflowResponse>) responseObserver);
          break;
        case METHODID_SET_WORKFLOW_FILE:
          serviceImpl.setWorkflowFile((config.Config.SetWorkflowRequest) request,
              (io.grpc.stub.StreamObserver<config.Config.SetWorkflowResponse>) responseObserver);
          break;
        case METHODID_GET_SERVICES:
          serviceImpl.getServices((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<config.Config.GetServicesResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getGetWorkflowFileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              config.Config.GetWorkflowResponse>(
                service, METHODID_GET_WORKFLOW_FILE)))
        .addMethod(
          getSetWorkflowFileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              config.Config.SetWorkflowRequest,
              config.Config.SetWorkflowResponse>(
                service, METHODID_SET_WORKFLOW_FILE)))
        .addMethod(
          getGetServicesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              config.Config.GetServicesResponse>(
                service, METHODID_GET_SERVICES)))
        .build();
  }

  private static abstract class ConfigServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ConfigServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return config.Config.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ConfigService");
    }
  }

  private static final class ConfigServiceFileDescriptorSupplier
      extends ConfigServiceBaseDescriptorSupplier {
    ConfigServiceFileDescriptorSupplier() {}
  }

  private static final class ConfigServiceMethodDescriptorSupplier
      extends ConfigServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    ConfigServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (ConfigServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ConfigServiceFileDescriptorSupplier())
              .addMethod(getGetWorkflowFileMethod())
              .addMethod(getSetWorkflowFileMethod())
              .addMethod(getGetServicesMethod())
              .build();
        }
      }
    }
    return result;
  }
}
