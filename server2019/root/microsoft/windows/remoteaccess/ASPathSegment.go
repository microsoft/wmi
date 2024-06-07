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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ASPathSegment struct
type ASPathSegment struct { 
	*cim.WmiInstance

	// 
	ASN []uint32

	// 
	PathType string
}

	func NewASPathSegmentEx1(instance *cim.WmiInstance) (newInstance *ASPathSegment, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ASPathSegment {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewASPathSegmentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ASPathSegment, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ASPathSegment {
	WmiInstance: tmp,
	}
	return
	}
	

// SetASN sets the value of ASN for the instance
func (instance *ASPathSegment) SetPropertyASN(value []uint32) (err error) { 
    return instance.SetProperty("ASN", (value))
}

// GetASN gets the value of ASN for the instance
func (instance *ASPathSegment) GetPropertyASN()(value []uint32, err error) { 
    retValue, err := instance.GetProperty("ASN")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint32(valuetmp))
    }

    return
}

// SetPathType sets the value of PathType for the instance
func (instance *ASPathSegment) SetPropertyPathType(value string) (err error) { 
    return instance.SetProperty("PathType", (value))
}

// GetPathType gets the value of PathType for the instance
func (instance *ASPathSegment) GetPropertyPathType()(value string, err error) { 
    retValue, err := instance.GetProperty("PathType")
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

