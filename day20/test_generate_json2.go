package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server3 struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

func main() {
	s := Server3{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	fmt.Println(s) // {3 Go "1.0"  Go "1.0"  }

	b, _ := json.Marshal(s)
	os.Stdout.Write(b) // {"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}

}
