package stego.hub.grpc.wrapper;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.68.1)",
    comments = "Source: steg_service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class StegServiceGrpc {

  private StegServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "steg_service.StegService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest,
      stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse> getExecuteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Execute",
      requestType = stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest.class,
      responseType = stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest,
      stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse> getExecuteMethod() {
    io.grpc.MethodDescriptor<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest, stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse> getExecuteMethod;
    if ((getExecuteMethod = StegServiceGrpc.getExecuteMethod) == null) {
      synchronized (StegServiceGrpc.class) {
        if ((getExecuteMethod = StegServiceGrpc.getExecuteMethod) == null) {
          StegServiceGrpc.getExecuteMethod = getExecuteMethod =
              io.grpc.MethodDescriptor.<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest, stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Execute"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse.getDefaultInstance()))
              .setSchemaDescriptor(new StegServiceMethodDescriptorSupplier("Execute"))
              .build();
        }
      }
    }
    return getExecuteMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo> getGetStegServiceInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetStegServiceInfo",
      requestType = com.google.protobuf.Empty.class,
      responseType = stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.google.protobuf.Empty,
      stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo> getGetStegServiceInfoMethod() {
    io.grpc.MethodDescriptor<com.google.protobuf.Empty, stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo> getGetStegServiceInfoMethod;
    if ((getGetStegServiceInfoMethod = StegServiceGrpc.getGetStegServiceInfoMethod) == null) {
      synchronized (StegServiceGrpc.class) {
        if ((getGetStegServiceInfoMethod = StegServiceGrpc.getGetStegServiceInfoMethod) == null) {
          StegServiceGrpc.getGetStegServiceInfoMethod = getGetStegServiceInfoMethod =
              io.grpc.MethodDescriptor.<com.google.protobuf.Empty, stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetStegServiceInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo.getDefaultInstance()))
              .setSchemaDescriptor(new StegServiceMethodDescriptorSupplier("GetStegServiceInfo"))
              .build();
        }
      }
    }
    return getGetStegServiceInfoMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static StegServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StegServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StegServiceStub>() {
        @java.lang.Override
        public StegServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StegServiceStub(channel, callOptions);
        }
      };
    return StegServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static StegServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StegServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StegServiceBlockingStub>() {
        @java.lang.Override
        public StegServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StegServiceBlockingStub(channel, callOptions);
        }
      };
    return StegServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static StegServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StegServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StegServiceFutureStub>() {
        @java.lang.Override
        public StegServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StegServiceFutureStub(channel, callOptions);
        }
      };
    return StegServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void execute(stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest request,
        io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getExecuteMethod(), responseObserver);
    }

    /**
     */
    default void getStegServiceInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetStegServiceInfoMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service StegService.
   */
  public static abstract class StegServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return StegServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service StegService.
   */
  public static final class StegServiceStub
      extends io.grpc.stub.AbstractAsyncStub<StegServiceStub> {
    private StegServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StegServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StegServiceStub(channel, callOptions);
    }

    /**
     */
    public void execute(stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest request,
        io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getExecuteMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getStegServiceInfo(com.google.protobuf.Empty request,
        io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetStegServiceInfoMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service StegService.
   */
  public static final class StegServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<StegServiceBlockingStub> {
    private StegServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StegServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StegServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse execute(stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getExecuteMethod(), getCallOptions(), request);
    }

    /**
     */
    public stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo getStegServiceInfo(com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetStegServiceInfoMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service StegService.
   */
  public static final class StegServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<StegServiceFutureStub> {
    private StegServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StegServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StegServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse> execute(
        stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getExecuteMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo> getStegServiceInfo(
        com.google.protobuf.Empty request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetStegServiceInfoMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_EXECUTE = 0;
  private static final int METHODID_GET_STEG_SERVICE_INFO = 1;

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
        case METHODID_EXECUTE:
          serviceImpl.execute((stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest) request,
              (io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse>) responseObserver);
          break;
        case METHODID_GET_STEG_SERVICE_INFO:
          serviceImpl.getStegServiceInfo((com.google.protobuf.Empty) request,
              (io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo>) responseObserver);
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
          getExecuteMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceRequest,
              stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceResponse>(
                service, METHODID_EXECUTE)))
        .addMethod(
          getGetStegServiceInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.google.protobuf.Empty,
              stego.hub.grpc.wrapper.StegServiceOuterClass.StegServiceInfo>(
                service, METHODID_GET_STEG_SERVICE_INFO)))
        .build();
  }

  private static abstract class StegServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    StegServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return stego.hub.grpc.wrapper.StegServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("StegService");
    }
  }

  private static final class StegServiceFileDescriptorSupplier
      extends StegServiceBaseDescriptorSupplier {
    StegServiceFileDescriptorSupplier() {}
  }

  private static final class StegServiceMethodDescriptorSupplier
      extends StegServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    StegServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (StegServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new StegServiceFileDescriptorSupplier())
              .addMethod(getExecuteMethod())
              .addMethod(getGetStegServiceInfoMethod())
              .build();
        }
      }
    }
    return result;
  }
}
