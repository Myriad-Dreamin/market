//go:generate stringer -type=MethodType
package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type MethodType uint8

const (
	Any MethodType = iota
	GET
	POST
	DELETE
	PATCH
	PUT
	OPTIONS
	HEAD
	Static
	StaticFile
	StaticFS
)

type LeafRouter struct {
	leafType        MethodType
	rawRelativePath string
	option          interface{}
	middleware      []HandlerFunc
	afterHandlers   []HandlerFunc
}

func (l *LeafRouter) Use(handlers ...HandlerFunc) *LeafRouter {
	l.middleware = append(l.middleware, handlers...)
	return l
}

func (l *LeafRouter) Build(g IRouter) {

	switch l.leafType {
	case GET:
		g.GET(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case POST:
		g.POST(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case DELETE:
		g.DELETE(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case PATCH:
		g.PATCH(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case PUT:
		g.PUT(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case OPTIONS:
		g.OPTIONS(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case HEAD:
		g.HEAD(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case Static:
		g.Group("/", l.middleware...).Static(l.rawRelativePath, l.option.(string))
	case StaticFile:
		g.Group("/", l.middleware...).StaticFile(l.rawRelativePath, l.option.(string))
	case StaticFS:
		g.Group("/", l.middleware...).StaticFS(l.rawRelativePath, l.option.(http.FileSystem))
	default:
		panic(fmt.Errorf("bad leaf type %v", l.leafType))
	}
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func (l *LeafRouter) RouteInfo(path string) RouteInfo {
	return RouteInfo{
		Method:      l.leafType.String(),
		Path:        strings.ReplaceAll(path, "\\", "/"),
		Handler:     nameOfFunction(l.Last()),
		HandlerFunc: l.Last(),
	}
}

func (l *LeafRouter) Last() HandlerFunc {
	if len(l.afterHandlers) == 0 {
		return nil
	}
	return l.afterHandlers[len(l.afterHandlers)-1]
}

func NewLeafRouter(leafType MethodType, path string, handlers ...HandlerFunc) *LeafRouter {
	return &LeafRouter{
		leafType:        leafType,
		rawRelativePath: path,
		afterHandlers:   handlers,
	}
}

func NewLeafRouterOp(leafType MethodType, path string, option interface{}, handlers ...HandlerFunc) *LeafRouter {
	switch leafType {
	case Static:
		if _, ok := option.(string); !ok {
			panic(fmt.Errorf("must serve as router.Static(path, root string)"))
		}
	case StaticFile:
		if _, ok := option.(string); !ok {
			panic(fmt.Errorf("must serve as router.Static(path, filepath string)"))
		}
	case StaticFS:
		if _, ok := option.(http.FileSystem); !ok {
			panic(fmt.Errorf("must serve as router.Static(path, filesystem http.FileSystem)"))
		}
	}
	return &LeafRouter{
		leafType:        leafType,
		rawRelativePath: path,
		option:          option,
		middleware:      handlers,
	}
}

type Router struct {
	handlers           []HandlerFunc
	SubRouter          map[string]*Router
	SameLevelSubRouter map[string]*Router
	Leafs              map[string]*LeafRouter
}

func NewRouterGroup(handlers ...HandlerFunc) *Router {
	return &Router{
		handlers: handlers,
	}
}

func (c *Router) Build(g IRouter) {
	g.Use(c.handlers...)
	for k, sr := range c.SubRouter {
		sr.Build(g.Group(k))
	}
	for _, sr := range c.SameLevelSubRouter {
		sr.Build(g.Group(""))
	}
	for _, leaf := range c.Leafs {
		leaf.Build(g)
	}
}

func (c *Router) Group(relativePath string, handlers ...HandlerFunc) (r *Router) {
	if c.SubRouter == nil {
		c.SubRouter = make(map[string]*Router)
	} else if _, ok := c.SubRouter[relativePath]; ok {
		panic(fmt.Errorf("must conflict path %v", relativePath))
	}
	r = NewRouterGroup(handlers...)
	c.SubRouter[relativePath] = r
	return
}

func (c *Router) Extend(name string, handlers ...HandlerFunc) (r *Router) {
	if c.SameLevelSubRouter == nil {
		c.SameLevelSubRouter = make(map[string]*Router)
	} else if _, ok := c.SameLevelSubRouter[name]; ok {
		panic(fmt.Errorf("must conflict name %v", name))
	}
	r = NewRouterGroup(handlers...)
	c.SameLevelSubRouter[name] = r
	return
}

func (c *Router) Use(handlers ...HandlerFunc) *Router {
	c.handlers = append(c.handlers, handlers...)
	return c
}

func (c *Router) leaf(leafType MethodType, relativePath string, handlers ...HandlerFunc) (r *LeafRouter) {
	if c.Leafs == nil {
		c.Leafs = make(map[string]*LeafRouter)
	} else if _, ok := c.Leafs[relativePath+string(leafType)]; ok {
		panic(fmt.Errorf("must conflict path %v", relativePath+string(leafType)))
	}
	r = NewLeafRouter(leafType, relativePath, handlers...)
	c.Leafs[relativePath+string(leafType)] = r
	return
}

func (c *Router) leafOp(leafType MethodType, relativePath string, op interface{}, handlers ...HandlerFunc) (r *LeafRouter) {
	if c.Leafs == nil {
		c.Leafs = make(map[string]*LeafRouter)
	} else if _, ok := c.Leafs[relativePath]; ok {
		panic(fmt.Errorf("must conflict path %v", relativePath))
	}
	r = NewLeafRouterOp(leafType, relativePath, op, handlers...)
	c.Leafs[relativePath] = r
	return
}

func (c *Router) GET(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(GET, relativePath, handlers...)
}

func (c *Router) POST(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(POST, relativePath, handlers...)
}

func (c *Router) DELETE(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(DELETE, relativePath, handlers...)
}

func (c *Router) PATCH(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(PATCH, relativePath, handlers...)
}

func (c *Router) PUT(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(PUT, relativePath, handlers...)
}

func (c *Router) OPTIONS(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(OPTIONS, relativePath, handlers...)
}

func (c *Router) HEAD(relativePath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leaf(HEAD, relativePath, handlers...)
}

func (c *Router) StaticFile(relativePath string, filepath string, handlers ...HandlerFunc) *LeafRouter {
	return c.leafOp(HEAD, relativePath, filepath, handlers...)
}

func (c *Router) Static(relativePath string, root string, handlers ...HandlerFunc) *LeafRouter {
	return c.leafOp(HEAD, relativePath, root, handlers...)
}

func (c *Router) StaticFS(relativePath string, fs http.FileSystem, handlers ...HandlerFunc) *LeafRouter {
	return c.leafOp(HEAD, relativePath, fs, handlers...)
}

// RouteInfo represents a request route's specification which contains method and path and its handler.
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

// RoutesInfo defines a RouteInfo array.
type RoutesInfo []RouteInfo

// Routes returns a slice of registered routes, including some useful information, such as:
// the http method, path and the handler name.
func (engine *Router) routes(path string, routes RoutesInfo) RoutesInfo {
	for spath, router := range engine.SubRouter {
		routes = router.routes(filepath.Join(path, spath), routes)
	}
	for _, router := range engine.SameLevelSubRouter {
		routes = router.routes(path, routes)
	}
	for spath, leaf := range engine.Leafs {
		routes = append(routes, leaf.RouteInfo(filepath.Join(path, spath[:len(spath)-1])))
	}
	return routes
}

func (engine *Router) Routes() (routes RoutesInfo) {
	return engine.routes("", nil)
}
