package stego.hub.grpc.wrapper;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.68.1)",
    comments = "Source: steg_analysis.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class StegAnalysisServiceGrpc {

  private StegAnalysisServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "steg_analysis.StegAnalysisService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest,
      stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse> getExecuteMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Execute",
      requestType = stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest.class,
      responseType = stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest,
      stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse> getExecuteMethod() {
    io.grpc.MethodDescriptor<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest, stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse> getExecuteMethod;
    if ((getExecuteMethod = StegAnalysisServiceGrpc.getExecuteMethod) == null) {
      synchronized (StegAnalysisServiceGrpc.class) {
        if ((getExecuteMethod = StegAnalysisServiceGrpc.getExecuteMethod) == null) {
          StegAnalysisServiceGrpc.getExecuteMethod = getExecuteMethod =
              io.grpc.MethodDescriptor.<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest, stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Execute"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse.getDefaultInstance()))
              .setSchemaDescriptor(new StegAnalysisServiceMethodDescriptorSupplier("Execute"))
              .build();
        }
      }
    }
    return getExecuteMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static StegAnalysisServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StegAnalysisServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StegAnalysisServiceStub>() {
        @java.lang.Override
        public StegAnalysisServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StegAnalysisServiceStub(channel, callOptions);
        }
      };
    return StegAnalysisServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static StegAnalysisServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StegAnalysisServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StegAnalysisServiceBlockingStub>() {
        @java.lang.Override
        public StegAnalysisServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StegAnalysisServiceBlockingStub(channel, callOptions);
        }
      };
    return StegAnalysisServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static StegAnalysisServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<StegAnalysisServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<StegAnalysisServiceFutureStub>() {
        @java.lang.Override
        public StegAnalysisServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new StegAnalysisServiceFutureStub(channel, callOptions);
        }
      };
    return StegAnalysisServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void execute(stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest request,
        io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getExecuteMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service StegAnalysisService.
   */
  public static abstract class StegAnalysisServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return StegAnalysisServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service StegAnalysisService.
   */
  public static final class StegAnalysisServiceStub
      extends io.grpc.stub.AbstractAsyncStub<StegAnalysisServiceStub> {
    private StegAnalysisServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StegAnalysisServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StegAnalysisServiceStub(channel, callOptions);
    }

    /**
     */
    public void execute(stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest request,
        io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getExecuteMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service StegAnalysisService.
   */
  public static final class StegAnalysisServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<StegAnalysisServiceBlockingStub> {
    private StegAnalysisServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StegAnalysisServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StegAnalysisServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse execute(stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getExecuteMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service StegAnalysisService.
   */
  public static final class StegAnalysisServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<StegAnalysisServiceFutureStub> {
    private StegAnalysisServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StegAnalysisServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new StegAnalysisServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse> execute(
        stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getExecuteMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_EXECUTE = 0;

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
          serviceImpl.execute((stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest) request,
              (io.grpc.stub.StreamObserver<stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse>) responseObserver);
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
              stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisRequest,
              stego.hub.grpc.wrapper.StegAnalysis.StegAnalysisResponse>(
                service, METHODID_EXECUTE)))
        .build();
  }

  private static abstract class StegAnalysisServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    StegAnalysisServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return stego.hub.grpc.wrapper.StegAnalysis.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("StegAnalysisService");
    }
  }

  private static final class StegAnalysisServiceFileDescriptorSupplier
      extends StegAnalysisServiceBaseDescriptorSupplier {
    StegAnalysisServiceFileDescriptorSupplier() {}
  }

  private static final class StegAnalysisServiceMethodDescriptorSupplier
      extends StegAnalysisServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    StegAnalysisServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (StegAnalysisServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new StegAnalysisServiceFileDescriptorSupplier())
              .addMethod(getExecuteMethod())
              .build();
        }
      }
    }
    return result;
  }
}
