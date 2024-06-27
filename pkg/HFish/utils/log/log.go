package log

import (
	"fmt"
	"os"
	"time"
)

func Pr(typex string, ip string, text string, a ...interface{}) {
	//fmt.Fprintln(gin.DefaultWriter, "["+typex+"] "+ip+" - ["+time.Now().Format("2006-01-02 15:04:05")+"] "+text+" ", a)
	fmt.Fprintln(os.Stdout, "["+typex+"] "+ip+" - ["+time.Now().Format("2006-01-02 15:04:05")+"] "+text+" ", a)

}
