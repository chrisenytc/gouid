/*
 * gouid
 * https://github.com/chrisenytc/gouid
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package main

import (
	"github.com/chrisenytc/gouid"
	"testing"
)

func TestNewUid(t *testing.T) {
	n := gouid.UId{8}
	n.SetSeed()
	testRand := n.NewUId()
	testRand2 := n.NewUId()
	if testRand == testRand2 {
		t.Errorf("'%s' should not be equal to '%s'", testRand, testRand2)
	}
}

func TestNewUidLength(t *testing.T) {
	n := gouid.UId{8}
	testRand := n.NewUId()
	if len(testRand) != 8 {
		t.Errorf("'%s' length is not equal to 8", testRand)
	}
}
