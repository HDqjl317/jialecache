/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-03-26 19:46:09
 * @LastEditTime: 2023-03-26 19:55:07
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */

package main

import (
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

func main() {
	jialecache.NewGroup("scores", 2<<10, jialecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := httpdb[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := jialecache.NewHTTPPool(addr)
	log.Println("jialecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
