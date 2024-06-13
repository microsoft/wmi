// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ClusterSetFaultDomain struct
type MSFT_ClusterSetFaultDomain struct {
	*cim.WmiInstance

	//
	ClusterName []string

	//
	Description string

	//
	FaultDomainType uint32

	//
	FDName string

	//
	Id uint64
}

func NewMSFT_ClusterSetFaultDomainEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetFaultDomain, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetFaultDomain{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ClusterSetFaultDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetFaultDomain, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetFaultDomain{
		WmiInstance: tmp,
	}
	return
}

// SetClusterName sets the value of ClusterName for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyClusterName(value []string) (err error) {
	return instance.SetProperty("ClusterName", (value))
}

// GetClusterName gets the value of ClusterName for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyClusterName() (value []string, err error) {
	retValue, err := instance.GetProperty("ClusterName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetFaultDomainType sets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyFaultDomainType(value uint32) (err error) {
	return instance.SetProperty("FaultDomainType", (value))
}

// GetFaultDomainType gets the value of FaultDomainType for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyFaultDomainType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultDomainType")
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

// SetFDName sets the value of FDName for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyFDName(value string) (err error) {
	return instance.SetProperty("FDName", (value))
}

// GetFDName gets the value of FDName for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyFDName() (value string, err error) {
	retValue, err := instance.GetProperty("FDName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetFaultDomain) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetFaultDomain) GetPropertyId() (value uint64, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

//

// <param name="ClusterName" type="string []"></param>
// <param name="Description" type="string "></param>
// <param name="FDName" type="string "></param>
// <param name="FDType" type="uint32 "></param>
// <param name="Flags" type="uint32 "></param>

// <param name="CreatedFaultDomain" type="MSFT_ClusterSetFaultDomain "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) CreateFaultDomain( /* IN */ FDName string,
	/* IN */ ClusterName []string,
	/* IN */ FDType uint32,
	/* IN */ Description string,
	/* IN */ Flags uint32,
	/* OUT */ CreatedFaultDomain MSFT_ClusterSetFaultDomain) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateFaultDomain", FDName, ClusterName, FDType, Description, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) RemoveFaultDomain( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveFaultDomain", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterName" type="string []"></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) AddMembers( /* IN */ ClusterName []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddMembers", ClusterName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClusterName" type="string []"></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ClusterSetFaultDomain) RemoveMembers( /* IN */ ClusterName []string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveMembers", ClusterName, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
