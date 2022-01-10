include "Upf.thrift"

namespace go Upf

struct PortStats {
    1:required i32 rxPktCount;
    2:required i32 txPktCount;
}
typedef PortStats Stats

struct StatsResponse {
    1:required i32 errCode; 
    2:required string errMsg;
    3:required Stats data;
}

service UpfService {
    StatsResponse GetStats(
        1:required i32 port
    )
    void ClearStats(
        1:required i32 port
    )
}