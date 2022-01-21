namespace go UpfGtp5g

struct PdrInfo {
    1:required i32 id;
	2:required i32 precdence;  
	3:required i32 Ohr;
	4:required i32 roleAddr4;
	5:required i32 pdiUeAddr4;          
	6:required i32 pdiFTeid;  
	7:required i32 pdiGtpuAddr4;
	8:required i32 farId;
    9:required list<i32>  qerId;
    10:required i32 ulDropCount;
	11:required i32 dlDropCount;
}

struct GetPdrResponse {
    1:required i32 errCode; 
    2:required string errMsg;
    3:required PdrInfo pdr;
}

struct FarInfo {
    1:required i32 id;
	2:required i32 action;  
	3:required i32 description;
	4:required i32 teid;
	5:required i32 peerAddr4;          
}

struct GetFarResponse {
    1:required i32 errCode; 
    2:required string errMsg;
    3:required FarInfo far;
}

struct QerInfo {
    1:required i32 id;
	2:required i32 qfi;           
}

struct GetQerResponse {
    1:required i32 errCode; 
    2:required string errMsg;
    3:required FarInfo far;
}

service UpfGtp5gService {
    GetPdrResponse GetPdr(
        1:required string iface
        2:required i32 id
    )
    GetFarResponse GetFar(
        1:required string iface
        2:required i32 id
    )
    GetQerResponse GetQer(
        1:required string iface
        2:required i32 id
    )
}