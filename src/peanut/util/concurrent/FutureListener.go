package concurrent

type FutureListener interface {
	OperationComplete(param interface{})
}
