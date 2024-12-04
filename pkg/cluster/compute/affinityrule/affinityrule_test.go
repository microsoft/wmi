// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package affinityrule

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

func TestGetAffinityRule(t *testing.T) {
	cn, err := GetAffinityRule(whost, "Cloud Agent-CNO")
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer cn.Close()
}

func TestGetAffinityRules(t *testing.T) {
	nc, err := GetAffinityRules(whost)
	if err != nil {
		t.Fatal("Failed " + err.Error())
		return
	}
	defer nc.Close()
	t.Logf("Nodes returned %d\n", len(nc))

	for _, affinityRule := range nc {
		affinityRuleName, err := affinityRule.GetPropertyName()
		if err != nil {
			t.Fatal("Failed " + err.Error())
			return
		}
		t.Logf("NoodeName : %s\n", affinityRuleName)
	}
}
