// +build windows
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package monitor

import (
	"fmt"
	"sync"

	"github.com/microsoft/wmi/pkg/base/event"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Monitor is a generic monitor to subscribe to Wmi Events based on a Query String
type Monitor struct {
	mux sync.Mutex
	// eventSinks for each of the entities that are being monitored here
	eventSinks          map[string]*wmi.WmiEventSink
	propertyNameToQuery string
	wmiNamespaceName    string
	queryString         string
	callbackContext     interface{}
	callbackFunction    func(interface{}, string)
}

func (c *Monitor) onModified(instance *wmi.WmiInstance, cbData string) {
	/*
		prop, err := instance.GetProperty(c.propertyNameToQuery)
		if err != nil {
			fmt.Printf("Err: %v - [%s]\n", err, c.propertyNameToQuery)
			return
		}
		propVal, ok := prop.(string)
		if !ok {
			return
		}
	*/
	c.callbackFunction(c.callbackContext, cbData)
	return
}

// CreateMonitor createa a new Monitor
func CreateMonitor(wmiNamespace string, callbackContext interface{},
	callback func(interface{}, string), queryString, propertyToQuery string) *Monitor {
	return &Monitor{
		wmiNamespaceName:    wmiNamespace,
		eventSinks:          map[string]*wmi.WmiEventSink{},
		callbackFunction:    callback,
		callbackContext:     callbackContext,
		propertyNameToQuery: propertyToQuery,
		queryString:         queryString,
	}
}

// AddEntity would add the entity to be monitored for changes
func (c *Monitor) AddEntity(entityName string) (err error) {
	if _, ok := c.eventSinks[entityName]; ok {
		return nil // error
	}
	c.mux.Lock()
	defer c.mux.Unlock()

	c.eventSinks[entityName], err = c.getEventSink(entityName)
	return
}

// RemoveEntity to remove the entity being monitored for changes
func (c *Monitor) RemoveEntity(entityName string) (err error) {
	if _, ok := c.eventSinks[entityName]; !ok {
		return nil // error NOT Found
	}
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.eventSinks, entityName)
	return
}

// Close the monitor
func (c *Monitor) Close() {
	for k := range c.eventSinks {
		c.eventSinks[k].Close()
		delete(c.eventSinks, k)
	}
}

// getEventSink
func (c *Monitor) getEventSink(entityName string) (*wmi.WmiEventSink, error) {
	qString := fmt.Sprintf("%s AND TargetInstance.ElementName ='%s'", c.queryString, entityName)
	context := event.NewCallbackContext(c.onModified, entityName)
	return event.RegisterWmiCallback(context, c.wmiNamespaceName, constant.HostName, qString)
}
