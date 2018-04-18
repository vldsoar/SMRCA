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

//export callbackControllere225c2_Constructor
func callbackControllere225c2_Constructor(ptr unsafe.Pointer) {
	this := NewControllerFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackControllere225c2_UpdateCounter
func callbackControllere225c2_UpdateCounter(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateCounter"); signal != nil {
		signal.(func(string))(cGoUnpackString(data))
	}

}

func (ptr *Controller) ConnectUpdateCounter(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateCounter") {
			C.Controllere225c2_ConnectUpdateCounter(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateCounter"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateCounter", func(data string) {
				signal.(func(string))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateCounter", f)
		}
	}
}

func (ptr *Controller) DisconnectUpdateCounter() {
	if ptr.Pointer() != nil {
		C.Controllere225c2_DisconnectUpdateCounter(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateCounter")
	}
}

func (ptr *Controller) UpdateCounter(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.Controllere225c2_UpdateCounter(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackControllere225c2_UpdateFlowRate
func callbackControllere225c2_UpdateFlowRate(ptr unsafe.Pointer, data C.int) {
	if signal := qt.GetSignal(ptr, "updateFlowRate"); signal != nil {
		signal.(func(int))(int(int32(data)))
	}

}

func (ptr *Controller) ConnectUpdateFlowRate(f func(data int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateFlowRate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "updateFlowRate", func(data int) {
				signal.(func(int))(data)
				f(data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateFlowRate", f)
		}
	}
}

func (ptr *Controller) DisconnectUpdateFlowRate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateFlowRate")
	}
}

func (ptr *Controller) UpdateFlowRate(data int) {
	if ptr.Pointer() != nil {
		C.Controllere225c2_UpdateFlowRate(ptr.Pointer(), C.int(int32(data)))
	}
}

func Controller_QRegisterMetaType() int {
	return int(int32(C.Controllere225c2_Controllere225c2_QRegisterMetaType()))
}

func (ptr *Controller) QRegisterMetaType() int {
	return int(int32(C.Controllere225c2_Controllere225c2_QRegisterMetaType()))
}

func Controller_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Controllere225c2_Controllere225c2_QRegisterMetaType2(typeNameC)))
}

func (ptr *Controller) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Controllere225c2_Controllere225c2_QRegisterMetaType2(typeNameC)))
}

func Controller_QmlRegisterType() int {
	return int(int32(C.Controllere225c2_Controllere225c2_QmlRegisterType()))
}

func (ptr *Controller) QmlRegisterType() int {
	return int(int32(C.Controllere225c2_Controllere225c2_QmlRegisterType()))
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
	return int(int32(C.Controllere225c2_Controllere225c2_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.Controllere225c2_Controllere225c2_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Controller) __children_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controllere225c2___children_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __children_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controllere225c2___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __children_newList() unsafe.Pointer {
	return C.Controllere225c2___children_newList(ptr.Pointer())
}

func (ptr *Controller) __dynamicPropertyNames_atList(i int, p unsafe.Pointer) *std_core.QByteArray {
	tmpValue := std_core.NewQByteArrayFromPointer(C.Controllere225c2___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i)), p))
	runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
	return tmpValue
}

func (ptr *Controller) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF, p unsafe.Pointer) {
	C.Controllere225c2___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i), p)
}

func (ptr *Controller) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.Controllere225c2___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *Controller) __findChildren_atList2(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controllere225c2___findChildren_atList2(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __findChildren_setList2(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controllere225c2___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __findChildren_newList2() unsafe.Pointer {
	return C.Controllere225c2___findChildren_newList2(ptr.Pointer())
}

func (ptr *Controller) __findChildren_atList3(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controllere225c2___findChildren_atList3(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __findChildren_setList3(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controllere225c2___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __findChildren_newList3() unsafe.Pointer {
	return C.Controllere225c2___findChildren_newList3(ptr.Pointer())
}

func (ptr *Controller) __findChildren_atList(i int, p unsafe.Pointer) *std_core.QObject {
	tmpValue := std_core.NewQObjectFromPointer(C.Controllere225c2___findChildren_atList(ptr.Pointer(), C.int(int32(i)), p))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Controller) __findChildren_setList(i std_core.QObject_ITF, p unsafe.Pointer) {
	C.Controllere225c2___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i), p)
}

func (ptr *Controller) __findChildren_newList() unsafe.Pointer {
	return C.Controllere225c2___findChildren_newList(ptr.Pointer())
}

func NewController(parent std_core.QObject_ITF) *Controller {
	tmpValue := NewControllerFromPointer(C.Controllere225c2_NewController(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackControllere225c2_DestroyController
func callbackControllere225c2_DestroyController(ptr unsafe.Pointer) {
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
		C.Controllere225c2_DestroyController(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Controller) DestroyControllerDefault() {
	if ptr.Pointer() != nil {
		C.Controllere225c2_DestroyControllerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackControllere225c2_TimerEvent
func callbackControllere225c2_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewControllerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Controller) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Controllere225c2_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackControllere225c2_ChildEvent
func callbackControllere225c2_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewControllerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Controller) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Controllere225c2_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackControllere225c2_ConnectNotify
func callbackControllere225c2_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewControllerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Controller) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Controllere225c2_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackControllere225c2_CustomEvent
func callbackControllere225c2_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewControllerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Controller) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Controllere225c2_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackControllere225c2_DeleteLater
func callbackControllere225c2_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Controller) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Controllere225c2_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackControllere225c2_Destroyed
func callbackControllere225c2_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackControllere225c2_DisconnectNotify
func callbackControllere225c2_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewControllerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Controller) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Controllere225c2_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackControllere225c2_Event
func callbackControllere225c2_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewControllerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Controller) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Controllere225c2_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackControllere225c2_EventFilter
func callbackControllere225c2_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewControllerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Controller) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Controllere225c2_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackControllere225c2_ObjectNameChanged
func callbackControllere225c2_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}
