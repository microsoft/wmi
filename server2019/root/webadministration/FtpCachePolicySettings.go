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

// FtpCachePolicySettings struct
type FtpCachePolicySettings struct { 
	*EmbeddedObject

	// 
	PolicyLevel int32
}

	func NewFtpCachePolicySettingsEx1(instance *cim.WmiInstance) (newInstance *FtpCachePolicySettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpCachePolicySettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpCachePolicySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpCachePolicySettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpCachePolicySettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetPolicyLevel sets the value of PolicyLevel for the instance
func (instance *FtpCachePolicySettings) SetPropertyPolicyLevel(value int32) (err error) { 
    return instance.SetProperty("PolicyLevel", (value))
}

// GetPolicyLevel gets the value of PolicyLevel for the instance
func (instance *FtpCachePolicySettings) GetPropertyPolicyLevel()(value int32, err error) { 
    retValue, err := instance.GetProperty("PolicyLevel")
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

