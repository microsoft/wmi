// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageObject struct
type MSFT_StorageObject struct {
	*cim.WmiInstance

	// ObjectId is a mandatory property that is used to opaquely and uniquely identify an instance of a class. ObjectIds must be unique within the scope of the management server (which is hosting the provider). The ObjectId is created and maintained for use of the Storage Management Providers and their clients to track instances of objects. If an object is visible through two different paths (for example: there are two separate Storage Management Providers that point to the same storage subsystem) then the same object may appear with two different ObjectIds. For determining if two object instances are the same object, refer to the UniqueId property.
	ObjectId string

	// PassThroughClass is the WBEM class name of the proprietary storage provider object.
	PassThroughClass string

	// PassThroughIds is a comma-separated list of all implementation specific keys. It is used by storage management applications to access the vendor proprietary object model. This field should be in the form: key1='value1',key2='value2'.
	PassThroughIds string

	// PassThroughNamespace is the WBEM namespace that contains the proprietary storage provider classes.
	PassThroughNamespace string

	// PassThroughServer is the name or address of the computer system hosting the proprietary storage provider classes.
	PassThroughServer string

	// UniqueId is a mandatory property that is used to uniquely identify a logical instance of a storage subsystem's object. This value must be the same for an object viewed by two or more provider instances (even if they are running on seperate management servers). UniqueId can be any globally unique, opaque value unless otherwise specified by a derived class.
	UniqueId string
}

func NewMSFT_StorageObjectEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageObject, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageObject{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_StorageObjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageObject, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageObject{
		WmiInstance: tmp,
	}
	return
}

// SetObjectId sets the value of ObjectId for the instance
func (instance *MSFT_StorageObject) SetPropertyObjectId(value string) (err error) {
	return instance.SetProperty("ObjectId", (value))
}

// GetObjectId gets the value of ObjectId for the instance
func (instance *MSFT_StorageObject) GetPropertyObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("ObjectId")
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

// SetPassThroughClass sets the value of PassThroughClass for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughClass(value string) (err error) {
	return instance.SetProperty("PassThroughClass", (value))
}

// GetPassThroughClass gets the value of PassThroughClass for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughClass() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughClass")
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

// SetPassThroughIds sets the value of PassThroughIds for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughIds(value string) (err error) {
	return instance.SetProperty("PassThroughIds", (value))
}

// GetPassThroughIds gets the value of PassThroughIds for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughIds() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughIds")
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

// SetPassThroughNamespace sets the value of PassThroughNamespace for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughNamespace(value string) (err error) {
	return instance.SetProperty("PassThroughNamespace", (value))
}

// GetPassThroughNamespace gets the value of PassThroughNamespace for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughNamespace")
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

// SetPassThroughServer sets the value of PassThroughServer for the instance
func (instance *MSFT_StorageObject) SetPropertyPassThroughServer(value string) (err error) {
	return instance.SetProperty("PassThroughServer", (value))
}

// GetPassThroughServer gets the value of PassThroughServer for the instance
func (instance *MSFT_StorageObject) GetPropertyPassThroughServer() (value string, err error) {
	retValue, err := instance.GetProperty("PassThroughServer")
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

// SetUniqueId sets the value of UniqueId for the instance
func (instance *MSFT_StorageObject) SetPropertyUniqueId(value string) (err error) {
	return instance.SetProperty("UniqueId", (value))
}

// GetUniqueId gets the value of UniqueId for the instance
func (instance *MSFT_StorageObject) GetPropertyUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueId")
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
