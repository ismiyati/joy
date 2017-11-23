package rtcdtmftonechangeevent

import (
	"github.com/matthewmueller/golly/dom/rtcdtmftonechangeeventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *rtcdtmftonechangeeventinit.RTCDTMFToneChangeEventInit) *RTCDTMFToneChangeEvent {
	js.Rewrite("RTCDTMFToneChangeEvent")
	return &RTCDTMFToneChangeEvent{}
}

// RTCDTMFToneChangeEvent struct
// js:"RTCDTMFToneChangeEvent,omit"
type RTCDTMFToneChangeEvent struct {
	window.Event
}

// Tone prop
func (*RTCDTMFToneChangeEvent) Tone() (tone string) {
	js.Rewrite("$<.tone")
	return tone
}
