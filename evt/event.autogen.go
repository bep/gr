//go:generate go run generate.go

// Package evt defines markup to bind DOM events.
//
// Generated from "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.
package evt

import "github.com/bep/gr"

// Abort gets notified when a transaction has been aborted.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
func Abort(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAbort", listener)
}

// AfterPrint gets notified when the associated document has started printing or the print preview has been closed.
//
// https://developer.mozilla.org/docs/Web/Events/afterprint
func AfterPrint(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAfterPrint", listener)
}

// AnimationEnd gets notified when a CSS animation has completed.
//
// https://developer.mozilla.org/docs/Web/Events/animationend
func AnimationEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAnimationEnd", listener)
}

// AnimationIteration gets notified when a CSS animation is repeated.
//
// https://developer.mozilla.org/docs/Web/Events/animationiteration
func AnimationIteration(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAnimationIteration", listener)
}

// AnimationStart gets notified when a CSS animation has started.
//
// https://developer.mozilla.org/docs/Web/Events/animationstart
func AnimationStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAnimationStart", listener)
}

// AudioEnd gets notified when the user agent has finished capturing audio for speech recognition.
//
// https://developer.mozilla.org/docs/Web/Events/audioend
func AudioEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAudioEnd", listener)
}

// AudioProcess gets notified when the input buffer of a ScriptProcessorNode is ready to be processed.
//
// https://developer.mozilla.org/docs/Web/Events/audioprocess
func AudioProcess(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAudioProcess", listener)
}

// AudioStart gets notified when the user agent has started to capture audio for speech recognition.
//
// https://developer.mozilla.org/docs/Web/Events/audiostart
func AudioStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAudioStart", listener)
}

// BeforePrint gets notified when the associated document is about to be printed or previewed for printing.
//
// https://developer.mozilla.org/docs/Web/Events/beforeprint
func BeforePrint(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBeforePrint", listener)
}

// BeforeUnload gets notified when (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/beforeunload
func BeforeUnload(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBeforeUnload", listener)
}

// BeginEvent gets notified when a SMIL animation element begins.
//
// https://developer.mozilla.org/docs/Web/Events/beginEvent
func BeginEvent(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBeginEvent", listener)
}

// Blocked gets notified when an open connection to a database is blocking a versionchange transaction on the same database.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
func Blocked(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBlocked", listener)
}

// Blur gets notified when an element has lost focus (does not bubble).
//
// https://developer.mozilla.org/docs/Web/Events/blur
func Blur(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBlur", listener)
}

// Boundary gets notified when the spoken utterance reaches a word or sentence boundary
//
// https://developer.mozilla.org/docs/Web/Events/boundary
func Boundary(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBoundary", listener)
}

// Cached gets notified when the resources listed in the manifest have been downloaded, and the application is now cached.
//
// https://developer.mozilla.org/docs/Web/Events/cached
func Cached(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCached", listener)
}

// CanPlay gets notified when the user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
//
// https://developer.mozilla.org/docs/Web/Events/canplay
func CanPlay(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCanPlay", listener)
}

// CanPlayThrough gets notified when the user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
//
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
func CanPlayThrough(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCanPlayThrough", listener)
}

// Change gets notified when the change event is fired for <input>, <select>, and <textarea> elements when a change to the element's value is committed by the user.
//
// https://developer.mozilla.org/docs/Web/Events/change
func Change(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChange", listener)
}

// ChargingChange gets notified when the battery begins or stops charging.
//
// https://developer.mozilla.org/docs/Web/Events/chargingchange
func ChargingChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChargingChange", listener)
}

// ChargingTimeChange gets notified when the chargingTime attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
func ChargingTimeChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChargingTimeChange", listener)
}

// Checking gets notified when the user agent is checking for an update, or attempting to download the cache manifest for the first time.
//
// https://developer.mozilla.org/docs/Web/Events/checking
func Checking(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChecking", listener)
}

// Click gets notified when a pointing device button has been pressed and released on an element.
//
// https://developer.mozilla.org/docs/Web/Events/click
func Click(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onClick", listener)
}

// Close gets notified when a WebSocket connection has been closed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/close_websocket
func Close(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onClose", listener)
}

// Complete gets notified when the rendering of an OfflineAudioContext is terminated.
//
// https://developer.mozilla.org/docs/Web/Events/complete
func Complete(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onComplete", listener)
}

// CompositionEnd gets notified when the composition of a passage of text has been completed or canceled.
//
// https://developer.mozilla.org/docs/Web/Events/compositionend
func CompositionEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCompositionEnd", listener)
}

// CompositionStart gets notified when the composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition).
//
// https://developer.mozilla.org/docs/Web/Events/compositionstart
func CompositionStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCompositionStart", listener)
}

// CompositionUpdate gets notified when a character is added to a passage of text being composed.
//
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
func CompositionUpdate(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCompositionUpdate", listener)
}

// ContextMenu gets notified when the right button of the mouse is clicked (before the context menu is displayed).
//
// https://developer.mozilla.org/docs/Web/Events/contextmenu
func ContextMenu(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onContextMenu", listener)
}

// Copy gets notified when the text selection has been added to the clipboard.
//
// https://developer.mozilla.org/docs/Web/Events/copy
func Copy(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCopy", listener)
}

// Cut gets notified when the text selection has been removed from the document and added to the clipboard.
//
// https://developer.mozilla.org/docs/Web/Events/cut
func Cut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCut", listener)
}

// DOMContentLoaded gets notified when the document has finished loading (but not its dependent resources).
//
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
func DOMContentLoaded(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDOMContentLoaded", listener)
}

// DeviceLight gets notified when fresh data is available from a light sensor.
//
// https://developer.mozilla.org/docs/Web/Events/devicelight
func DeviceLight(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceLight", listener)
}

// DeviceMotion gets notified when fresh data is available from a motion sensor.
//
// https://developer.mozilla.org/docs/Web/Events/devicemotion
func DeviceMotion(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceMotion", listener)
}

// DeviceOrientation gets notified when fresh data is available from an orientation sensor.
//
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
func DeviceOrientation(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceOrientation", listener)
}

// DeviceProximity gets notified when fresh data is available from a proximity sensor (indicates an approximated distance between the device and a nearby object).
//
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
func DeviceProximity(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceProximity", listener)
}

// DischargingTimeChange gets notified when the dischargingTime attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
func DischargingTimeChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDischargingTimeChange", listener)
}

// DoubleClick gets notified when a pointing device button is clicked twice on an element.
//
// https://developer.mozilla.org/docs/Web/Events/dblclick
func DoubleClick(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDoubleClick", listener)
}

// Downloading gets notified when the user agent has found an update and is fetching it, or is downloading the resources listed by the cache manifest for the first time.
//
// https://developer.mozilla.org/docs/Web/Events/downloading
func Downloading(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDownloading", listener)
}

// Drag gets notified when an element or text selection is being dragged (every 350ms).
//
// https://developer.mozilla.org/docs/Web/Events/drag
func Drag(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDrag", listener)
}

// DragEnd gets notified when a drag operation is being ended (by releasing a mouse button or hitting the escape key).
//
// https://developer.mozilla.org/docs/Web/Events/dragend
func DragEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragEnd", listener)
}

// DragEnter gets notified when a dragged element or text selection enters a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/dragenter
func DragEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragEnter", listener)
}

// DragLeave gets notified when a dragged element or text selection leaves a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/dragleave
func DragLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragLeave", listener)
}

// DragOver gets notified when an element or text selection is being dragged over a valid drop target (every 350ms).
//
// https://developer.mozilla.org/docs/Web/Events/dragover
func DragOver(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragOver", listener)
}

// DragStart gets notified when the user starts dragging an element or text selection.
//
// https://developer.mozilla.org/docs/Web/Events/dragstart
func DragStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragStart", listener)
}

// Drop gets notified when an element is dropped on a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/drop
func Drop(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDrop", listener)
}

// DurationChange gets notified when the duration attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/durationchange
func DurationChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDurationChange", listener)
}

// Emptied gets notified when the media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.
//
// https://developer.mozilla.org/docs/Web/Events/emptied
func Emptied(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEmptied", listener)
}

// End gets notified when the utterance has finished being spoken.
//
// https://developer.mozilla.org/docs/Web/Events/end_(SpeechSynthesis)
func End(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEnd", listener)
}

// EndEvent gets notified when a SMIL animation element ends.
//
// https://developer.mozilla.org/docs/Web/Events/endEvent
func EndEvent(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEndEvent", listener)
}

// Ended gets notified when (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
func Ended(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEnded", listener)
}

// Error gets notified when an error occurs that prevents the utterance from being successfully spoken.
//
// https://developer.mozilla.org/docs/Web/Events/error_(SpeechSynthesisError)
func Error(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onError", listener)
}

// Focus gets notified when an element has received focus (does not bubble).
//
// https://developer.mozilla.org/docs/Web/Events/focus
func Focus(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFocus", listener)
}

// FocusIn gets notified when an element is about to receive focus (bubbles).
//
// https://developer.mozilla.org/docs/Web/Events/focusin
func FocusIn(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFocusIn", listener)
}

// FocusOut gets notified when an element is about to lose focus (bubbles).
//
// https://developer.mozilla.org/docs/Web/Events/focusout
func FocusOut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFocusOut", listener)
}

// FullScreenChange gets notified when an element was turned to fullscreen mode or back to normal mode.
//
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
func FullScreenChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFullScreenChange", listener)
}

// FullScreenError gets notified when it was impossible to switch to fullscreen mode for technical reasons or because the permission was denied.
//
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
func FullScreenError(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFullScreenError", listener)
}

// GamepadConnected gets notified when a gamepad has been connected.
//
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
func GamepadConnected(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onGamepadConnected", listener)
}

// GamepadDisconnected gets notified when a gamepad has been disconnected.
//
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
func GamepadDisconnected(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onGamepadDisconnected", listener)
}

// GotPointerCapture gets notified when element receives pointer capture.
//
// https://developer.mozilla.org/docs/Web/Events/gotpointercapture
func GotPointerCapture(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onGotPointerCapture", listener)
}

// HashChange gets notified when the fragment identifier of the URL has changed (the part of the URL after the #).
//
// https://developer.mozilla.org/docs/Web/Events/hashchange
func HashChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onHashChange", listener)
}

// Input gets notified when the value of an element changes or the content of an element with the attribute contenteditable is modified.
//
// https://developer.mozilla.org/docs/Web/Events/input
func Input(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onInput", listener)
}

// Invalid gets notified when a submittable element has been checked and doesn't satisfy its constraints.
//
// https://developer.mozilla.org/docs/Web/Events/invalid
func Invalid(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onInvalid", listener)
}

// KeyDown gets notified when a key is pressed down.
//
// https://developer.mozilla.org/docs/Web/Events/keydown
func KeyDown(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onKeyDown", listener)
}

// KeyPress gets notified when a key is pressed down and that key normally produces a character value (use input instead).
//
// https://developer.mozilla.org/docs/Web/Events/keypress
func KeyPress(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onKeyPress", listener)
}

// KeyUp gets notified when a key is released.
//
// https://developer.mozilla.org/docs/Web/Events/keyup
func KeyUp(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onKeyUp", listener)
}

// LanguageChange gets notified when (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/languagechange
func LanguageChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLanguageChange", listener)
}

// LevelChange gets notified when the level attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/levelchange
func LevelChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLevelChange", listener)
}

// Load gets notified when progression has been successful.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
func Load(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoad", listener)
}

// LoadEnd gets notified when progress has stopped (after "error", "abort" or "load" have been dispatched).
//
// https://developer.mozilla.org/docs/Web/Events/loadend
func LoadEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadEnd", listener)
}

// LoadStart gets notified when progress has begun.
//
// https://developer.mozilla.org/docs/Web/Events/loadstart
func LoadStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadStart", listener)
}

// LoadedData gets notified when the first frame of the media has finished loading.
//
// https://developer.mozilla.org/docs/Web/Events/loadeddata
func LoadedData(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadedData", listener)
}

// LoadedMetadata gets notified when the metadata has been loaded.
//
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
func LoadedMetadata(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadedMetadata", listener)
}

// LostPointerCapture gets notified when element lost pointer capture.
//
// https://developer.mozilla.org/docs/Web/Events/lostpointercapture
func LostPointerCapture(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLostPointerCapture", listener)
}

// Mark gets notified when the spoken utterance reaches a named SSML "mark" tag.
//
// https://developer.mozilla.org/docs/Web/Events/mark
func Mark(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMark", listener)
}

// Message gets notified when a message is received from a service worker, or a message is received in a service worker from another context.
//
// https://developer.mozilla.org/docs/Web/Events/message_(ServiceWorker)
func Message(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMessage", listener)
}

// MouseDown gets notified when a pointing device button (usually a mouse) is pressed on an element.
//
// https://developer.mozilla.org/docs/Web/Events/mousedown
func MouseDown(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseDown", listener)
}

// MouseEnter gets notified when a pointing device is moved onto the element that has the listener attached.
//
// https://developer.mozilla.org/docs/Web/Events/mouseenter
func MouseEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseEnter", listener)
}

// MouseLeave gets notified when a pointing device is moved off the element that has the listener attached.
//
// https://developer.mozilla.org/docs/Web/Events/mouseleave
func MouseLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseLeave", listener)
}

// MouseMove gets notified when a pointing device is moved over an element.
//
// https://developer.mozilla.org/docs/Web/Events/mousemove
func MouseMove(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseMove", listener)
}

// MouseOut gets notified when a pointing device is moved off the element that has the listener attached or off one of its children.
//
// https://developer.mozilla.org/docs/Web/Events/mouseout
func MouseOut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseOut", listener)
}

// MouseOver gets notified when a pointing device is moved onto the element that has the listener attached or onto one of its children.
//
// https://developer.mozilla.org/docs/Web/Events/mouseover
func MouseOver(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseOver", listener)
}

// MouseUp gets notified when a pointing device button is released over an element.
//
// https://developer.mozilla.org/docs/Web/Events/mouseup
func MouseUp(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseUp", listener)
}

// NoMatch gets notified when the speech recognition service returns a final result with no significant recognition.
//
// https://developer.mozilla.org/docs/Web/Events/nomatch
func NoMatch(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onNoMatch", listener)
}

// NoUpdate gets notified when the manifest hadn't changed.
//
// https://developer.mozilla.org/docs/Web/Events/noupdate
func NoUpdate(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onNoUpdate", listener)
}

// NotificationClick gets notified when a system notification spawned by ServiceWorkerRegistration.showNotification() has been clicked.
//
// https://developer.mozilla.org/docs/Web/Events/notificationclick
func NotificationClick(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onNotificationClick", listener)
}

// Obsolete gets notified when the manifest was found to have become a 404 or 410 page, so the application cache is being deleted.
//
// https://developer.mozilla.org/docs/Web/Events/obsolete
func Obsolete(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onObsolete", listener)
}

// Offline gets notified when the browser has lost access to the network.
//
// https://developer.mozilla.org/docs/Web/Events/offline
func Offline(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOffline", listener)
}

// Online gets notified when the browser has gained access to the network (but particular websites might be unreachable).
//
// https://developer.mozilla.org/docs/Web/Events/online
func Online(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOnline", listener)
}

// Open gets notified when an event source connection has been established.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
func Open(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOpen", listener)
}

// OrientationChange gets notified when the orientation of the device (portrait/landscape) has changed
//
// https://developer.mozilla.org/docs/Web/Events/orientationchange
func OrientationChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOrientationChange", listener)
}

// PageHide gets notified when a session history entry is being traversed from.
//
// https://developer.mozilla.org/docs/Web/Events/pagehide
func PageHide(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPageHide", listener)
}

// PageShow gets notified when a session history entry is being traversed to.
//
// https://developer.mozilla.org/docs/Web/Events/pageshow
func PageShow(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPageShow", listener)
}

// Paste gets notified when data has been transferred from the system clipboard to the document.
//
// https://developer.mozilla.org/docs/Web/Events/paste
func Paste(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPaste", listener)
}

// Pause gets notified when the utterance is paused part way through.
//
// https://developer.mozilla.org/docs/Web/Events/pause_(SpeechSynthesis)
func Pause(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPause", listener)
}

// Play gets notified when playback has begun.
//
// https://developer.mozilla.org/docs/Web/Events/play
func Play(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPlay", listener)
}

// Playing gets notified when playback is ready to start after having been paused or delayed due to lack of data.
//
// https://developer.mozilla.org/docs/Web/Events/playing
func Playing(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPlaying", listener)
}

// PointerCancel gets notified when the pointer is unlikely to produce any more events.
//
// https://developer.mozilla.org/docs/Web/Events/pointercancel
func PointerCancel(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerCancel", listener)
}

// PointerDown gets notified when the pointer enters the active buttons state.
//
// https://developer.mozilla.org/docs/Web/Events/pointerdown
func PointerDown(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerDown", listener)
}

// PointerEnter gets notified when pointing device is moved inside the hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerenter
func PointerEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerEnter", listener)
}

// PointerLeave gets notified when pointing device is moved out of the hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerleave
func PointerLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerLeave", listener)
}

// PointerLockChange gets notified when the pointer was locked or released.
//
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
func PointerLockChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerLockChange", listener)
}

// PointerLockError gets notified when it was impossible to lock the pointer for technical reasons or because the permission was denied.
//
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
func PointerLockError(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerLockError", listener)
}

// PointerMove gets notified when the pointer changed coordinates.
//
// https://developer.mozilla.org/docs/Web/Events/pointermove
func PointerMove(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerMove", listener)
}

// PointerOut gets notified when the pointing device moved out of hit-testing boundary or leaves detectable hover range.
//
// https://developer.mozilla.org/docs/Web/Events/pointerout
func PointerOut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerOut", listener)
}

// PointerOver gets notified when the pointing device is moved into the hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerover
func PointerOver(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerOver", listener)
}

// PointerUp gets notified when the pointer leaves the active buttons state.
//
// https://developer.mozilla.org/docs/Web/Events/pointerup
func PointerUp(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerUp", listener)
}

// PopState gets notified when a session history entry is being navigated to (in certain cases).
//
// https://developer.mozilla.org/docs/Web/Events/popstate
func PopState(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPopState", listener)
}

// Progress gets notified when the user agent is downloading resources listed by the manifest.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
func Progress(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onProgress", listener)
}

// Push gets notified when a Service Worker has received a push message.
//
// https://developer.mozilla.org/docs/Web/Events/push
func Push(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPush", listener)
}

// PushSubscriptionChange gets notified when a PushSubscription has expired.
//
// https://developer.mozilla.org/docs/Web/Events/pushsubscriptionchange
func PushSubscriptionChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPushSubscriptionChange", listener)
}

// RateChange gets notified when the playback rate has changed.
//
// https://developer.mozilla.org/docs/Web/Events/ratechange
func RateChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onRateChange", listener)
}

// ReadyStateChange gets notified when the readyState attribute of a document has changed.
//
// https://developer.mozilla.org/docs/Web/Events/readystatechange
func ReadyStateChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onReadyStateChange", listener)
}

// RepeatEvent gets notified when a SMIL animation element is repeated.
//
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
func RepeatEvent(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onRepeatEvent", listener)
}

// Reset gets notified when a form is reset.
//
// https://developer.mozilla.org/docs/Web/Events/reset
func Reset(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onReset", listener)
}

// Resize gets notified when the document view has been resized.
//
// https://developer.mozilla.org/docs/Web/Events/resize
func Resize(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResize", listener)
}

// ResourceTimingBufferFull gets notified when the browser's resource timing buffer is full.
//
// https://developer.mozilla.org/docs/Web/Events/resourcetimingbufferfull
func ResourceTimingBufferFull(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResourceTimingBufferFull", listener)
}

// Result gets notified when the speech recognition service returns a result — a word or phrase has been positively recognized and this has been communicated back to the app.
//
// https://developer.mozilla.org/docs/Web/Events/result
func Result(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResult", listener)
}

// Resume gets notified when a paused utterance is resumed.
//
// https://developer.mozilla.org/docs/Web/Events/resume
func Resume(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResume", listener)
}

// SVGAbort gets notified when page loading has been stopped before the SVG was loaded.
//
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
func SVGAbort(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGAbort", listener)
}

// SVGError gets notified when an error has occurred before the SVG was loaded.
//
// https://developer.mozilla.org/docs/Web/Events/SVGError
func SVGError(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGError", listener)
}

// SVGLoad gets notified when an SVG document has been loaded and parsed.
//
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
func SVGLoad(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGLoad", listener)
}

// SVGResize gets notified when an SVG document is being resized.
//
// https://developer.mozilla.org/docs/Web/Events/SVGResize
func SVGResize(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGResize", listener)
}

// SVGScroll gets notified when an SVG document is being scrolled.
//
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
func SVGScroll(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGScroll", listener)
}

// SVGUnload gets notified when an SVG document has been removed from a window or frame.
//
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
func SVGUnload(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGUnload", listener)
}

// SVGZoom gets notified when an SVG document is being zoomed.
//
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
func SVGZoom(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGZoom", listener)
}

// Scroll gets notified when the document view or an element has been scrolled.
//
// https://developer.mozilla.org/docs/Web/Events/scroll
func Scroll(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onScroll", listener)
}

// Seeked gets notified when a seek operation completed.
//
// https://developer.mozilla.org/docs/Web/Events/seeked
func Seeked(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSeeked", listener)
}

// Seeking gets notified when a seek operation began.
//
// https://developer.mozilla.org/docs/Web/Events/seeking
func Seeking(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSeeking", listener)
}

// Select gets notified when some text is being selected.
//
// https://developer.mozilla.org/docs/Web/Events/select
func Select(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSelect", listener)
}

// SelectStart gets notified when a selection just started.
//
// https://developer.mozilla.org/docs/Web/Events/selectstart
func SelectStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSelectStart", listener)
}

// SelectionChange gets notified when the selection in the document has been changed.
//
// https://developer.mozilla.org/docs/Web/Events/selectionchange
func SelectionChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSelectionChange", listener)
}

// Show gets notified when a contextmenu event was fired on/bubbled to an element that has a contextmenu attribute
//
// https://developer.mozilla.org/docs/Web/Events/show
func Show(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onShow", listener)
}

// SoundEnd gets notified when any sound — recognisable speech or not — has stopped being detected.
//
// https://developer.mozilla.org/docs/Web/Events/soundend
func SoundEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSoundEnd", listener)
}

// SoundStart gets notified when any sound — recognisable speech or not — has been detected.
//
// https://developer.mozilla.org/docs/Web/Events/soundstart
func SoundStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSoundStart", listener)
}

// SpeechEnd gets notified when speech recognised by the speech recognition service has stopped being detected.
//
// https://developer.mozilla.org/docs/Web/Events/speechend
func SpeechEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSpeechEnd", listener)
}

// SpeechStart gets notified when sound that is recognised by the speech recognition service as speech has been detected.
//
// https://developer.mozilla.org/docs/Web/Events/speechstart
func SpeechStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSpeechStart", listener)
}

// Stalled gets notified when the user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
//
// https://developer.mozilla.org/docs/Web/Events/stalled
func Stalled(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onStalled", listener)
}

// Start gets notified when the utterance has begun to be spoken.
//
// https://developer.mozilla.org/docs/Web/Events/start_(SpeechSynthesis)
func Start(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onStart", listener)
}

// Storage gets notified when a storage area (localStorage or sessionStorage) has changed.
//
// https://developer.mozilla.org/docs/Web/Events/storage
func Storage(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onStorage", listener)
}

// Submit gets notified when a form is submitted.
//
// https://developer.mozilla.org/docs/Web/Events/submit
func Submit(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSubmit", listener)
}

// Success gets notified when a request successfully completed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
func Success(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSuccess", listener)
}

// Suspend gets notified when media data loading has been suspended.
//
// https://developer.mozilla.org/docs/Web/Events/suspend
func Suspend(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSuspend", listener)
}

// TimeUpdate gets notified when the time indicated by the currentTime attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/timeupdate
func TimeUpdate(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTimeUpdate", listener)
}

// Timeout gets notified when (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/timeout
func Timeout(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTimeout", listener)
}

// TouchCancel gets notified when a touch point has been disrupted in an implementation-specific manners (too many touch points for example).
//
// https://developer.mozilla.org/docs/Web/Events/touchcancel
func TouchCancel(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchCancel", listener)
}

// TouchEnd gets notified when a touch point is removed from the touch surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchend
func TouchEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchEnd", listener)
}

// TouchEnter gets notified when a touch point is moved onto the interactive area of an element.
//
// https://developer.mozilla.org/docs/Web/Events/touchenter
func TouchEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchEnter", listener)
}

// TouchLeave gets notified when a touch point is moved off the interactive area of an element.
//
// https://developer.mozilla.org/docs/Web/Events/touchleave
func TouchLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchLeave", listener)
}

// TouchMove gets notified when a touch point is moved along the touch surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchmove
func TouchMove(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchMove", listener)
}

// TouchStart gets notified when a touch point is placed on the touch surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchstart
func TouchStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchStart", listener)
}

// TransitionEnd gets notified when a CSS transition has completed.
//
// https://developer.mozilla.org/docs/Web/Events/transitionend
func TransitionEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTransitionEnd", listener)
}

// Unload gets notified when the document or a dependent resource is being unloaded.
//
// https://developer.mozilla.org/docs/Web/Events/unload
func Unload(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUnload", listener)
}

// UpdateReady gets notified when the resources listed in the manifest have been newly redownloaded, and the script can use swapCache() to switch to the new cache.
//
// https://developer.mozilla.org/docs/Web/Events/updateready
func UpdateReady(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUpdateReady", listener)
}

// UpgradeNeeded gets notified when an attempt was made to open a database with a version number higher than its current version. A versionchange transaction has been created.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
func UpgradeNeeded(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUpgradeNeeded", listener)
}

// UserProximity gets notified when fresh data is available from a proximity sensor (indicates whether the nearby object is near the device or not).
//
// https://developer.mozilla.org/docs/Web/Events/userproximity
func UserProximity(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUserProximity", listener)
}

// VersionChange gets notified when a versionchange transaction completed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
func VersionChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVersionChange", listener)
}

// VisibilityChange gets notified when the content of a tab has become visible or has been hidden.
//
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
func VisibilityChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVisibilityChange", listener)
}

// VoicesChanged gets notified when the list of SpeechSynthesisVoice objects that would be returned by the SpeechSynthesis.getVoices() method has changed (when the voiceschanged event fires.)
//
// https://developer.mozilla.org/docs/Web/Events/voiceschanged
func VoicesChanged(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVoicesChanged", listener)
}

// VolumeChange gets notified when the volume has changed.
//
// https://developer.mozilla.org/docs/Web/Events/volumechange
func VolumeChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVolumeChange", listener)
}

// Waiting gets notified when playback has stopped because of a temporary lack of data.
//
// https://developer.mozilla.org/docs/Web/Events/waiting
func Waiting(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onWaiting", listener)
}

// Wheel gets notified when a wheel button of a pointing device is rotated in any direction.
//
// https://developer.mozilla.org/docs/Web/Events/wheel
func Wheel(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onWheel", listener)
}
