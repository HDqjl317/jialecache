/*
 * @Author: jiale_quan jiale_quan@ustc.edu
 * @Date: 2023-03-27 13:52:24
 * @LastEditTime: 2023-03-27 13:53:06
 * @Description:
 * Copyright © jiale_quan, All Rights Reserved
 */
package singleflight

import "testing"

func TestDo(t *testing.T) {
	var g Group
	v, err := g.Do("key", func() (interface{}, error) {
		return "bar", nil
	})

	if v != "bar" || err != nil {
		t.Errorf("Do v = %v, error = %v", v, err)
	}
}
