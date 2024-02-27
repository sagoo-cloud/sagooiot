package sse

import (
	"github.com/gogf/gf/v2/net/ghttp"
	sseserver "github.com/xinjiayu/sse"
	"sagooiot/internal/sse/sysenv"
	"time"
)

var redisInfoEvent = sseserver.NewServer()

func runRedisInfo() {
	redisInfoData := sysenv.GetRedisInfo()

	cpuMsg := sseserver.SSEMessage{}
	cpuMsg.Event = "cpu"
	cpuMsg.Data = redisInfoData["CPU"]
	cpuMsg.Namespace = "/redisinfo/cpu"
	redisInfoEvent.Broadcast <- cpuMsg

	redisMsg := sseserver.SSEMessage{}
	redisMsg.Event = "server"
	redisMsg.Data = redisInfoData["Server"]
	redisMsg.Namespace = "/redisinfo/server"
	redisInfoEvent.Broadcast <- redisMsg

	clientsMsg := sseserver.SSEMessage{}
	clientsMsg.Event = "clients"
	clientsMsg.Data = redisInfoData["Clients"]
	clientsMsg.Namespace = "/redisinfo/clients"
	redisInfoEvent.Broadcast <- clientsMsg

	errorstatsMsg := sseserver.SSEMessage{}
	errorstatsMsg.Event = "errorstats"
	errorstatsMsg.Data = redisInfoData["Errorstats"]
	errorstatsMsg.Namespace = "/redisinfo/errorstats"
	redisInfoEvent.Broadcast <- errorstatsMsg

	keyspaceMsg := sseserver.SSEMessage{}
	keyspaceMsg.Event = "keyspace"
	keyspaceMsg.Data = redisInfoData["Keyspace"]
	keyspaceMsg.Namespace = "/redisinfo/keyspace"
	redisInfoEvent.Broadcast <- keyspaceMsg

	clusterMsg := sseserver.SSEMessage{}
	clusterMsg.Event = "cluster"
	clusterMsg.Data = redisInfoData["Cluster"]
	clusterMsg.Namespace = "/redisinfo/cluster"
	redisInfoEvent.Broadcast <- clusterMsg

	memoryMsg := sseserver.SSEMessage{}
	memoryMsg.Event = "memory"
	memoryMsg.Data = redisInfoData["Memory"]
	memoryMsg.Namespace = "/redisinfo/memory"
	redisInfoEvent.Broadcast <- memoryMsg

	statsMsg := sseserver.SSEMessage{}
	statsMsg.Event = "stats"
	statsMsg.Data = redisInfoData["Stats"]
	statsMsg.Namespace = "/redisinfo/stats"
	redisInfoEvent.Broadcast <- statsMsg

	replicationMsg := sseserver.SSEMessage{}
	replicationMsg.Event = "replication"
	replicationMsg.Data = redisInfoData["Replication"]
	replicationMsg.Namespace = "/redisinfo/replication"
	redisInfoEvent.Broadcast <- replicationMsg

	time.Sleep(time.Duration(1) * time.Second)
}

func RedisInfoMessageEvent(r *ghttp.Request) {
	redisInfoEvent.ServeHTTP(r.Response.RawWriter(), r.Request)
}
