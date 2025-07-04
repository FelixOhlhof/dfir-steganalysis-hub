// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: steg_service.proto
#ifndef GRPC_steg_5fservice_2eproto__INCLUDED
#define GRPC_steg_5fservice_2eproto__INCLUDED

#include "steg_service.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace steg_service {

class StegService final {
 public:
  static constexpr char const* service_full_name() {
    return "steg_service.StegService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status Execute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::steg_service::StegServiceResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceResponse>> AsyncExecute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceResponse>>(AsyncExecuteRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceResponse>> PrepareAsyncExecute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceResponse>>(PrepareAsyncExecuteRaw(context, request, cq));
    }
    virtual ::grpc::Status GetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::steg_service::StegServiceInfo* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceInfo>> AsyncGetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceInfo>>(AsyncGetStegServiceInfoRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceInfo>> PrepareAsyncGetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceInfo>>(PrepareAsyncGetStegServiceInfoRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void Execute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest* request, ::steg_service::StegServiceResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void Execute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest* request, ::steg_service::StegServiceResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      virtual void GetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty* request, ::steg_service::StegServiceInfo* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty* request, ::steg_service::StegServiceInfo* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceResponse>* AsyncExecuteRaw(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceResponse>* PrepareAsyncExecuteRaw(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceInfo>* AsyncGetStegServiceInfoRaw(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::steg_service::StegServiceInfo>* PrepareAsyncGetStegServiceInfoRaw(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status Execute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::steg_service::StegServiceResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceResponse>> AsyncExecute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceResponse>>(AsyncExecuteRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceResponse>> PrepareAsyncExecute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceResponse>>(PrepareAsyncExecuteRaw(context, request, cq));
    }
    ::grpc::Status GetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::steg_service::StegServiceInfo* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceInfo>> AsyncGetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceInfo>>(AsyncGetStegServiceInfoRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceInfo>> PrepareAsyncGetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceInfo>>(PrepareAsyncGetStegServiceInfoRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void Execute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest* request, ::steg_service::StegServiceResponse* response, std::function<void(::grpc::Status)>) override;
      void Execute(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest* request, ::steg_service::StegServiceResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void GetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty* request, ::steg_service::StegServiceInfo* response, std::function<void(::grpc::Status)>) override;
      void GetStegServiceInfo(::grpc::ClientContext* context, const ::google::protobuf::Empty* request, ::steg_service::StegServiceInfo* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceResponse>* AsyncExecuteRaw(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceResponse>* PrepareAsyncExecuteRaw(::grpc::ClientContext* context, const ::steg_service::StegServiceRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceInfo>* AsyncGetStegServiceInfoRaw(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::steg_service::StegServiceInfo>* PrepareAsyncGetStegServiceInfoRaw(::grpc::ClientContext* context, const ::google::protobuf::Empty& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_Execute_;
    const ::grpc::internal::RpcMethod rpcmethod_GetStegServiceInfo_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status Execute(::grpc::ServerContext* context, const ::steg_service::StegServiceRequest* request, ::steg_service::StegServiceResponse* response);
    virtual ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* context, const ::google::protobuf::Empty* request, ::steg_service::StegServiceInfo* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_Execute : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_Execute() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_Execute() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Execute(::grpc::ServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestExecute(::grpc::ServerContext* context, ::steg_service::StegServiceRequest* request, ::grpc::ServerAsyncResponseWriter< ::steg_service::StegServiceResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_GetStegServiceInfo : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_GetStegServiceInfo() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_GetStegServiceInfo() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetStegServiceInfo(::grpc::ServerContext* context, ::google::protobuf::Empty* request, ::grpc::ServerAsyncResponseWriter< ::steg_service::StegServiceInfo>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_Execute<WithAsyncMethod_GetStegServiceInfo<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_Execute : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_Execute() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::steg_service::StegServiceRequest, ::steg_service::StegServiceResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::steg_service::StegServiceRequest* request, ::steg_service::StegServiceResponse* response) { return this->Execute(context, request, response); }));}
    void SetMessageAllocatorFor_Execute(
        ::grpc::MessageAllocator< ::steg_service::StegServiceRequest, ::steg_service::StegServiceResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::steg_service::StegServiceRequest, ::steg_service::StegServiceResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_Execute() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Execute(::grpc::ServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Execute(
      ::grpc::CallbackServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_GetStegServiceInfo : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_GetStegServiceInfo() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::google::protobuf::Empty, ::steg_service::StegServiceInfo>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::google::protobuf::Empty* request, ::steg_service::StegServiceInfo* response) { return this->GetStegServiceInfo(context, request, response); }));}
    void SetMessageAllocatorFor_GetStegServiceInfo(
        ::grpc::MessageAllocator< ::google::protobuf::Empty, ::steg_service::StegServiceInfo>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::google::protobuf::Empty, ::steg_service::StegServiceInfo>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_GetStegServiceInfo() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetStegServiceInfo(
      ::grpc::CallbackServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_Execute<WithCallbackMethod_GetStegServiceInfo<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_Execute : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_Execute() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_Execute() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Execute(::grpc::ServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_GetStegServiceInfo : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_GetStegServiceInfo() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_GetStegServiceInfo() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_Execute : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_Execute() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_Execute() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Execute(::grpc::ServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestExecute(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetStegServiceInfo : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_GetStegServiceInfo() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_GetStegServiceInfo() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetStegServiceInfo(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_Execute : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_Execute() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->Execute(context, request, response); }));
    }
    ~WithRawCallbackMethod_Execute() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status Execute(::grpc::ServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* Execute(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_GetStegServiceInfo : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_GetStegServiceInfo() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->GetStegServiceInfo(context, request, response); }));
    }
    ~WithRawCallbackMethod_GetStegServiceInfo() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetStegServiceInfo(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_Execute : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_Execute() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::steg_service::StegServiceRequest, ::steg_service::StegServiceResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::steg_service::StegServiceRequest, ::steg_service::StegServiceResponse>* streamer) {
                       return this->StreamedExecute(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_Execute() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status Execute(::grpc::ServerContext* /*context*/, const ::steg_service::StegServiceRequest* /*request*/, ::steg_service::StegServiceResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedExecute(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::steg_service::StegServiceRequest,::steg_service::StegServiceResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetStegServiceInfo : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_GetStegServiceInfo() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::google::protobuf::Empty, ::steg_service::StegServiceInfo>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::google::protobuf::Empty, ::steg_service::StegServiceInfo>* streamer) {
                       return this->StreamedGetStegServiceInfo(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_GetStegServiceInfo() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetStegServiceInfo(::grpc::ServerContext* /*context*/, const ::google::protobuf::Empty* /*request*/, ::steg_service::StegServiceInfo* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetStegServiceInfo(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::google::protobuf::Empty,::steg_service::StegServiceInfo>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_Execute<WithStreamedUnaryMethod_GetStegServiceInfo<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_Execute<WithStreamedUnaryMethod_GetStegServiceInfo<Service > > StreamedService;
};

}  // namespace steg_service


#endif  // GRPC_steg_5fservice_2eproto__INCLUDED
