// +build windows
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package monitor

import (
	"fmt"
	"sync"

	"github.com/microsoft/wmi/pkg/base/event"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Monitor is a generic monitor to subscribe to Wmi Events based on a Query String
type Monitor struct {
	mux sync.Mutex
	// eventSinks for each of the entities that are being monitored here
	// multiple event sinks can be setup for a single entity
	eventSinks          map[string][]*wmi.WmiEventSink
	propertyNameToQuery string
	wmiNamespaceName    string
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
	callback func(interface{}, string)) *Monitor {
	return &Monitor{
		wmiNamespaceName: wmiNamespace,
		eventSinks:       map[string][]*wmi.WmiEventSink{},
		callbackFunction: callback,
		callbackContext:  callbackContext,
	}
}

// AddEntityWithFilter would add the entity to be monitored for changes
func (c *Monitor) AddEntityWithFilter(entityName, wqlQueryString string, filters query.WmiQueryFilterCollection) (err error) {
	if _, ok := c.eventSinks[entityName]; ok {
		return nil // error
	}
	qString := fmt.Sprintf("%s %s", wqlQueryString, filters.String())
	fmt.Printf("Event Query [%s]\n", qString)
	esink, err := c.getEventSink(entityName, qString)
	if err != nil {
		return
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	c.eventSinks[entityName] = append(c.eventSinks[entityName], esink)
	return
}

// RemoveEntity to remove the entity being monitored for changes
func (c *Monitor) RemoveEntity(entityName string) (err error) {
	val, ok := c.eventSinks[entityName]
	if !ok {
		return nil // error NOT Found
	}
	for _, es := range val {
		es.Close()
	}

	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.eventSinks, entityName)
	return
}

// Close the monitor
func (c *Monitor) Close() error {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k := range c.eventSinks {
		for _, s := range c.eventSinks[k] {
			s.Close()
		}
		// TODO: There may be pending/ongoing events - delay cleanp?
		delete(c.eventSinks, k)
	}
	return nil
}

// getEventSink
func (c *Monitor) getEventSink(entityName, wqlQueryString string) (*wmi.WmiEventSink, error) {
	//qString := fmt.Sprintf("%s %s", queryString, additionalFilter)
	context := event.NewCallbackContext(c.onModified, entityName)
	return event.RegisterWmiCallback(context, c.wmiNamespaceName, constant.HostName, wqlQueryString)
}
