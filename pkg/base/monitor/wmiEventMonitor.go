// +build windows
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

// TODO: merge the Monitor and the WmiNotificationMonitor classes together.
// The WmiNotificationMonitor is a superset of the Monitor class but approaches things slightly differently.
// It should be possible to make the monitor class internally use WmiNotificationMonitor

package monitor

import (
	"context"
	"sync"

	"github.com/microsoft/wmi/pkg/base/session"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

type WmiEventMessage interface {
	ToString() string
	AsFilter() interface{}
}

type WmiEventMessageFactory interface {
	NewWmiEventMessage(wmiInstances []*cim.WmiInstance) (WmiEventMessage, error)
}

type WmiNotificationLockedCallback struct {
	callback        func(context.Context, WmiEventMessage, interface{}) error
	callbackContext interface{}
	lock            sync.Mutex
}

type WmiNotificationMonitor struct {
	ctrlchannel       chan bool
	wmiNamespace      string
	wmiObjectName     string
	wmiHostDomainName string
	queryString       string
	wmiEventFactory   WmiEventMessageFactory
	wmiEventTriggered chan WmiEventMessage
	callbacks         map[interface{}]map[interface{}]*WmiNotificationLockedCallback
}

type WmiNotificationMonitorInstance struct {
	session   *cim.WmiSession
	eventSink *cim.WmiEventSink
}

func (c *WmiNotificationMonitorInstance) Close() {
	c.eventSink.Close()
}

func NewWmiNotificationMonitor(wmiNamespace string, wmiObjectName string, wmiHostDomainName string, queryString string, wmiEventFactory WmiEventMessageFactory) *WmiNotificationMonitor {
	return &WmiNotificationMonitor{
		wmiNamespace:      wmiNamespace,
		wmiObjectName:     wmiObjectName,
		wmiHostDomainName: wmiHostDomainName,
		wmiEventFactory:   wmiEventFactory,
		wmiEventTriggered: make(chan WmiEventMessage, 2),
		queryString:       queryString,
		ctrlchannel:       make(chan bool),
		callbacks:         map[interface{}]map[interface{}]*WmiNotificationLockedCallback{},
	}
}

// AddCallback adds a corresponding callback to be called when a notification comes in for the corresponding operation
func (n *WmiNotificationMonitor) AddCallback(filter interface{}, cb func(context.Context, WmiEventMessage, interface{}) error, callbackContext interface{}) (err error) {
	if n.callbacks[filter] == nil {
		n.callbacks[filter] = map[interface{}]*WmiNotificationLockedCallback{}
	}
	callback := &WmiNotificationLockedCallback{
		callback:        cb,
		callbackContext: callbackContext,
	}

	n.callbacks[filter][cb] = callback

	return nil
}

// RemoveCallback removes the corresponding callback
func (n *WmiNotificationMonitor) RemoveCallback(filter interface{}, cb func(context.Context, WmiEventMessage, interface{}) error) {
	delete(n.callbacks[filter], cb)

	if len(n.callbacks[filter]) == 0 {
		delete(n.callbacks, filter)
	}
}

// Stop
func (n *WmiNotificationMonitor) Stop() (err error) {
	n.ctrlchannel <- true
	return nil
}

// Start
func (n *WmiNotificationMonitor) Start() error {
	wmiCallbackInstance, err := n.registerWMICallback(n.wmiNamespace, n.wmiObjectName, n.wmiHostDomainName, n.queryString)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case message := <-n.ctrlchannel:
				if message {
					wmiCallbackInstance.Close()
					// control message came in to stop this go routine
					return
				}
			case wmiEventMessage, ok := <-n.wmiEventTriggered:
				if !ok {
					wmiCallbackInstance.Close()
					// Log the error and close this
					return
				}

				n.handleNotification(wmiEventMessage)
			}
		}
	}()

	return nil
}

// handleNotification
func (n *WmiNotificationMonitor) handleNotification(wmiEventMessage WmiEventMessage) {
	ctx := context.Background()

	notifCallbackList, ok := n.callbacks[wmiEventMessage.AsFilter()]
	if !ok {
		return
	}

	for _, notifCallback := range notifCallbackList {
		go func() {
			notifCallback.lock.Lock()
			defer func() {
				notifCallback.lock.Unlock()
			}()

			err := notifCallback.callback(ctx, wmiEventMessage, notifCallback.callbackContext)
			if err != nil {
				// TODO: log something here
			}
		}()
	}

	return
}

func onObjectReady(ctx interface{}, wmiInstances []*cim.WmiInstance) {
	if len(wmiInstances) < 1 {
		return
	}

	t := ctx.(*WmiNotificationMonitor)

	wmiEventMessage, err := t.wmiEventFactory.NewWmiEventMessage(wmiInstances)
	if err != nil {
		return
	}

	if wmiEventMessage != nil {
		select {
		case t.wmiEventTriggered <- wmiEventMessage:
			break
		default:
			// The queue is full, just drop the last queued event and enqueue the next one
			// IOW, we are only interested in keeping the last relevant node
			_ = <-t.wmiEventTriggered
			t.wmiEventTriggered <- wmiEventMessage
			break
		}
	}
}

func onCompleted(ctx interface{}, wmiInstances []*cim.WmiInstance) {
	// Not used
}

func onProgress(ctx interface{}, wmiInstances []*cim.WmiInstance) {
	// Not used
}

func onObjectPut(ctx interface{}, wmiInstances []*cim.WmiInstance) {
	// Not used
}

func (n *WmiNotificationMonitor) registerWMICallback(wmiNamespace string, wmiObjectName string, wmiHostDomainName string, queryString string) (*WmiNotificationMonitorInstance, error) {
	wmiSession, err := session.GetSession(wmiNamespace, wmiObjectName, wmiHostDomainName, "", "")
	if err != nil {
		return nil, err
	}

	eventSink, err := cim.CreateWmiEventSink(wmiSession, n, onObjectReady, onCompleted, onProgress, onObjectPut)
	if err != nil {
		return nil, err
	}

	_, err = eventSink.Connect()
	if err != nil {
		return nil, err
	}

	_, err = wmiSession.ExecNotificationQueryAsync(eventSink, queryString)
	if err != nil {
		return nil, err
	}

	return &WmiNotificationMonitorInstance{
		session:   wmiSession,
		eventSink: eventSink,
	}, nil
}
