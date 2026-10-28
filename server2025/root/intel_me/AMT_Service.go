// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// AMT_Service struct
type AMT_Service struct {
	*CIM_Service
}

func NewAMT_ServiceEx1(instance *cim.WmiInstance) (newInstance *AMT_Service, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AMT_Service{
		CIM_Service: tmp,
	}
	return
}

func NewAMT_ServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AMT_Service, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AMT_Service{
		CIM_Service: tmp,
	}
	return
}

// Indicates whether the WebUI capability is currently enabled

// <param name="enabled" type="bool ">Indicates whether the WebUI capability is currently enabled</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *AMT_Service) isWebUIEnabled( /* OUT */ enabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("isWebUIEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Indicates whether SOL is currently enabled / active

// <param name="active" type="bool ">Indicates whether SOL session is currently open</param>
// <param name="hardEnabled" type="Service_hardEnabled ">Indicates whether SOL is enabled by MEBX</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="softEnabled" type="Service_softEnabled ">Indicates whether SOL is enabled by WS-Management interface (using AMT_RedirectionService class)</param>
func (instance *AMT_Service) getSOLState( /* OUT */ hardEnabled Service_hardEnabled,
	/* OUT */ softEnabled Service_softEnabled,
	/* OUT */ active bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getSOLState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Indicates whether IDER is currently enabled / active

// <param name="active" type="bool ">Indicates whether IDER session is currently open</param>
// <param name="hardEnabled" type="Service_hardEnabled ">Indicates whether IDER is enabled by MEBX</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="softEnabled" type="Service_softEnabled ">Indicates whether IDER is enabled by WS-Management interface (using AMT_RedirectionService class)</param>
func (instance *AMT_Service) getIDERState( /* OUT */ hardEnabled Service_hardEnabled,
	/* OUT */ softEnabled Service_softEnabled,
	/* OUT */ active bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getIDERState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Indicates whether KVM is currently enabled / active

// <param name="active" type="bool ">Indicates whether KVM session is currently open</param>
// <param name="hardEnabled" type="Service_hardEnabled ">Indicates whether KVM is enabled by MEBX</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="softEnabled" type="Service_softEnabled ">Indicates whether KVM is enabled by WS-Management interface (using CIM_KVMRedirectionSAP class)</param>
func (instance *AMT_Service) getKVMState( /* OUT */ hardEnabled Service_hardEnabled,
	/* OUT */ softEnabled Service_softEnabled,
	/* OUT */ active bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("getKVMState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Terminates KVM session

// <param name="ReturnValue" type="uint32 "></param>
func (instance *AMT_Service) TerminateKVMSession() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TerminateKVMSession")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Allows user to specify Sprite zoom level

// <param name="zoom" type="Service_zoom ">Zoom level</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *AMT_Service) setSpriteZoom( /* IN */ zoom Service_zoom) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("setSpriteZoom", zoom)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Allows user to specify Sprite locale

// <param name="locale" type="Service_locale ">Sprite locale</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *AMT_Service) setSpriteLocale( /* IN */ locale Service_locale) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("setSpriteLocale", locale)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
