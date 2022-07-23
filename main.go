package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node struct {
	NodeId        string
	NextMessageId int
}

type EchoInitResponse struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Body struct {
		MsgId     int    `json:"msg_id"`
		InReplyTo int    `json:"in_reply_to"`
		Msgtype   string `json:"type"`
	} `json:"body"`
}

type EchoResponse struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Body struct {
		MsgId     int    `json:"msg_id"`
		InReplyTo int    `json:"in_reply_to"`
		Msgtype   string `json:"type"`
		EchoMsg   string `json:"echo"`
	} `json:"body"`
}

func main() {
	var node Node
	reader := bufio.NewReader(os.Stdin)
	for {
		rawLine, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line := strings.TrimSpace(rawLine)
		msgtype := messageType(line)
		switch msgtype {
		case "init":
			handleInit(line, &node)
		case "echo":
			handleEcho(line, &node)
		default:
			fmt.Fprintf(os.Stderr, "unknown message type: %v\n", msgtype)
		}
	}
}

func messageType(msg string) string {
	parsed, err := UnmarshalJson(msg)
	if err != nil {
		return err.Error()
	}
	msgtype, err := ParseJson[string](parsed, "body", "type")
	if err == nil {
		return msgtype
	}
	return err.Error()
}

// handle an 'init' message and print an appropriate response
func handleInit(msg string, node *Node) {
	parsed, err := UnmarshalJson(msg)
	if err != nil {
		return
	}
	nodeId, e1 := ParseJson[string](parsed, "body", "node_id")
	// int values are stored as float64's in JSON
	msgId, e2 := ParseJson[float64](parsed, "body", "msg_id")
	src, e3 := ParseJson[string](parsed, "src")
	if e1 != nil || e2 != nil || e3 != nil {
		fmt.Fprintf(os.Stderr, "error picking values from JSON\n")
		return
	}

	node.NodeId = nodeId

	var resp EchoInitResponse
	resp.Body.Msgtype = "init_ok"
	resp.Src = nodeId
	resp.Dest = src
	resp.Body.InReplyTo = int(msgId)
	resp.Body.MsgId = node.NextMessageId
	node.NextMessageId++

	output, err := json.Marshal(&resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Println(string(output))
}

// handle an 'echo' message and print an appropriate response
func handleEcho(msg string, node *Node) {
	parsed, err := UnmarshalJson(msg)
	if err != nil {
		return
	}
	message, e1 := ParseJson[string](parsed, "body", "echo")
	msgId, e2 := ParseJson[float64](parsed, "body", "msg_id")
	src, e3 := ParseJson[string](parsed, "src")
	if e1 != nil || e2 != nil || e3 != nil {
		fmt.Fprintf(os.Stderr, "error picking values from JSON\n")
		return
	}

	var resp EchoResponse
	resp.Body.Msgtype = "echo_ok"
	resp.Src = node.NodeId
	resp.Dest = src
	resp.Body.EchoMsg = message
	resp.Body.InReplyTo = int(msgId)
	resp.Body.MsgId = node.NextMessageId
	node.NextMessageId++

	output, err := json.Marshal(&resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Println(string(output))
}
