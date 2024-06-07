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

// MDM_Policy_Result01_SmartScreen02 struct
type MDM_Policy_Result01_SmartScreen02 struct { 
	*cim.WmiInstance

	// 
	EnableAppInstallControl int32

	// 
	EnableSmartScreenInShell int32

	// 
	InstanceID string

	// 
	ParentID string

	// 
	PreventOverrideForFilesInShell int32
}

	func NewMDM_Policy_Result01_SmartScreen02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_SmartScreen02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_SmartScreen02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_Result01_SmartScreen02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_Result01_SmartScreen02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_SmartScreen02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetEnableAppInstallControl sets the value of EnableAppInstallControl for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyEnableAppInstallControl(value int32) (err error) { 
    return instance.SetProperty("EnableAppInstallControl", (value))
}

// GetEnableAppInstallControl gets the value of EnableAppInstallControl for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyEnableAppInstallControl()(value int32, err error) { 
    retValue, err := instance.GetProperty("EnableAppInstallControl")
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

// SetEnableSmartScreenInShell sets the value of EnableSmartScreenInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyEnableSmartScreenInShell(value int32) (err error) { 
    return instance.SetProperty("EnableSmartScreenInShell", (value))
}

// GetEnableSmartScreenInShell gets the value of EnableSmartScreenInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyEnableSmartScreenInShell()(value int32, err error) { 
    retValue, err := instance.GetProperty("EnableSmartScreenInShell")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyInstanceID()(value string, err error) { 
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
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyParentID()(value string, err error) { 
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

// SetPreventOverrideForFilesInShell sets the value of PreventOverrideForFilesInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) SetPropertyPreventOverrideForFilesInShell(value int32) (err error) { 
    return instance.SetProperty("PreventOverrideForFilesInShell", (value))
}

// GetPreventOverrideForFilesInShell gets the value of PreventOverrideForFilesInShell for the instance
func (instance *MDM_Policy_Result01_SmartScreen02) GetPropertyPreventOverrideForFilesInShell()(value int32, err error) { 
    retValue, err := instance.GetProperty("PreventOverrideForFilesInShell")
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

