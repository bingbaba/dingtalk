package dingtalk

import (
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	corpid := os.Getenv("DT_ID")
	corpSecret := os.Getenv("DT_S")
	if corpid == "" || corpSecret == "" {
		return
	}

	client, client_err := NewDTalkClient(corpid, corpSecret)
	if client_err != nil {
		t.Fatalf("new client error:%s", client_err.Error())
	}

	//测试access_token
	access_token, at_err := client.GetAccessToken()
	if at_err != nil {
		t.Fatalf("get accesstoken error:%s", at_err)
	}
	fmt.Println("AccessToken:" + access_token)

	//测试用户
	users, user_err := client.GetAlluser()
	if user_err != nil {
		t.Fatalf("get all user failed:%s", user_err.Error())
	}

	for _, user := range users {
		fmt.Printf("%v\n", user)
	}

	//测试部门
	departments, depart_err := client.GetAllDeparts()
	if depart_err != nil {
		t.Fatalf("get all department failed:%s", depart_err.Error())
	}

	department := departments.GetDepartByName("仅供测试")
	if department == nil {
		t.Fatalf("now found the department \"仅供测试\"!")
	}

	fmt.Printf("%v\n", department)

	//发送消息
	message := NewPartyMessage("26829076", fmt.Sprintf("%d", department.ID), "亲爱的朋友们，我们来相聚！")
	send_err := client.SendToCompany(message)
	if send_err != nil {
		t.Fatalf("send err:%s", send_err.Error())
	}
}
