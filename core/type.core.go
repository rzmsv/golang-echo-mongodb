package core

type ErrorMsgType struct {
	status  int
	code    string
	message interface{}
}

type ResponseOkType struct {
	status  int
	code    string
	message interface{}
}
