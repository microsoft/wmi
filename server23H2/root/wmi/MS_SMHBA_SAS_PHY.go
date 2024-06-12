// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MS_SMHBA_SAS_PHY struct
type MS_SMHBA_SAS_PHY struct {
	*cim.WmiInstance

	//
	domainPortWWN []uint8

	//
	HardwareMaxLinkRate uint32

	//
	HardwareMinLinkRate uint32

	//
	NegotiatedLinkRate uint32

	//
	PhyIdentifier uint8

	//
	ProgrammedMaxLinkRate uint32

	//
	ProgrammedMinLinkRate uint32
}

func NewMS_SMHBA_SAS_PHYEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_SAS_PHY, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_SAS_PHY{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_SAS_PHYEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_SAS_PHY, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_SAS_PHY{
		WmiInstance: tmp,
	}
	return
}

// SetdomainPortWWN sets the value of domainPortWWN for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertydomainPortWWN(value []uint8) (err error) {
	return instance.SetProperty("domainPortWWN", (value))
}

// GetdomainPortWWN gets the value of domainPortWWN for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertydomainPortWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("domainPortWWN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetHardwareMaxLinkRate sets the value of HardwareMaxLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertyHardwareMaxLinkRate(value uint32) (err error) {
	return instance.SetProperty("HardwareMaxLinkRate", (value))
}

// GetHardwareMaxLinkRate gets the value of HardwareMaxLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertyHardwareMaxLinkRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardwareMaxLinkRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetHardwareMinLinkRate sets the value of HardwareMinLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertyHardwareMinLinkRate(value uint32) (err error) {
	return instance.SetProperty("HardwareMinLinkRate", (value))
}

// GetHardwareMinLinkRate gets the value of HardwareMinLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertyHardwareMinLinkRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardwareMinLinkRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNegotiatedLinkRate sets the value of NegotiatedLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertyNegotiatedLinkRate(value uint32) (err error) {
	return instance.SetProperty("NegotiatedLinkRate", (value))
}

// GetNegotiatedLinkRate gets the value of NegotiatedLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertyNegotiatedLinkRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("NegotiatedLinkRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPhyIdentifier sets the value of PhyIdentifier for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertyPhyIdentifier(value uint8) (err error) {
	return instance.SetProperty("PhyIdentifier", (value))
}

// GetPhyIdentifier gets the value of PhyIdentifier for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertyPhyIdentifier() (value uint8, err error) {
	retValue, err := instance.GetProperty("PhyIdentifier")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetProgrammedMaxLinkRate sets the value of ProgrammedMaxLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertyProgrammedMaxLinkRate(value uint32) (err error) {
	return instance.SetProperty("ProgrammedMaxLinkRate", (value))
}

// GetProgrammedMaxLinkRate gets the value of ProgrammedMaxLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertyProgrammedMaxLinkRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProgrammedMaxLinkRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetProgrammedMinLinkRate sets the value of ProgrammedMinLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) SetPropertyProgrammedMinLinkRate(value uint32) (err error) {
	return instance.SetProperty("ProgrammedMinLinkRate", (value))
}

// GetProgrammedMinLinkRate gets the value of ProgrammedMinLinkRate for the instance
func (instance *MS_SMHBA_SAS_PHY) GetPropertyProgrammedMinLinkRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProgrammedMinLinkRate")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
