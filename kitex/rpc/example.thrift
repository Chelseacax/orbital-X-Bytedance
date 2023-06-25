// File1：example.thrift
namespace go test
include "base.thrift"

struct MyReq{
    1:required string input
    2:required base.BaseReq baseReq
}

service MyService{
    string Hello(1:required MyReq req)
}

// File2：base.thrift
struct BaseReq{
    1:required string name
}

