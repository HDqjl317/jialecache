/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-03-26 16:02:46
 * @LastEditTime: 2023-03-26 16:07:53
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package jialecache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func (v ByteView) String() string {
	return string(v.b)
}
