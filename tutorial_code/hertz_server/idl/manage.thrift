namespace go user.management

struct QueryUserRequest {
    1: string Num;
}

struct QueryUserResponse {
    1: bool   Exist;
    2: string Num;
    3: string Name;
    4: string Gender;
}

struct AddUserRequest {
    1: string Num;
    2: string Name;
    3: string Gender;
}

struct AddUserResponse {
    1: bool Ok;
    2: string Msg;
}

service UserManagement {
    QueryUserResponse queryUser(1: QueryUserRequest req);
    AddUserResponse AddUser(1: AddUserRequest req);
}