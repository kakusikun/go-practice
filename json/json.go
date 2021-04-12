package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Data1 string          `json:"data1"`
	Data2 json.RawMessage `json:"data2"`
}

type message2 struct {
	Data3 string `json:"data3"`
}

type message3 struct {
	To   string          `json:"to"`
	From string          `json:"from"`
	Api  string          `json:"api"`
	Data json.RawMessage `json:"data"`
}

func buildMessage3(data interface{}) message3 {
	switch v := data.(type) {
	case string:
		msg := message3{
			To:   "A",
			From: "B",
			Api:  "C",
			Data: []byte(v),
		}
		fmt.Println("......string", string(msg.Data))
		return msg
	case []byte:
		msg := message3{
			To:   "A",
			From: "B",
			Api:  "C",
			Data: v,
		}
		fmt.Println("......byte", msg)
		return msg
	}
	return message3{}
}

func main() {
	data := []byte(`{"data1": "shit", "data2": {"data3": "wtf"}}`)
	fmt.Println("raw", data)
	msg := message{}
	json.Unmarshal(data, &msg)
	fmt.Println("Unmarshal 1 level", msg)
	strMsg := string(msg.Data2)
	fmt.Println("string", strMsg)
	msg2 := message2{}
	json.Unmarshal(msg.Data2, &msg2)
	fmt.Println("Unmarshal 1 level", msg2)

	if rawMsg, err := json.Marshal(msg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Marshal 1 level", rawMsg)
		strMsg := string(rawMsg)
		fmt.Println("string", strMsg)
	}

	msg = message{}
	json.Unmarshal(data, &msg)
	fmt.Println("Unmarshal the marshal", msg)

	data2 := []byte(`{"data1": "shit", "data2": "wtf"}`)
	json.Unmarshal(data2, &msg)
	fmt.Println("Unmarshal", msg.Data2)
	fmt.Println("Unmarshal data2", string(msg.Data2))

	msg3 := message{
		Data1: "wtf",
		Data2: []byte(`"haha"`),
	}

	if rawMsg, err := json.Marshal(msg3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Marshal from struct", rawMsg)
		strMsg := string(rawMsg)
		fmt.Println("string", strMsg)
	}

	msg4 := buildMessage3("\"client does not exist: thoth\"")
	if rawMsg, err := json.Marshal(msg4); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Marshal from struct", rawMsg)
		strMsg := string(rawMsg)
		fmt.Println("string", strMsg)
	}

}
