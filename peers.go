/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-03-27 10:45:52
 * @LastEditTime: 2023-03-27 10:46:03
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package jialecache

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
