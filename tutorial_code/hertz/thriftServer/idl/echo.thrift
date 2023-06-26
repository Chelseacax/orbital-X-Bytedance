namespace go api

struct EchoReq {
	1: string message
}

struct EchoResp {
	1: string message
}

service Echo {
    EchoResp echo(1: EchoReq req)
}
