// Package repl 这个包一般用于接受用户的输入，将用户的输入通过语法分析器lexer，转换成对应的词法单元token
package repl

import (
	"bufio"
	"fmt"
	"io"
	"monky/lexer"
	"monky/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// 从输入流中读取文本
	scanner := bufio.NewScanner(in)
	for {
		// 输出到>>到控制台
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		// 返回当前行的文本
		line := scanner.Text()
		l := lexer.New(line)
		//一直循环，直到;为止
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%v\n", tok)
		}
	}
}
