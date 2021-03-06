// Package headlessexperimental provides the Chrome Debugging Protocol
// commands, types, and events for the HeadlessExperimental domain.
//
// This domain provides experimental commands only supported in headless
// mode.
//
// Generated by the cdproto-gen command.
package headlessexperimental

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"
	"encoding/base64"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
)

// BeginFrameParams sends a BeginFrame to the target and returns when the
// frame was completed. Optionally captures a screenshot from the resulting
// frame. Requires that the target was created with enabled BeginFrameControl.
// Designed for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gl/3zHXhB for more background.
type BeginFrameParams struct {
	FrameTime        *runtime.Timestamp `json:"frameTime,omitempty"`        // Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used unless frameTicks is specified.
	FrameTimeTicks   float64            `json:"frameTimeTicks,omitempty"`   // Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used unless frameTime is specified.
	Deadline         *runtime.Timestamp `json:"deadline,omitempty"`         // Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval unless deadlineTicks is specified.
	DeadlineTicks    float64            `json:"deadlineTicks,omitempty"`    // Deadline of this BeginFrame in Renderer TimeTicks  (milliseconds of uptime). If not set, the deadline will be calculated from the frameTime and interval unless deadline is specified.
	Interval         float64            `json:"interval,omitempty"`         // The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	NoDisplayUpdates bool               `json:"noDisplayUpdates,omitempty"` // Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	Screenshot       *ScreenshotParams  `json:"screenshot,omitempty"`       // If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
}

// BeginFrame sends a BeginFrame to the target and returns when the frame was
// completed. Optionally captures a screenshot from the resulting frame.
// Requires that the target was created with enabled BeginFrameControl. Designed
// for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gl/3zHXhB for more background.
//
// parameters:
func BeginFrame() *BeginFrameParams {
	return &BeginFrameParams{}
}

// WithFrameTime timestamp of this BeginFrame (milliseconds since epoch). If
// not set, the current time will be used unless frameTicks is specified.
func (p BeginFrameParams) WithFrameTime(frameTime *runtime.Timestamp) *BeginFrameParams {
	p.FrameTime = frameTime
	return &p
}

// WithFrameTimeTicks timestamp of this BeginFrame in Renderer TimeTicks
// (milliseconds of uptime). If not set, the current time will be used unless
// frameTime is specified.
func (p BeginFrameParams) WithFrameTimeTicks(frameTimeTicks float64) *BeginFrameParams {
	p.FrameTimeTicks = frameTimeTicks
	return &p
}

// WithDeadline deadline of this BeginFrame (milliseconds since epoch). If
// not set, the deadline will be calculated from the frameTime and interval
// unless deadlineTicks is specified.
func (p BeginFrameParams) WithDeadline(deadline *runtime.Timestamp) *BeginFrameParams {
	p.Deadline = deadline
	return &p
}

// WithDeadlineTicks deadline of this BeginFrame in Renderer TimeTicks
// (milliseconds of uptime). If not set, the deadline will be calculated from
// the frameTime and interval unless deadline is specified.
func (p BeginFrameParams) WithDeadlineTicks(deadlineTicks float64) *BeginFrameParams {
	p.DeadlineTicks = deadlineTicks
	return &p
}

// WithInterval the interval between BeginFrames that is reported to the
// compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e.
// about 16.666 milliseconds.
func (p BeginFrameParams) WithInterval(interval float64) *BeginFrameParams {
	p.Interval = interval
	return &p
}

// WithNoDisplayUpdates whether updates should not be committed and drawn
// onto the display. False by default. If true, only side effects of the
// BeginFrame will be run, such as layout and animations, but any visual updates
// may not be visible on the display or in screenshots.
func (p BeginFrameParams) WithNoDisplayUpdates(noDisplayUpdates bool) *BeginFrameParams {
	p.NoDisplayUpdates = noDisplayUpdates
	return &p
}

// WithScreenshot if set, a screenshot of the frame will be captured and
// returned in the response. Otherwise, no screenshot will be captured. Note
// that capturing a screenshot can fail, for example, during renderer
// initialization. In such a case, no screenshot data will be returned.
func (p BeginFrameParams) WithScreenshot(screenshot *ScreenshotParams) *BeginFrameParams {
	p.Screenshot = screenshot
	return &p
}

// BeginFrameReturns return values.
type BeginFrameReturns struct {
	HasDamage      bool   `json:"hasDamage,omitempty"`      // Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
	ScreenshotData string `json:"screenshotData,omitempty"` // Base64-encoded image data of the screenshot, if one was requested and successfully taken.
}

// Do executes HeadlessExperimental.beginFrame against the provided context.
//
// returns:
//   hasDamage - Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
//   screenshotData - Base64-encoded image data of the screenshot, if one was requested and successfully taken.
func (p *BeginFrameParams) Do(ctxt context.Context, h cdp.Executor) (hasDamage bool, screenshotData []byte, err error) {
	// execute
	var res BeginFrameReturns
	err = h.Execute(ctxt, CommandBeginFrame, p, &res)
	if err != nil {
		return false, nil, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.ScreenshotData)
	if err != nil {
		return false, nil, err
	}
	return res.HasDamage, dec, nil
}

// EnterDeterministicModeParams puts the browser into deterministic mode.
// Only effective for subsequently created web contents. Only supported in
// headless mode. Once set there's no way of leaving deterministic mode.
type EnterDeterministicModeParams struct {
	InitialDate float64 `json:"initialDate,omitempty"` // Number of seconds since the Epoch
}

// EnterDeterministicMode puts the browser into deterministic mode. Only
// effective for subsequently created web contents. Only supported in headless
// mode. Once set there's no way of leaving deterministic mode.
//
// parameters:
func EnterDeterministicMode() *EnterDeterministicModeParams {
	return &EnterDeterministicModeParams{}
}

// WithInitialDate number of seconds since the Epoch.
func (p EnterDeterministicModeParams) WithInitialDate(initialDate float64) *EnterDeterministicModeParams {
	p.InitialDate = initialDate
	return &p
}

// Do executes HeadlessExperimental.enterDeterministicMode against the provided context.
func (p *EnterDeterministicModeParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnterDeterministicMode, p, nil)
}

// DisableParams disables headless events for the target.
type DisableParams struct{}

// Disable disables headless events for the target.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes HeadlessExperimental.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables headless events for the target.
type EnableParams struct{}

// Enable enables headless events for the target.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes HeadlessExperimental.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnable, nil, nil)
}

// Command names.
const (
	CommandBeginFrame             = "HeadlessExperimental.beginFrame"
	CommandEnterDeterministicMode = "HeadlessExperimental.enterDeterministicMode"
	CommandDisable                = "HeadlessExperimental.disable"
	CommandEnable                 = "HeadlessExperimental.enable"
)
