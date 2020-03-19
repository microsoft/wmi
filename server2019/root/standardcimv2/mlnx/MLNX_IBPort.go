// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_IBPort struct
type MLNX_IBPort struct {
	*CIM_IBPort

	//
	MaxMsgSize uint64

	//
	MaxVls uint16

	//
	NumGids uint16

	//
	NumPkeys uint16

	//
	PciLocation string

	//
	QkeyCtr uint16

	//
	SmSl uint8

	//
	SubnetTimeout uint8

	//
	Transport string
}

func NewMLNX_IBPortEx1(instance *cim.WmiInstance) (newInstance *MLNX_IBPort, err error) {
	tmp, err := NewCIM_IBPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_IBPort{
		CIM_IBPort: tmp,
	}
	return
}

func NewMLNX_IBPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_IBPort, err error) {
	tmp, err := NewCIM_IBPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_IBPort{
		CIM_IBPort: tmp,
	}
	return
}

// SetMaxMsgSize sets the value of MaxMsgSize for the instance
func (instance *MLNX_IBPort) SetPropertyMaxMsgSize(value uint64) (err error) {
	return instance.SetProperty("MaxMsgSize", value)
}

// GetMaxMsgSize gets the value of MaxMsgSize for the instance
func (instance *MLNX_IBPort) GetPropertyMaxMsgSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxMsgSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxVls sets the value of MaxVls for the instance
func (instance *MLNX_IBPort) SetPropertyMaxVls(value uint16) (err error) {
	return instance.SetProperty("MaxVls", value)
}

// GetMaxVls gets the value of MaxVls for the instance
func (instance *MLNX_IBPort) GetPropertyMaxVls() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaxVls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumGids sets the value of NumGids for the instance
func (instance *MLNX_IBPort) SetPropertyNumGids(value uint16) (err error) {
	return instance.SetProperty("NumGids", value)
}

// GetNumGids gets the value of NumGids for the instance
func (instance *MLNX_IBPort) GetPropertyNumGids() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumGids")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumPkeys sets the value of NumPkeys for the instance
func (instance *MLNX_IBPort) SetPropertyNumPkeys(value uint16) (err error) {
	return instance.SetProperty("NumPkeys", value)
}

// GetNumPkeys gets the value of NumPkeys for the instance
func (instance *MLNX_IBPort) GetPropertyNumPkeys() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumPkeys")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPciLocation sets the value of PciLocation for the instance
func (instance *MLNX_IBPort) SetPropertyPciLocation(value string) (err error) {
	return instance.SetProperty("PciLocation", value)
}

// GetPciLocation gets the value of PciLocation for the instance
func (instance *MLNX_IBPort) GetPropertyPciLocation() (value string, err error) {
	retValue, err := instance.GetProperty("PciLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQkeyCtr sets the value of QkeyCtr for the instance
func (instance *MLNX_IBPort) SetPropertyQkeyCtr(value uint16) (err error) {
	return instance.SetProperty("QkeyCtr", value)
}

// GetQkeyCtr gets the value of QkeyCtr for the instance
func (instance *MLNX_IBPort) GetPropertyQkeyCtr() (value uint16, err error) {
	retValue, err := instance.GetProperty("QkeyCtr")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmSl sets the value of SmSl for the instance
func (instance *MLNX_IBPort) SetPropertySmSl(value uint8) (err error) {
	return instance.SetProperty("SmSl", value)
}

// GetSmSl gets the value of SmSl for the instance
func (instance *MLNX_IBPort) GetPropertySmSl() (value uint8, err error) {
	retValue, err := instance.GetProperty("SmSl")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubnetTimeout sets the value of SubnetTimeout for the instance
func (instance *MLNX_IBPort) SetPropertySubnetTimeout(value uint8) (err error) {
	return instance.SetProperty("SubnetTimeout", value)
}

// GetSubnetTimeout gets the value of SubnetTimeout for the instance
func (instance *MLNX_IBPort) GetPropertySubnetTimeout() (value uint8, err error) {
	retValue, err := instance.GetProperty("SubnetTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTransport sets the value of Transport for the instance
func (instance *MLNX_IBPort) SetPropertyTransport(value string) (err error) {
	return instance.SetProperty("Transport", value)
}

// GetTransport gets the value of Transport for the instance
func (instance *MLNX_IBPort) GetPropertyTransport() (value string, err error) {
	retValue, err := instance.GetProperty("Transport")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
