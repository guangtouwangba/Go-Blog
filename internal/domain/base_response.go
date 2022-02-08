package domain

type IBaseResponse interface {
	GetStatus() int
	GetMessage() string
	GetData() interface{}
}
