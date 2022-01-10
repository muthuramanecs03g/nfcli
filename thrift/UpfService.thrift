include "Upf.thrift"

namespace go Upf

typedef map<string, string> Data

struct Response {
    1:required i32 errCode; 
    2:required string errMsg;
    3:required Data data;
}

service Greeter {
    Response SayHello(
        1:required Upf.User user
    )

    Response GetUser(
        1:required i32 uid
    )
}