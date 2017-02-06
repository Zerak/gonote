package main

import (
	"bytes"
	"fmt"
	//"strconv"
	"strings"
	"time"

	"live_server/common"
	"strconv"
)

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // temporary bytes.Buffer
			b := bytes.Buffer{}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 4: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

		if i == n-1 {
			if index == 4 { // for stable bytes.Buffer
				s = buf.String()
			}
			fmt.Println(len(s)) // consume s to avoid compiler optimization
		}
	}
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}

func main() {
	//k := 5
	//d := [5]time.Duration{}
	//for i := 0; i < k; i++ {
	//	d[i] = benchmarkStringFunction(10000, i)
	//}
	//
	//for i := 0; i < k-1; i++ {
	//	fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	//}

	// process br
	contents := "[url:rtmp://pushws.hifun.mobi/live/10016589_zz3VJVUVbKdev?type=1][lIP:192.168.2.133][rIP:111.202.74.130][cT:229(1)][fA:196 fV:255][fAK=0 fVK=0][aCT:0 vCT:0][aF:22 vF:10][aIT:1464289][vIT:1464099][aDT:0 vDT:0][aPT:0 vPT:0][aC:1 vC:0][dA:0 dV:0][bR:290k/s][AVS:190][eSPSPPS:1][aBC:0 vBC:0][t:0][sys am:0, bg:0, pause:0, cameradev:1, micphonedev:1]"
	//contents := ""
	prefix := "bR:"
	brIdx := strings.Index(contents, "bR:")
	fmt.Printf("index :%v\n", brIdx)

	ksidx := strings.Index(contents, "k/s")
	fmt.Printf("k/s index:%v\n", ksidx)

	str := common.SubString(contents, brIdx+len(prefix), ksidx-brIdx-len(prefix))
	fmt.Printf("str :%v\n", str)

	spd, _ := strconv.Atoi(str)
	fmt.Printf("spd:%v\n", spd)

	// process liveid
	var liveStr string
	//liveid := "liveid_10016589_zz3VJVUVbKdev?type=1"
	liveid := "liveid_10016589_zz3VJVUVbKdev"
	idIdx := strings.Index(liveid, "?type")
	if idIdx != -1 {
		liveStr = common.SubString(liveid, 0, idIdx)
	} else {
		liveStr = liveid
	}
	fmt.Printf("liveid:%v\n", liveStr)

	str2 := "192.168.1.200:10080"
	if strings.Contains(str2, ":") {
		str2 = common.SubString(str2, 0, strings.Index(str2, ":"))
	}
	fmt.Printf("ip:%v\n", str2)

	timestr := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Printf("time fmt:%v\n", timestr)

	by := `[123 34 116 121 112 101 34 58 34 102 108 118 34 44 34 115 116 114 101 97 109 95 101 118 101 110 116 34 58 34 102 108 118 34 44 34 111 114 105 95 117 114 108 34 58 34 114 116 109 112 58 47 47 112 117 115 104 121 102 46 104 105 102 117 110 46 109 111 98 105 47 108 105 118 101 47 49 48 48 49 54 53 57 51 95 66 106 105 111 56 97 104 52 79 88 100 101 118 34 44 34 115 116 97 114 116 95 116 105 109 101 34 58 49 52 55 48 50 57 54 55 48 54 44 34 115 116 111 112 95 116 105 109 101 34 58 48 44 34 100 117 114 97 116 105 111 110 34 58 48 44 34 99 100 110 95 117 114 108 34 58 34 104 116 116 112 58 47 47 34 125]`
	strby := string(by)
	fmt.Printf("by:%v\n", strby)

}
