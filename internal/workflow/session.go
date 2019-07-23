// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package workflow

import (
	"reflect"

	"github.com/honeydipper/honeydipper/internal/config"
	"github.com/honeydipper/honeydipper/pkg/dipper"
)

// Session is the data structure about a running workflow and its definition.
type Session struct {
	ID                string
	workflow          *config.Workflow
	current           int32 // current thread or step
	iteration         int32 // current item in the iteration list
	loopCount         int   // counter for looping
	parent            string
	ctx               map[string]interface{}
	event             map[string]interface{}
	exported          map[string]interface{}
	elseBranch        *config.Workflow
	collapsedFunction *config.CollapsedFunction
	store             *SessionStore
	loadedContexts    []string
	currentHook       string
	savedMsg          *dipper.Message
}

const (
	// SessionStatusSuccess means the workflow executed successfully
	SessionStatusSuccess = "success"
	// SessionStatusFailure means the workflow completed with failure
	SessionStatusFailure = "failure"
	// SessionStatusError means the workflow ran into error, and was not able to complete
	SessionStatusError = "error"

	// SessionContextDefault is a builtin context for all workflows
	SessionContextDefault = "_default"
	// SessionContextEvents is a context with default values for all directly event triggered workflows
	SessionContextEvents = "_events"
	// SessionContextHooks is a context for all hooks workflows
	SessionContextHooks = "_hooks"
)

// buildEnvData builds a map of environmental data for interpolation
func (w *Session) buildEnvData(msg *dipper.Message) map[string]interface{} {
	data := msg.Payload
	if data == nil {
		data = map[string]interface{}{}
	}

	envData := map[string]interface{}{
		"data":   data,
		"labels": msg.Labels,
		"event":  w.event,
		"ctx":    w.ctx,
	}

	return envData
}

// save will store the session in the memory
func (w *Session) save() string {
	if w.ID == "" {
		w.ID = dipper.IDMapPut(&w.store.sessions, w)
		dipper.Logger.Infof("[workflow] session with parent [%s] saved as [%s]", w.parent, w.ID)
	}
	return w.ID
}

// isLoop checks if the workflow uses looping statements while and until
func (w *Session) isLoop() bool {
	return len(w.workflow.While) > 0 || len(w.workflow.Until) > 0 || len(w.workflow.WhileAny) > 0 || len(w.workflow.UntilAny) > 0
}

// isIteration checks if the workflow needs to iterate through a list
func (w *Session) isIteration() bool {
	return w.workflow.Iterate != nil || w.workflow.IterateParallel != nil
}

// lenOfIterate gives the length of the iteration list
func (w *Session) lenOfIterate() int {
	var it reflect.Value
	if w.workflow.Iterate != nil {
		it = reflect.ValueOf(w.workflow.Iterate)
	} else {
		it = reflect.ValueOf(w.workflow.IterateParallel)
	}

	return it.Len()
}

// isFunction checks if the workflow is a simple function call
func (w *Session) isFunction() bool {
	return w.workflow.Function.Driver != "" || w.workflow.Function.Target.System != ""
}

// injectMsg injects the dipper message data into the session as event
func (w *Session) injectMsg(msg *dipper.Message) {
	if w.parent == "" {
		data, _ := dipper.GetMapData(msg.Payload, "data")
		if data != nil {
			d := data.(map[string]interface{})
			w.event = d
		} else {
			w.event = map[string]interface{}{}
		}
	}
}

// injectNamedCTX inject a named context into the workflow
func (w *Session) injectNamedCTX(name string) {
	var err error
	var contexts = w.store.GetConfig().DataSet.Contexts

	namedCTXs, ok := contexts[name]
	if name[0] != '_' && !ok {
		dipper.Logger.Panicf("[workflow] named workflow %s not defined", name)
	}
	if namedCTXs != nil {
		ctx, ok := namedCTXs["*"]
		if ok {
			ctx, err = dipper.DeepCopy(ctx)
			if err != nil {
				panic(err)
			}
			w.ctx = dipper.MergeMap(w.ctx, ctx)
			dipper.Logger.Infof("merged global values (*) from named context %s to workflow", name)
		}
		if w.workflow.Name != "" {
			ctx, ok := namedCTXs[w.workflow.Name]
			if ok {
				ctx, err = dipper.DeepCopy(ctx)
				if err != nil {
					panic(err)
				}
				w.ctx = dipper.MergeMap(w.ctx, ctx)
				dipper.Logger.Infof("[workflow] merged named context %s to workflow %s", name, w.workflow.Name)
			}
		}
	}
}

// initCTX initialize the contextual data used in this workflow
func (w *Session) initCTX() {
	w.injectNamedCTX(SessionContextDefault)
	if w.parent == "" {
		w.injectNamedCTX(SessionContextEvents)
	}

	var isHook bool

	for _, name := range w.loadedContexts {
		if name == SessionContextHooks {
			isHook = true
		}
		w.injectNamedCTX(name)
	}

	if w.workflow.Context != "" {
		if w.workflow.Context == SessionContextHooks {
			isHook = true
		}
		w.injectNamedCTX(w.workflow.Context)
		w.loadedContexts = append(w.loadedContexts, w.workflow.Context)
	} else {
		for _, name := range w.workflow.Contexts {
			if name == SessionContextHooks {
				isHook = true
			}
			w.injectNamedCTX(name)
			w.loadedContexts = append(w.loadedContexts, name)
		}
	}

	if isHook {
		// avoid hook in hook
		delete(w.ctx, "hooks")
	} else {
		w.ctx["_meta_name"] = w.workflow.Name
		w.ctx["_meta_desc"] = w.workflow.Description
	}
}

// injectEventCTX injects the contextual data from the event into the workflow
func (w *Session) injectEventCTX(ctx map[string]interface{}) {
	if ctx != nil {
		w.ctx = dipper.MergeMap(w.ctx, ctx)
	}
}

// injectLocalCTX injects the workflow local context data
func (w *Session) injectLocalCTX(msg *dipper.Message) {
	if w.workflow.Local != nil {
		envData := w.buildEnvData(msg)
		locals := dipper.Interpolate(w.workflow.Local, envData)

		w.ctx = dipper.MergeMap(w.ctx, locals)
	}
}

// interpolateWorkflow creates a copy of the workflow and interpolates it with envData
func (w *Session) interpolateWorkflow(msg *dipper.Message) {
	v := w.workflow
	envData := w.buildEnvData(msg)
	ret := config.Workflow{}

	ret.Name = dipper.InterpolateStr(v.Name, envData)
	ret.If = dipper.Interpolate(v.If, envData).([]string)
	ret.IfAny = dipper.Interpolate(v.IfAny, envData).([]string)
	ret.Unless = dipper.Interpolate(v.Unless, envData).([]string)
	ret.UnlessAny = dipper.Interpolate(v.UnlessAny, envData).([]string)
	ret.Match = dipper.Interpolate(v.Match, envData).(map[string]interface{})
	ret.Iterate = dipper.Interpolate(v.Iterate, envData)
	ret.IterateParallel = dipper.Interpolate(v.IterateParallel, envData)
	ret.Retry = dipper.InterpolateStr(v.Retry, envData)
	ret.Backoff = dipper.InterpolateStr(v.Backoff, envData)
	ret.Wait = dipper.InterpolateStr(v.Wait, envData)

	ret.While = v.While                     // repeatedly interpolated later
	ret.WhileAny = v.WhileAny               // repeatedly interpolated later
	ret.Until = v.Until                     // repeatedly interpolated later
	ret.UntilAny = v.UntilAny               // repeatedly interpolated later
	ret.Else = v.Else                       // delayed
	ret.Workflow = v.Workflow               // delayed
	ret.Function = v.Function               // delayed
	ret.CallFunc = v.CallFunc               // delayed
	ret.Steps = v.Steps                     // delayed
	ret.Threads = v.Threads                 // delayed
	ret.Export = v.Export                   // delayed
	ret.ExportOnSuccess = v.ExportOnSuccess // delayed
	ret.ExportOnFailure = v.ExportOnFailure // delayed
	ret.Switch = v.Switch                   // delayed
	ret.Cases = v.Cases                     // delayed
	ret.Default = v.Default                 // delayed

	ret.Context = v.Context         // no interpolation
	ret.Contexts = v.Contexts       // no interpolation
	ret.NoExport = v.NoExport       // no interpolation
	ret.IterateAs = v.IterateAs     // no interpolation
	ret.Description = v.Description // no interpolation

	w.workflow = &ret
}

// createChildSession creates a child workflow session
func (w *Session) createChildSession(wf *config.Workflow, msg *dipper.Message) *Session {
	child := w.store.newSession(w.ID, wf)
	child.injectMsg(msg)
	child.initCTX()
	child.injectLocalCTX(msg)
	child.interpolateWorkflow(msg)
	return child
}

// createChildSessionWithName creates a child workflow session
func (w *Session) createChildSessionWithName(name string, msg *dipper.Message) *Session {
	src := w.store.GetConfig().DataSet.Workflows[name]
	if src.Name == "" {
		src.Name = name
	}
	wf := &src
	return w.createChildSession(wf, msg)
}
