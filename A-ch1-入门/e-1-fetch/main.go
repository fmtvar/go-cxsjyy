package main

import (
	"fmt"
	_ "io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		fmt.Println("resp=", resp.StatusCode)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
	}
}

//1.7课后练习题

//函数io.Copy(dst,src)从src读,并且写入dst.使用它代替ioutil.ReadAll来复制响应内容到os.StdOut,这样不需要装下整个数据流的缓冲区.确保检查io.Copy返回的错误结果
//io.Copy代替ioutil.ReadAll更省内存，同时也可以把它看做是一个同步的操作，返回了最终传输数据的大小

// func main() {
// 	start := time.Now()
// 	ch := make(chan string)
// 	for _, url := range os.Args[1:] {
// 		go fetch(url, ch)
// 	}
// 	for range os.Args[1:] {
// 		fmt.Println(<-ch)
// 	}
// 	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
// }
// func fetch(url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprint(err)
// 		return
// 	}
// 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
// }

//1.8 课后习题
//修改fetch程序添加一个http://前缀(假如该URL参数缺失协议前缀).可能会用到strings.HasPrefix
// func TestAddPrefix(t *testing.T) {
// 	inputURL := "z.cn"
// 	if ! strings.HasPrefix(inputURL, "http") {
// 		inputURL = "http://" + inputURL
// 	}
// 	fmt.Println(inputURL)
// }

//1.9 课后习题
// 修改fetch来输出http状态码,可以在resp.Status中找到它
// func TestOutputStatusCode(t *testing.T) {
// 	resp, _ := http.Get("http://z.cn")
// 	fmt.Println("StatusCode: ", resp.StatusCode)
// 	io.Copy(os.Stdout, resp.Body)
// 	resp.Body.Close()
// }
