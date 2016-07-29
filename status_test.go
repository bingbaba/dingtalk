package dingtalk

import (
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	DefalutNoticeConf = InitNoticeConf(9, 30, 23, 8, 3)
	if DefalutNoticeConf.hourPoint != "09:00" {
		t.Fatalf("expect \"09:00\", but get \"%s\"\n", DefalutNoticeConf.hourPoint)
	}

	curtime := time.Now().Unix()
	DefalutNoticeConf.hourPoint = time.Unix(curtime, 0).Format("15:04")
	s := &CheckStatus{Status: "S"}
	if !s.Set("S", curtime) {
		t.Fatalf("expect true, but get false\n")
	}

}
