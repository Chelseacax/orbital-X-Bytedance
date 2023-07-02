namespace go example

struct Variable {
    1: i64 FirstInt
    2: i64 SecondInt
}

struct Summer {
    1: i64 Sum
}

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}

service ExampleService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    Summer Add(1: Variable inputs) (api.post="/add");
}