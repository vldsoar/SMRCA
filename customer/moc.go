package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type Client_ITF interface {
	std_core.QObject_ITF
	Client_PTR() *Client
}

func (ptr *Client) Client_PTR() *Client {
	return ptr
}

func (ptr *Client) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Client) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromClient(ptr Client_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Client_PTR().Pointer()
	}
	return nil
}

func NewClientFromPointer(ptr unsafe.Pointer) (n *Client) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Client)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Client:
			n = deduced

		case *std_core.QObject:
			n = &Client{QObject: *deduced}

		default:
			n = new(Client)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackClient6356b5_Constructor
func callbackClient6356b5_Constructor(ptr unsafe.Pointer) {
	this := NewClientFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackClient6356b5_State
func callbackClient6356b5_State(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "state"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewClientFromPointer(ptr).StateDefault()))
}

func (ptr *Client) ConnectState(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "state"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "state", func() int {
				signal.(func() int)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "state", f)
		}
	}
}

func (ptr *Client) DisconnectState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "state")
	}
}

func (ptr *Client) State() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Client6356b5_State(ptr.Pointer())))
	}
	return 0
}

func (ptr *Client) StateDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Client6356b5_StateDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackClient6356b5_SetState
func callbackClient6356b5_SetState(ptr unsafe.Pointer, state C.int) {
	if signal := qt.GetSignal(ptr, "setState"); signal != nil {
		signal.(func(int))(int(int32(state)))
	} else {
		NewClientFromPointer(ptr).SetStateDefault(int(int32(state)))
	}
}

func (ptr *Client) ConnectSetState(f func(state int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setState"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setState", func(state int) {
				signal.(func(int))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setState", f)
		}
	}
}

func (ptr *Client) DisconnectSetState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setState")
	}
}

func (ptr *Client) SetState(state int) {
	if ptr.Pointer() != nil {
		C.Client6356b5_SetState(ptr.Pointer(), C.int(int32(state)))
	}
}

func (ptr *Client) SetStateDefault(state int) {
	if ptr.Pointer() != nil {
		C.Client6356b5_SetStateDefault(ptr.Pointer(), C.int(int32(state)))
	}
}

//export callbackClient6356b5_StateChanged
func callbackClient6356b5_StateChanged(ptr unsafe.Pointer, state C.int) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(int))(int(int32(state)))
	}

}

func (ptr *Client) ConnectStateChanged(f func(state int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.Client6356b5_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state int) {
				signal.(func(int))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *Client) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.Client6356b5_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *Client) StateChanged(state int) {
	if ptr.Pointer() != nil {
		C.Client6356b5_StateChanged(ptr.Pointer(), C.int(int32(state)))
	}
}

//export callbackClient6356b5_Data
func callbackClient6356b5_Data(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQJsonObject(signal.(func() *std_core.QJsonObject)())
	}

	return std_core.PointerFromQJsonObject(NewClientFromPointer(ptr).DataDefault())
}

func (ptr *Client) ConnectData(f func() *std_core.QJsonObject) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "data"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "data", func() *std_core.QJsonObject {
				signal.(func() *std_core.QJsonObject)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "data", f)
		}
	}
}

func (ptr *Client) DisconnectData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "data")
	}
}

func (ptr *Client) Data() *std_core.QJsonObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQJsonObjectFromPointer(C.Client6356b5_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QJsonObject).DestroyQJsonObject)
		return tmpValue
	}
	return nil
}

func (ptr *Client) DataDefault() *std_core.QJsonObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQJsonObjectFromPointer(C.Client6356b5_DataDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QJsonObject).DestroyQJsonObject)
		return tmpValue
	}
	return nil
}

//export callbackClient6356b5_SetData
func callbackClient6356b5_SetData(ptr unsafe.Pointer, data unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		signal.(func(*std_core.QJsonObject))(std_core.NewQJsonObjectFromPointer(data))
	} else {
		NewClientFromPointer(ptr).SetDataDefault(std_core.NewQJsonObjectFromPointer(data))
	}
}

func (ptr *Client) ConnectSetData(f func(data *std_core.QJsonObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setData", func(data *std_core.QJsonObject) {
				signal.(func(*std_core.QJsonObject))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setData", f)
		}
	}
}

func (ptr *Client) DisconnectSetData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setData")
	}
}

func (ptr *Client) SetData(data std_core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_SetData(ptr.Pointer(), std_core.PointerFromQJsonObject(data))
	}
}

func (ptr *Client) SetDataDefault(data std_core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_SetDataDefault(ptr.Pointer(), std_core.PointerFromQJsonObject(data))
	}
}

//export callbackClient6356b5_DataChanged
func callbackClient6356b5_DataChanged(ptr unsafe.Pointer, data unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		signal.(func(*std_core.QJsonObject))(std_core.NewQJsonObjectFromPointer(data))
	}

}

func (ptr *Client) ConnectDataChanged(f func(data *std_core.QJsonObject)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataChanged") {
			C.Client6356b5_ConnectDataChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataChanged", func(data *std_core.QJsonObject) {
				signal.(func(*std_core.QJsonObject))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataChanged", f)
		}
	}
}

func (ptr *Client) DisconnectDataChanged() {
	if ptr.Pointer() != nil {
		C.Client6356b5_DisconnectDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataChanged")
	}
}

func (ptr *Client) DataChanged(data std_core.QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_DataChanged(ptr.Pointer(), std_core.PointerFromQJsonObject(data))
	}
}

func Client_QRegisterMetaType() int {
	return int(int32(C.Client6356b5_Client6356b5_QRegisterMetaType()))
}

func (ptr *Client) QRegisterMetaType() int {
	return int(int32(C.Client6356b5_Client6356b5_QRegisterMetaType()))
}

func Client_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Client6356b5_Client6356b5_QRegisterMetaType2(typeNameC)))
}

func (ptr *Client) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Client6356b5_Client6356b5_QRegisterMetaType2(typeNameC)))
}

func Client_QmlRegisterType() int {
	return int(int32(C.Client6356b5_Client6356b5_QmlRegisterType()))
}

func (ptr *Client) QmlRegisterType() int {
	return int(int32(C.Client6356b5_Client6356b5_QmlRegisterType()))
}

func Client_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Client6356b5_Client6356b5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Client) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Client6356b5_Client6356b5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Client) __children_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Client6356b5___children_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Client) __children_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Client6356b5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Client) __children_newList() unsafe.Pointer {
	return C.Client6356b5___children_newList(ptr.Pointer())
}

func (ptr *Client) __dynamicPropertyNames_atList(i int, p unsafe.Pointer) *std_core.QByteArray {
	tmpValue := std_core.NewQByteArrayFromPointer(C.Client6356b5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i)), p))
	runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *Client) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF, p unsafe.Pointer) {
	C.Client6356b5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i), p)
}

func (ptr *Client) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Client6356b5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Client) __findChildren_atList2(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Client6356b5___findChildren_atList2(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Client) __findChildren_setList2(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Client6356b5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Client) __findChildren_newList2() unsafe.Pointer {
	return C.Client6356b5___findChildren_newList2(ptr.Pointer())
}

func (ptr *Client) __findChildren_atList3(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Client6356b5___findChildren_atList3(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Client) __findChildren_setList3(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Client6356b5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Client) __findChildren_newList3() unsafe.Pointer {
	return C.Client6356b5___findChildren_newList3(ptr.Pointer())
}

func (ptr *Client) __findChildren_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Client6356b5___findChildren_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Client) __findChildren_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Client6356b5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Client) __findChildren_newList() unsafe.Pointer {
	return C.Client6356b5___findChildren_newList(ptr.Pointer())
}

func NewClient(parent std_core.QObject_ITF) *Client {
	tmpValue := NewClientFromPointer(C.Client6356b5_NewClient(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackClient6356b5_DestroyClient
func callbackClient6356b5_DestroyClient(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Client"); signal != nil {
		signal.(func())()
	} else {
		NewClientFromPointer(ptr).DestroyClientDefault()
	}
}

func (ptr *Client) ConnectDestroyClient(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Client"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Client", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Client", f)
		}
	}
}

func (ptr *Client) DisconnectDestroyClient() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Client")
	}
}

func (ptr *Client) DestroyClient() {
	if ptr.Pointer() != nil {
		C.Client6356b5_DestroyClient(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Client) DestroyClientDefault() {
	if ptr.Pointer() != nil {
		C.Client6356b5_DestroyClientDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackClient6356b5_TimerEvent
func callbackClient6356b5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewClientFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Client) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackClient6356b5_ChildEvent
func callbackClient6356b5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewClientFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Client) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackClient6356b5_ConnectNotify
func callbackClient6356b5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewClientFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Client) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackClient6356b5_CustomEvent
func callbackClient6356b5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewClientFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Client) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackClient6356b5_DeleteLater
func callbackClient6356b5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewClientFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Client) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Client6356b5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackClient6356b5_Destroyed
func callbackClient6356b5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackClient6356b5_DisconnectNotify
func callbackClient6356b5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewClientFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Client) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Client6356b5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackClient6356b5_Event
func callbackClient6356b5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewClientFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Client) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Client6356b5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackClient6356b5_EventFilter
func callbackClient6356b5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewClientFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Client) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Client6356b5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackClient6356b5_ObjectNameChanged
func callbackClient6356b5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

type Controller_ITF interface {
	std_core.QObject_ITF
	Controller_PTR() *Controller
}

func (ptr *Controller) Controller_PTR() *Controller {
	return ptr
}

func (ptr *Controller) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Controller) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromController(ptr Controller_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Controller_PTR().Pointer()
	}
	return nil
}

func NewControllerFromPointer(ptr unsafe.Pointer) (n *Controller) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Controller)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Controller:
			n = deduced

		case *std_core.QObject:
			n = &Controller{QObject: *deduced}

		default:
			n = new(Controller)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackController6356b5_Constructor
func callbackController6356b5_Constructor(ptr unsafe.Pointer) {
	this := NewControllerFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackController6356b5_Login
func callbackController6356b5_Login(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "login"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *Controller) ConnectLogin(f func(data string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "login"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "login", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "login", f)
		}
	}
}

func (ptr *Controller) DisconnectLogin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "login")
	}
}

func (ptr *Controller) Login(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.Controller6356b5_Login(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackController6356b5_Logout
func callbackController6356b5_Logout(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "logout"); signal != nil {
		signal.(func())()
	}

}

func (ptr *Controller) ConnectLogout(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "logout"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "logout", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "logout", f)
		}
	}
}

func (ptr *Controller) DisconnectLogout() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "logout")
	}
}

func (ptr *Controller) Logout() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_Logout(ptr.Pointer())
	}
}

//export callbackController6356b5_GetConsumptionValue
func callbackController6356b5_GetConsumptionValue(ptr unsafe.Pointer, userID C.int) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getConsumptionValue"); signal != nil {
		tempVal := signal.(func(int) string)(int(int32(userID)))
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *Controller) ConnectGetConsumptionValue(f func(userID int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getConsumptionValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getConsumptionValue", func(userID int) string {
				signal.(func(int) string)(userID)
				return f(userID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getConsumptionValue", f)
		}
	}
}

func (ptr *Controller) DisconnectGetConsumptionValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getConsumptionValue")
	}
}

func (ptr *Controller) GetConsumptionValue(userID int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.Controller6356b5_GetConsumptionValue(ptr.Pointer(), C.int(int32(userID))))
	}
	return ""
}

//export callbackController6356b5_GetConsumptions
func callbackController6356b5_GetConsumptions(ptr unsafe.Pointer, userID C.int, start C.struct_Moc_PackedString, end C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "getConsumptions"); signal != nil {
		signal.(func(int, string, string))(int(int32(userID)), cGoUnpackString(start), cGoUnpackString(end))
	}

}

func (ptr *Controller) ConnectGetConsumptions(f func(userID int, start string, end string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getConsumptions"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getConsumptions", func(userID int, start string, end string) {
				signal.(func(int, string, string))(userID, start, end)
				f(userID, start, end)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getConsumptions", f)
		}
	}
}

func (ptr *Controller) DisconnectGetConsumptions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getConsumptions")
	}
}

func (ptr *Controller) GetConsumptions(userID int, start string, end string) {
	if ptr.Pointer() != nil {
		var startC *C.char
		if start != "" {
			startC = C.CString(start)
			defer C.free(unsafe.Pointer(startC))
		}
		var endC *C.char
		if end != "" {
			endC = C.CString(end)
			defer C.free(unsafe.Pointer(endC))
		}
		C.Controller6356b5_GetConsumptions(ptr.Pointer(), C.int(int32(userID)), C.struct_Moc_PackedString{data: startC, len: C.longlong(len(start))}, C.struct_Moc_PackedString{data: endC, len: C.longlong(len(end))})
	}
}

//export callbackController6356b5_GetConsumptionsZone
func callbackController6356b5_GetConsumptionsZone(ptr unsafe.Pointer, zone C.int, start C.struct_Moc_PackedString, end C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "getConsumptionsZone"); signal != nil {
		signal.(func(int, string, string))(int(int32(zone)), cGoUnpackString(start), cGoUnpackString(end))
	}

}

func (ptr *Controller) ConnectGetConsumptionsZone(f func(zone int, start string, end string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getConsumptionsZone"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getConsumptionsZone", func(zone int, start string, end string) {
				signal.(func(int, string, string))(zone, start, end)
				f(zone, start, end)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getConsumptionsZone", f)
		}
	}
}

func (ptr *Controller) DisconnectGetConsumptionsZone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getConsumptionsZone")
	}
}

func (ptr *Controller) GetConsumptionsZone(zone int, start string, end string) {
	if ptr.Pointer() != nil {
		var startC *C.char
		if start != "" {
			startC = C.CString(start)
			defer C.free(unsafe.Pointer(startC))
		}
		var endC *C.char
		if end != "" {
			endC = C.CString(end)
			defer C.free(unsafe.Pointer(endC))
		}
		C.Controller6356b5_GetConsumptionsZone(ptr.Pointer(), C.int(int32(zone)), C.struct_Moc_PackedString{data: startC, len: C.longlong(len(start))}, C.struct_Moc_PackedString{data: endC, len: C.longlong(len(end))})
	}
}

//export callbackController6356b5_SessionAuthenticated
func callbackController6356b5_SessionAuthenticated(ptr unsafe.Pointer, reply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sessionAuthenticated"); signal != nil {
		signal.(func(string))(cGoUnpackString(reply))
	}

}

func (ptr *Controller) ConnectSessionAuthenticated(f func(reply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sessionAuthenticated") {
			C.Controller6356b5_ConnectSessionAuthenticated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sessionAuthenticated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticated", func(reply string) {
				signal.(func(string))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticated", f)
		}
	}
}

func (ptr *Controller) DisconnectSessionAuthenticated() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DisconnectSessionAuthenticated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sessionAuthenticated")
	}
}

func (ptr *Controller) SessionAuthenticated(reply string) {
	if ptr.Pointer() != nil {
		var replyC *C.char
		if reply != "" {
			replyC = C.CString(reply)
			defer C.free(unsafe.Pointer(replyC))
		}
		C.Controller6356b5_SessionAuthenticated(ptr.Pointer(), C.struct_Moc_PackedString{data: replyC, len: C.longlong(len(reply))})
	}
}

//export callbackController6356b5_SessionAuthenticationError
func callbackController6356b5_SessionAuthenticationError(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sessionAuthenticationError"); signal != nil {
		signal.(func())()
	}

}

func (ptr *Controller) ConnectSessionAuthenticationError(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sessionAuthenticationError") {
			C.Controller6356b5_ConnectSessionAuthenticationError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sessionAuthenticationError"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticationError", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sessionAuthenticationError", f)
		}
	}
}

func (ptr *Controller) DisconnectSessionAuthenticationError() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DisconnectSessionAuthenticationError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sessionAuthenticationError")
	}
}

func (ptr *Controller) SessionAuthenticationError() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_SessionAuthenticationError(ptr.Pointer())
	}
}

//export callbackController6356b5_SessionTerminated
func callbackController6356b5_SessionTerminated(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sessionTerminated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *Controller) ConnectSessionTerminated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sessionTerminated") {
			C.Controller6356b5_ConnectSessionTerminated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sessionTerminated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sessionTerminated", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sessionTerminated", f)
		}
	}
}

func (ptr *Controller) DisconnectSessionTerminated() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DisconnectSessionTerminated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sessionTerminated")
	}
}

func (ptr *Controller) SessionTerminated() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_SessionTerminated(ptr.Pointer())
	}
}

//export callbackController6356b5_SessionLoader
func callbackController6356b5_SessionLoader(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "sessionLoader"); signal != nil {
		signal.(func())()
	}

}

func (ptr *Controller) ConnectSessionLoader(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sessionLoader") {
			C.Controller6356b5_ConnectSessionLoader(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sessionLoader"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "sessionLoader", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sessionLoader", f)
		}
	}
}

func (ptr *Controller) DisconnectSessionLoader() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DisconnectSessionLoader(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sessionLoader")
	}
}

func (ptr *Controller) SessionLoader() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_SessionLoader(ptr.Pointer())
	}
}

//export callbackController6356b5_GetConsumption
func callbackController6356b5_GetConsumption(ptr unsafe.Pointer, consumptions C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "getConsumption"); signal != nil {
		signal.(func(string))(cGoUnpackString(consumptions))
	}

}

func (ptr *Controller) ConnectGetConsumption(f func(consumptions string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "getConsumption") {
			C.Controller6356b5_ConnectGetConsumption(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "getConsumption"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "getConsumption", func(consumptions string) {
				signal.(func(string))(consumptions)
				f(consumptions)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getConsumption", f)
		}
	}
}

func (ptr *Controller) DisconnectGetConsumption() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DisconnectGetConsumption(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "getConsumption")
	}
}

func (ptr *Controller) GetConsumption(consumptions string) {
	if ptr.Pointer() != nil {
		var consumptionsC *C.char
		if consumptions != "" {
			consumptionsC = C.CString(consumptions)
			defer C.free(unsafe.Pointer(consumptionsC))
		}
		C.Controller6356b5_GetConsumption(ptr.Pointer(), C.struct_Moc_PackedString{data: consumptionsC, len: C.longlong(len(consumptions))})
	}
}

func Controller_QRegisterMetaType() int {
	return int(int32(C.Controller6356b5_Controller6356b5_QRegisterMetaType()))
}

func (ptr *Controller) QRegisterMetaType() int {
	return int(int32(C.Controller6356b5_Controller6356b5_QRegisterMetaType()))
}

func Controller_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Controller6356b5_Controller6356b5_QRegisterMetaType2(typeNameC)))
}

func (ptr *Controller) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Controller6356b5_Controller6356b5_QRegisterMetaType2(typeNameC)))
}

func Controller_QmlRegisterType() int {
	return int(int32(C.Controller6356b5_Controller6356b5_QmlRegisterType()))
}

func (ptr *Controller) QmlRegisterType() int {
	return int(int32(C.Controller6356b5_Controller6356b5_QmlRegisterType()))
}

func Controller_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Controller6356b5_Controller6356b5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Controller) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Controller6356b5_Controller6356b5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Controller) __children_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controller6356b5___children_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __children_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controller6356b5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __children_newList() unsafe.Pointer {
	return C.Controller6356b5___children_newList(ptr.Pointer())
}

func (ptr *Controller) __dynamicPropertyNames_atList(i int, p unsafe.Pointer) *std_core.QByteArray {
	tmpValue := std_core.NewQByteArrayFromPointer(C.Controller6356b5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i)), p))
	runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *Controller) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF, p unsafe.Pointer) {
	C.Controller6356b5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i), p)
}

func (ptr *Controller) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Controller6356b5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Controller) __findChildren_atList2(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controller6356b5___findChildren_atList2(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __findChildren_setList2(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controller6356b5___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __findChildren_newList2() unsafe.Pointer {
	return C.Controller6356b5___findChildren_newList2(ptr.Pointer())
}

func (ptr *Controller) __findChildren_atList3(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controller6356b5___findChildren_atList3(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __findChildren_setList3(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controller6356b5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __findChildren_newList3() unsafe.Pointer {
	return C.Controller6356b5___findChildren_newList3(ptr.Pointer())
}

func (ptr *Controller) __findChildren_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controller6356b5___findChildren_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __findChildren_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controller6356b5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __findChildren_newList() unsafe.Pointer {
	return C.Controller6356b5___findChildren_newList(ptr.Pointer())
}

func NewController(parent std_core.QObject_ITF) *Controller {
	tmpValue := NewControllerFromPointer(C.Controller6356b5_NewController(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackController6356b5_DestroyController
func callbackController6356b5_DestroyController(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Controller"); signal != nil {
		signal.(func())()
	} else {
		NewControllerFromPointer(ptr).DestroyControllerDefault()
	}
}

func (ptr *Controller) ConnectDestroyController(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Controller"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Controller", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Controller", f)
		}
	}
}

func (ptr *Controller) DisconnectDestroyController() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Controller")
	}
}

func (ptr *Controller) DestroyController() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DestroyController(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Controller) DestroyControllerDefault() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DestroyControllerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackController6356b5_TimerEvent
func callbackController6356b5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewControllerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Controller) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Controller6356b5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackController6356b5_ChildEvent
func callbackController6356b5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewControllerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Controller) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Controller6356b5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackController6356b5_ConnectNotify
func callbackController6356b5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewControllerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Controller) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Controller6356b5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackController6356b5_CustomEvent
func callbackController6356b5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewControllerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Controller) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Controller6356b5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackController6356b5_DeleteLater
func callbackController6356b5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Controller) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackController6356b5_Destroyed
func callbackController6356b5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackController6356b5_DisconnectNotify
func callbackController6356b5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewControllerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Controller) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Controller6356b5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackController6356b5_Event
func callbackController6356b5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewControllerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Controller) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Controller6356b5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackController6356b5_EventFilter
func callbackController6356b5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewControllerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Controller) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Controller6356b5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackController6356b5_ObjectNameChanged
func callbackController6356b5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}
