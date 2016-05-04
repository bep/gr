// +build ignore

// Portions Copyright (c) 2016 The Vecty Authors. All rights reserved.
// See https://github.com/gopherjs/vecty for the origin of this clever
// code generator.

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

type Event struct {
	Name string
	Link string
	Desc string
}

func main() {
	nameMap := map[string]string{
		"afterprint":               "AfterPrint",
		"animationend":             "AnimationEnd",
		"animationiteration":       "AnimationIteration",
		"animationstart":           "AnimationStart",
		"audioprocess":             "AudioProcess",
		"audioend":                 "AudioEnd",
		"audiostart":               "AudioStart",
		"beforeprint":              "BeforePrint",
		"beforeunload":             "BeforeUnload",
		"canplay":                  "CanPlay",
		"canplaythrough":           "CanPlayThrough",
		"chargingchange":           "ChargingChange",
		"chargingtimechange":       "ChargingTimeChange",
		"compassneedscalibration":  "CompassNeedsCalibration",
		"compositionend":           "CompositionEnd",
		"compositionstart":         "CompositionStart",
		"compositionupdate":        "CompositionUpdate",
		"contextmenu":              "ContextMenu",
		"dblclick":                 "DoubleClick",
		"devicelight":              "DeviceLight",
		"devicemotion":             "DeviceMotion",
		"deviceorientation":        "DeviceOrientation",
		"deviceproximity":          "DeviceProximity",
		"dischargingtimechange":    "DischargingTimeChange",
		"dragend":                  "DragEnd",
		"dragenter":                "DragEnter",
		"dragleave":                "DragLeave",
		"dragover":                 "DragOver",
		"dragstart":                "DragStart",
		"durationchange":           "DurationChange",
		"focusin":                  "FocusIn",
		"focusout":                 "FocusOut",
		"fullscreenchange":         "FullScreenChange",
		"fullscreenerror":          "FullScreenError",
		"gamepadconnected":         "GamepadConnected",
		"gamepaddisconnected":      "GamepadDisconnected",
		"gotpointercapture":        "GotPointerCapture",
		"hashchange":               "HashChange",
		"keydown":                  "KeyDown",
		"keypress":                 "KeyPress",
		"keyup":                    "KeyUp",
		"languagechange":           "LanguageChange",
		"levelchange":              "LevelChange",
		"loadeddata":               "LoadedData",
		"loadedmetadata":           "LoadedMetadata",
		"loadend":                  "LoadEnd",
		"loadstart":                "LoadStart",
		"lostpointercapture":       "LostPointerCapture",
		"mousedown":                "MouseDown",
		"mouseenter":               "MouseEnter",
		"mouseleave":               "MouseLeave",
		"mousemove":                "MouseMove",
		"mouseout":                 "MouseOut",
		"mouseover":                "MouseOver",
		"mouseup":                  "MouseUp",
		"noupdate":                 "NoUpdate",
		"nomatch":                  "NoMatch",
		"notificationclick":        "NotificationClick",
		"orientationchange":        "OrientationChange",
		"pagehide":                 "PageHide",
		"pageshow":                 "PageShow",
		"pointercancel":            "PointerCancel",
		"pointerdown":              "PointerDown",
		"pointerenter":             "PointerEnter",
		"pointerleave":             "PointerLeave",
		"pointerlockchange":        "PointerLockChange",
		"pointerlockerror":         "PointerLockError",
		"pointermove":              "PointerMove",
		"pointerout":               "PointerOut",
		"pointerover":              "PointerOver",
		"pointerup":                "PointerUp",
		"popstate":                 "PopState",
		"pushsubscriptionchange":   "PushSubscriptionChange",
		"ratechange":               "RateChange",
		"readystatechange":         "ReadyStateChange",
		"resourcetimingbufferfull": "ResourceTimingBufferFull",
		"selectstart":              "SelectStart",
		"selectionchange":          "SelectionChange",
		"soundend":                 "SoundEnd",
		"soundstart":               "SoundStart",
		"speechend":                "SpeechEnd",
		"speechstart":              "SpeechStart",
		"timeupdate":               "TimeUpdate",
		"touchcancel":              "TouchCancel",
		"touchend":                 "TouchEnd",
		"touchenter":               "TouchEnter",
		"touchleave":               "TouchLeave",
		"touchmove":                "TouchMove",
		"touchstart":               "TouchStart",
		"transitionend":            "TransitionEnd",
		"updateready":              "UpdateReady",
		"upgradeneeded":            "UpgradeNeeded",
		"userproximity":            "UserProximity",
		"versionchange":            "VersionChange",
		"visibilitychange":         "VisibilityChange",
		"voiceschanged":            "VoicesChanged",
		"volumechange":             "VolumeChange",
	}

	doc, err := goquery.NewDocument("https://developer.mozilla.org/en-US/docs/Web/Events")
	if err != nil {
		panic(err)
	}

	events := make(map[string]*Event)

	doc.Find(".standard-table").Eq(0).Find("tr").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		if cols.Length() == 0 || cols.Find(".icon-thumbs-down-alt").Length() != 0 {
			return
		}
		link := cols.Eq(0).Find("a").Eq(0)
		var e Event
		e.Name = link.Text()
		e.Link, _ = link.Attr("href")
		e.Desc = strings.TrimSpace(cols.Eq(3).Text())
		if e.Desc == "" {
			e.Desc = "(no documentation)"
		}

		funName := nameMap[e.Name]
		if funName == "" {
			funName = capitalize(e.Name)
		}

		events[funName] = &e
	})

	var names []string
	for name := range events {
		names = append(names, name)
	}
	sort.Strings(names)

	file, err := os.Create("event.autogen.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, `//go:generate go run generate.go

// Package evt defines markup to bind DOM events.
//
// Generated from "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.
package evt

import "github.com/bep/gr"
`)

	for _, name := range names {
		e := events[name]
		fmt.Fprintf(file, `
// %s gets notified when %s
//
// https://developer.mozilla.org%s
func %s(listener gr.Listener) *gr.EventListener {
	return gr.NewEventListener("on%s", listener)
}
`, name, firstToLower(e.Desc), e.Link[6:], name, name)
	}
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func firstToLower(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
