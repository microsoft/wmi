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

// CIM_Card struct
type CIM_Card struct {
	*CIM_PhysicalPackage

	//
	HostingBoard bool

	//
	OperatingVoltages []int16

	//
	RequirementsDescription string

	//
	RequiresDaughterBoard bool

	//
	SlotLayout string

	//
	SpecialRequirements bool
}

func NewCIM_CardEx1(instance *cim.WmiInstance) (newInstance *CIM_Card, err error) {
	tmp, err := NewCIM_PhysicalPackageEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Card{
		CIM_PhysicalPackage: tmp,
	}
	return
}

func NewCIM_CardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Card, err error) {
	tmp, err := NewCIM_PhysicalPackageEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Card{
		CIM_PhysicalPackage: tmp,
	}
	return
}

// SetHostingBoard sets the value of HostingBoard for the instance
func (instance *CIM_Card) SetPropertyHostingBoard(value bool) (err error) {
	return instance.SetProperty("HostingBoard", value)
}

// GetHostingBoard gets the value of HostingBoard for the instance
func (instance *CIM_Card) GetPropertyHostingBoard() (value bool, err error) {
	retValue, err := instance.GetProperty("HostingBoard")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperatingVoltages sets the value of OperatingVoltages for the instance
func (instance *CIM_Card) SetPropertyOperatingVoltages(value []int16) (err error) {
	return instance.SetProperty("OperatingVoltages", value)
}

// GetOperatingVoltages gets the value of OperatingVoltages for the instance
func (instance *CIM_Card) GetPropertyOperatingVoltages() (value []int16, err error) {
	retValue, err := instance.GetProperty("OperatingVoltages")
	if err != nil {
		return
	}
	value, ok := retValue.([]int16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequirementsDescription sets the value of RequirementsDescription for the instance
func (instance *CIM_Card) SetPropertyRequirementsDescription(value string) (err error) {
	return instance.SetProperty("RequirementsDescription", value)
}

// GetRequirementsDescription gets the value of RequirementsDescription for the instance
func (instance *CIM_Card) GetPropertyRequirementsDescription() (value string, err error) {
	retValue, err := instance.GetProperty("RequirementsDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiresDaughterBoard sets the value of RequiresDaughterBoard for the instance
func (instance *CIM_Card) SetPropertyRequiresDaughterBoard(value bool) (err error) {
	return instance.SetProperty("RequiresDaughterBoard", value)
}

// GetRequiresDaughterBoard gets the value of RequiresDaughterBoard for the instance
func (instance *CIM_Card) GetPropertyRequiresDaughterBoard() (value bool, err error) {
	retValue, err := instance.GetProperty("RequiresDaughterBoard")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSlotLayout sets the value of SlotLayout for the instance
func (instance *CIM_Card) SetPropertySlotLayout(value string) (err error) {
	return instance.SetProperty("SlotLayout", value)
}

// GetSlotLayout gets the value of SlotLayout for the instance
func (instance *CIM_Card) GetPropertySlotLayout() (value string, err error) {
	retValue, err := instance.GetProperty("SlotLayout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecialRequirements sets the value of SpecialRequirements for the instance
func (instance *CIM_Card) SetPropertySpecialRequirements(value bool) (err error) {
	return instance.SetProperty("SpecialRequirements", value)
}

// GetSpecialRequirements gets the value of SpecialRequirements for the instance
func (instance *CIM_Card) GetPropertySpecialRequirements() (value bool, err error) {
	retValue, err := instance.GetProperty("SpecialRequirements")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Connector" type="CIM_PhysicalConnector "></param>
// <param name="PoweredOn" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Card) ConnectorPower( /* IN */ Connector CIM_PhysicalConnector,
	/* IN */ PoweredOn bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ConnectorPower", Connector, PoweredOn)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
