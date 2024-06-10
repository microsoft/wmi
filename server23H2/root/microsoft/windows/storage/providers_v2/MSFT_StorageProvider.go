// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Storage.Providers_v2
//
// ////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageProvider struct
type MSFT_StorageProvider struct {
	*MSFT_StorageObject

	// A string indicating the manufacturer of the SMP software.
	Manufacturer string

	// A user-friendly name of the storage provider.
	Name string

	//
	RemoteSubsystemCacheMode StorageProvider_RemoteSubsystemCacheMode

	// Denotes the caching modes this provider supports. The modes are 'Disabled' and 'Manual-Discovery'.
	SupportedRemoteSubsystemCacheModes []StorageProvider_SupportedRemoteSubsystemCacheModes

	// Denotes whether this provider supports remote registration and management.
	SupportsSubsystemRegistration bool

	// Denotes whether the provider is a stand-alone SMP provider or an SMIS provider that uses the SMIS proxy SMP.
	Type StorageProvider_Type

	// A version string used by the SMP manufacturer to differentiate between software versions.
	Version string
}

func NewMSFT_StorageProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageProvider, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageProvider{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_StorageProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageProvider, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageProvider{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_StorageProvider) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_StorageProvider) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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
func (instance *MSFT_StorageProvider) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageProvider) GetPropertyName() (value string, err error) {
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

// SetRemoteSubsystemCacheMode sets the value of RemoteSubsystemCacheMode for the instance
func (instance *MSFT_StorageProvider) SetPropertyRemoteSubsystemCacheMode(value StorageProvider_RemoteSubsystemCacheMode) (err error) {
	return instance.SetProperty("RemoteSubsystemCacheMode", (value))
}

// GetRemoteSubsystemCacheMode gets the value of RemoteSubsystemCacheMode for the instance
func (instance *MSFT_StorageProvider) GetPropertyRemoteSubsystemCacheMode() (value StorageProvider_RemoteSubsystemCacheMode, err error) {
	retValue, err := instance.GetProperty("RemoteSubsystemCacheMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = StorageProvider_RemoteSubsystemCacheMode(valuetmp)

	return
}

// SetSupportedRemoteSubsystemCacheModes sets the value of SupportedRemoteSubsystemCacheModes for the instance
func (instance *MSFT_StorageProvider) SetPropertySupportedRemoteSubsystemCacheModes(value []StorageProvider_SupportedRemoteSubsystemCacheModes) (err error) {
	return instance.SetProperty("SupportedRemoteSubsystemCacheModes", (value))
}

// GetSupportedRemoteSubsystemCacheModes gets the value of SupportedRemoteSubsystemCacheModes for the instance
func (instance *MSFT_StorageProvider) GetPropertySupportedRemoteSubsystemCacheModes() (value []StorageProvider_SupportedRemoteSubsystemCacheModes, err error) {
	retValue, err := instance.GetProperty("SupportedRemoteSubsystemCacheModes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageProvider_SupportedRemoteSubsystemCacheModes(valuetmp))
	}

	return
}

// SetSupportsSubsystemRegistration sets the value of SupportsSubsystemRegistration for the instance
func (instance *MSFT_StorageProvider) SetPropertySupportsSubsystemRegistration(value bool) (err error) {
	return instance.SetProperty("SupportsSubsystemRegistration", (value))
}

// GetSupportsSubsystemRegistration gets the value of SupportsSubsystemRegistration for the instance
func (instance *MSFT_StorageProvider) GetPropertySupportsSubsystemRegistration() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsSubsystemRegistration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_StorageProvider) SetPropertyType(value StorageProvider_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_StorageProvider) GetPropertyType() (value StorageProvider_Type, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = StorageProvider_Type(valuetmp)

	return
}

// SetVersion sets the value of Version for the instance
func (instance *MSFT_StorageProvider) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MSFT_StorageProvider) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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

// This method is used when a user wants to explicitly discover or re-enumerate objects owned by the storage provider. A call to this method could result in a full or partial cache invalidation and over-the-wire calls to the storage subsystem(s) in order to discover new or updated objects. As this is a time consuming task, it should be used sparingly by the user.
///The scope of the discovery operation is controlled by the DiscoveryLevel and RootObject parameters. DiscoveryLevel controls the depth of the object discovery. RootObject defines the starting point from which discovery will happen.

// <param name="DiscoveryLevel" type="StorageProvider_DiscoveryLevel ">This field denotes the level (or depth) of discovery that should be performed. This parameter can only be specified if the root object is a storage provider, storage subsystem, or NULL. When specified, the storage provider will discover objects starting from Level 0 and continuing until the specified level is reached. Associations between objects (within the discovered levels) will also be discovered. The discovery levels are defined as follows:
///0 - 'Level 0': The storage provider and storage subsystem objects will be discovered.
///1 - 'Level 1': Storage pools, resiliency settings, target ports, target portals, and initiator ids will be discovered.
///2 - 'Level 2': Virtual disks and masking sets will be discovered.
///3 - 'Level 3': Physical disks will be discovered.</param>
// <param name="RootObject" type="MSFT_StorageObject ">If this parameter is set, discovery will begin from this object. When DiscoveryLevel is NULL, well-defined actions will be taken depending on the type of object specified by RootObject:
///Storage subsystem: All associated objects will be discovered.
///Storage pool: The pool, along with any associated resiliency settings, virtual disks, and physical disks will be discovered.
///Masking set: The masking set, along with any associated target ports, initiator ids, and virtual disks will be discovered.
///For all other objects: Only that object will be discovered / refreshed.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">This parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) Discover( /* IN */ DiscoveryLevel StorageProvider_DiscoveryLevel,
	/* IN */ RootObject MSFT_StorageObject,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Discover", DiscoveryLevel, RootObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method returns the security descriptor that controls access to this specific object instance.

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SecurityDescriptor" type="string ">A Security Descriptor Definition Language (SDDL) formed string describing the access control list of the object.</param>
func (instance *MSFT_StorageProvider) GetSecurityDescriptor( /* OUT */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a user with sufficient privileges to set the security descriptor that control access to this specific object instance. If the call is not made in the context of a user specified in the security descriptor's access control list, this method will fail with 40001 - 'Access Denied'. If an empty security descriptor is passed to this function, the behavior is left to the specific implementation so long as there is some user context (typically domain administrators) that can access and administer the object.

// <param name="SecurityDescriptor" type="string ">A Security Descriptor Definition Language (SDDL) formed string describing the desired access control list for this object.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) SetSecurityDescriptor( /* IN */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetSecurityDescriptor", SecurityDescriptor)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method registers a subsystem to be managed by this provider. Note that the subsystem must be compatible with the provider software.

// <param name="ComputerName" type="string "></param>
// <param name="Credential" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="RegisteredSubsystem" type="MSFT_StorageSubSystem "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) RegisterSubsystem( /* IN */ ComputerName string,
	/* IN */ Credential string,
	/* OUT */ RegisteredSubsystem MSFT_StorageSubSystem,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterSubsystem", ComputerName, Credential)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method unregisters a subsystem. The provider will no longer manage this subsystem.

// <param name="Force" type="bool "></param>
// <param name="StorageSubSystemUniqueId" type="string "></param>
// <param name="Subsystem" type="MSFT_StorageSubSystem "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) UnregisterSubsystem( /* IN */ Subsystem MSFT_StorageSubSystem,
	/* IN */ StorageSubSystemUniqueId string,
	/* IN */ Force bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UnregisterSubsystem", Subsystem, StorageSubSystemUniqueId, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method sets the attributes of the provider.

// <param name="RemoteSubsystemCacheMode" type="StorageProvider_RemoteSubsystemCacheMode ">If set to 3, caching for all the registered remote subsystem is enabled. If set to 2, caching for all the registered remote subsystem is disabled. This API only effects the remote subsystem registered and local Subsystem requests are not cached and reported live always.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageProvider) SetAttributes( /* IN */ RemoteSubsystemCacheMode StorageProvider_RemoteSubsystemCacheMode,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", RemoteSubsystemCacheMode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
