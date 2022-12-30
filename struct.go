package main

type PROCESS struct {
	Rate    string `json:"Rate"`
	Percent string `json:"Percent"`
	Time    string `json:"Time"`
	Found   string `json:"Found"`
}

type DISCOVERD struct {
	IP   string `json:"IP"`
	Port string `json:"Port"`
}
type BANNER struct {
	IP     string `json:"IP"`
	Port   string `json:"Port"`
	Proto  string `json:"Proto"`
	Banner string `json:"Banner"`
}

type MASSCAN_STATUS struct {
	PID        string      `json:"PID"`
	Args       string      `json:"Args"`
	Process    PROCESS     `json:"Process"`
	Discoverds []DISCOVERD `json:"Discoverds"`
	Banners    []BANNER    `json:"Banners"`
	Status     string      `json:"Status"`
}

type MASSCAN []MASSCAN_STATUS

type PID_TEMP struct {
	Subscript int    `json:"Subscript"`
	Type      string `json:"Type"`
	PID       string `json:"PID"`
}

type Args_TEMP struct {
	PID       string `json:"PID"`
	Subscript int    `json:"Subscript"`
	Type      string `json:"Type"`
	Args      BANNER `json:"Args"`
}
type PROCESS_TEMP struct {
	PID       string  `json:"PID"`
	Subscript int     `json:"Subscript"`
	Type      string  `json:"Type"`
	Process   PROCESS `json:"Process"`
}

type DISCOVERD_TEMP struct {
	PID       string    `json:"PID"`
	Subscript int       `json:"Subscript"`
	Type      string    `json:"Type"`
	Discoverd DISCOVERD `json:"Discoverd"`
}
type BANNER_TEMP struct {
	PID       string `json:"PID"`
	Subscript int    `json:"Subscript"`
	Type      string `json:"Type"`
	Banner    BANNER `json:"Banner"`
}
type Status_TEMP struct {
	PID       string `json:"PID"`
	Subscript int    `json:"Subscript"`
	Type      string `json:"Type"`
	Status    string `json:"Status"`
}

// メッセージ用構造体
type Message struct {
	Cmd string `json:"cmd"`
}
