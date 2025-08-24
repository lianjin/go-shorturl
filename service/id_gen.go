package service

import (
	"gsurl/log"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

var (
	counter *snowflake.Node
)

func InitIdGenerator() {
	nodeIdStr := os.Getenv("SNOWFLAKE_NODE_ID")
	nodeId, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		log.Logger.Errorf("Invalid SNOWFLAKE_NODE_ID: %s", nodeIdStr)
		panic(err)
	}
	node, err := snowflake.NewNode(int64(nodeId))
	if err != nil {
		log.Logger.Errorf("Failed to create snowflake node: %v", err)
		panic(err)
	}
	counter = node
}

func GenId() int64 {
	return counter.Generate().Int64()
}
