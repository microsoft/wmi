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

// SocketSettings struct
type SocketSettings struct { 
	*EmbeddedObject

	// 
	AlwaysUseCompletionPortsForAccept bool

	// 
	AlwaysUseCompletionPortsForConnect bool
}

	func NewSocketSettingsEx1(instance *cim.WmiInstance) (newInstance *SocketSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &SocketSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewSocketSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SocketSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SocketSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAlwaysUseCompletionPortsForAccept sets the value of AlwaysUseCompletionPortsForAccept for the instance
func (instance *SocketSettings) SetPropertyAlwaysUseCompletionPortsForAccept(value bool) (err error) { 
    return instance.SetProperty("AlwaysUseCompletionPortsForAccept", (value))
}

// GetAlwaysUseCompletionPortsForAccept gets the value of AlwaysUseCompletionPortsForAccept for the instance
func (instance *SocketSettings) GetPropertyAlwaysUseCompletionPortsForAccept()(value bool, err error) { 
    retValue, err := instance.GetProperty("AlwaysUseCompletionPortsForAccept")
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

// SetAlwaysUseCompletionPortsForConnect sets the value of AlwaysUseCompletionPortsForConnect for the instance
func (instance *SocketSettings) SetPropertyAlwaysUseCompletionPortsForConnect(value bool) (err error) { 
    return instance.SetProperty("AlwaysUseCompletionPortsForConnect", (value))
}

// GetAlwaysUseCompletionPortsForConnect gets the value of AlwaysUseCompletionPortsForConnect for the instance
func (instance *SocketSettings) GetPropertyAlwaysUseCompletionPortsForConnect()(value bool, err error) { 
    retValue, err := instance.GetProperty("AlwaysUseCompletionPortsForConnect")
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

