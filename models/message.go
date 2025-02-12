package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// Message 消息
type Message struct {
	gorm.Model
	Amount      int    `json:"amount"`    // 文件大小
	Type        int    `json:"type"`      // 发送类型 群聊 私聊 广播
	Media       int    `json:"media"`     // 消息类型 文字 图片 音频
	FormID      int64  `json:"form_id"`   // 发送者
	TargetID    int64  `json:"target_id"` // 接受者
	Pic         string `json:"pic"`       // 图片
	Url         string `json:"url"`
	Content     string `json:"content"` // 消息类型
	Description string `json:"description"`
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

func (m *Message) TableName() string {
	return "message"
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node)

// 读写锁
var rwLocker sync.RWMutex

// Chat 需要：发送者ID，接受者ID，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 校验token等信息
	// token := query.Get("token")
	query := request.URL.Query()
	ID := query.Get("userId")
	userID, _ := strconv.ParseInt(ID, 10, 64)
	/*	SendType := query.Get("sendType")
		targetId := query.Get("targetId")
		content := query.Get("content")*/
	isValida := true // checkToken()
	conn, err := (&websocket.Upgrader{
		// token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 3.用户关系
	// 4.userID和node绑定并加锁
	rwLocker.Lock()
	clientMap[userID] = node
	rwLocker.Unlock()
	// 5.完成发送逻辑
	go sendProc(node)
	// 6.完成接受逻辑
	go recvProc(node)
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println(string(data))
	}
}

var udpSendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpSendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 25),
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpSendChan:
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}
}

// 完成udp数据接受协程
func udpRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: // 私信
		sendMsg(msg.TargetID, data)
	case 2: // 群发
	case 3: // 广播
	case 4:
	}
}

func sendMsg(userID int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userID]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
