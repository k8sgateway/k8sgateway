package storage

import "github.com/solo-io/gloo/pkg/api/types/v1"

type Watcher struct {
	runFunc func(stop <-chan struct{}, errs chan error)
}

func NewWatcher(runFunc func(stop <-chan struct{}, errs chan error)) *Watcher {
	return &Watcher{runFunc: runFunc}
}

func (w *Watcher) Run(stop <-chan struct{}, errs chan error) {
	w.runFunc(stop, errs)
}

type UpstreamEventHandler interface {
	OnAdd(updatedList []*v1.Upstream, obj *v1.Upstream)
	OnUpdate(updatedList []*v1.Upstream, newObj *v1.Upstream)
	OnDelete(updatedList []*v1.Upstream, obj *v1.Upstream)
}

type VirtualServiceEventHandler interface {
	OnAdd(updatedList []*v1.VirtualService, obj *v1.VirtualService)
	OnUpdate(updatedList []*v1.VirtualService, newObj *v1.VirtualService)
	OnDelete(updatedList []*v1.VirtualService, obj *v1.VirtualService)
}

type RoleEventHandler interface {
	OnAdd(updatedList []*v1.Role, obj *v1.Role)
	OnUpdate(updatedList []*v1.Role, newObj *v1.Role)
	OnDelete(updatedList []*v1.Role, obj *v1.Role)
}

type AttributeEventHandler interface {
	OnAdd(updatedList []*v1.Attribute, obj *v1.Attribute)
	OnUpdate(updatedList []*v1.Attribute, newObj *v1.Attribute)
	OnDelete(updatedList []*v1.Attribute, obj *v1.Attribute)
}

// UpstreamEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// UpstreamEventHandler.
type UpstreamEventHandlerFuncs struct {
	AddFunc    func(updatedList []*v1.Upstream, obj *v1.Upstream)
	UpdateFunc func(updatedList []*v1.Upstream, newObj *v1.Upstream)
	DeleteFunc func(updatedList []*v1.Upstream, obj *v1.Upstream)
}

// OnAdd calls AddFunc if it's not nil.
func (r UpstreamEventHandlerFuncs) OnAdd(updatedList []*v1.Upstream, obj *v1.Upstream) {
	if r.AddFunc != nil {
		r.AddFunc(updatedList, obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r UpstreamEventHandlerFuncs) OnUpdate(updatedList []*v1.Upstream, newObj *v1.Upstream) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(updatedList, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r UpstreamEventHandlerFuncs) OnDelete(updatedList []*v1.Upstream, obj *v1.Upstream) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(updatedList, obj)
	}
}

// VirtualServiceEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// VirtualServiceEventHandler.
type VirtualServiceEventHandlerFuncs struct {
	AddFunc    func(updatedList []*v1.VirtualService, obj *v1.VirtualService)
	UpdateFunc func(updatedList []*v1.VirtualService, newObj *v1.VirtualService)
	DeleteFunc func(updatedList []*v1.VirtualService, obj *v1.VirtualService)
}

// OnAdd calls AddFunc if it's not nil.
func (r VirtualServiceEventHandlerFuncs) OnAdd(updatedList []*v1.VirtualService, obj *v1.VirtualService) {
	if r.AddFunc != nil {
		r.AddFunc(updatedList, obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r VirtualServiceEventHandlerFuncs) OnUpdate(updatedList []*v1.VirtualService, newObj *v1.VirtualService) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(updatedList, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r VirtualServiceEventHandlerFuncs) OnDelete(updatedList []*v1.VirtualService, obj *v1.VirtualService) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(updatedList, obj)
	}
}

// RoleEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// RoleEventHandler.
type RoleEventHandlerFuncs struct {
	AddFunc    func(updatedList []*v1.Role, obj *v1.Role)
	UpdateFunc func(updatedList []*v1.Role, newObj *v1.Role)
	DeleteFunc func(updatedList []*v1.Role, obj *v1.Role)
}

// OnAdd calls AddFunc if it's not nil.
func (r RoleEventHandlerFuncs) OnAdd(updatedList []*v1.Role, obj *v1.Role) {
	if r.AddFunc != nil {
		r.AddFunc(updatedList, obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r RoleEventHandlerFuncs) OnUpdate(updatedList []*v1.Role, newObj *v1.Role) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(updatedList, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r RoleEventHandlerFuncs) OnDelete(updatedList []*v1.Role, obj *v1.Role) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(updatedList, obj)
	}
}

// AttributeEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// AttributeEventHandler.
type AttributeEventHandlerFuncs struct {
	AddFunc    func(updatedList []*v1.Attribute, obj *v1.Attribute)
	UpdateFunc func(updatedList []*v1.Attribute, newObj *v1.Attribute)
	DeleteFunc func(updatedList []*v1.Attribute, obj *v1.Attribute)
}

// OnAdd calls AddFunc if it's not nil.
func (r AttributeEventHandlerFuncs) OnAdd(updatedList []*v1.Attribute, obj *v1.Attribute) {
	if r.AddFunc != nil {
		r.AddFunc(updatedList, obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r AttributeEventHandlerFuncs) OnUpdate(updatedList []*v1.Attribute, newObj *v1.Attribute) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(updatedList, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r AttributeEventHandlerFuncs) OnDelete(updatedList []*v1.Attribute, obj *v1.Attribute) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(updatedList, obj)
	}
}
