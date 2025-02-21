package cim

import (
	"testing"

	"github.com/microsoft/wmi/go/wmi"
	"github.com/stretchr/testify/assert"
)

func Test_Type(t *testing.T) {
	sm := NewWmiSessionManager()
	defer sm.Close()
	defer sm.Dispose()
	session, err := sm.GetLocalSession("root\\cimv2")
	if err != nil {
		t.Errorf("Could not get session %v", err)
		return
	}
	_, err = session.Connect()
	if err != nil {
		t.Errorf("Could not connect session %v", err)
		return
	}
	defer session.Close()
	defer session.Dispose()

	query := "SELECT * FROM Win32_OperatingSystem"
	rows, err := session.QueryInstances(query)

	if err != nil {
		t.Logf("Error could not fetch rows")
	}

	if len(rows) > 0 {
		r := rows[0]
		sysProp, err := r.GetSystemProperty("__Path")
		if err != nil {
			t.Logf("Cannot fetch __Path")
		}
		// This is causing panic on 64-bit before fixing #150
		myType := sysProp.Type()

		// This line is not executed before fixing #150
		assert.Equal(t, myType, wmi.WbemCimtypeString, "__Path should be of type string")

	}
}
