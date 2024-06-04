package utlog

type IElastic interface {
	SendLog(app string, message interface{}) error
}
