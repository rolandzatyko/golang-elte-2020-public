#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>

#include "query.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using query::QueryRequest;
using query::QueryResponse;
using query::QueryService;

class QueryServiceImpl final : public QueryService::Service {
  Status Query(ServerContext* context, const QueryRequest *request, QueryResponse *response) override {
    std::string prefix("r-");
    // TODO: complete the body of the server
    return Status::OK;
  }
};

void RunServer() {
  std::string server_address("0.0.0.0:54321");
  QueryServiceImpl service;

  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  ServerBuilder builder;
  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<Server> server(builder.BuildAndStart());
  std::cout << "Server listening on " << server_address << std::endl;
  server->Wait();
}

int main(int argc, char** argv) {
  RunServer();

  return 0;
}
