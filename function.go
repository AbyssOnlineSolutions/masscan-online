package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func StartScan(input string) {
	log.Printf("called StartScan()")

	//cmdstr := "masscan -p 80-1000 **********/24 --source-port 61000 --banners --rate=3000 "
	// シェルで実行するコマンドを作成
	println(input)
	cmd := exec.Command("sh", "-c", input)

	var status_new MASSCAN_STATUS
	var status_PID PID_TEMP
	var status_Status Status_TEMP

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	mu.Lock()

	status_new.PID = strconv.Itoa(cmd.Process.Pid)
	status_new.Status = "Running"
	data = append(data, status_new)
	var element = len(data) - 1

	status_PID.PID = strconv.Itoa(cmd.Process.Pid)
	status_Status.Status = "Running"
	status_PID.Subscript = element
	status_Status.Subscript = element
	status_PID.Type = "PID"
	status_Status.Type = "Status"
	SendBroadcast(status_PID)
	SendBroadcast(status_Status)

	mu.Unlock()
	data[element].Args = input
	go BufRead(stdout, element, strconv.Itoa(cmd.Process.Pid))
	go BufRead(stderr, element, strconv.Itoa(cmd.Process.Pid))

	cmd.Wait()

	mu.Lock()
	data[element].Status = "Finished"
	status_Status.Status = "Finished"
	SendBroadcast(status_Status)

	mu.Unlock()
}

func DataInsert(element int, stdout string, PID string) {
	var stdoutback = stdout
	stdoutarray := strings.Split(stdout, " ")

	if strings.Contains(stdoutarray[0], "Discovered") {
		var send_data DISCOVERD_TEMP
		var temp DISCOVERD
		send_data.PID = PID
		temp.IP = stdoutarray[5]
		temp.Port = stdoutarray[3]
		data[element].Discoverds = append(data[element].Discoverds, temp)
		send_data.Type = "Discovered"
		send_data.Discoverd = temp
		send_data.Subscript = element
		SendBroadcast(send_data)

	} else if strings.Contains(stdoutarray[0], "rate:") {
		var send_data PROCESS_TEMP
		var temp PROCESS
		temp.Rate = stdoutarray[2]
		if stdoutarray[3] == "" {
			temp.Percent = stdoutarray[4]
			temp.Found = stdoutarray[10]
			temp.Time = stdoutarray[8]
		} else {
			temp.Percent = stdoutarray[3]
			if temp.Percent == "100.00%" {
				temp.Found = stdoutarray[7]
				temp.Time = stdoutarray[5] + stdoutarray[6]
			} else {
				temp.Found = stdoutarray[9]
				temp.Time = stdoutarray[7]
			}
		}

		send_data.PID = PID
		send_data.Type = "Process"
		send_data.Process = temp
		data[element].Process = temp
		send_data.Subscript = element
		SendBroadcast(send_data)

	} else if strings.Contains(stdoutarray[0], "Banner") {
		var send_data BANNER_TEMP
		var temp BANNER

		temp.IP = stdoutarray[5]
		temp.Port = stdoutarray[3]
		temp.Proto = stdoutarray[6]

		var str = ""
		for i := 0; i < 7; i++ {
			str += stdoutarray[i] + " "
		}
		temp.Banner = strings.Replace(stdoutback, str, "", 1)

		data[element].Banners = append(data[element].Banners, temp)
		send_data.PID = PID
		send_data.Type = "Banner"
		send_data.Banner = temp
		send_data.Subscript = element
		SendBroadcast(send_data)
	}
}

func BufRead(r io.Reader, element int, PID string) {
	scanner := bufio.NewScanner(r)
	scanner.Split(CustomScan)
	for scanner.Scan() {
		mu.Lock()
		line := scanner.Text()
		DataInsert(element, line, PID)
		mu.Unlock()
	}
}

func SendBroadcast(v interface{}) {
	for client := range clients {
		err := client.WriteJSON(v)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func CustomScan(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	var i int
	if i = bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, DropCR(data[0:i]), nil
	}
	if i = bytes.IndexByte(data, '\r'); i >= 0 {
		// ここを追加した。（CR があったら、そこまでのデータを返そう）
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), DropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func DropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
