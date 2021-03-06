//练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，
// 中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，
// 根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var (
	testI   int
	depth   int
	problem string = "tto"
)

func main() {
	s := "<p id='tto'>ddd<a>aaaaa</a></p><p>cccc</p>"
	read := strings.NewReader(s)
	doc, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	visit(doc, start, end)
}
func visit(n *html.Node, start, end func(n *html.Node) bool) {
	if er := start(n); !er {
		return
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(i, start, end)
	}
	if er := end(n); !er {
		return
	}
}

func start(n *html.Node) bool {
	switch n.Type {
	case html.ElementNode:
		var str string
		for _, v := range n.Attr {
			str = fmt.Sprintf("<%s %s='%s'></%s>", n.Data, v.Key, v.Val, n.Data)
		}
		for _, v := range n.Attr {
			if v.Key == "id" {
				if v.Val == problem {
					fmt.Println(str)
					os.Exit(0)
				}

			}
		}

		fmt.Printf("%*s<%s>%d\n", depth*2, "", n.Data, depth)
		depth++
	case html.TextNode:
		fmt.Printf("%*s%s%d\n", depth*2, "", n.Data, depth)
	case html.CommentNode:
		fmt.Println(n.Data)
	case html.DoctypeNode:
		fmt.Println(n.Data)
	case html.ErrorNode:
		fmt.Println(n.Data)
	case html.DocumentNode:
		fmt.Println(n.Data)
	}
	return true
}
func end(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>%d\n", depth*2, "", n.Data, depth)
	}
	return true
}

//   g
//go run main.go
//
//{0xc0000a8000 0xc0000a80e0 0xc0000a8150 <nil> <nil> 3 html html  []}
//<html>0
//{0xc0000a8070 <nil> <nil> <nil> 0xc0000a8150 3 head head  []}
//  <head>1
//  </head>1
//{0xc0000a8070 0xc0000a81c0 0xc0000a8380 0xc0000a80e0 <nil> 3 body body  []}
//  <body>1
//{0xc0000a8150 0xc0000a8230 0xc0000a82a0 <nil> 0xc0000a8380 3 p p  []}
//    <p>2
//{0xc0000a81c0 <nil> <nil> <nil> 0xc0000a82a0 1  ddd  []}
//      ddd3
//{0xc0000a81c0 0xc0000a8310 0xc0000a8310 0xc0000a8230 <nil> 3 a a  []}
//      <a>3
//{0xc0000a82a0 <nil> <nil> <nil> <nil> 1  aaaaa  []}
//        aaaaa4
//      </a>3
//    </p>2
//{0xc0000a8150 0xc0000a83f0 0xc0000a83f0 0xc0000a81c0 <nil> 3 p p  []}
//    <p>2
//{0xc0000a8380 <nil> <nil> <nil> <nil> 1  cccc  []}
//      cccc3
//    </p>2
//  </body>1
//</html>0

//根据解析结果，如果是同一个级别的数据，其实他们是同时运行同时生成了数据，但是到了输出的时候确实先让一个数据输出然后子数据输出，然后在兄弟
//数据，其实兄弟数据是同时生成的这也就说明了为什么同一层的p的depth值是一样的。只是它们的子不一样。所以我们可以这样理解
// 0xc0000a8150  0xc0000a8150  他们一样是8150 证明 他们的地址是一样的证明他们同时运行。只是不是同时输出而已。真的是nice啊。
// 1 同一级的是同时生成的
//1 同一级的虽然是同时生成的但是却是按照先输出再输出子的顺序往外输出的。
//根据打印结果，我们可以直接想象成，虽然实质是 一层一层的生成，但是打印结果却是从上到下依次输出。

//比如某一层栈内同时生成了是 head和body 那么推出来的时候 也是 head和body。
