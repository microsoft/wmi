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

// OOB_Service struct
type OOB_Service struct {
	*CIM_Service
}

func NewOOB_ServiceEx1(instance *cim.WmiInstance) (newInstance *OOB_Service, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OOB_Service{
		CIM_Service: tmp,
	}
	return
}

func NewOOB_ServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OOB_Service, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OOB_Service{
		CIM_Service: tmp,
	}
	return
}

// Returns the current provisioning (activation) mode of the device. This command is deprecated since Intel(R) AMT 7.0

// <param name="mode" type="Service_mode ">Provisioning mode</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetAMTProvisioningMode( /* OUT */ mode Service_mode) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAMTProvisioningMode")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Indicates if Remote Configuration is enabled in the MEBx

// <param name="enabled" type="bool ">Indicates if Remote Configuration is enabled in the MEBx</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) isRemoteConfigEnabled( /* OUT */ enabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("isRemoteConfigEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Indicates whether the device is using TLS-PSK or TLS-PKI activation mode

// <param name="mode" type="Service_mode_2 ">TLS activation mode</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetActivationTLSMode( /* OUT */ mode Service_mode_2) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetActivationTLSMode")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// When the device is configured in TLS-PSK mode, this functions returns the PID value of the device, this can be used by Provisioning servers to associate a PSK value (deprecated)

// <param name="PID" type="string ">PID value</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetPID( /* OUT */ PID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPID")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Returns the information about the server which Intel(R) AMT will send Hello packets to. The OOB server (aka Intel(R) AMT) sends 'Hello packet' messages to this address as a mean to trigger remote configuration

// <param name="Address" type="string ">returns the address information (IPv4, IPv6 or FQDN) of the configuration server</param>
// <param name="ConfigServerListeningPort" type="uint16 ">TCP Port of the configuration server to send Hello packets to</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetHelloPacketDestInfo( /* OUT */ Address string,
	/* OUT */ ConfigServerListeningPort uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetHelloPacketDestInfo")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Intel(R) AMT's configured Fully Qualified Domain Name

// <param name="FQDN" type="string ">Intel(R) AMT's configured Fully Qualified Domain Name</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetAMTFQDN( /* OUT */ FQDN string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAMTFQDN")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Activate Intel(R) AMT using Remote Configuration (Deprecated)

// <param name="OTP" type="string ">Optional One-Time-Password value the remote configuration service can use to authenticate the device</param>
// <param name="PKIDNSSuffix" type="string ">Optional DNSSuffix to match</param>

// <param name="provTLSMode" type="Service_provTLSMode ">Intel(R) AMT's TLS activation mode</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) Activate( /* IN */ OTP string,
	/* IN */ PKIDNSSuffix string,
	/* OUT */ provTLSMode Service_provTLSMode) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Activate", OTP, PKIDNSSuffix)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Stops Intel(R) AMT Activation if it is currently enabled, this will prevent Intel(R) AMT from being present on the Network for Bare-Metal activation (Deprected)

// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) CancelActivation() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CancelActivation")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// GetProvisioningState

// <param name="ReturnValue" type="uint32 "></param>
// <param name="state" type="Service_state ">Provisioning state</param>
func (instance *OOB_Service) GetProvisioningState( /* OUT */ state Service_state) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetProvisioningState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// GetProvisioningInfo returns the FQDN of configuration server and the PKI FQDN suffix

// <param name="ConfigServerFQDN" type="string ">The ConfigurationServerFQDN, set to null if not specified</param>
// <param name="PKIDNSSuffix" type="string ">PKIDNSSuffix Specified</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetProvisioningInfo( /* OUT */ PKIDNSSuffix string,
	/* OUT */ ConfigServerFQDN string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetProvisioningInfo")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Specifies whether Intel(R) AMT TLS port is enabled

// <param name="enabled" type="bool ">Specifies whether Intel(R) AMT TLS port is enabled</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) isTLSEnabled( /* OUT */ enabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("isTLSEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// CIRA: User request to initiate a CIRA connection

// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) OpenUserInitiatedConnection() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("OpenUserInitiatedConnection")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// CIRA: Request Intel(R) ME to close an existing CIRA connection

// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) CloseUserInitiatedConnection() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloseUserInitiatedConnection")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// CIRA: Retrieves information of an active CIRA connection

// <param name="ConnectionTrigger" type="Service_ConnectionTrigger ">Indicates the trigger for the remote access establishment in case a connection status is connected/connecting</param>
// <param name="MPshostName" type="string ">The MPS hostname that Intel(R) AMT is connected or trying to connect to</param>
// <param name="NetworkConStatus" type="Service_NetworkConStatus ">Indicates the network connection status from Intel(R) AMT standpoint. Starting in v7.0, if the network cable is unplugged, an Unknown value appears.</param>
// <param name="RemoteAccessConStatus" type="Service_RemoteAccessConStatus ">The status of the CIRA connection (not connected, connecting, connected). Starting in v7.0, this value indicates the CIRA connection.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) GetRemoteAccessConnectionStatus( /* OUT */ NetworkConStatus Service_NetworkConStatus,
	/* OUT */ ConnectionTrigger Service_ConnectionTrigger,
	/* OUT */ MPshostName string,
	/* OUT */ RemoteAccessConStatus Service_RemoteAccessConStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetRemoteAccessConnectionStatus")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// Un-configure the AMT system

// <param name="ReturnValue" type="uint32 "></param>
func (instance *OOB_Service) Unconfigure() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Unconfigure")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Get the AMT local Administrator credentials. A local application or a service running as a Windows local administrator, can use this method to retrieve credentials from AMT. This application can use these credentials to authenticate to AMT during WSMAN commands.

// <param name="Password" type="string ">Local administrator password</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="Username" type="string ">Local administrator username</param>
func (instance *OOB_Service) GetLocalAdminCredentials( /* OUT */ Username string,
	/* OUT */ Password string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLocalAdminCredentials")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
