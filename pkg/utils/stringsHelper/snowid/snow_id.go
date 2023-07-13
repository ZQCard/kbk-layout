package snowid

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func GenSnowId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Generate a snowflake ID.
	// type ID int64 这是id的类型，下面定义了很多方法， 获取Base64呀，或者是打印它的时间戳等
	id := node.Generate()
	return id.Int64()
}
