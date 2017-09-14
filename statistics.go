// Copyright 漏 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	//以下的常量定义，其实是根据一个html来定义的
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	//引用了error定义的风格
	anError = `<p class="error">%s</p>`
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}
func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()        // Must be called before writing response
	fmt.Fprint(writer, pageTop, form) // 行头
	//如果出错，则会在打印出错误信息
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)             //取得排序后和计算好的数据
			fmt.Fprint(writer, formatStats(stats)) // 格式化后数据输出
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom) //行尾
}
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64                                                 //定义成切片
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 { //取得number项
		text := strings.Replace(slice[0], ",", " ", -1) // slice[0] ?
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil { //转换成浮点数
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x) //用内置函数append追加到切片中
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}
func formatStats(stats statistics) string {
	//格式化输出
	return fmt.Sprintf(`<table border="1">  
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers                           // go语言中的 返回值技巧
	sort.Float64s(stats.numbers)                      //排序包
	stats.mean = sum(numbers) / float64(len(numbers)) //平均值时，先转换成float64，不然有可以是整数相除，取整了
	stats.median = median(numbers)
	return stats
}
func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}
func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
