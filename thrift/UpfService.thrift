namespace go Upf

struct PortStats {
    1:required list<i32> rxPktCnt;
	2:required list<i32> rxPktDropCnt;  
	3:required i32 txNQfiPktCnt;
	4:required i32 txNQfiPktDropCnt;
	5:required list<i32> txPktCnt;          
	6:required list<i32> txPktDropCnt;  
	7:required list<list<i32>> dropperPkts;
	8:required list<list<i32>> dropperDrops;
	9:required list<i32> schedulerQDropCnt;
    10:required i32 nffGoSchedulerDropCnt;
    11:required i32 classifierDropCnt;
    12:required i32 workerDropCnt;
}

struct StatsResponse {
    1:required i32 errCode; 
    2:required string errMsg;
    3:required PortStats stats;
}

service UpfService {
    StatsResponse GetStats(
        1:required i32 port
    )
    void ClearStats(
        1:required i32 port
    )
    void PcapStart(
       1:required i32 port
    )
    void PcapStop(
       1:required i32 port
    )
}