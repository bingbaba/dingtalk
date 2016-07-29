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

	//每天通告一次
	curtime := time.Now().Unix()
	DefalutNoticeConf.hourPoint = time.Unix(curtime, 0).Format("15:04")
	s := &CheckStatus{Status: "S"}
	if !s.Set("S", curtime) {
		t.Fatalf("expect true, but get false\n")
	}
	DefalutNoticeConf.hourPoint = "09:00"

	//第二次异常，需告警
	s = &CheckStatus{Status: "E", Tally: 1}
	if s.Set("E", curtime) == false {
		t.Fatalf("\"tally == 1\": should notice, but can't get the task of notice\n")
	}

	//已告警三次，且不超过30分钟，无需告警
	s.Tally = 3
	s.AlarmTime = time.Now().Unix() - 60
	if s.Set("E", curtime) == true {
		t.Fatalf("\"tally == 3\" and \"alarmtime < 30min\": should not notice, but get the task of notice\n")
	}

	//已告警三次，单已超过30分钟，需告警
	s.Tally = 3
	s.AlarmTime = time.Now().Unix() - 1801
	if s.Set("E", curtime) == false {
		t.Fatalf("\"tally == 3\" and \"alarmtime >= 30min\": should notice, but can't get the task of notice\n")
	}
}
