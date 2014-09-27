package common

import (
	"bytes"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func init() {
	c참 = &sC참거짓{true}
	c거짓 = &sC참거짓{false}
	
	후보값_문자열 := "1234567890" +
			"abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"~!@#$%^&*()_+|';:/?.,<>`" +
			"가나다라마바사하자차카타파하" // 한자도 포함시킬 것.
	
	문자열_후보값_모음 = strings.Split(후보값_문자열, "")
}

var c참, c거짓 *sC참거짓
var 문자열_후보값_모음 []string

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

// 시점 (time.Time)
type s시점 struct{ 값 time.Time }
func (s *s시점) G값() time.Time  { return s.값 }
func (s *s시점) G변수형() V시점 { return NV시점(s.값) }
func (s *s시점) String() string { return s.값.String() }

func NC시점(값 time.Time) C시점 { return &sC시점{&s시점{값}} }
type sC시점 struct{ *s시점 }
func (s *sC시점) 상수형임() {}
func (s *sC시점) G변수형() V시점 { return NV시점(s.s시점.값) }
func (s *sC시점) Generate(임의값_생성기 *rand.Rand, 
						크기 int) reflect.Value {
	연도 := int(1900 + 임의값_생성기.Int31n(300))
	월, _ := F정수2월(int(1 + 임의값_생성기.Int31n(11)))
	일 := int(1 + 임의값_생성기.Int31n(30))
	시 := int(임의값_생성기.Int31n(24))
	분 := int(임의값_생성기.Int31n(59))
	초 := int(임의값_생성기.Int31n(59))
	나노초 := 임의값_생성기.Int()
	
	값 := time.Date(연도, 월, 일, 시, 분, 초, 나노초, time.Now().Location())
	
	return reflect.ValueOf(NC시점(값))							
}

func NV시점(값 time.Time) V시점 { return &sV시점{&s시점{값}} }
type sV시점 struct{ *s시점 }
func (s *sV시점) 변수형임()          {}
func (s *sV시점) S값(값 time.Time) { s.s시점.값 = 값 }
func (s *sV시점) G상수형() C시점      { return NC시점(s.s시점.값) }
func (s *sV시점) Generate(임의값_생성기 *rand.Rand, 
						크기 int) reflect.Value {
	연도 := int(1900 + 임의값_생성기.Int31n(300))
	월, _ := F정수2월(int(1 + 임의값_생성기.Int31n(11)))
	일 := int(1 + 임의값_생성기.Int31n(30))
	시 := int(임의값_생성기.Int31n(24))
	분 := int(임의값_생성기.Int31n(59))
	초 := int(임의값_생성기.Int31n(59))
	나노초 := 임의값_생성기.Int()
	
	값 := time.Date(연도, 월, 일, 시, 분, 초, 나노초, time.Now().Location())
	
	return reflect.ValueOf(NV시점(값))							
}

// 큰 정수 (*big.Int)
type s큰정수 struct{ 값 *big.Int }
func (s *s큰정수) G값() *big.Int   { return F큰정수_복사(s.값) }
func (s *s큰정수) G정수() int64    { return s.값.Int64() }
func (s *s큰정수) G큰정수() *big.Int { return s.G값() }
func (s *s큰정수) G실수() float64 { return F큰정수2실수(s.값) }
func (s *s큰정수) G정밀수() *big.Rat { return new(big.Rat).SetInt(s.값) }
func (s *s큰정수) G문자열() string { return s.값.String() }
func (s *s큰정수) String() string { return s.값.String() }

func NC큰정수(값 int64) C큰정수 { return NC큰정수Big(F정수2큰정수(값)) }
func NC큰정수Big(값 *big.Int) C큰정수 { return &sC큰정수{&s큰정수{F큰정수_복사(값)}} }
type sC큰정수 struct{ *s큰정수 }
func (s *sC큰정수) 상수형임() {}
func (s *sC큰정수) G변수형() V큰정수 { return NV큰정수Big(s.s큰정수.값) }
func (s *sC큰정수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	값 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = 값 * -1
	}
	
	return reflect.ValueOf(NC큰정수(값))							
}

func NV큰정수(값 int64) V큰정수 { return NV큰정수Big(F정수2큰정수(값)) }
func NV큰정수Big(값 *big.Int) V큰정수 { return &sV큰정수{&s큰정수{F큰정수_복사(값)}} }
type sV큰정수 struct{ *s큰정수 }
func (s *sV큰정수) 변수형임() {}
func (s *sV큰정수) G상수형() C큰정수 { return NC큰정수Big(s.s큰정수.값) }
func (s *sV큰정수) S값(값 int64) { s.s큰정수.값.SetInt64(값) }
func (s *sV큰정수) S값Big(값 *big.Int) { s.s큰정수.값.Set(값) }

func (s *sV큰정수) S절대값(큰정수 I큰정수) V큰정수 {
	s.S값Big(F큰정수_절대값(큰정수.G값())); return s
}
func (s *sV큰정수) S더하기(큰정수1 I큰정수, 큰정수2 I큰정수) V큰정수 {
	s.S값Big(F큰정수_더하기(큰정수1.G값(), 큰정수2.G값())); return s
}
func (s *sV큰정수) S빼기(큰정수1 I큰정수, 큰정수2 I큰정수) V큰정수 {
	s.S값Big(F큰정수_빼기(큰정수1.G값(), 큰정수2.G값())); return s
}
func (s *sV큰정수) S곱하기(큰정수1 I큰정수, 큰정수2 I큰정수) V큰정수 {
	s.S값Big(F큰정수_곱하기(큰정수1.G값(), 큰정수2.G값())); return s
}
func (s *sV큰정수) S나누기(분자 I큰정수, 분모 I큰정수) (V큰정수, error) {
	값, 에러 := F큰정수_나누기(분자.G값(), 분모.G값())
	
	if 에러 != nil {
		return nil, 에러
	}
	
	s.S값Big(값)
	
	return s, nil
}
func (s *sV큰정수) S반대부호값(큰정수 I큰정수) V큰정수 {
	s.S값Big(F큰정수_반대부호값(큰정수.G값())); return s
}

func (s *sV큰정수) S셀프_절대값() V큰정수 {
	s.S값Big(F큰정수_절대값(s.s큰정수.값)); return s
}
func (s *sV큰정수) S셀프_더하기(큰정수 I큰정수) V큰정수 {
	s.S값Big(F큰정수_더하기(s.s큰정수.값, 큰정수.G값())); return s
}
func (s *sV큰정수) S셀프_빼기(큰정수 I큰정수) V큰정수 {
	s.S값Big(F큰정수_빼기(s.s큰정수.값, 큰정수.G값())); return s
}
func (s *sV큰정수) S셀프_곱하기(큰정수 I큰정수) V큰정수 {
	s.S값Big(F큰정수_곱하기(s.s큰정수.값, 큰정수.G값())); return s
}
func (s *sV큰정수) S셀프_나누기(큰정수 I큰정수) (V큰정수, error) {
	값, 에러 := F큰정수_나누기(s.s큰정수.값, 큰정수.G값())
	
	if 에러 != nil {
		return nil, 에러
	}
	
	s.S값Big(값)
	
	return s, nil
}
func (s *sV큰정수) S셀프_반대부호값() V큰정수 {
	s.S값Big(F큰정수_반대부호값(s.s큰정수.값)); return s
}
func (s *sV큰정수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	값 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = 값 * -1
	}
	
	return reflect.ValueOf(NV큰정수(값))							
}

type s정밀수 struct{ 값 *big.Rat }
func (s *s정밀수) G값() *big.Rat  { return F정밀수_복사(s.값) }
func (s *s정밀수) G실수() float64 { return F정밀수2실수(s.값) }
func (s *s정밀수) G정밀수() *big.Rat  { return s.G값() }
func (s *s정밀수) G문자열() string { return F정밀수2문자열(s.값) }
func (s *s정밀수) G반올림_실수(소숫점_이하_자릿수 int) float64 {
	return F정밀수_반올림값(s.값, 소숫점_이하_자릿수)
}
func (s *s정밀수) G반올림_정밀수(소숫점_이하_자릿수 int) *big.Rat {
	return F정밀수_반올림값Big(s.값, 소숫점_이하_자릿수)
}
func (s *s정밀수) G반올림_문자열(소숫점_이하_자릿수 int) string {
	return F정밀수_반올림_문자열(s.값, 소숫점_이하_자릿수)
}
//func (s *s정밀수) G부호() int { return s.값.Sign() }
func (s *s정밀수) String() string { return F정밀수2문자열(s.값) }


func NC정밀수(값 float64) C정밀수 { return &sC정밀수{&s정밀수{F실수2정밀수(값)}} }
func NC정밀수Big(값 *big.Rat) C정밀수 { return &sC정밀수{&s정밀수{F정밀수_복사(값)}} }

type sC정밀수 struct { *s정밀수 }

func (s *sC정밀수) 상수형임() {}
func (s *sC정밀수) G변수형() V정밀수 { return NV정밀수Big(s.s정밀수.값) }
func (s *sC정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}
	
	return reflect.ValueOf(NC정밀수Big(big.NewRat(분자, 분모)))						
}


func NV정밀수(값 float64) V정밀수 { return &sV정밀수{&s정밀수{F실수2정밀수(값)}} }
func NV정밀수Big(값 *big.Rat) V정밀수 { return &sV정밀수{&s정밀수{F정밀수_복사(값)}} }

type sV정밀수 struct { *s정밀수 }

func (s *sV정밀수) 변수형임() {}
func (s *sV정밀수) G상수형() C정밀수 { return NC정밀수Big(s.s정밀수.값) }
func (s *sV정밀수) S값(값 float64) { s.s정밀수.값.Set(F실수2정밀수(값)) }
func (s *sV정밀수) S값Big(값 *big.Rat) { s.s정밀수.값.Set(값) }

func (s *sV정밀수) S절대값(값 I정밀수) V정밀수 {
	s.S값Big(F정밀수_절대값(값.G값())); return s
}
func (s *sV정밀수) S더하기(값1 I정밀수, 값2 I정밀수) V정밀수 {
	s.S값Big(F정밀수_더하기(값1.G정밀수(), 값2.G정밀수())); return s
}
func (s *sV정밀수) S빼기(값1 I정밀수, 값2 I정밀수) V정밀수 {
	s.S값Big(F정밀수_빼기(값1.G정밀수(), 값2.G정밀수())); return s
}
func (s *sV정밀수) S곱하기(값1 I정밀수, 값2 I정밀수) V정밀수 {
	s.S값Big(F정밀수_곱하기(값1.G정밀수(), 값2.G정밀수())); return s
}
func (s *sV정밀수) S나누기(분자 I정밀수, 분모 I정밀수) (V정밀수, error) {
	정밀수, 에러 := F정밀수_나누기(분자.G정밀수(), 분모.G정밀수())
	
	if 에러 != nil {
		return nil, 에러
	}
	
	s.S값Big(정밀수); return s, 에러
}
func (s *sV정밀수) S역수(값 I정밀수) (V정밀수, error) {
	정밀수, 에러 := F정밀수_역수(값.G값())
	
	if 에러 != nil {
		return nil, 에러
	}
	
	s.S값Big(정밀수); return s, nil
}
func (s *sV정밀수) S반대부호값(값 I정밀수) V정밀수 {
	s.S값Big(F정밀수_반대부호값(값.G정밀수())); return s
}

func (s *sV정밀수) S셀프_절대값() V정밀수 {
	s.S값Big(F정밀수_절대값(s.s정밀수.값)); return s
}
func (s *sV정밀수) S셀프_더하기(값 I정밀수) V정밀수 {
	s.S값Big(F정밀수_더하기(s.s정밀수.값, 값.G정밀수())); return s
}
func (s *sV정밀수) S셀프_빼기(값 I정밀수) V정밀수 {
	s.S값Big(F정밀수_빼기(s.s정밀수.값, 값.G정밀수())); return s
}
func (s *sV정밀수) S셀프_곱하기(값 I정밀수) V정밀수 {
	s.S값Big(F정밀수_곱하기(s.s정밀수.값, 값.G정밀수())); return s
}
func (s *sV정밀수) S셀프_나누기(값 I정밀수) (V정밀수, error) {
	정밀수, 에러 := F정밀수_나누기(s.s정밀수.값, 값.G정밀수())
	
	if 에러 != nil {
		return nil, 에러
	}
	
	s.S값Big(정밀수); return s, nil
}
func (s *sV정밀수) S셀프_역수() (V정밀수, error) {
	정밀수, 에러 := F정밀수_역수(s.s정밀수.값)
	
	if 에러 != nil {
		return nil, 에러
	}
	
	s.S값Big(정밀수); return s, nil
}
func (s *sV정밀수) S셀프_반대부호값() V정밀수 {
	s.S값Big(F정밀수_반대부호값(s.s정밀수.값)); return s
}
func (s *sV정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}
	
	return reflect.ValueOf(NV정밀수Big(big.NewRat(분자, 분모)))						
}