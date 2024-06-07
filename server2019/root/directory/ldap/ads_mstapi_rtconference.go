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

// ads_mstapi_rtconference struct
type ads_mstapi_rtconference struct { 
	*ds_top

	// 
	DS_msTAPI_ConferenceBlob Uint8Array

	// 
	DS_msTAPI_ProtocolId string

	// 
	DS_msTAPI_uid string
}

	func Newads_mstapi_rtconferenceEx1(instance *cim.WmiInstance) (newInstance *ads_mstapi_rtconference, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_mstapi_rtconference {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_mstapi_rtconferenceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_mstapi_rtconference, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_mstapi_rtconference {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msTAPI_ConferenceBlob sets the value of DS_msTAPI_ConferenceBlob for the instance
func (instance *ads_mstapi_rtconference) SetPropertyDS_msTAPI_ConferenceBlob(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msTAPI_ConferenceBlob", (value))
}

// GetDS_msTAPI_ConferenceBlob gets the value of DS_msTAPI_ConferenceBlob for the instance
func (instance *ads_mstapi_rtconference) GetPropertyDS_msTAPI_ConferenceBlob()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msTAPI_ConferenceBlob")
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

// SetDS_msTAPI_ProtocolId sets the value of DS_msTAPI_ProtocolId for the instance
func (instance *ads_mstapi_rtconference) SetPropertyDS_msTAPI_ProtocolId(value string) (err error) { 
    return instance.SetProperty("DS_msTAPI_ProtocolId", (value))
}

// GetDS_msTAPI_ProtocolId gets the value of DS_msTAPI_ProtocolId for the instance
func (instance *ads_mstapi_rtconference) GetPropertyDS_msTAPI_ProtocolId()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msTAPI_ProtocolId")
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

// SetDS_msTAPI_uid sets the value of DS_msTAPI_uid for the instance
func (instance *ads_mstapi_rtconference) SetPropertyDS_msTAPI_uid(value string) (err error) { 
    return instance.SetProperty("DS_msTAPI_uid", (value))
}

// GetDS_msTAPI_uid gets the value of DS_msTAPI_uid for the instance
func (instance *ads_mstapi_rtconference) GetPropertyDS_msTAPI_uid()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msTAPI_uid")
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

