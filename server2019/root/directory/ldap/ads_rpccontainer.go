// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_rpccontainer struct
type ads_rpccontainer struct { 
	*ads_container

	// 
	DS_nameServiceFlags int32
}

	func Newads_rpccontainerEx1(instance *cim.WmiInstance) (newInstance *ads_rpccontainer, err error) {tmp, err := Newads_containerEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_rpccontainer {
	ads_container: tmp,
	}
	return
	}
	

	func Newads_rpccontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_rpccontainer, err error) {tmp, err := Newads_containerEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_rpccontainer {
	ads_container: tmp,
	}
	return
	}
	

// SetDS_nameServiceFlags sets the value of DS_nameServiceFlags for the instance
func (instance *ads_rpccontainer) SetPropertyDS_nameServiceFlags(value int32) (err error) { 
    return instance.SetProperty("DS_nameServiceFlags", (value))
}

// GetDS_nameServiceFlags gets the value of DS_nameServiceFlags for the instance
func (instance *ads_rpccontainer) GetPropertyDS_nameServiceFlags()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_nameServiceFlags")
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

