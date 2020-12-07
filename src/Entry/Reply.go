package Entry

import "encoding/json"

type Reply struct {
	Code string `json:"Code"`
	Message string `json:"Message"`
	Data interface{} `json:"Data"`
	//Data *Balance `json:"Data"`
	//Data map[string]interface{} `json:"Data"`
}

func NewReply(structInfo interface{}) *Reply {
	return &Reply{
		Data: &structInfo,
	}
}

func ResponseResult(content string, structInfo interface{}) (reply *Reply,err error) {
	reply = NewReply(structInfo)
	//reply.Data = balance
	err = json.Unmarshal([]byte(content), reply)

	return reply, err
}