package common

import (
	"bytes"
	//"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
	//"time"
)

func init() {
	후보값_문자열 := "1234567890" +
			"abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"~!@#$%^&*()_+|';:/?.,<>`" +
			"가나다라마바사하자차카타파하" // 한자도 포함시킬 것.
	
	문자열_후보값_모음 = strings.Split(후보값_문자열, "")
	
	c참 = &sC참거짓{true}
	c거짓 = &sC참거짓{false}
}

var 문자열_후보값_모음 []string

var c참, c거짓 *sC참거짓

// 상수형이 immutable 하기 위해서는 생성할 때 입력되는 참조형 값이
// 적절하게 복사되는 것을 보장해야 함.
// 이를 위해서는 new()를 사용해서 상수형 자료형을 생성할 수 없도록 확실히 해야함.
// 상수형 구조체를 외부에 공개하지 않으면, new()를 통해서 생성할 수 없게 됨.
// 그리고, 적절한 초기화 및 값 복사를 거치는 NC 생성자를 통해서만 생성할 수 있도록 함.
// 이렇게 생성된 상수형 구조체는 외부에 공개된 인터페이스를 통해서 사용하면 됨.
// (Go언어에 생성자 기능이 없으니 이렇게라도 구현할 수 밖에 없음.
//  Go언어 특유의 단순함과 빠른 컴파일 속도를 위해서는 감수해야 하는 부분임.)
// 상수형 인터페이스를 구현하면서도 mutable한 사용자 정의 자료형을 걸러내는 작업은
// 가칭, F공유해도_안전함() 혹은 F고정값임()을 통해서 해결함.

func NC정수(값 int64) C정수 { return &sC정수64{값} }
type sC정수64 struct{ 값 int64 }
func (s *sC정수64) 상수형임()  {}
func (s *sC정수64) G값() int64 { return s.값 }
func (s *sC정수64) G정수() int64 { return s.값 }
func (s *sC정수64) G큰정수() *big.Int { return F정수2큰정수(s.값) }
func (s *sC정수64) G실수() float64 { return float64(s.값) }
func (s *sC정수64) G정밀수() *big.Rat { return F정수2정밀수(s.값) }
func (s *sC정수64) G문자열() string { return F정수2문자열(s.값) }
func (s *sC정수64) String() string { return F정수2문자열(s.값) }
func (s *sC정수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	값 := 임의값_생성기.Int63() // 0 혹은 양의 정수.
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = 값 * -1
	}
	
	return reflect.ValueOf(NC정수(값))
}

func NC부호없는_정수(값 uint64) C부호없는_정수 { return &sC부호없는_정수64{값} }
type sC부호없는_정수64 struct { 값 uint64 }
func (s *sC부호없는_정수64) 상수형임()          {}
func (s *sC부호없는_정수64) G값() uint64     { return s.값 }
func (s *sC부호없는_정수64) G정수() int64 { return int64(s.값) }
func (s *sC부호없는_정수64) G큰정수() *big.Int { return F부호없는_정수2큰정수(s.값) }
func (s *sC부호없는_정수64) G실수() float64 { return float64(s.값) }
func (s *sC부호없는_정수64) G정밀수() *big.Rat { return F부호없는_정수2정밀수(s.값) }
func (s *sC부호없는_정수64) G문자열() string { return F부호없는_정수2문자열(s.값) }
func (s *sC부호없는_정수64) String() string { return F부호없는_정수2문자열(s.값) }
func (s *sC부호없는_정수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	return reflect.ValueOf(NC부호없는_정수(uint64(임의값_생성기.Uint32())))
}

func NC실수(값 float64) C실수   { return &sC실수64{값} }
type sC실수64 struct{ 값 float64 }
func (s *sC실수64) 상수형임()          {}
func (s *sC실수64) G값() float64    { return s.값 }
func (s *sC실수64) G실수() float64 { return s.값 }
func (s *sC실수64) G정밀수() *big.Rat { return F실수2정밀수(s.값) }
func (s *sC실수64) G문자열() string { return F실수2문자열(s.값) }
func (s *sC실수64) String() string { return F실수2문자열(s.값) }
func (s *sC실수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	값 := 임의값_생성기.NormFloat64() 
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = 값 * -1.0
	}
	
	return reflect.ValueOf(NC실수(값))
}

func NC참거짓(값 bool) C참거짓 {
	if 값 {
		return c참
	} else {
		return c거짓
	}
}
type sC참거짓 struct{ 값 bool }
func (s *sC참거짓) 상수형임()          {}
func (s *sC참거짓) G값() bool       { return s.값 }
func (s *sC참거짓) G문자열() string { return F참거짓2문자열(s.값) }
func (s *sC참거짓) String() string { return F참거짓2문자열(s.값) }
func (s *sC참거짓) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	값 := true						
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = false
	}
	
	return reflect.ValueOf(NC참거짓(값))							
}

func NC문자열(값 string) C문자열 { return &sC문자열{값} }
type sC문자열 struct{ 값 string }
func (s *sC문자열) 상수형임()          {}
func (s *sC문자열) G값() string     { return s.값 }
func (s *sC문자열) G문자열() string { return s.값 }
func (s *sC문자열) String() string { return s.값 }
func (s *sC문자열) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	후보값_수량 := int32(len(문자열_후보값_모음) - 1)
	임의_문자열 := new(bytes.Buffer)
	
	for 반복횟수 := 0 ; 반복횟수 < 크기 ; 반복횟수++ {
		슬라이스_인덱스 := int(임의값_생성기.Int31n(후보값_수량))
		임의_문자열.WriteString(문자열_후보값_모음[슬라이스_인덱스])
	}
	
	return reflect.ValueOf(NC문자열(임의_문자열.String()))
}