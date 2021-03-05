type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = make(map[int]string)

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.printf("错误码 %d 已经存在，请更换一个", code))
	}

	codes[code] = msg

    return &Error{
        code: code,
        msg: msg
    }
}

func (e *Error)Error() string {
    
}
