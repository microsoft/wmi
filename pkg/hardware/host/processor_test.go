// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package host

import (
	"github.com/stretchr/testify/assert"
	"testing"

	wmihost "github.com/microsoft/wmi/pkg/base/host"
	_ "github.com/microsoft/wmi/pkg/base/session"
)

func TestGetProcessor(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	proc, err := GetTotalProcessor(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("Total Processor [%+v]\n", proc)
}

func TestGetProcessorInfo(t *testing.T) {
	whost := wmihost.NewWmiLocalHost()
	process, err := GetProcessor(whost)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	t.Logf("ProcessorInfo [%+v]\n", process)
}

func TestGetSpecificProcessorInfo(t *testing.T) {
	assert := assert.New(t)
	whost := wmihost.NewWmiLocalHost()
	parameters := []string{"Manufacturer", "CurrentClockSpeed", "Architecture"}
	process, err := GetSpecificProcessorInfo(whost, parameters)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	numberOfCores, err := process.GetProperty("NumberOfCores")
	assert.Errorf(err, "Failed to get property NumberOfCores: %v", err)
	assert.Equal(numberOfCores, nil, "NumberOfCores should be empty")

	manuf, err := process.GetProperty("Manufacturer")
	assert.NoErrorf(err, "Failed to get property Manufacturer: %v", err)
	assert.NotEqual(manuf, nil, "Manufacturer should not be empty")

	curr, err := process.GetProperty("CurrentClockSpeed")
	assert.NoErrorf(err, "Failed to get property CurrentClockSpeed: %v", err)
	assert.NotEqual(curr, nil, "CurrentClockSpeed should not be empty")

	archi, err := process.GetProperty("Architecture")
	assert.NoErrorf(err, "Failed to get property Architecture: %v", err)
	assert.NotEqual(archi, nil, "Architecture should not be empty")

	t.Logf("ProcessorInfo With Specific Query Parameters [%+v]\n", process)
}

func TestGetSpecificProcessorInfoWithEmptyParameters(t *testing.T) {
	assert := assert.New(t)
	whost := wmihost.NewWmiLocalHost()
	parameters := []string{}
	process, err := GetSpecificProcessorInfo(whost, parameters)
	if err != nil {
		t.Fatalf("[%+v]", err)
	}

	numberOfCores, err := process.GetProperty("NumberOfCores")
	assert.NoErrorf(err, "Failed to get property NumberOfCores: %v", err)
	assert.NotEqual(numberOfCores, nil, "NumberOfCores should not be empty")

	manuf, err := process.GetProperty("Manufacturer")
	assert.NoErrorf(err, "Failed to get property Manufacturer: %v", err)
	assert.NotEqual(manuf, nil, "Manufacturer should not be empty")

	curr, err := process.GetProperty("CurrentClockSpeed")
	assert.NoErrorf(err, "Failed to get property CurrentClockSpeed: %v", err)
	assert.NotEqual(curr, nil, "CurrentClockSpeed should not be empty")

	archi, err := process.GetProperty("Architecture")
	assert.NoErrorf(err, "Failed to get property Architecture: %v", err)
	assert.NotEqual(archi, nil, "Architecture should not be empty")

	t.Logf("ProcessorInfo With Specific Query Parameters [%+v]\n", process)
}
