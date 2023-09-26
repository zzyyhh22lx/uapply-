package response

type ResCode int64

// CodeSuccess For each enumeration here, you need to add an error message to codeMsgMap
const (
	CodeSuccess ResCode = 1000 + iota
	CodeParamsInvalid
	CodeParamsFalse
	CodeSystemBusy
	CodeNotRoot
	CodeTokenInvalid
	CodeNoContent
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "Success",
	CodeParamsInvalid: "Invalid params",
	CodeParamsFalse:   "The account number or password is incorrect",
	CodeSystemBusy:    "The system is busy",
	CodeNotRoot:       "have not right",
	CodeTokenInvalid:  "token is invalid",
	CodeNoContent:     "record not found",
}

func (c ResCode) Msg() string {
	if s, ok := codeMsgMap[c]; ok {
		return s
	}
	return ""
}
