package common

import (
	"strings"
)

func init() {
	F나를_위한_문구()
	
	c참 = &sC참거짓{true}
	c거짓 = &sC참거짓{false}

	F_TODO("문자열 후보값에 한자도 포함시킬 것.")
	문자열_후보값 := "1234567890" +
		"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"~!@#$%^&*()_+|';:/?.,<>`" +
		"가나다라마바사하자차카타파하"

	문자열_후보값_모음 = strings.Split(문자열_후보값, "")
}

const P시점_포맷 string = "2006-01-02 15:04:05 (MST) Mon -0700"
const P일자_포맷 string = "2006-01-02"
const P차이_한도 string = "1/1000000000000000000000000000000000000"
const asc코드_0 uint8 = uint8(48)
const asc코드_소숫점 uint8 = uint8(46)


// 통화
type P통화 int

const (
	KRW P통화 = iota
	USD
	CNY
	EUR
)

var 통화종류_문자열_모음 = [...]string{ "KRW", "USD", "CNY", "EUR"}

func (p P통화) String() string { return 통화종류_문자열_모음[p] }
