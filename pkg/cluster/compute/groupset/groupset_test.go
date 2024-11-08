// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package groupset

import (
	"github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
	"testing"
)

var (
	whost *host.WmiHost
)

func init() {
	whost = host.NewWmiLocalHost()
}

/*
func TestGetGroupSet(t *testing.T) {
	cn, err := GetGroupSet(whost, "*")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}
*/
func TestGetGroupSets(t *testing.T) {
	nc, err := GetGroupSets(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("Nodes returned %d\n", len(nc))

	for _, gpSet := range nc {
		grpSetName, err := gpSet.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("NoodeName : %s\n", grpSetName)
	}
}
