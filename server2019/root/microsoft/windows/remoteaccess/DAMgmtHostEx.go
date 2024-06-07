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

// DAMgmtHostEx struct
type DAMgmtHostEx struct { 
	*DAMgmtHost

	// 
	Type []string
}

	func NewDAMgmtHostExEx1(instance *cim.WmiInstance) (newInstance *DAMgmtHostEx, err error) {tmp, err := NewDAMgmtHostEx1(instance)
		
	if err != nil { return }
	newInstance = &DAMgmtHostEx {
	DAMgmtHost: tmp,
	}
	return
	}
	

	func NewDAMgmtHostExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAMgmtHostEx, err error) {tmp, err := NewDAMgmtHostEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAMgmtHostEx {
	DAMgmtHost: tmp,
	}
	return
	}
	

// SetType sets the value of Type for the instance
func (instance *DAMgmtHostEx) SetPropertyType(value []string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *DAMgmtHostEx) GetPropertyType()(value []string, err error) { 
    retValue, err := instance.GetProperty("Type")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

