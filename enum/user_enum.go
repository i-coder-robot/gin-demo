package enum

type ResponseType int

const (
	Operate_OK ResponseType = 200
	Operate_Fail ResponseType = 500
)

func (p ResponseType) String() string {
	switch (p) {
	case Operate_OK: return "Ok"
	case Operate_Fail: return "Fail"
	default:         return "UNKNOWN"
	}
}
