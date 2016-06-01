package dingtalk

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client, client_err := NewDTalkClient(`ding52450ad5a38b178f`, `HAN6AFVmmgE-R9df-j0ITn00l7lFd84NCRVw5fO6T6hOct_myavo-_wmkMIpyzHj`)
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

	department := departments.GetDepartByName("产品研发")
	if department == nil {
		t.Fatalf("now found the department \"产品研发\"!")
	}

	fmt.Printf("%v\n", department)

	//发送消息
	message := NewPartyMessage("26829076", fmt.Sprintf("%d", department.ID), "亲爱的朋友们，我们来相聚！")
	send_err := client.SendToCompany(message)
	if send_err != nil {
		t.Fatalf("send err:%s", send_err.Error())
	}
}
