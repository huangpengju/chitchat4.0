package trace

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-logr/logr"
)

// Field 是一个键值对，它提供有关跟踪的其他详细信息。
type Field struct {
	Key   string
	Value interface{}
}

// traceItem 踪迹项目
type traceItem interface {
	// 当跟踪记录为完成时，时间返回。
	time() time.Time
	//writeItem将traceItem输出到缓冲区。如果stepThreshold为非零，则仅输出
	//traceItem，如果其持续时间超过stepThreshold。
	//输出的每一行都以格式化程序为前缀，以直观地缩进嵌套项。
	writeItem(b *bytes.Buffer, formatter string, startTime time.Time, stepThreshold *time.Duration)
}

type traceStep struct {
	stepTime time.Time
	msg      string
	fields   []Field
}

func (s traceStep) time() time.Time {
	return s.stepTime
}
func (s traceStep) writeItem(b *bytes.Buffer, formatter string, startTime time.Time, stepThreshold *time.Duration) {
	stepDuration := s.stepTime.Sub(startTime)

	if stepThreshold == nil || *stepThreshold == 0 || stepDuration >= *stepThreshold {
		b.WriteString(fmt.Sprintf("%s---", formatter))
		writeTraceItemSummary(b, s.msg, stepDuration, s.stepTime, s.fields)
	}
}

func writeTraceItemSummary(b *bytes.Buffer, msg string, totalTime time.Duration, startTime time.Time, fields []Field) {
	b.WriteString(fmt.Sprintf("%q ", msg))
	if len(fields) > 0 {
		writeFields(b, fields)
		b.WriteString(" ")
	}
	b.WriteString(fmt.Sprintf("%vms (%v)", durationToMilliseconds(totalTime), startTime.Format("15:04:00.000")))
}
func durationToMilliseconds(timeDuration time.Duration) int64 {
	return timeDuration.Nanoseconds() / 1e6
}
func writeFields(b *bytes.Buffer, l []Field) {
	for i, f := range l {
		b.WriteString(f.format())
		if i < len(l)-1 {
			b.WriteString(",")
		}
	}
}
func (f Field) format() string {
	return fmt.Sprintf("%s:%v", f.Key, f.Value)
}

// Trace 跟踪一组“步骤”，并允许我们记录特定的
// 如果花费的时间超过其在总允许时间中所占的份额，则执行步骤
type Trace struct {
	name        string
	fields      []Field
	threshold   *time.Duration
	startTime   time.Time
	endTime     *time.Time
	traceItems  []traceItem
	parentTrace *Trace
	logger      logr.Logger
}

func (t *Trace) time() time.Time {
	if t.endTime != nil {
		return *t.endTime
	}
	return t.startTime
}

func (t *Trace) writeItem(b *bytes.Buffer, formatter string, startTime time.Time, stepThreshold *time.Duration) {
	if t.durationIsWithinThreshold() || t.logger.V(4).Enabled() {
		b.WriteString(fmt.Sprintf("%v[", formatter))
		writeTraceItemSummary(b, t.name, t.TotalTime(), t.startTime, t.fields)
		if st := t.calculateStepThreshold(); st != nil {
			stepThreshold = st
		}
		t.writeTraceSteps(b, formatter+" ", stepThreshold)
		b.WriteString("]")
		return
	}
	for _, s := range t.traceItems {
		if nestedTrace, ok := s.(*Trace); ok {
			nestedTrace.writeItem(b, formatter, startTime, stepThreshold)
		}
	}
}

func (t *Trace) Step(msg string, fields ...Field) {
	if t.traceItems == nil {
		t.traceItems = make([]traceItem, 0, 6)
	}
	t.traceItems = append(t.traceItems, traceStep{stepTime: time.Now(), msg: msg, fields: fields})
}

func (t *Trace) TotalTime() time.Duration {
	return time.Since(t.startTime)
}
func (t *Trace) calculateStepThreshold() *time.Duration {
	if t.threshold == nil {
		return nil
	}
	lenTrace := len(t.traceItems) + 1
	traceThreshold := *t.threshold
	for _, s := range t.traceItems {
		netstedTrace, ok := s.(*Trace)
		if ok && netstedTrace.threshold != nil {
			traceThreshold = traceThreshold - *netstedTrace.threshold
			lenTrace--
		}
	}
	limitThreshold := *t.threshold / 4
	if traceThreshold < limitThreshold {
		traceThreshold = limitThreshold
		lenTrace = len(t.traceItems) + 1
	}
	stepThreshold := traceThreshold / time.Duration(lenTrace)
	return &stepThreshold
}
func (t *Trace) logTrace() {
	if t.durationIsWithinThreshold() {
		var buffer bytes.Buffer
		traceNum := rand.Int31()

		totalTime := t.endTime.Sub(t.startTime)
		buffer.WriteString(fmt.Sprintf("Trace[%d]:%q ", traceNum, t.name))
		if len(t.fields) > 0 {
			writeFields(&buffer, t.fields)
			buffer.WriteString(" ")
		}
		buffer.WriteString(fmt.Sprintf("(%v) (total time: %vms):", t.startTime.Format("02-Jan-2006 15:04:05.000"), totalTime.Milliseconds()))
		stepThreshold := t.calculateStepThreshold()
		t.writeTraceSteps(&buffer, fmt.Sprintf("\nTrace[%d]:", traceNum), stepThreshold)
		buffer.WriteString(fmt.Sprintf("\nTrace[%d]: [%v] [%v] END\n", traceNum, t.endTime.Sub(t.startTime), totalTime))

		t.logger.Info(buffer.String())
		return
	}
	for _, s := range t.traceItems {
		if nestedTrace, ok := s.(*Trace); ok {
			nestedTrace.logTrace()
		}
	}
}
func (t *Trace) writeTraceSteps(b *bytes.Buffer, formatter string, stepThreshold *time.Duration) {
	lastStepTime := t.startTime
	for _, stepOrTrace := range t.traceItems {
		stepOrTrace.writeItem(b, formatter, lastStepTime, stepThreshold)
		lastStepTime = stepOrTrace.time()
	}
}
func (t *Trace) durationIsWithinThreshold() bool {
	if t.endTime == nil {
		return false
	}
	return t.threshold == nil || *t.threshold == 0 || t.endTime.Sub(t.startTime) >= *t.threshold
}
