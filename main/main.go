/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-03-26 19:46:09
 * @LastEditTime: 2023-03-27 12:35:16
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package main

import (
	"flag"
	"fmt"
	"jialecache"
	"log"
	"net/http"
)

var httpdb = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *jialecache.Group {
	return jialecache.NewGroup("scores", 2<<10, jialecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := httpdb[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, jiale *jialecache.Group) {
	peers := jialecache.NewHTTPPool(addr)
	peers.Set(addrs...)
	jiale.RegisterPeers(peers)
	log.Println("jialecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, jiale *jialecache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := jiale.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		},
	))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "jialecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	jiale := createGroup()
	if api {
		go startAPIServer(apiAddr, jiale)
	}
	startCacheServer(addrMap[port], addrs, jiale)
}
