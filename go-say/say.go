package main

import (
	"os/exec"
	"regexp"
	"unicode"
)

const (
	EN = iota
	ZH
	JP
)

var (
	regJp = regexp.MustCompile("([あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽアイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤ(イ)ユェヨラリルレロワウィ(ウ)ウェヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ])")
)

func Say(str string) {
	switch charIS(str) {
	case ZH:
		exec.Command("say", "-v", "ting", str).Run()
		break
	case JP:
		exec.Command("say", "-v", "otoya", str).Run()
		break
	default:
		exec.Command("say", str).Run()
		break
	}
}

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
