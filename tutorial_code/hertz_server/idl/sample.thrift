namespace go api

struct QueryUserRequest {
    1: i32 Num (api.query="num", api.vd="$<100; msg:'num must less than 100'"); 
}

struct QueryUserResponse {
    1: string Num;
    2: string Name;
    3: string Gender;
    4: string Msg;
}

struct AddUserRequest {
    1: string Num (api.form="num");
    2: string Name (api.form="name");
    3: string Gender (api.form="gender");
}

struct AddUserResponse {
    1: bool Ok; 
    2: string Msg; 
}

service UserApi {
   
   QueryUserResponse queryUser(1: QueryUserRequest req) (api.get="User/query");
   AddUserResponse AddUser(1: AddUserRequest req) (api.post="User/Add");
}
