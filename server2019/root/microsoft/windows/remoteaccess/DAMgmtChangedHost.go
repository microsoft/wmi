// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DAMgmtChangedHost struct
type DAMgmtChangedHost struct { 
	*DAMgmtHostEx

	// 
	Change string
}

	func NewDAMgmtChangedHostEx1(instance *cim.WmiInstance) (newInstance *DAMgmtChangedHost, err error) {tmp, err := NewDAMgmtHostExEx1(instance)
		
	if err != nil { return }
	newInstance = &DAMgmtChangedHost {
	DAMgmtHostEx: tmp,
	}
	return
	}
	

	func NewDAMgmtChangedHostEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAMgmtChangedHost, err error) {tmp, err := NewDAMgmtHostExEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAMgmtChangedHost {
	DAMgmtHostEx: tmp,
	}
	return
	}
	

// SetChange sets the value of Change for the instance
func (instance *DAMgmtChangedHost) SetPropertyChange(value string) (err error) { 
    return instance.SetProperty("Change", (value))
}

// GetChange gets the value of Change for the instance
func (instance *DAMgmtChangedHost) GetPropertyChange()(value string, err error) { 
    retValue, err := instance.GetProperty("Change")
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

