package main

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	// TODO: add type checking props
	getDefaultProps = "getDefaultProps"

	// Mounting
	getInitialState          = "getInitialState"
	getDerivedStateFromProps = "getDerivedStateFromProps"
	render                   = "render"
	componentDidMount        = "componentDidMount"

	// Updating
	shouldComponentUpdate   = "shouldComponentUpdate"
	getSnapshotBeforeUpdate = "getSnapshotBeforeUpdate"
	componentDidUpdate      = "componentDidUpdate"

	// Unmounting
	componentWillUnmount = "componentWillUnmount"
)

// SetGetDefaultProps sets the getDefaultProps method.
func (def ClassDef) SetGetDefaultProps(f func(this *js.Object) interface{}) {
	def.SetMethod(getDefaultProps, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		return SToMap(f(this))
	})
}

// SetGetInitialState sets the getInitialState method.
// Note: It is usually not recommended to use the props when setting the state.
func (def ClassDef) SetGetInitialState(f func(this *js.Object, props Map) interface{}) {
	def.SetMethod(getInitialState, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		return SToMap(f(this, props))
	})
}

// SetGetDerivedStateFromProps sets the getDerivedStateFromProps class method.
// See: https://reactjs.org/blog/2018/06/07/you-probably-dont-need-derived-state.html
func (def ClassDef) SetGetDerivedStateFromProps(f func(nextProps, prevState Map) interface{}) {

	def.setMethod(true, getDerivedStateFromProps, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {

		nextProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		prevState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}

		return SToMap(f(nextProps, prevState))
	})
}

// SetComponentDidMount sets the componentDidMount method.
// See: https://reactjs.org/docs/react-component.html#componentdidmount
func (def ClassDef) SetComponentDidMount(f func(this *js.Object, props, state Map, setState SetState)) {
	def.SetMethod(componentDidMount, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		f(this, props, state, setState)
		return nil
	})
}

// SetComponentWillUnmount sets the componentWillUnmount method.
// See: https://reactjs.org/docs/react-component.html#componentwillunmount
func (def ClassDef) SetComponentWillUnmount(f func(this *js.Object, props, state Map)) {
	def.SetMethod(componentWillUnmount, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		f(this, props, state)
		return nil
	})
}

// SetShouldComponentUpdate sets the shouldComponentUpdate method.
// See: https://reactjs.org/docs/react-component.html#shouldcomponentupdate
func (def ClassDef) SetShouldComponentUpdate(f func(this *js.Object, props, nextProps, state, nextState Map) bool) {
	def.SetMethod(shouldComponentUpdate, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		nextProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		nextState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}
		return f(this, props, nextProps, state, nextState)
	})
}

// SetGetSnapshotBeforeUpdate sets the getSnapshotBeforeUpdate method.
// See: https://reactjs.org/docs/react-component.html#getsnapshotbeforeupdate
func (def ClassDef) SetGetSnapshotBeforeUpdate(f func(this *js.Object, prevProps, props, prevState, state Map) interface{}) {
	def.SetMethod(getSnapshotBeforeUpdate, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		prevProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		prevState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}

		ret := f(this, prevProps, props, prevState, state)
		if ret == nil {
			return nil
		} else if isStruct(ret) {
			return convertStruct(ret)
		} else {
			return ret
		}
	})
}

// SetComponentDidUpdate sets the componentDidUpdate method.
// See: https://reactjs.org/docs/react-component.html#componentdidupdate
func (def ClassDef) SetComponentDidUpdate(f func(this *js.Object, prevProps, props, prevState, state Map, setState SetState, snapshot *js.Object)) {
	def.SetMethod(componentDidUpdate, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		snapshot := arguments[2]
		prevProps := func(key string) *js.Object {
			return arguments[0].Get(key)
		}
		prevState := func(key string) *js.Object {
			return arguments[1].Get(key)
		}
		f(this, prevProps, props, prevState, state, setState, snapshot)
		return nil
	})
}

// SetRender sets the render method.
func (def ClassDef) SetRender(f func(this *js.Object, props, state Map) interface{}) {
	def.SetMethod(render, func(this *js.Object, props, state Map, setState SetState, arguments []*js.Object) interface{} {
		return f(this, props, state)
	})
}