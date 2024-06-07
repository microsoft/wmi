// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// WebDavFileSystemSettings struct
type WebDavFileSystemSettings struct { 
	*EmbeddedObject

	// 
	AllowHiddenFiles bool

	// 
	HideChildVirtualDirectories bool

	// 
	UseTransactionalIo bool
}

	func NewWebDavFileSystemSettingsEx1(instance *cim.WmiInstance) (newInstance *WebDavFileSystemSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &WebDavFileSystemSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewWebDavFileSystemSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WebDavFileSystemSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WebDavFileSystemSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAllowHiddenFiles sets the value of AllowHiddenFiles for the instance
func (instance *WebDavFileSystemSettings) SetPropertyAllowHiddenFiles(value bool) (err error) { 
    return instance.SetProperty("AllowHiddenFiles", (value))
}

// GetAllowHiddenFiles gets the value of AllowHiddenFiles for the instance
func (instance *WebDavFileSystemSettings) GetPropertyAllowHiddenFiles()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowHiddenFiles")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetHideChildVirtualDirectories sets the value of HideChildVirtualDirectories for the instance
func (instance *WebDavFileSystemSettings) SetPropertyHideChildVirtualDirectories(value bool) (err error) { 
    return instance.SetProperty("HideChildVirtualDirectories", (value))
}

// GetHideChildVirtualDirectories gets the value of HideChildVirtualDirectories for the instance
func (instance *WebDavFileSystemSettings) GetPropertyHideChildVirtualDirectories()(value bool, err error) { 
    retValue, err := instance.GetProperty("HideChildVirtualDirectories")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetUseTransactionalIo sets the value of UseTransactionalIo for the instance
func (instance *WebDavFileSystemSettings) SetPropertyUseTransactionalIo(value bool) (err error) { 
    return instance.SetProperty("UseTransactionalIo", (value))
}

// GetUseTransactionalIo gets the value of UseTransactionalIo for the instance
func (instance *WebDavFileSystemSettings) GetPropertyUseTransactionalIo()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseTransactionalIo")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

