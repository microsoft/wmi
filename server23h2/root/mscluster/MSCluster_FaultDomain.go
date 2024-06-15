// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_FaultDomain struct
type MSCluster_FaultDomain struct {
	*cim.WmiInstance

	//
	Description string

	//
	Id string

	//
	Location string

	//
	Name string

	//
	Type uint32
}

func NewMSCluster_FaultDomainEx1(instance *cim.WmiInstance) (newInstance *MSCluster_FaultDomain, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSCluster_FaultDomain{
		WmiInstance: tmp,
	}
	return
}

func NewMSCluster_FaultDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_FaultDomain, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_FaultDomain{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSCluster_FaultDomain) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSCluster_FaultDomain) GetPropertyDescription() (value string, err error) {
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

// SetId sets the value of Id for the instance
func (instance *MSCluster_FaultDomain) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_FaultDomain) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetLocation sets the value of Location for the instance
func (instance *MSCluster_FaultDomain) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *MSCluster_FaultDomain) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
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

// SetName sets the value of Name for the instance
func (instance *MSCluster_FaultDomain) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSCluster_FaultDomain) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetType sets the value of Type for the instance
func (instance *MSCluster_FaultDomain) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSCluster_FaultDomain) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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

//

// <param name="Description" type="string "></param>
// <param name="FaultDomain" type="string "></param>
// <param name="FaultDomainType" type="uint32 "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Location" type="string "></param>
// <param name="Name" type="string "></param>

// <param name="CreatedFaultDomain" type="MSCluster_FaultDomain "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) CreateFaultDomain( /* IN */ Name string,
	/* IN */ FaultDomain string,
	/* IN */ FaultDomainType uint32,
	/* IN */ Description string,
	/* IN */ Location string,
	/* IN */ Flags uint32,
	/* OUT */ CreatedFaultDomain MSCluster_FaultDomain) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateFaultDomain", Name, FaultDomain, FaultDomainType, Description, Location, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="XML" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) SetFaultDomainXML( /* IN */ XML string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetFaultDomainXML", XML, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="XML" type="string "></param>
func (instance *MSCluster_FaultDomain) GetFaultDomainXML( /* OUT */ XML string,
	/* OPTIONAL IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetFaultDomainXML", Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FaultDomain" type="string "></param>
// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) MoveFaultDomain( /* IN */ FaultDomain string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("MoveFaultDomain", FaultDomain, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="Children" type="MSCluster_FaultDomain []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) GetChildren( /* OUT */ Children []MSCluster_FaultDomain,
	/* OPTIONAL IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetChildren", Flags)
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
// <param name="StorageNodes" type="MSCluster_StorageNode []"></param>
func (instance *MSCluster_FaultDomain) GetStorageNodes( /* OUT */ StorageNodes []MSCluster_StorageNode,
	/* OPTIONAL IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetStorageNodes", Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="Parent" type="MSCluster_FaultDomain "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) GetParent( /* OUT */ Parent MSCluster_FaultDomain,
	/* OPTIONAL IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetParent", Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Description" type="string "></param>
// <param name="FaultDomain" type="string "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Location" type="string "></param>
// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) SetFaultDomain( /* IN */ NewName string,
	/* IN */ FaultDomain string,
	/* IN */ Description string,
	/* IN */ Location string,
	/* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetFaultDomain", NewName, FaultDomain, Description, Location, Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_FaultDomain) RemoveFaultDomain( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveFaultDomain", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
