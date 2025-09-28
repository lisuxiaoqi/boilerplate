package net

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"testing"
)

func Test_NormalWS(t *testing.T) {
	// 直接写完整的 wss URL
	wsURL := "wss://gl-exp-api-m.gatescan.org/socket/v2/websocket?vsn=2.0.0"
	log.Printf("连接到 %s", wsURL)

	// 建立连接
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 捕获 Ctrl+C 退出
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 发消息
	msg := []interface{}{"6", "6", "blocks:new_block", "phx_join", map[string]interface{}{}}
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("json marshal:", err)
	}

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Fatal("write message:", err)
	}
	log.Printf("已发送: %s\n", data)

	// 启动一个 goroutine 不断读取消息
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("收到消息: %s", message)
		}
	}()

	// 主协程等待退出
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("中断，关闭连接")
			// 优雅关闭
			err := c.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}
}
