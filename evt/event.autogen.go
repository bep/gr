//go:generate go run generate.go

// Package evt defines markup to bind DOM events.
//
// Generated from "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.
package evt

import "github.com/bep/gr"

// A transaction has been aborted.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
func Abort(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAbort", listener)
}

// The associated document has started printing or the print preview has been closed.
//
// https://developer.mozilla.org/docs/Web/Events/afterprint
func AfterPrint(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAfterPrint", listener)
}

// A CSS animation has completed.
//
// https://developer.mozilla.org/docs/Web/Events/animationend
func AnimationEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAnimationEnd", listener)
}

// A CSS animation is repeated.
//
// https://developer.mozilla.org/docs/Web/Events/animationiteration
func AnimationIteration(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAnimationIteration", listener)
}

// A CSS animation has started.
//
// https://developer.mozilla.org/docs/Web/Events/animationstart
func AnimationStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAnimationStart", listener)
}

// The user agent has finished capturing audio for speech recognition.
//
// https://developer.mozilla.org/docs/Web/Events/audioend
func AudioEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAudioEnd", listener)
}

// The input buffer of a ScriptProcessorNode is ready to be processed.
//
// https://developer.mozilla.org/docs/Web/Events/audioprocess
func AudioProcess(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAudioProcess", listener)
}

// The user agent has started to capture audio for speech recognition.
//
// https://developer.mozilla.org/docs/Web/Events/audiostart
func AudioStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onAudioStart", listener)
}

// The associated document is about to be printed or previewed for printing.
//
// https://developer.mozilla.org/docs/Web/Events/beforeprint
func BeforePrint(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBeforePrint", listener)
}

// (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/beforeunload
func BeforeUnload(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBeforeUnload", listener)
}

// A SMIL animation element begins.
//
// https://developer.mozilla.org/docs/Web/Events/beginEvent
func BeginEvent(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBeginEvent", listener)
}

// An open connection to a database is blocking a versionchange transaction on the same database.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
func Blocked(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBlocked", listener)
}

// An element has lost focus (does not bubble).
//
// https://developer.mozilla.org/docs/Web/Events/blur
func Blur(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBlur", listener)
}

// The spoken utterance reaches a word or sentence boundary
//
// https://developer.mozilla.org/docs/Web/Events/boundary
func Boundary(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onBoundary", listener)
}

// The resources listed in the manifest have been downloaded, and the application is now cached.
//
// https://developer.mozilla.org/docs/Web/Events/cached
func Cached(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCached", listener)
}

// The user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
//
// https://developer.mozilla.org/docs/Web/Events/canplay
func CanPlay(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCanPlay", listener)
}

// The user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
//
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
func CanPlayThrough(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCanPlayThrough", listener)
}

// The change event is fired for <input>, <select>, and <textarea> elements when a change to the element's value is committed by the user.
//
// https://developer.mozilla.org/docs/Web/Events/change
func Change(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChange", listener)
}

// The battery begins or stops charging.
//
// https://developer.mozilla.org/docs/Web/Events/chargingchange
func ChargingChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChargingChange", listener)
}

// The chargingTime attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
func ChargingTimeChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChargingTimeChange", listener)
}

// The user agent is checking for an update, or attempting to download the cache manifest for the first time.
//
// https://developer.mozilla.org/docs/Web/Events/checking
func Checking(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onChecking", listener)
}

// A pointing device button has been pressed and released on an element.
//
// https://developer.mozilla.org/docs/Web/Events/click
func Click(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onClick", listener)
}

// A WebSocket connection has been closed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/close_websocket
func Close(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onClose", listener)
}

// The rendering of an OfflineAudioContext is terminated.
//
// https://developer.mozilla.org/docs/Web/Events/complete
func Complete(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onComplete", listener)
}

// The composition of a passage of text has been completed or canceled.
//
// https://developer.mozilla.org/docs/Web/Events/compositionend
func CompositionEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCompositionEnd", listener)
}

// The composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition).
//
// https://developer.mozilla.org/docs/Web/Events/compositionstart
func CompositionStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCompositionStart", listener)
}

// A character is added to a passage of text being composed.
//
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
func CompositionUpdate(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCompositionUpdate", listener)
}

// The right button of the mouse is clicked (before the context menu is displayed).
//
// https://developer.mozilla.org/docs/Web/Events/contextmenu
func ContextMenu(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onContextMenu", listener)
}

// The text selection has been added to the clipboard.
//
// https://developer.mozilla.org/docs/Web/Events/copy
func Copy(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCopy", listener)
}

// The text selection has been removed from the document and added to the clipboard.
//
// https://developer.mozilla.org/docs/Web/Events/cut
func Cut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onCut", listener)
}

// The document has finished loading (but not its dependent resources).
//
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
func DOMContentLoaded(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDOMContentLoaded", listener)
}

// Fresh data is available from a light sensor.
//
// https://developer.mozilla.org/docs/Web/Events/devicelight
func DeviceLight(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceLight", listener)
}

// Fresh data is available from a motion sensor.
//
// https://developer.mozilla.org/docs/Web/Events/devicemotion
func DeviceMotion(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceMotion", listener)
}

// Fresh data is available from an orientation sensor.
//
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
func DeviceOrientation(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceOrientation", listener)
}

// Fresh data is available from a proximity sensor (indicates an approximated distance between the device and a nearby object).
//
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
func DeviceProximity(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDeviceProximity", listener)
}

// The dischargingTime attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
func DischargingTimeChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDischargingTimeChange", listener)
}

// A pointing device button is clicked twice on an element.
//
// https://developer.mozilla.org/docs/Web/Events/dblclick
func DoubleClick(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDoubleClick", listener)
}

// The user agent has found an update and is fetching it, or is downloading the resources listed by the cache manifest for the first time.
//
// https://developer.mozilla.org/docs/Web/Events/downloading
func Downloading(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDownloading", listener)
}

// An element or text selection is being dragged (every 350ms).
//
// https://developer.mozilla.org/docs/Web/Events/drag
func Drag(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDrag", listener)
}

// A drag operation is being ended (by releasing a mouse button or hitting the escape key).
//
// https://developer.mozilla.org/docs/Web/Events/dragend
func DragEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragEnd", listener)
}

// A dragged element or text selection enters a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/dragenter
func DragEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragEnter", listener)
}

// A dragged element or text selection leaves a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/dragleave
func DragLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragLeave", listener)
}

// An element or text selection is being dragged over a valid drop target (every 350ms).
//
// https://developer.mozilla.org/docs/Web/Events/dragover
func DragOver(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragOver", listener)
}

// The user starts dragging an element or text selection.
//
// https://developer.mozilla.org/docs/Web/Events/dragstart
func DragStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDragStart", listener)
}

// An element is dropped on a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/drop
func Drop(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDrop", listener)
}

// The duration attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/durationchange
func DurationChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onDurationChange", listener)
}

// The media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.
//
// https://developer.mozilla.org/docs/Web/Events/emptied
func Emptied(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEmptied", listener)
}

// The utterance has finished being spoken.
//
// https://developer.mozilla.org/docs/Web/Events/end_(SpeechSynthesis)
func End(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEnd", listener)
}

// A SMIL animation element ends.
//
// https://developer.mozilla.org/docs/Web/Events/endEvent
func EndEvent(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEndEvent", listener)
}

// (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
func Ended(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onEnded", listener)
}

// An error occurs that prevents the utterance from being succesfully spoken.
//
// https://developer.mozilla.org/docs/Web/Events/error_(SpeechSynthesisError)
func Error(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onError", listener)
}

// An element has received focus (does not bubble).
//
// https://developer.mozilla.org/docs/Web/Events/focus
func Focus(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFocus", listener)
}

// An element is about to receive focus (bubbles).
//
// https://developer.mozilla.org/docs/Web/Events/focusin
func FocusIn(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFocusIn", listener)
}

// An element is about to lose focus (bubbles).
//
// https://developer.mozilla.org/docs/Web/Events/focusout
func FocusOut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFocusOut", listener)
}

// An element was turned to fullscreen mode or back to normal mode.
//
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
func FullScreenChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFullScreenChange", listener)
}

// It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied.
//
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
func FullScreenError(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onFullScreenError", listener)
}

// A gamepad has been connected.
//
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
func GamepadConnected(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onGamepadConnected", listener)
}

// A gamepad has been disconnected.
//
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
func GamepadDisconnected(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onGamepadDisconnected", listener)
}

// Element receives pointer capture.
//
// https://developer.mozilla.org/docs/Web/Events/gotpointercapture
func GotPointerCapture(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onGotPointerCapture", listener)
}

// The fragment identifier of the URL has changed (the part of the URL after the #).
//
// https://developer.mozilla.org/docs/Web/Events/hashchange
func HashChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onHashChange", listener)
}

// The value of an element changes or the content of an element with the attribute contenteditable is modified.
//
// https://developer.mozilla.org/docs/Web/Events/input
func Input(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onInput", listener)
}

// A submittable element has been checked and doesn't satisfy its constraints.
//
// https://developer.mozilla.org/docs/Web/Events/invalid
func Invalid(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onInvalid", listener)
}

// A key is pressed down.
//
// https://developer.mozilla.org/docs/Web/Events/keydown
func KeyDown(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onKeyDown", listener)
}

// A key is pressed down and that key normally produces a character value (use input instead).
//
// https://developer.mozilla.org/docs/Web/Events/keypress
func KeyPress(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onKeyPress", listener)
}

// A key is released.
//
// https://developer.mozilla.org/docs/Web/Events/keyup
func KeyUp(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onKeyUp", listener)
}

// (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/languagechange
func LanguageChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLanguageChange", listener)
}

// The level attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/levelchange
func LevelChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLevelChange", listener)
}

// Progression has been successful.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
func Load(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoad", listener)
}

// Progress has stopped (after "error", "abort" or "load" have been dispatched).
//
// https://developer.mozilla.org/docs/Web/Events/loadend
func LoadEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadEnd", listener)
}

// Progress has begun.
//
// https://developer.mozilla.org/docs/Web/Events/loadstart
func LoadStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadStart", listener)
}

// The first frame of the media has finished loading.
//
// https://developer.mozilla.org/docs/Web/Events/loadeddata
func LoadedData(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadedData", listener)
}

// The metadata has been loaded.
//
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
func LoadedMetadata(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLoadedMetadata", listener)
}

// Element lost pointer capture.
//
// https://developer.mozilla.org/docs/Web/Events/lostpointercapture
func LostPointerCapture(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onLostPointerCapture", listener)
}

// The spoken utterance reaches a named SSML "mark" tag.
//
// https://developer.mozilla.org/docs/Web/Events/mark
func Mark(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMark", listener)
}

// A message is received from a service worker, or a message is received in a service worker from another context.
//
// https://developer.mozilla.org/docs/Web/Events/message_(ServiceWorker)
func Message(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMessage", listener)
}

// A pointing device button (usually a mouse) is pressed on an element.
//
// https://developer.mozilla.org/docs/Web/Events/mousedown
func MouseDown(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseDown", listener)
}

// A pointing device is moved onto the element that has the listener attached.
//
// https://developer.mozilla.org/docs/Web/Events/mouseenter
func MouseEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseEnter", listener)
}

// A pointing device is moved off the element that has the listener attached.
//
// https://developer.mozilla.org/docs/Web/Events/mouseleave
func MouseLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseLeave", listener)
}

// A pointing device is moved over an element.
//
// https://developer.mozilla.org/docs/Web/Events/mousemove
func MouseMove(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseMove", listener)
}

// A pointing device is moved off the element that has the listener attached or off one of its children.
//
// https://developer.mozilla.org/docs/Web/Events/mouseout
func MouseOut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseOut", listener)
}

// A pointing device is moved onto the element that has the listener attached or onto one of its children.
//
// https://developer.mozilla.org/docs/Web/Events/mouseover
func MouseOver(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseOver", listener)
}

// A pointing device button is released over an element.
//
// https://developer.mozilla.org/docs/Web/Events/mouseup
func MouseUp(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onMouseUp", listener)
}

// The speech recognition service returns a final result with no significant recognition.
//
// https://developer.mozilla.org/docs/Web/Events/nomatch
func NoMatch(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onNoMatch", listener)
}

// The manifest hadn't changed.
//
// https://developer.mozilla.org/docs/Web/Events/noupdate
func NoUpdate(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onNoUpdate", listener)
}

// A system notification spawned by ServiceWorkerRegistration.showNotification() has been clicked.
//
// https://developer.mozilla.org/docs/Web/Events/notificationclick
func NotificationClick(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onNotificationClick", listener)
}

// The manifest was found to have become a 404 or 410 page, so the application cache is being deleted.
//
// https://developer.mozilla.org/docs/Web/Events/obsolete
func Obsolete(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onObsolete", listener)
}

// The browser has lost access to the network.
//
// https://developer.mozilla.org/docs/Web/Events/offline
func Offline(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOffline", listener)
}

// The browser has gained access to the network (but particular websites might be unreachable).
//
// https://developer.mozilla.org/docs/Web/Events/online
func Online(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOnline", listener)
}

// An event source connection has been established.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
func Open(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOpen", listener)
}

// The orientation of the device (portrait/landscape) has changed
//
// https://developer.mozilla.org/docs/Web/Events/orientationchange
func OrientationChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onOrientationChange", listener)
}

// A session history entry is being traversed from.
//
// https://developer.mozilla.org/docs/Web/Events/pagehide
func PageHide(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPageHide", listener)
}

// A session history entry is being traversed to.
//
// https://developer.mozilla.org/docs/Web/Events/pageshow
func PageShow(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPageShow", listener)
}

// Data has been transfered from the system clipboard to the document.
//
// https://developer.mozilla.org/docs/Web/Events/paste
func Paste(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPaste", listener)
}

// The utterance is paused part way through.
//
// https://developer.mozilla.org/docs/Web/Events/pause_(SpeechSynthesis)
func Pause(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPause", listener)
}

// Playback has begun.
//
// https://developer.mozilla.org/docs/Web/Events/play
func Play(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPlay", listener)
}

// Playback is ready to start after having been paused or delayed due to lack of data.
//
// https://developer.mozilla.org/docs/Web/Events/playing
func Playing(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPlaying", listener)
}

// The pointer is unlikely to produce any more events.
//
// https://developer.mozilla.org/docs/Web/Events/pointercancel
func PointerCancel(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerCancel", listener)
}

// The pointer enters the active buttons state.
//
// https://developer.mozilla.org/docs/Web/Events/pointerdown
func PointerDown(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerDown", listener)
}

// Pointing device is moved inside the hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerenter
func PointerEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerEnter", listener)
}

// Pointing device is moved out of the hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerleave
func PointerLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerLeave", listener)
}

// The pointer was locked or released.
//
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
func PointerLockChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerLockChange", listener)
}

// It was impossible to lock the pointer for technical reasons or because the permission was denied.
//
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
func PointerLockError(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerLockError", listener)
}

// The pointer changed coordinates.
//
// https://developer.mozilla.org/docs/Web/Events/pointermove
func PointerMove(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerMove", listener)
}

// The pointing device moved out of hit-testing boundary or leaves detectable hover range.
//
// https://developer.mozilla.org/docs/Web/Events/pointerout
func PointerOut(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerOut", listener)
}

// The pointing device is moved into the hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerover
func PointerOver(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerOver", listener)
}

// The pointer leaves the active buttons state.
//
// https://developer.mozilla.org/docs/Web/Events/pointerup
func PointerUp(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPointerUp", listener)
}

// A session history entry is being navigated to (in certain cases).
//
// https://developer.mozilla.org/docs/Web/Events/popstate
func PopState(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPopState", listener)
}

// The user agent is downloading resources listed by the manifest.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
func Progress(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onProgress", listener)
}

// A Service Worker has received a push message.
//
// https://developer.mozilla.org/docs/Web/Events/push
func Push(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPush", listener)
}

// A PushSubscription has expired.
//
// https://developer.mozilla.org/docs/Web/Events/pushsubscriptionchange
func PushSubscriptionChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onPushSubscriptionChange", listener)
}

// The playback rate has changed.
//
// https://developer.mozilla.org/docs/Web/Events/ratechange
func RateChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onRateChange", listener)
}

// The readyState attribute of a document has changed.
//
// https://developer.mozilla.org/docs/Web/Events/readystatechange
func ReadyStateChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onReadyStateChange", listener)
}

// A SMIL animation element is repeated.
//
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
func RepeatEvent(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onRepeatEvent", listener)
}

// A form is reset.
//
// https://developer.mozilla.org/docs/Web/Events/reset
func Reset(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onReset", listener)
}

// The document view has been resized.
//
// https://developer.mozilla.org/docs/Web/Events/resize
func Resize(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResize", listener)
}

// The browser's resource timing buffer is full.
//
// https://developer.mozilla.org/docs/Web/Events/resourcetimingbufferfull
func ResourceTimingBufferFull(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResourceTimingBufferFull", listener)
}

// The speech recognition service returns a result — a word or phrase has been positively recognized and this has been communicated back to the app.
//
// https://developer.mozilla.org/docs/Web/Events/result
func Result(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResult", listener)
}

// A paused utterance is resumed.
//
// https://developer.mozilla.org/docs/Web/Events/resume
func Resume(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onResume", listener)
}

// Page loading has been stopped before the SVG was loaded.
//
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
func SVGAbort(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGAbort", listener)
}

// An error has occurred before the SVG was loaded.
//
// https://developer.mozilla.org/docs/Web/Events/SVGError
func SVGError(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGError", listener)
}

// An SVG document has been loaded and parsed.
//
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
func SVGLoad(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGLoad", listener)
}

// An SVG document is being resized.
//
// https://developer.mozilla.org/docs/Web/Events/SVGResize
func SVGResize(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGResize", listener)
}

// An SVG document is being scrolled.
//
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
func SVGScroll(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGScroll", listener)
}

// An SVG document has been removed from a window or frame.
//
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
func SVGUnload(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGUnload", listener)
}

// An SVG document is being zoomed.
//
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
func SVGZoom(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSVGZoom", listener)
}

// The document view or an element has been scrolled.
//
// https://developer.mozilla.org/docs/Web/Events/scroll
func Scroll(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onScroll", listener)
}

// A seek operation completed.
//
// https://developer.mozilla.org/docs/Web/Events/seeked
func Seeked(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSeeked", listener)
}

// A seek operation began.
//
// https://developer.mozilla.org/docs/Web/Events/seeking
func Seeking(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSeeking", listener)
}

// Some text is being selected.
//
// https://developer.mozilla.org/docs/Web/Events/select
func Select(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSelect", listener)
}

// A selection just started.
//
// https://developer.mozilla.org/docs/Web/Events/selectstart
func SelectStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSelectStart", listener)
}

// The selection in the document has been changed.
//
// https://developer.mozilla.org/docs/Web/Events/selectionchange
func SelectionChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSelectionChange", listener)
}

// A contextmenu event was fired on/bubbled to an element that has a contextmenu attribute
//
// https://developer.mozilla.org/docs/Web/Events/show
func Show(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onShow", listener)
}

// Any sound — recognisable speech or not — has stopped being detected.
//
// https://developer.mozilla.org/docs/Web/Events/soundend
func SoundEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSoundEnd", listener)
}

// Any sound — recognisable speech or not — has been detected.
//
// https://developer.mozilla.org/docs/Web/Events/soundstart
func SoundStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSoundStart", listener)
}

// Speech recognised by the speech recognition service has stopped being detected.
//
// https://developer.mozilla.org/docs/Web/Events/speechend
func SpeechEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSpeechEnd", listener)
}

// Sound that is recognised by the speech recognition service as speech has been detected.
//
// https://developer.mozilla.org/docs/Web/Events/speechstart
func SpeechStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSpeechStart", listener)
}

// The user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
//
// https://developer.mozilla.org/docs/Web/Events/stalled
func Stalled(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onStalled", listener)
}

// The utterance has begun to be spoken.
//
// https://developer.mozilla.org/docs/Web/Events/start_(SpeechSynthesis)
func Start(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onStart", listener)
}

// A storage area (localStorage or sessionStorage) has changed.
//
// https://developer.mozilla.org/docs/Web/Events/storage
func Storage(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onStorage", listener)
}

// A form is submitted.
//
// https://developer.mozilla.org/docs/Web/Events/submit
func Submit(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSubmit", listener)
}

// A request successfully completed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
func Success(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSuccess", listener)
}

// Media data loading has been suspended.
//
// https://developer.mozilla.org/docs/Web/Events/suspend
func Suspend(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onSuspend", listener)
}

// The time indicated by the currentTime attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/timeupdate
func TimeUpdate(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTimeUpdate", listener)
}

// (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/timeout
func Timeout(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTimeout", listener)
}

// A touch point has been disrupted in an implementation-specific manners (too many touch points for example).
//
// https://developer.mozilla.org/docs/Web/Events/touchcancel
func TouchCancel(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchCancel", listener)
}

// A touch point is removed from the touch surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchend
func TouchEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchEnd", listener)
}

// A touch point is moved onto the interactive area of an element.
//
// https://developer.mozilla.org/docs/Web/Events/touchenter
func TouchEnter(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchEnter", listener)
}

// A touch point is moved off the interactive area of an element.
//
// https://developer.mozilla.org/docs/Web/Events/touchleave
func TouchLeave(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchLeave", listener)
}

// A touch point is moved along the touch surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchmove
func TouchMove(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchMove", listener)
}

// A touch point is placed on the touch surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchstart
func TouchStart(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTouchStart", listener)
}

// A CSS transition has completed.
//
// https://developer.mozilla.org/docs/Web/Events/transitionend
func TransitionEnd(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onTransitionEnd", listener)
}

// The document or a dependent resource is being unloaded.
//
// https://developer.mozilla.org/docs/Web/Events/unload
func Unload(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUnload", listener)
}

// The resources listed in the manifest have been newly redownloaded, and the script can use swapCache() to switch to the new cache.
//
// https://developer.mozilla.org/docs/Web/Events/updateready
func UpdateReady(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUpdateReady", listener)
}

// An attempt was made to open a database with a version number higher than its current version. A versionchange transaction has been created.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
func UpgradeNeeded(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUpgradeNeeded", listener)
}

// Fresh data is available from a proximity sensor (indicates whether the nearby object is near the device or not).
//
// https://developer.mozilla.org/docs/Web/Events/userproximity
func UserProximity(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onUserProximity", listener)
}

// A versionchange transaction completed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
func VersionChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVersionChange", listener)
}

// The content of a tab has become visible or has been hidden.
//
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
func VisibilityChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVisibilityChange", listener)
}

// The list of SpeechSynthesisVoice objects that would be returned by the SpeechSynthesis.getVoices() method has changed (when the voiceschanged event fires.)
//
// https://developer.mozilla.org/docs/Web/Events/voiceschanged
func VoicesChanged(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVoicesChanged", listener)
}

// The volume has changed.
//
// https://developer.mozilla.org/docs/Web/Events/volumechange
func VolumeChange(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onVolumeChange", listener)
}

// Playback has stopped because of a temporary lack of data.
//
// https://developer.mozilla.org/docs/Web/Events/waiting
func Waiting(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onWaiting", listener)
}

// A wheel button of a pointing device is rotated in any direction.
//
// https://developer.mozilla.org/docs/Web/Events/wheel
func Wheel(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("onWheel", listener)
}
