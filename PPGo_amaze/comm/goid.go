package comm

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func ExtractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

// Parse the goid from runtime.Stack() output. Slow, but it works.
func GetSlow() int64 {
	var buf [64]byte
	return ExtractGID(buf[:runtime.Stack(buf[:], false)])
}
