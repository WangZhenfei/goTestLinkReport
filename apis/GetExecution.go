package apis

import (
	"encoding/json"
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/alphazero/Go-Redis"
	"github.com/royburns/goTestLinkReport/models"
	// "strconv"
	// "strings"
)

func (this *ApiController) GetLastExecution() {

	expiration := int64(60 * 60)
	spec := redis.DefaultSpec().Db(0)
	client, err := redis.NewSynchClientWithSpec(spec)
	if err != nil {
		fmt.Println("Failed to create the client: ", err)
		return
	}

	fmt.Println("...")
	// var executions []*V_testlink_testexecution_tree
	executions := []models.V_testlink_testexecution_tree{}
	testplan := this.Input().Get("testplan")
	fmt.Println(testplan)
	if testplan != "" {
		key := testplan
		value, err := client.Get(key)
		if err != nil {
			fmt.Println("Failed to get value: ", err)
			return
		}

		if value == nil {
			fmt.Println("The plan is not exists. We will query it from mysql and then store them in redis.")

			// Get Executions
			maxTPNum := 10000
			testexecution_where, err := models.GetAllExecutionsWhere(testplan)
			tpCount := len(testexecution_where)
			if err != nil || tpCount > maxTPNum {
				testplan = ""
			}

			executions = testexecution_where
			value, err := json.Marshal(testexecution_where)
			if err != nil {
				fmt.Println("Failed to marshal: ", err)
			}
			client.Set(key, value)
			client.Expire(key, expiration)
		} else {
			fmt.Println("The plan is exists. We will unmarshal them.")
			// var temp planinfo
			err := json.Unmarshal(value, &executions)
			if err != nil {
				fmt.Println("Failed to unmarshal: ", err)
				return
			}
		}
	}

	this.Data["json"] = &executions
	this.ServeJson()

	// this.TplNames = "report.tpl"
}
