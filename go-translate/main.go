package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode"

	"github.com/antonholmquist/jason"
	"github.com/astaxie/bat/httplib"
	"github.com/ttacon/chalk"
)

const (
	URL_dict      = "http://apis.baidu.com/apistore/tranlateservice/dictionary"
	URL_translate = "http://apis.baidu.com/apistore/tranlateservice/translate"
	APIKEY        = "4338b2d51fa348a355d789b994a58ea4"
)

const (
	EN = iota
	ZH
	JP
)

func getflag() (str string) {
	// var str string

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {

		if len(os.Args) == 1 {
			fmt.Printf("[Err] Usage: %s content", os.Args[0])
			os.Exit(1)
		}

		for _, s := range os.Args[1:] {
			str = str + s + " "
		}

	} else if info.Size() > 0 {
		reader := bufio.NewReader(os.Stdin)
		// match(pattern, reader)

		for {
			input, err := reader.ReadString('\n')
			if err != nil && err == io.EOF {
				break
			}
			str = str + string(input)
		}
	}

	return
}

func main() {
	str := getflag()
	getdict := httplib.Get(URL_translate)
	getdict.Header("apikey", APIKEY)

	getdict.Param("query", str)
	switch charIS(str) {
	case JP:
		getdict.Param("from", "jp")
		getdict.Param("to", "zh")
	case EN:
		getdict.Param("from", "en")
		getdict.Param("to", "zh")
	case ZH:
		getdict.Param("from", "zh")
		getdict.Param("to", "en")
	default:
		getdict.Param("from", "zh")
		getdict.Param("to", "en")
	}

	//resp, _ := getdict.Response()
	//fmt.Println("status:", resp.Status)

	// respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody))

	body, _ := getdict.Bytes()
	// s, _ := strconv.Unquote(string(body))
	data, _ := jason.NewObjectFromBytes(body)
	// errstr, _ := data.GetString("errMsg")
	transData, _ := data.GetObjectArray("retData", "trans_result")
	// fmt.Println(data)

	for k, va := range transData {
		s1, _ := va.GetString("src")
		s2, _ := va.GetString("dst")
		fmt.Printf("%ssrc :%s %s %s\n", chalk.Red, chalk.Reset, chalk.Cyan, s1)
		fmt.Printf("%strans-%d :%s %s\n", chalk.Blue, k, chalk.Reset, s2)
	}

}

// engine

var (
	regJp = regexp.MustCompile("([あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽアイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤ(イ)ユェヨラリルレロワウィ(ウ)ウェヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ])")
)

/*
判断字符串是否包含中文字符
*/
func charIS(str string) int {
	l := EN
	if regJp.MatchString(str) {
		return JP
	}
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) && l != JP {
			l = ZH
		}
		if unicode.Is(unicode.Scripts["Hiragana"], r) {
			l = JP
			return l
		}
		// for k, _ := range unicode.Scripts {
		// 	if unicode.Is(unicode.Scripts[k], r) {
		// 		fmt.Printf("%s is %s\n", string(r), k)
		// 	}
		// }
	}
	return l
}
