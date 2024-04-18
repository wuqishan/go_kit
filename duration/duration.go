package duration

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Duration struct {
	useFile  string
	useFunc  string
	line     int
	start    time.Time
	currTime time.Time
	dots     []*dot
	logged   bool
}

type dot struct {
	flag string
	dur  time.Duration
}

func NewDuration() *Duration {
	return (&Duration{
		start:    time.Now(),
		currTime: time.Now(),
		dots:     []*dot{},
	}).init()
}

func (d *Duration) init() *Duration {
	funcName := "unknown"
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			funcName = fn.Name()
		}
	}
	useFile := file
	funcNameArr := strings.Split(funcName, ".")
	funcName = funcNameArr[len(funcNameArr)-1]
	d.useFile = useFile
	d.useFunc = funcName
	d.line = line
	return d
}

func (d *Duration) Dot(pos ...string) *Duration {
	_, _, line, ok := runtime.Caller(1)
	p := d.useFunc
	if ok {
		p = fmt.Sprintf("%s_%d", d.useFunc, line)
	}
	if len(pos) > 0 && pos[0] != "" {
		p = pos[0]
	}
	d.dots = append(d.dots, &dot{
		flag: p,
		dur:  time.Since(d.currTime),
	})
	d.currTime = time.Now()
	return d
}

func (d *Duration) Info(ctx context.Context) string {
	if d.logged {
		return ""
	}
	res := make([]string, 0, len(d.dots)+1)
	res = append(res, fmt.Sprintf("Start: %s", d.start.Format(time.DateTime)))
	for _, v := range d.dots {
		res = append(res, fmt.Sprintf("%s(%v)", v.flag, v.dur))
	}
	d.logged = true
	return fmt.Sprintf("[%s:%d] ", d.useFile, d.line) + " [" + strings.Join(res, " - ") + "]"
}
