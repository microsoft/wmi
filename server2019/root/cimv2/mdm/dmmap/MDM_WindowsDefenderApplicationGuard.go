// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_WindowsDefenderApplicationGuard struct
type MDM_WindowsDefenderApplicationGuard struct { 
	*cim.WmiInstance

	// 
	InstallWindowsDefenderApplicationGuard string

	// 
	InstanceID string

	// 
	ParentID string

	// 
	PlatformStatus int32

	// 
	Status int32
}

	func NewMDM_WindowsDefenderApplicationGuardEx1(instance *cim.WmiInstance) (newInstance *MDM_WindowsDefenderApplicationGuard, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_WindowsDefenderApplicationGuard {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_WindowsDefenderApplicationGuardEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_WindowsDefenderApplicationGuard, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_WindowsDefenderApplicationGuard {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInstallWindowsDefenderApplicationGuard sets the value of InstallWindowsDefenderApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyInstallWindowsDefenderApplicationGuard(value string) (err error) { 
    return instance.SetProperty("InstallWindowsDefenderApplicationGuard", (value))
}

// GetInstallWindowsDefenderApplicationGuard gets the value of InstallWindowsDefenderApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyInstallWindowsDefenderApplicationGuard()(value string, err error) { 
    retValue, err := instance.GetProperty("InstallWindowsDefenderApplicationGuard")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetPlatformStatus sets the value of PlatformStatus for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyPlatformStatus(value int32) (err error) { 
    return instance.SetProperty("PlatformStatus", (value))
}

// GetPlatformStatus gets the value of PlatformStatus for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyPlatformStatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("PlatformStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) SetPropertyStatus(value int32) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_WindowsDefenderApplicationGuard) GetPropertyStatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("Status")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// 

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WindowsDefenderApplicationGuard) InstallWindowsDefenderApplicationGuardMethod( /* IN */ param string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("InstallWindowsDefenderApplicationGuardMethod" , param);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


