package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type taxinfo struct {
	gz float64 `json:"gz"`
	sb float64 `json:"sb"`
	zx float64 `json:"zx"`
}

//计算个人所得税
func taxcal(w http.ResponseWriter, r *http.Request) {
	//读取body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println(string(body))

	//解析json
	t := taxinfo{}
	fmt.Println(t)
	json.Unmarshal(body, &t)
	fmt.Println(t)
	ysje := t.gz - t.sb - t.zx - 5000.00

	//计算个税
	var total float64
	var ljysje, sl, kcs, gsje [12]float64
	for i := 0; i < 12; i++ {
		ljysje[i] = ysje * float64(i+1)
		sl[i], kcs[i] = slb(ljysje[i])
		gsje[i] = ljysje[i]*sl[i] - kcs[i] - total
		total += gsje[i]
		buf := fmt.Sprintf("[%02d]月份\t累计应税金额[%0.2f]\t税率[%0.2f]\t扣除数[%0.2f]\t应缴个税[%0.2f]\n",
			i+1, ljysje[i], sl[i], kcs[i], gsje[i])
		io.WriteString(w, buf)
	}
}

//税率表
func slb(ysje float64) (sl, kcs float64) {
	switch {
	case ysje <= 36000:
		sl, kcs = 0.03, 0
	case ysje <= 144000:
		sl, kcs = 0.1, 2520
	case ysje <= 300000:
		sl, kcs = 0.2, 16920
	case ysje <= 420000:
		sl, kcs = 0.25, 31920
	case ysje <= 660000:
		sl, kcs = 0.3, 52920
	case ysje <= 960000:
		sl, kcs = 0.35, 85920
	default:
		sl, kcs = 0.45, 181920
	}
	return
}

func main() {
	http.HandleFunc("/", taxcal)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务启动失败")
	}
}
