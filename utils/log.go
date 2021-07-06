package utils

var (
	logErr  func(err error)
	logInfo func(info string)
)

func InitLog(le func(error), li func(string)) {
	logErr = le
	logInfo = li
}

func Error(err error) {
	if logErr != nil {
		logErr(err)
	}
}

func Info(info string) {
	if logInfo != nil {
		logInfo(info)
	}
}
