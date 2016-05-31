package dingtalk

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	client, client_err := NewDTalkClient(`ding52450ad5a38b178f`, `HAN6AFVmmgE-R9df-j0ITn00l7lFd84NCRVw5fO6T6hOct_myavo-_wmkMIpyzHj`)
	if client_err != nil {
		t.Fatalf("new client error:%s", client_err.Error())
	}

	access_token, at_err := client.GetAccessToken()
	if at_err != nil {
		t.Fatalf("get accesstoken error:%s", at_err)
	}
	fmt.Println("AccessToken:" + access_token)

	users, user_err := client.GetAlluser()
	if user_err != nil {
		t.Fatalf("get all user failed:%s", user_err.Error())
	}

	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}
