/*
created by xffish@126.com
本示例可以解决
1 +
2 -
3 *
4 /
5 ()
5种运算符
并能解决多位数字的运算，比如123+321
算法核心：逆波兰式，数字栈，符号栈
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

//Stack 栈结构定义
type Stack struct {
	i    int
	data []string
}

//Push 入栈
func (s *Stack) push(num string) {
	s.data = append(s.data, num)
	s.i++
}

// Pop 出栈
func (s *Stack) pop() (temp string) {
	s.i--
	temp = s.data[s.i]
	s.data = s.data[:s.i]
	return
}

//合法运算符 必须要让乘法除法优先级一样，可以解决6/2*3先算谁的问题
var cal = map[string]int{")": 0, "(": 10, "+": 1, "-": 1, "*": 2, "/": 2}

//list1出栈两次，list2出栈一次，运算结果压回list1
func calcOnce(list1, list2 *Stack) {
	tempSymbol := list2.pop()
	tempNum1 := list1.pop()
	tempNum2 := list1.pop()
	//需要让后出栈的放在左操作数的位置上
	temp := easyCaculate(tempNum2, tempSymbol, tempNum1)
	list1.push(strconv.FormatFloat(temp, 'f', -1, 64))
}

func plus(x, y float64) float64 {
	return x + y
}

func minus(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

var operator = map[string]func(float64, float64) float64{"+": plus, "-": minus, "*": multiply, "/": divide}

func easyCaculate(left, oper, right string) (temp float64) {
	tempLeft, _ := strconv.ParseFloat(left, 64)
	tempRight, _ := strconv.ParseFloat(right, 64)
	temp = operator[oper](tempLeft, tempRight)
	return
}

func (s *Stack) String() (str string) {
	str = "长度为" + strconv.Itoa(s.i) + ","
	for i := 0; i < s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + s.data[i] + "]"
	}
	return
}

func caculate(source string) string {
	numStack := new(Stack)    //数字栈
	symbolStack := new(Stack) //符号栈

	//是否需要继续解析数字
	flag := false
	//扫描原始数学表达式
	for _, item := range source {
		strTemp := string(item)
		//如果这是数字
		if unicode.IsDigit(item) || item == '.' {
			temp := strTemp
			if flag {
				//说明数字尚未解析完毕
				temp = numStack.pop() + temp
			} else {
				flag = true
			}
			numStack.push(temp)
			continue
		}
		//故意用continue而不是一个if-elif，就是为了这里能在这里统一地设置flag而不是在每个分支里写flag=false
		//能到这里来说明一个数字解析完了
		flag = false

		value, ok := cal[strTemp]
		//如果符号栈为空且是合法运算符
		if symbolStack.i == 0 && ok {
			symbolStack.push(strTemp)
		} else if symbolStack.i != 0 && ok {
			//如果符号栈不为空且是合法运算符 这是程序的核心
			//得到当前符号栈栈顶元素
			topSymbolStack := symbolStack.data[symbolStack.i-1]
			//计算优先级
			priorityTop := cal[topSymbolStack]
			priorityItem := value
			//如果栈顶是“（”那就直接入栈
			if strTemp == ")" {
				if topSymbolStack == "(" {
					symbolStack.pop()
				} else {
					//倒序遍历
					//死循环
					for {
						//如果此时栈顶是“（”字符，那就到头了，简单出栈即可
						if symbolStack.data[symbolStack.i-1] == "(" {
							symbolStack.pop()
							break
						}
						calcOnce(numStack, symbolStack)
					}
				}

			} else if topSymbolStack == "(" {
				symbolStack.push(strTemp)
			} else if priorityItem > priorityTop {
				//如果当前元素比栈顶元素优先级高，那就入栈
				symbolStack.push(strTemp)
			} else {
				calcOnce(numStack, symbolStack)
				//然后把这个符号入栈
				symbolStack.push(strTemp)
			}
		}
	}
	//扫描完了还没算完，还要继续扫面现有的符号栈
	for {
		//如果栈空那就推出循环
		if symbolStack.i == 0 {
			break
		}
		calcOnce(numStack, symbolStack)
	}
	//现在numStack里的值是结果
	if numStack.i != 1 {
		fmt.Printf("不对呀，怎么还有%d个数", numStack.i)
		fmt.Println(numStack)
	}
	return numStack.data[0]
}

func main() {
	// sourcestr := "5+(10+22*34+4444)*2-6/2"
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	sourcestr, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	fmt.Println(caculate(sourcestr))

}
