// @Title
// @Description
// @Author  2022/11/24
package enuls

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"net/http"
	"reflect"
	"time"
)

const blackListUrl = "https://assets.nabox.io/api/dict/enuls-blacklist"

var BlackList []string

var started bool = false

func initENULSScheduler() {
	if started {
		return
	}
	started = true
	// 每 60 秒钟时执行一次
	ticker := time.NewTicker(60 * time.Second) // 创建一个定时器
	go func() {                                // 用新协程去执行定时任务
		freshBlackList()
		for { // 用上一个死循环，不停地执行，否则只会执行一次
			select {
			case <-ticker.C: // 时间到了就会触发这个分支的执行，其实时间到了定时器会往ticker.C 这个 channel 中写一条数据，随后被 select 捕捉到channel中有数据可读，就读取channel数据，执行相应分支的语句
				freshBlackList()
			}
		}
	}()
}

func Start() {
	initENULSScheduler()
}

// 如果通过返回true
func Filter(addr common.Address) bool {
	if Find(BlackList, addr.Hex()) {
		return false
	}
	return true
}

func freshBlackList() {
	resp, err := http.Get(blackListUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	list := []map[string]interface{}{}
	if resp.StatusCode != 200 {
		fmt.Println("failed")
		return
	}
	json.Unmarshal(body, &list)
	for _, m := range list {
		str := m["dictValue"].(string)
		if !Find(BlackList, str) {
			BlackList = append(BlackList, str)
		}
	}

	if len(BlackList) > len(list) {
		//证明有地址取消屏蔽了，此时应该从黑名单中删除
		newBlackList := []string{}
		for _, m := range list {
			str := m["dictValue"].(string)
			newBlackList = append(newBlackList, str)
		}
		BlackList = newBlackList

	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if reflect.DeepEqual(item, val) {
			return true
		}
	}
	return false
}
