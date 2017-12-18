package tmpllib

import "github.com/KristinaEtc/slf"

var pwdCurr string = "KristinaEtc/tmpl"
var log = slf.WithContext(pwdCurr)

func PrintsTmpl(str string) {
	log.Info(str)
}
