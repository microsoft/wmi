// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetTeredoConfiguration struct
type MSFT_NetTeredoConfiguration struct {
	MSFT_NetSettingData

	//
	ClientPort uint32

	//
	DefaultQualified bool

	//
	PolicyStore string

	//
	RefreshInterval uint32

	//
	ServerName string

	//
	ServerShunt bool

	//
	ServerVirtualIP string

	//
	Type uint32
}

// SetClientPort sets the value of ClientPort for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyClientPort(value uint32) (err error) {
	return instance.SetProperty("ClientPort", value)
}

// GetClientPort gets the value of ClientPort for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyClientPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultQualified sets the value of DefaultQualified for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyDefaultQualified(value bool) (err error) {
	return instance.SetProperty("DefaultQualified", value)
}

// GetDefaultQualified gets the value of DefaultQualified for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyDefaultQualified() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQualified")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", value)
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRefreshInterval sets the value of RefreshInterval for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyRefreshInterval(value uint32) (err error) {
	return instance.SetProperty("RefreshInterval", value)
}

// GetRefreshInterval gets the value of RefreshInterval for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyRefreshInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("RefreshInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerName sets the value of ServerName for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", value)
}

// GetServerName gets the value of ServerName for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerShunt sets the value of ServerShunt for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyServerShunt(value bool) (err error) {
	return instance.SetProperty("ServerShunt", value)
}

// GetServerShunt gets the value of ServerShunt for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyServerShunt() (value bool, err error) {
	retValue, err := instance.GetProperty("ServerShunt")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerVirtualIP sets the value of ServerVirtualIP for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyServerVirtualIP(value string) (err error) {
	return instance.SetProperty("ServerVirtualIP", value)
}

// GetServerVirtualIP gets the value of ServerVirtualIP for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyServerVirtualIP() (value string, err error) {
	retValue, err := instance.GetProperty("ServerVirtualIP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_NetTeredoConfiguration) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_NetTeredoConfiguration) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ClientPort" type="bool "></param>
// <param name="DefaultQualified" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RefreshInterval" type="bool "></param>
// <param name="ServerName" type="bool "></param>
// <param name="ServerShunt" type="bool "></param>
// <param name="ServerVirtualIP" type="bool "></param>
// <param name="Type" type="bool "></param>

// <param name="OutputObject" type="MSFT_NetTeredoConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetTeredoConfiguration) Reset( /* IN */ Type bool,
	/* IN */ ServerName bool,
	/* IN */ RefreshInterval bool,
	/* IN */ ClientPort bool,
	/* IN */ ServerVirtualIP bool,
	/* IN */ DefaultQualified bool,
	/* IN */ ServerShunt bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetTeredoConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", Type, ServerName, RefreshInterval, ClientPort, ServerVirtualIP, DefaultQualified, ServerShunt, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
