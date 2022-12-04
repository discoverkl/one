package ha

import (
	"log"
	"sync"
	"time"
)

// FailoverHandle provide failover feature for a given resource.
type FailoverHandle[T any] struct {
	FailOverTime time.Duration		// perform failover if current downtime is bigger than this value
	FailOverOperationCount int		// perform failover if Ret method is called certain times

	handle  *T
	factory func() (*T, error)
	close func(*T)					// optional close operation

	closeLock sync.Mutex
	getLock   sync.Mutex

	operationCount int
	failoverCount int
	failStartTime time.Time
}

func NewFailoverHandle[T any](factory func() (*T, error)) (*FailoverHandle[T], error) {
	return NewFailoverHandleCloser(factory, nil)
}

func NewFailoverHandleCloser[T any](factory func() (*T, error), close func(t *T)) (*FailoverHandle[T], error) {
	if factory == nil {
		panic("factory can't be nil")
	}
	np := &FailoverHandle[T]{
		factory: factory,
		close: close,
		FailOverTime: time.Second * 10,
		FailOverOperationCount: 0,
	}
	producer, err := factory()
	if err != nil {
		return nil, err
	}
	np.handle = producer
	return np, nil
}

// Get return the target object. User should not copy this value to another variable.
func (n *FailoverHandle[T]) Get() *T {
	n.getLock.Lock()
	defer n.getLock.Unlock()
	return n.handle
}

// Close is used when target object need cleanup.
func (n *FailoverHandle[T]) Close() error {
	n.closeLock.Lock()
	defer n.closeLock.Unlock()
	if n.close != nil {
		n.close(n.handle)
	}
	n.handle = nil // we can do this, only inside closeLock
	return nil
}

// Ret should be called after any method call on the target object. This is 
// necessary for computing downtime.
func (n *FailoverHandle[T]) Ret(ok bool) {
	n.operationCount++
	if ok {
		n.failStartTime = time.Time{}
	} else {
		if n.failStartTime.IsZero() {
			n.failStartTime = time.Now()
		}

		// failover by downtime
		if n.downtime() >= n.FailOverTime {
			n.Failover()
		}
	}
	// failover by count
	if n.FailOverOperationCount > 0 && n.operationCount >= n.FailOverOperationCount {
		n.operationCount = 0
		n.Failover()
	}
}

func (n *FailoverHandle[T]) downtime() time.Duration {
	if n.failStartTime.IsZero() {
		return 0
	}
	return time.Now().Sub(n.failStartTime)
}

// Failover create another target object using a factory method, call close callback
// passing the old target object if necessary.
func (n *FailoverHandle[T]) Failover() {
	n.failoverCount++
	log.Printf("failover %d for %T.", n.failoverCount, n.handle)
	producer, err := n.factory()
	if err != nil {
		log.Printf("failover failed: %v", err)
		return
	}
	n.getLock.Lock()
	defer n.getLock.Unlock()
	n.Close()
	n.handle = producer
}