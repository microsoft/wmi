// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpStatistics struct
type BgpStatistics struct {
	*cim.WmiInstance

	//
	IPv4Route BgpRouteCounter

	//
	IPv6Route BgpRouteCounter

	//
	KeepAliveMessage BgpMessageStatistics

	//
	NotificationMessage BgpMessageStatistics

	//
	OpenMessage BgpMessageStatistics

	//
	PeerName string

	//
	RouteRefreshMessage BgpMessageStatistics

	//
	RoutingDomain string

	//
	TcpConnectionClosed string

	//
	TcpConnectionEstablished string

	//
	UpdateMessage BgpMessageStatistics
}

func NewBgpStatisticsEx1(instance *cim.WmiInstance) (newInstance *BgpStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewBgpStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetIPv4Route sets the value of IPv4Route for the instance
func (instance *BgpStatistics) SetPropertyIPv4Route(value BgpRouteCounter) (err error) {
	return instance.SetProperty("IPv4Route", (value))
}

// GetIPv4Route gets the value of IPv4Route for the instance
func (instance *BgpStatistics) GetPropertyIPv4Route() (value BgpRouteCounter, err error) {
	retValue, err := instance.GetProperty("IPv4Route")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpRouteCounter)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpRouteCounter is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpRouteCounter(valuetmp)

	return
}

// SetIPv6Route sets the value of IPv6Route for the instance
func (instance *BgpStatistics) SetPropertyIPv6Route(value BgpRouteCounter) (err error) {
	return instance.SetProperty("IPv6Route", (value))
}

// GetIPv6Route gets the value of IPv6Route for the instance
func (instance *BgpStatistics) GetPropertyIPv6Route() (value BgpRouteCounter, err error) {
	retValue, err := instance.GetProperty("IPv6Route")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpRouteCounter)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpRouteCounter is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpRouteCounter(valuetmp)

	return
}

// SetKeepAliveMessage sets the value of KeepAliveMessage for the instance
func (instance *BgpStatistics) SetPropertyKeepAliveMessage(value BgpMessageStatistics) (err error) {
	return instance.SetProperty("KeepAliveMessage", (value))
}

// GetKeepAliveMessage gets the value of KeepAliveMessage for the instance
func (instance *BgpStatistics) GetPropertyKeepAliveMessage() (value BgpMessageStatistics, err error) {
	retValue, err := instance.GetProperty("KeepAliveMessage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpMessageStatistics)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpMessageStatistics is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpMessageStatistics(valuetmp)

	return
}

// SetNotificationMessage sets the value of NotificationMessage for the instance
func (instance *BgpStatistics) SetPropertyNotificationMessage(value BgpMessageStatistics) (err error) {
	return instance.SetProperty("NotificationMessage", (value))
}

// GetNotificationMessage gets the value of NotificationMessage for the instance
func (instance *BgpStatistics) GetPropertyNotificationMessage() (value BgpMessageStatistics, err error) {
	retValue, err := instance.GetProperty("NotificationMessage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpMessageStatistics)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpMessageStatistics is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpMessageStatistics(valuetmp)

	return
}

// SetOpenMessage sets the value of OpenMessage for the instance
func (instance *BgpStatistics) SetPropertyOpenMessage(value BgpMessageStatistics) (err error) {
	return instance.SetProperty("OpenMessage", (value))
}

// GetOpenMessage gets the value of OpenMessage for the instance
func (instance *BgpStatistics) GetPropertyOpenMessage() (value BgpMessageStatistics, err error) {
	retValue, err := instance.GetProperty("OpenMessage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpMessageStatistics)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpMessageStatistics is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpMessageStatistics(valuetmp)

	return
}

// SetPeerName sets the value of PeerName for the instance
func (instance *BgpStatistics) SetPropertyPeerName(value string) (err error) {
	return instance.SetProperty("PeerName", (value))
}

// GetPeerName gets the value of PeerName for the instance
func (instance *BgpStatistics) GetPropertyPeerName() (value string, err error) {
	retValue, err := instance.GetProperty("PeerName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRouteRefreshMessage sets the value of RouteRefreshMessage for the instance
func (instance *BgpStatistics) SetPropertyRouteRefreshMessage(value BgpMessageStatistics) (err error) {
	return instance.SetProperty("RouteRefreshMessage", (value))
}

// GetRouteRefreshMessage gets the value of RouteRefreshMessage for the instance
func (instance *BgpStatistics) GetPropertyRouteRefreshMessage() (value BgpMessageStatistics, err error) {
	retValue, err := instance.GetProperty("RouteRefreshMessage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpMessageStatistics)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpMessageStatistics is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpMessageStatistics(valuetmp)

	return
}

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpStatistics) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpStatistics) GetPropertyRoutingDomain() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomain")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTcpConnectionClosed sets the value of TcpConnectionClosed for the instance
func (instance *BgpStatistics) SetPropertyTcpConnectionClosed(value string) (err error) {
	return instance.SetProperty("TcpConnectionClosed", (value))
}

// GetTcpConnectionClosed gets the value of TcpConnectionClosed for the instance
func (instance *BgpStatistics) GetPropertyTcpConnectionClosed() (value string, err error) {
	retValue, err := instance.GetProperty("TcpConnectionClosed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTcpConnectionEstablished sets the value of TcpConnectionEstablished for the instance
func (instance *BgpStatistics) SetPropertyTcpConnectionEstablished(value string) (err error) {
	return instance.SetProperty("TcpConnectionEstablished", (value))
}

// GetTcpConnectionEstablished gets the value of TcpConnectionEstablished for the instance
func (instance *BgpStatistics) GetPropertyTcpConnectionEstablished() (value string, err error) {
	retValue, err := instance.GetProperty("TcpConnectionEstablished")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetUpdateMessage sets the value of UpdateMessage for the instance
func (instance *BgpStatistics) SetPropertyUpdateMessage(value BgpMessageStatistics) (err error) {
	return instance.SetProperty("UpdateMessage", (value))
}

// GetUpdateMessage gets the value of UpdateMessage for the instance
func (instance *BgpStatistics) GetPropertyUpdateMessage() (value BgpMessageStatistics, err error) {
	retValue, err := instance.GetProperty("UpdateMessage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BgpMessageStatistics)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BgpMessageStatistics is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BgpMessageStatistics(valuetmp)

	return
}
