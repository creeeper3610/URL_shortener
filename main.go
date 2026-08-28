package main

import (
	"crypto/rand"
	"encoding/hex"
	"syscall/js"
)

// メモリ内のURLデータベース
var store = make(map[string]string)

// 6文字のランダムIDを生成
func generateID() string {
	b := make([]byte, 3)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// 短縮URL発行ハンドラー
func handleShorten(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return ""
	}
	longURL := args[0].String()
	id := generateID()
	store[id] = longURL
	return id
}

// URL復元ハンドラー
func handleResolve(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return ""
	}
	id := args[0].String()
	if target, exists := store[id]; exists {
		return target
	}
	return ""
}

func main() {
	// JavaScript側にサーバーの関数を公開
	js.Global().Set("__go_server_shorten", js.FuncOf(handleShorten))
	js.Global().Set("__go_server_resolve", js.FuncOf(handleResolve))

	// Goサーバーを常駐稼働
	select {}
}
