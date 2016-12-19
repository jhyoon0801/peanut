package concurrent

type PeanutFuture interface {
	IsSuccess() bool
	IsCancellable() bool
	//Throwable cause()
	AddListener(FutureListener) Future
	AddListeners([]FutureListener) Future
	RemoveListener(FutureListener) Future
	RemoveListeners([]FutureListener) Future
	Sync() Future
	SyncUninterruptibly() Future
	Await() Future
	AwaitUninterruptibly() Future
	AwaitWithTimeout(long, ...time.Duration) bool
	AwaitUninterruptiblyWithTimeout(long, ...time.Duration) bool
	getNow() interface{}
}
