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

// ads_msimaging_postscanprocess struct
type ads_msimaging_postscanprocess struct { 
	*ds_top

	// 
	DS_msImaging_PSPIdentifier Uint8Array

	// 
	DS_msImaging_PSPString string

	// 
	DS_serverName string
}

	func Newads_msimaging_postscanprocessEx1(instance *cim.WmiInstance) (newInstance *ads_msimaging_postscanprocess, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msimaging_postscanprocess {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msimaging_postscanprocessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msimaging_postscanprocess, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msimaging_postscanprocess {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msImaging_PSPIdentifier sets the value of DS_msImaging_PSPIdentifier for the instance
func (instance *ads_msimaging_postscanprocess) SetPropertyDS_msImaging_PSPIdentifier(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msImaging_PSPIdentifier", (value))
}

// GetDS_msImaging_PSPIdentifier gets the value of DS_msImaging_PSPIdentifier for the instance
func (instance *ads_msimaging_postscanprocess) GetPropertyDS_msImaging_PSPIdentifier()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msImaging_PSPIdentifier")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_msImaging_PSPString sets the value of DS_msImaging_PSPString for the instance
func (instance *ads_msimaging_postscanprocess) SetPropertyDS_msImaging_PSPString(value string) (err error) { 
    return instance.SetProperty("DS_msImaging_PSPString", (value))
}

// GetDS_msImaging_PSPString gets the value of DS_msImaging_PSPString for the instance
func (instance *ads_msimaging_postscanprocess) GetPropertyDS_msImaging_PSPString()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msImaging_PSPString")
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

// SetDS_serverName sets the value of DS_serverName for the instance
func (instance *ads_msimaging_postscanprocess) SetPropertyDS_serverName(value string) (err error) { 
    return instance.SetProperty("DS_serverName", (value))
}

// GetDS_serverName gets the value of DS_serverName for the instance
func (instance *ads_msimaging_postscanprocess) GetPropertyDS_serverName()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_serverName")
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

