include "base.thrift"
namespace go kitex.test.server

struct EchoReq {
    1: required string Msg,
    255: base.Base Base,
}
struct EchoResp {
    1: required string Msg,
    255: base.BaseResp BaseResp,
}
service EchoService {
    EchoResp EchoMethod(1: EchoReq req),
}

