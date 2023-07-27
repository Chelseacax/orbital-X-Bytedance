namespace go orbital

// for calculator's input / request 
struct Variable {
    1: i64 FirstInt
    2: i64 SecondInt
}

// for calculator's output / response
struct Answer {
    1: i64 Output
}

// Methods in "calculator" microservice
service CalculatorService {
    Answer Add(1: Variable Inputs) (api.post="/add");
    Answer Subtract(1: Variable Inputs) (api.post="/subtract");
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