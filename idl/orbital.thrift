// for calculator's input
struct Variable {
    1: i64 FirstInt
    2: i64 SecondInt
}

// for calculator's output
struct Result {
    1: i64 Output
}

// Methods in "calculator" microservice
service CalculatorService {
    Result Add(1: Variable Inputs) (api.post="/add");
}

// for HelloMethod's input
struct HelloRequest {
    1: string Name (api.query="name");
}

// for HelloMethod's output
struct HelloResponse {
    1: string ResponseBody;
}

// Methods in "hello" microservice
service HelloService {
    HelloResponse HelloMethod(1: HelloRequest Request) (api.get="/hello");
}