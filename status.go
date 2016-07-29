package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type CheckStatus struct {
	Bpfile    string `json:"-"`
	Status    string `json:"status"`
	AlarmTime int64  `json:"alarmtime"`
	StartTime int64  `json:"starttime"`
	Tally     int64  `json:"tally"`
}

type NoticeConf struct {
	hourPoint   string
	intervalMin int64
	exceptBegin int
	exceptEnd   int
	continuous  int64
	zero        bool
}

var (
	DefalutNoticeConf = InitNoticeConf(9, 30, 23, 8, 3)
)

func InitNoticeConf(h, i, b, e, c int) *NoticeConf {
	n := &NoticeConf{
		hourPoint:   fmt.Sprintf("%02d:00", h),
		intervalMin: int64(i*60 - 10),
		exceptBegin: b,
		exceptEnd:   e,
		continuous:  int64(c),
	}

	if n.exceptBegin > n.exceptEnd {
		n.zero = true
	}

	return n
}

func (status *CheckStatus) Set2(failed bool, ctime int64) bool {
	if failed {
		return status.Set("E", ctime)
	} else {
		return status.Set("S", ctime)
	}
}

func (status *CheckStatus) Set(s string, ctime int64) bool {

	//状态变化
	if status.Status != s {
		status.Tally = 1
		status.StartTime = ctime
		status.Status = s
		if s == "异常" || s == "F" || s == "E" {
			status.AlarmTime = ctime
		}

		return true
	} else {
		status.Tally++
		hour, _ := strconv.Atoi(time.Unix(ctime, 0).Format("15"))

		if time.Unix(ctime, 0).Format("15:04") == DefalutNoticeConf.hourPoint || //每天9点通告一次
			((s == "异常" || s == "F" || s == "E") &&
				(status.Tally <= DefalutNoticeConf.continuous || //异常且告警不超过3次
					(ctime-status.AlarmTime > DefalutNoticeConf.intervalMin &&
						(DefalutNoticeConf.zero && //跨越0点
							hour >= DefalutNoticeConf.exceptEnd &&
							hour <= DefalutNoticeConf.exceptBegin) ||
						(!DefalutNoticeConf.zero && //不跨越0点
							hour <= DefalutNoticeConf.exceptEnd &&
							hour >= DefalutNoticeConf.exceptBegin)))) {
			return true
		}
	}

	return false
}

func (status *CheckStatus) Save() error {
	f, f_err := os.OpenFile(status.Bpfile, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0755)
	if f_err != nil {
		return errors.New(fmt.Sprintf("save bp err:%v", f_err))
	}
	defer f.Close()

	bp_bytes, _ := json.Marshal(status)
	f.Write(bp_bytes)
	f.Sync()

	return nil
}

func (status *CheckStatus) Load() error {
	fi, file_err := os.Open(status.Bpfile)
	if file_err != nil {
		return file_err
	}
	defer fi.Close()

	bp_bytes, read_err := ioutil.ReadAll(fi)
	if read_err != nil {
		return read_err
	}

	err := json.Unmarshal(bp_bytes, status)
	return err
}
