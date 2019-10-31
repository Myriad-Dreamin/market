package rbac

import (
	_ "fmt"
	"github.com/casbin/casbin"
	"sync"
	"time"
)

var initalLifeTime = 10
var tickDuration = time.Second

var maxDescriptorCount int = 15
var allocDescriptorCount int

func SetInitalLifeTime(i int) {
	initalLifeTime = i
}

func SetTickDuration(t time.Duration) {
	tickDuration = t
}

func SetMaxDescriptorCount(c int) {
	maxDescriptorCount = c
}

var fileMutex sync.Mutex

type EnforceDescriptor struct {
	lifetime int
	reference int
	*casbin.SyncedEnforcer
}

func newDescriptor(e *casbin.SyncedEnforcer) *EnforceDescriptor {
	return &EnforceDescriptor{
		lifetime: initalLifeTime,
		reference: 1,
		SyncedEnforcer: e,
	}
}

var cachedEnforcers = make(map[string]*EnforceDescriptor)

func Accquire(path string) *casbin.SyncedEnforcer {
	fileMutex.Lock()
	if e, ok := cachedEnforcers[path]; ok {
		e.lifetime += 2
		if e.lifetime > initalLifeTime {
			e.lifetime = initalLifeTime
		}
		e.reference ++
		fileMutex.Unlock()
		return e.SyncedEnforcer
	}

	if maxDescriptorCount <= allocDescriptorCount {
		fileMutex.Unlock()
		// is it ok ?
		return casbin.NewSyncedEnforcer(rbacModel, path)
	} else {
		allocDescriptorCount ++
		e := newDescriptor(casbin.NewSyncedEnforcer(rbacModel, path))
		cachedEnforcers[path] = e
		fileMutex.Unlock()
		// fmt.Println("alocation...", e)
		go func() {
			ticker := time.NewTicker(tickDuration)
			var e *EnforceDescriptor
			for _ = range ticker.C {
				fileMutex.Lock()
				e = cachedEnforcers[path]
				e.lifetime--
				// fmt.Println(e)
				if e.reference <= 0 && e.lifetime <= 0 {
					allocDescriptorCount--
					delete(cachedEnforcers, path)
					// fmt.Println("released", e)
					fileMutex.Unlock()
					ticker.Stop()
					break
				}
				fileMutex.Unlock()
			}
		}()

		return e.SyncedEnforcer
	}
}

func Release(path string) {
	fileMutex.Lock()
	if e, ok := cachedEnforcers[path]; ok {
		e.lifetime--
		e.reference--
		fileMutex.Unlock()
		return
	}
	fileMutex.Unlock()
}







