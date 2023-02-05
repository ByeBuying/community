package recovery

import (
 "bytes"
 "community/model"
 "community/protocol"
 "community/util"
 "fmt"
 "github.com/gin-gonic/gin"
 "github.com/golang-jwt/jwt"
 "go-common/klay/elog"
 "io/ioutil"
 "net/http"
 "runtime"
 "strings"
 "time"
)

var (
 dunno     = []byte("???")
 centerDot = []byte("·")
 dot       = []byte(".")
 slash     = []byte("/")
)

const (
 AppName = "Gateway"
)

func GinRecovery(mod string) gin.HandlerFunc {
 return func(c *gin.Context) {
	defer func() {
	 if err := recover(); err != nil {
		url := c.Request.URL.String()
		method := c.Request.Method
		device := c.GetHeader("device")
		version := c.GetHeader("version")
		uuid := GetUUID(c)
		trace := fmt.Sprintf("panic: %+v \n\n%s", err, string(stack(3)))

		elog.Error(trace)

		message := util.NewMessage().
		 SetZone(mod).
		 SetMessageType(util.MessageTypePanic).
		 SetTitle("Panic Report").
		 AddKeyValueWidget("App", AppName).
		 AddKeyValueWidget("Zone", mod).
		 EndSection().
		 AddKeyValueWidget("URL", fmt.Sprintf("[%s] %s", method, url)).
		 AddKeyValueWidget("Device", device).
		 AddKeyValueWidget("Version", version).
		 AddKeyValueWidget("UUID", uuid)

		if b, err := ioutil.ReadAll(c.Request.Body); err == nil && len(b) > 0 {
		 message.EndSection().
			AddKeyValueWidget("Body", string(b))
		}

		message.EndSection().
		 AddTextParagraphWidget(trace)

		message.SendMessage()

		c.JSON(http.StatusInternalServerError, protocol.NewRespHeader(protocol.PanicError, "internal server error"))
		c.Abort()
	 }
	}()
	c.Next()
 }
}

func GetUUID(c *gin.Context) (uuid string) {
 token := ""
 bearToken := c.Request.Header.Get("Authorization")
 slice := strings.Split(bearToken, " ") //"bearer" 제거
 if len(slice) > 1 {
	token = slice[1]
 } else {
	return
 }
 claims := &model.UserClaims{}
 if _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
	return nil, nil
 }); err != nil {
	if strings.Contains(err.Error(), "token is expired by") {
	 uuid = "token is expired by"
	 return
	} else {
	 uuid = claims.UserID
	 return
	}
 }

 return
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
 buf := new(bytes.Buffer) // the returned data
 // As we loop, we open files and read them. These variables record the currently
 // loaded file.
 var lines [][]byte
 var lastFile string
 for i := skip; ; i++ { // Skip the expected number of frames
	pc, file, line, ok := runtime.Caller(i)
	if !ok {
	 break
	}
	// Print this much at least.  If we can't find the source, it won't show.
	fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	if file != lastFile {
	 data, err := ioutil.ReadFile(file)
	 if err != nil {
		continue
	 }
	 lines = bytes.Split(data, []byte{'\n'})
	 lastFile = file
	}
	fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
 }
 return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
 n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
 if n < 0 || n >= len(lines) {
	return dunno
 }
 return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
 fn := runtime.FuncForPC(pc)
 if fn == nil {
	return dunno
 }
 name := []byte(fn.Name())
 // The name includes the path name to the package, which is unnecessary
 // since the file name is already included.  Plus, it has center dots.
 // That is, we see
 //	runtime/debug.*T·ptrmethod
 // and want
 //	*T.ptrmethod
 // Also the package path might contains dot (e.g. code.google.com/...),
 // so first eliminate the path prefix
 if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
	name = name[lastSlash+1:]
 }
 if period := bytes.Index(name, dot); period >= 0 {
	name = name[period+1:]
 }
 name = bytes.Replace(name, centerDot, dot, -1)
 return name
}

func timeFormat(t time.Time) string {
 timeString := t.Format("2006/01/02 - 15:04:05")
 return timeString
}
