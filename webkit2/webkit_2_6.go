package webkit

// #cgo pkg-config: webkit2gtk-4.0
// #include <webkit2/webkit2.h>
// #include "webkit_2_6.go.h"
import "C"

import (
	//"errors"
	"runtime"
	"unsafe"

	//"github.com/andre-hub/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.webkit_user_style_level_get_type()), marshalUserStyleLevel},

		// Objects/Interfaces
		{glib.Type(C.webkit_user_content_manager_get_type()), marshalUserContentManager},

		// Boxed
		{glib.Type(C.webkit_user_script_get_type()), marshalUserScript},
		{glib.Type(C.webkit_user_style_sheet_get_type()), marshalUserStyleSheet},
	}
	glib.RegisterGValueMarshalers(tm)
}

type UserStyleLevel int

const (
	USER_STYLE_LEVEL_USER   UserStyleLevel = C.WEBKIT_USER_STYLE_LEVEL_USER
	USER_STYLE_LEVEL_AUTHOR UserStyleLevel = C.WEBKIT_USER_STYLE_LEVEL_AUTHOR
)

func marshalUserStyleLevel(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return UserStyleLevel(c), nil
}

/*
 * WebKitUserContentManager
 */

type UserContentManager struct {
	*glib.Object
}

// native returns a pointer to the underlying WebKitUserContentManager.
func (v *UserContentManager) native() *C.WebKitUserContentManager {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toWebKitUserContentManager(p)
}

func marshalUserContentManager(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	return wrapUserContentManager(obj), nil
}

func wrapUserContentManager(obj *glib.Object) *UserContentManager {
	return &UserContentManager{obj}
}

// UserContentManagerNew() is a wrapper around webkit_user_content_manager_new().
func UserContentManagerNew() (*UserContentManager, error) {
	c := C.webkit_user_content_manager_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapUserContentManager(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

func (v *UserContentManager) AddStyleSheet(sheet *UserStyleSheet) {
	C.webkit_user_content_manager_add_style_sheet(v.native(), sheet.native())
}

func (v *UserContentManager) RemoveAllStyleSheets() {
	C.webkit_user_content_manager_remove_all_style_sheets(v.native())
}

/*
 * WebKitWebView
 */

// WebViewNewWithSettings() is a wrapper around webkit_web_view_new_with_settings().
func WebViewNewWithSettings(settings *Settings) (*WebView, error) {
	c := C.webkit_web_view_new_with_settings(settings.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// WebViewNewWithUserContentManager() is a wrapper around webkit_web_view_new_with_user_content_manager().
func WebViewNewWithUserContentManager(manager *UserContentManager) (*WebView, error) {
	c := C.webkit_web_view_new_with_user_content_manager(manager.native())
	if c == nil {
		return nil, nilPtrErr
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	b := wrapWebView(obj)
	b.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return b, nil
}

// LoadBytes() is a wrapper around webkit_web_view_load_bytes().
func (v *WebView) LoadBytes(bytes []byte, mimeType, encoding, baseUri string) {
	gbytes := C.g_bytes_new(C.gconstpointer(&bytes[0]), C.gsize(len(bytes)))
	defer C.g_bytes_unref(gbytes)
	cmimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cmimeType))
	cencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cencoding))
	cbaseUri := C.CString(baseUri)
	defer C.free(unsafe.Pointer(cbaseUri))
	C.webkit_web_view_load_bytes(v.native(), gbytes, (*C.gchar)(cmimeType), (*C.gchar)(cencoding), (*C.gchar)(cbaseUri))
}

// GetUserContentManager() is a wrapper around webkit_web_view_get_user_content_manager().
func (v *WebView) GetUserContentManager() *UserContentManager {
	c := C.webkit_web_view_get_user_content_manager(v.native())
	if c == nil {
		return nil
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	w := wrapUserContentManager(obj)
	obj.RefSink()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)
	return w
}

/*
 * WebKitUserScript
 */

type UserScript struct {
	WebKitUserScript *C.WebKitUserScript
}

// native returns a pointer to the underlying WebKitUserScript.
func (v *UserScript) native() *C.WebKitUserScript {
	if v == nil {
		return nil
	}
	return v.WebKitUserScript
}

func marshalUserScript(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*UserScript)(unsafe.Pointer(c)), nil
}

func (v *UserScript) Unref() {
	C.webkit_user_script_unref(v.native())
}

func (v *UserScript) Ref() {
	C.webkit_user_script_ref(v.native())
}

/*
 * WebKitUserStyleSheet
 */

type UserStyleSheet struct {
	WebKitUserStyleSheet *C.WebKitUserStyleSheet
}

// native returns a pointer to the underlying WebKitUserStyleSheet.
func (v *UserStyleSheet) native() *C.WebKitUserStyleSheet {
	if v == nil {
		return nil
	}
	return v.WebKitUserStyleSheet
}

func marshalUserStyleSheet(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*UserStyleSheet)(unsafe.Pointer(c)), nil
}

func (v *UserStyleSheet) Unref() {
	C.webkit_user_style_sheet_unref(v.native())
}

func (v *UserStyleSheet) Ref() {
	C.webkit_user_style_sheet_ref(v.native())
}
