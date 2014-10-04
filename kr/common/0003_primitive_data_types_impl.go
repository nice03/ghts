package common

import (
	"bytes"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)

func init() {
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
// 또한, 공유되는 변수형의 경우에도 RWMutex로 보호해야 공유로 인한 문제를 줄일 수 있다.

type sC정수64 struct{ 값 int64 }
func (s *sC정수64) 상수형임()  {}
func (s *sC정수64) G값() int64 { return s.값 }
func (s *sC정수64) G정수() int64 { return s.값 }
func (s *sC정수64) G실수() float64 { return float64(s.값) }
func (s *sC정수64) G정밀수() C정밀수 { return NC정밀수(s.값) }
func (s *sC정수64) G변수형() V정수 { return NV정수(s.값) }
func (s *sC정수64) String() string { return F문자열(s.값) }
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

type sV정수64 struct {
	잠금 sync.RWMutex
	값 int64
}
func (s *sV정수64) 변수형임()  {}
func (s *sV정수64) G값() int64 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값
}
func (s *sV정수64) G정수() int64 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값
}
func (s *sV정수64) G실수() float64 {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return float64(값)
}
func (s *sV정수64) G정밀수() C정밀수  {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return NC정밀수(값)
}
func (s *sV정수64) G상수형() C정수  {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return NC정수(값)
}
func (s *sV정수64) S값(값 int64) V정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = 값; return s
}
func (s *sV정수64) S절대값() V정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	if s.값 < 0 { s.값 = s.값 * -1 }; return s
}
func (s *sV정수64) S더하기(값 int64) V정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 + 값; return s
}
func (s *sV정수64) S빼기(값 int64) V정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 - 값; return s
}
func (s *sV정수64) S곱하기(값 int64) V정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 * 값; return s
}
func (s *sV정수64) S나누기(값 int64) V정수 {
	if 값 == 0 {
		F문자열_출력("sV정수.S나누기() : 0으로 나눌 수 없음."); return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 / 값; return s
}
func (s *sV정수64) String() string {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return F문자열(값)
}
func (s *sV정수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	값 := 임의값_생성기.Int63() // 0 혹은 양의 정수.
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = 값 * -1
	}
	
	return reflect.ValueOf(NV정수(값))
}


type sC부호없는_정수64 struct { 값 uint64 }
func (s *sC부호없는_정수64) 상수형임()          {}
func (s *sC부호없는_정수64) G값() uint64     { return s.값 }
func (s *sC부호없는_정수64) G정수() int64 { return int64(s.값) }
func (s *sC부호없는_정수64) G실수() float64 { return float64(s.값) }
func (s *sC부호없는_정수64) G정밀수() C정밀수 { return NC정밀수(s.값) }
func (s *sC부호없는_정수64) G변수형() V부호없는_정수 { return NV부호없는_정수(s.값) }
func (s *sC부호없는_정수64) String() string { return F문자열(s.값) }
func (s *sC부호없는_정수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	return reflect.ValueOf(NC부호없는_정수(uint64(임의값_생성기.Uint32())))
}

type sV부호없는_정수64 struct {
	잠금 sync.RWMutex
	값 uint64
}
func (s *sV부호없는_정수64) 변수형임()  {}
func (s *sV부호없는_정수64) G값() uint64 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값
}
func (s *sV부호없는_정수64) G정수() int64 {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return new(big.Int).SetUint64(값).Int64()
}
func (s *sV부호없는_정수64) G실수() float64 {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return float64(값)
}
func (s *sV부호없는_정수64) G정밀수() C정밀수  {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return NC정밀수(값)
}
func (s *sV부호없는_정수64) G상수형() C부호없는_정수  {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return NC부호없는_정수(값)
}
func (s *sV부호없는_정수64) S값(값 uint64) V부호없는_정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값 = 값; return s
}
func (s *sV부호없는_정수64) S더하기(값 uint64) V부호없는_정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값 = s.값 + 값; return s
}
func (s *sV부호없는_정수64) S빼기(값 uint64) V부호없는_정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값 = s.값 - 값; return s
}
func (s *sV부호없는_정수64) S곱하기(값 uint64) V부호없는_정수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값 = s.값 * 값; return s
}
func (s *sV부호없는_정수64) S나누기(값 uint64) V부호없는_정수 {
	if 값 == 0 {
		F문자열_출력("sV부호없는_정수.S나누기() : 0으로 나눌 수 없음."); return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값 = s.값 / 값; return s
}
func (s *sV부호없는_정수64) String() string {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return F문자열(값)
}
func (s *sV부호없는_정수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	값 := 임의값_생성기.Int63() // 0 혹은 양의 부호없는_부호없는_정수.
	
	return reflect.ValueOf(NV부호없는_정수(uint64(값)))
}

type sC실수64 struct{ 값 float64 }
func (s *sC실수64) 상수형임()          {}
func (s *sC실수64) G값() float64    { return s.값 }
func (s *sC실수64) G실수() float64 { return s.값 }
func (s *sC실수64) G정밀수() C정밀수 { return NC정밀수(s.값) }
func (s *sC실수64) G변수형() V실수 { return NV실수(s.값) }
func (s *sC실수64) String() string { return F문자열(s.값) }
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

type sV실수64 struct {
	잠금 sync.RWMutex
	값 float64
}
func (s *sV실수64) 변수형임()  {}
func (s *sV실수64) G값() float64 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값
}
func (s *sV실수64) G실수() float64 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값  
}
func (s *sV실수64) G정밀수() C정밀수  {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return NC정밀수(값)
}
func (s *sV실수64) G상수형() C실수  {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return NC실수(값)
}
func (s *sV실수64) S값(값 float64) V실수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = 값; return s
}
func (s *sV실수64) S절대값() V실수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	if s.값 < 0 { s.값 = s.값 * -1 }; return s
}
func (s *sV실수64) S더하기(값 float64) V실수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 + 값; return s
}
func (s *sV실수64) S빼기(값 float64) V실수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 - 값; return s
}
func (s *sV실수64) S곱하기(값 float64) V실수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 * 값; return s
}
func (s *sV실수64) S나누기(값 float64) V실수 {
	if 값 == 0 {
		F문자열_출력("sV실수.S나누기() : 0으로 나눌 수 없음."); return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	
	s.값 = s.값 / 값; return s
}
func (s *sV실수64) String() string {
	s.잠금.RLock(); 값 := s.값; s.잠금.RUnlock()
	return F문자열(값)
}
func (s *sV실수64) Generate(
							임의값_생성기 *rand.Rand, 
							크기 int) reflect.Value {
	값 := 임의값_생성기.Float64() // 0 혹은 양의 정수.
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		값 = 값 * -1
	}
	
	return reflect.ValueOf(NV실수(값))
}

type sC참거짓 struct{ 값 bool }
func (s *sC참거짓) 상수형임()          {}
func (s *sC참거짓) G값() bool       { return s.값 }
func (s *sC참거짓) String() string { return F문자열(s.값) }
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

type sC문자열 struct{ 값 string }
func (s *sC문자열) 상수형임()          {}
func (s *sC문자열) G값() string     { return s.값 }
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
type sC시점 struct{ 값 time.Time }
func (s *sC시점) 상수형임() {}
func (s *sC시점) G값() time.Time  { return s.값 }
func (s *sC시점) G변수형() V시점 { return NV시점(s.값) }
func (s *sC시점) String() string { return F문자열(s.값) }
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

type sV시점 struct {
	잠금 sync.RWMutex
	값 time.Time
}
func (s *sV시점) 변수형임()          {}
func (s *sV시점) G값() time.Time  {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값
}
func (s *sV시점) G상수형() C시점 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return NC시점(s.값)
}
func (s *sV시점) S값(값 time.Time) V시점 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값 = 값; return s
}
func (s *sV시점) S날짜_더하기(연, 월, 일 int) V시점 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.AddDate(연, 월, 일); return s
}
func (s *sV시점) String() string {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.값.Format(P시점_포맷)
}
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

type sC정밀수 struct { 값 *big.Rat }
func (s *sC정밀수) 상수형임() {}
func (s *sC정밀수) G값() string { return s.String() }
func (s *sC정밀수) GRat() *big.Rat { return new(big.Rat).Set(s.값) }
func (s *sC정밀수) G실수() float64 {
	실수, 에러 := F문자열2실수(s.String())

	if 에러 != nil {
		실수, _ = s.값.Float64()
	}
	
	return 실수
}
func (s *sC정밀수) G정밀수() C정밀수  { return s }
func (s *sC정밀수) G같음(값 interface{}) bool {
	정밀수 := NC정밀수(값)
	
	if 정밀수 == nil { return false }
	
	차이_절대값 := new(big.Rat).Abs(new(big.Rat).Sub(s.GRat(), 정밀수.GRat()))
	
	if 차이_절대값.Cmp(NC정밀수(P차이_한도).GRat()) == -1 {
		return true
	}
	
	return false
}
	
func (s *sC정밀수) G비교(값 interface{}) int {
	if s.G같음(값) { return 0 }
	
	정밀수 := NC정밀수(값)
	
	if 정밀수  == nil { return -2 }
	
	return s.값.Cmp(정밀수.GRat())
}
func (s *sC정밀수) G변수형() V정밀수 { return NV정밀수(s.G값()) }
func (s *sC정밀수) String() string { { return F마지막_0_제거(s.GRat().FloatString(100)) } }
func (s *sC정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}
	
	return reflect.ValueOf(NC정밀수(big.NewRat(분자, 분모)))						
}

type sV정밀수 struct {
	잠금 sync.RWMutex
	값 *big.Rat
}
func (s *sV정밀수) 변수형임() {}
func (s *sV정밀수) G값() string { return s.String() }
func (s *sV정밀수) GRat() *big.Rat {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return new(big.Rat).Set(s.값)
}
func (s *sV정밀수) G실수() float64 {
	실수, 에러 := F문자열2실수(s.String())
	
	if 에러 != nil {
		s.잠금.RLock()
		실수, _ = s.값.Float64()
		s.잠금.RUnlock()
	}
	
	return 실수
}
func (s *sV정밀수) G정밀수() C정밀수  { return NC정밀수(s.G값()) }
func (s *sV정밀수) G상수형() C정밀수 { return NC정밀수(s.G값()) }
func (s *sV정밀수) G같음(값 interface{}) bool {
	정밀수 := NC정밀수(값)
	
	if 정밀수 == nil { return false }
	
	차이_절대값 := new(big.Rat).Abs(new(big.Rat).Sub(s.GRat(), 정밀수.GRat()))
	
	if 차이_절대값.Cmp(NC정밀수(P차이_한도).GRat()) == -1 {
		return true
	}
	
	return false
}
func (s *sV정밀수) G비교(값 interface{}) int {
	if s.G같음(값) { return 0 }
	
	정밀수 := NC정밀수(값)

	if 정밀수 == nil { return -2 }
	
	s.잠금.RLock(); s.잠금.RUnlock()
	return s.값.Cmp(정밀수.GRat())
}
func (s *sV정밀수) S값(값 interface{}) V정밀수 {
	if 값 == nil { return nil }
	
	정밀수 := NC정밀수(값)
	
	if 정밀수 == nil { return nil }
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Set(정밀수.GRat()); return s
}
func (s *sV정밀수) S반올림(소숫점_이하_자릿수 int) V정밀수 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	문자열 := s.값.FloatString(소숫점_이하_자릿수)
	s.값.Set(NC정밀수(문자열).GRat()); return s
}
func (s *sV정밀수) S절대값() V정밀수 { return s.S절대값2(s) }
func (s *sV정밀수) S더하기(값 interface{}) V정밀수 { return s.S더하기2(s, 값) }
func (s *sV정밀수) S빼기(값 interface{}) V정밀수 { return s.S빼기2(s, 값) }
func (s *sV정밀수) S곱하기(값 interface{}) V정밀수 { return s.S곱하기2(s, 값) }
func (s *sV정밀수) S나누기(값 interface{}) V정밀수 { return s.S나누기2(s, 값) }
func (s *sV정밀수) S역수() V정밀수 { return s.S역수2(s) }
func (s *sV정밀수) S반대부호값() V정밀수 { return s.S반대부호값2(s) }
func (s *sV정밀수) S절대값2(값 interface{}) V정밀수 {
	정밀수 := NC정밀수(값)
	
	if 정밀수 == nil {
		F문자열_출력("common.sV정밀수.S절대값2() : 정밀수 변환 에러. 값 %v.", 값)
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Abs(정밀수.GRat()); return s
}
func (s *sV정밀수) S더하기2(값1, 값2 interface{}) V정밀수 {	
	정밀수1, 정밀수2 := NC정밀수(값1), NC정밀수(값2)
	
	if 정밀수1 == nil || 정밀수2 == nil {
		F문자열_출력("common.sV정밀수.S더하기2() : 정밀수 변환 에러. 값1 %v, 값2 %v .", 
					값1, 값2)
		return nil
	}

	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Add(정밀수1.GRat(), 정밀수2.GRat()); return s
}
func (s *sV정밀수) S빼기2(값1, 값2 interface{}) V정밀수 {
	정밀수1, 정밀수2 := NC정밀수(값1), NC정밀수(값2)
	
	if 정밀수1 == nil || 정밀수2 == nil {
		F문자열_출력("common.sV정밀수.S빼기2() : 정밀수 변환 에러. 값1 %v, 값2 %v .", 
					값1, 값2)
		return nil
	}

	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Sub(정밀수1.GRat(), 정밀수2.GRat()); return s
}
func (s *sV정밀수) S곱하기2(값1, 값2 interface{}) V정밀수 {	
	정밀수1, 정밀수2 := NC정밀수(값1), NC정밀수(값2)
	
	if 정밀수1 == nil || 정밀수2 == nil {
		F문자열_출력("common.sV정밀수.S곱하기2() : 정밀수 변환 에러. 값1 %v, 값2 %v .", 
					값1, 값2)
		return nil
	}

	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Mul(정밀수1.GRat(), 정밀수2.GRat()); return s
}
func (s *sV정밀수) S나누기2(분자, 분모 interface{}) V정밀수 {
	정밀수1, 정밀수2 := NC정밀수(분자), NC정밀수(분모)
	
	if 정밀수1 == nil || 정밀수2 == nil {
		F문자열_출력("common.sV정밀수.S나누기2() : 정밀수 변환 에러. 분자 %v, 분모 %v .", 
					분자, 분모)
		return nil
	}
	
	if 정밀수2.G같음(0.0) {
		F문자열_출력("common.sV정밀수.S나누기2() : 분모가 0임. 분자 %v, 분모 %v .", 
					분자, 분모)
		return nil
	}

	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Quo(정밀수1.GRat(), 정밀수2.GRat()); return s
}
func (s *sV정밀수) S역수2(값 interface{}) V정밀수 {
	정밀수 := NC정밀수(값)
	
	if 정밀수 == nil {
		F문자열_출력("common.sV정밀수.S역수2() : 정밀수 변환 에러. 값 %v .", 값)
		return nil
	}
	
	if 정밀수.G같음(0.0) {
		F문자열_출력("common.sV정밀수.S역수2() : 0의 역수는 무한대임. 값 %v .", 값)
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Inv(정밀수.GRat()); return s
}
func (s *sV정밀수) S반대부호값2(값 interface{}) V정밀수 {
	정밀수 := NC정밀수(값)
	
	if 정밀수 == nil {
		F문자열_출력("common.sV정밀수.S반대부호값2() : 정밀수 변환 에러. 값 %v .", 값)
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.값.Neg(정밀수.GRat()); return s
}
func (s *sV정밀수) String() string { return F마지막_0_제거(s.GRat().FloatString(100)) }
func (s *sV정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}
	
	return reflect.ValueOf(NV정밀수(big.NewRat(분자, 분모)))						
}

// 통화
type sC통화 struct {
	종류 P통화
	금액 C정밀수
}
func (s *sC통화) 상수형임() {}
func (s *sC통화) G종류() P통화 { return s.종류 }
func (s *sC통화) G금액() C정밀수 { return s.금액 }
func (s *sC통화) G같음(값 I통화) bool {
	if 값 == nil { return false }
	
	if s.종류 == 값.G종류() && 
		s.금액.G같음(값.G금액()) {
		return true
	}
	
	return false
}
func (s *sC통화) G변수형() V통화 { return NV통화(s.종류, s.금액) }
func (s *sC통화) String() string {
	// TODO. 예 : KRW 1,000,000 (통화종류, 콤마로 분리된 금액)
	return string(s.종류) + " " + s.금액.String()
}
func (s *sC통화) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	종류 := F임의_통화종류()
	
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}
	
	return reflect.ValueOf(NC통화(종류, big.NewRat(분자, 분모)))
}

type sV통화 struct {
	잠금 sync.RWMutex
	종류 P통화
	금액 V정밀수
}
func (s *sV통화) 변수형임() {}
func (s *sV통화) G종류() P통화 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.종류
}
func (s *sV통화) G금액() C정밀수 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return s.금액.G상수형()
}
func (s *sV통화) G같음(값 I통화) bool {
	if 값 == nil { return false }
	
	s.잠금.RLock(); defer s.잠금.RUnlock()
	if s.종류 == 값.G종류() && 
		s.금액.G같음(값.G금액()) {
		return true
	}
	
	return false
}
func (s *sV통화) G상수형() C통화 {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	return NC통화(s.종류, s.금액)
}
func (s *sV통화) S종류(종류 P통화) {
	s.잠금.RLock(); defer s.잠금.RUnlock()
	s.종류 = 종류
}
func (s *sV통화) S금액(금액 I정밀수) V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액 = NV정밀수(금액); return s
}
func (s *sV통화) S절대값() V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액.S절대값(); return s
}
func (s *sV통화) S더하기(값 interface{}) V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액.S더하기(s); return s
}
func (s *sV통화) S빼기(값 interface{}) V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액.S빼기(s); return s
}
func (s *sV통화) S곱하기(값 interface{}) V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액.S곱하기(s); return s
}
func (s *sV통화) S나누기(값 interface{}) V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액.S나누기(s); return s
}
func (s *sV통화) S반대부호값() V통화 {
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.금액.S반대부호값(); return s
}
func (s *sV통화) S절대값2(값 I통화) V통화 {
	if 값 == nil || 값.G금액() == nil {
		F문자열_출력("common.sV통화.S절대값2() : nil 입력값."); return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.종류 = 값.G종류()
	s.금액.S절대값2(값.G금액())
	
	return s
}
func (s *sV통화) S더하기2(값1, 값2 interface{}) V통화 {
	if 값1 == nil || 값2 == nil {
		F문자열_출력("common.sV통화.S더하기2() : nil 입력값."); return nil
	}

	통화종류, 에러 := F통화_종류(값1, 값2)
	
	if 에러 != nil {
		F문자열_출력("common.sV통화.S더하기2() : %s. 값1 %v, 값2 %v", 
						에러.Error(), 값1, 값2)
						
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.종류 = 통화종류
	s.금액.S더하기2(값1, 값2)
	
	return s
}
func (s *sV통화) S빼기2(값1, 값2 interface{}) V통화 {
	if 값1 == nil || 값2 == nil {
		F문자열_출력("common.sV통화.S빼기2() : nil 입력값."); return nil
	}

	통화종류, 에러 := F통화_종류(값1, 값2)
	
	if 에러 != nil {
		F문자열_출력("common.sV통화.S빼기2() : %s. 값1 %v, 값2 %v", 
						에러.Error(), 값1, 값2)
						
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.종류 = 통화종류
	s.금액.S빼기2(값1, 값2)
	
	return s
}
func (s *sV통화) S곱하기2(값1, 값2 interface{}) V통화 {
	if 값1 == nil || 값2 == nil {
		F문자열_출력("common.sV통화.S곱하기2() : nil 입력값."); return nil
	}

	통화종류, 에러 := F통화_종류(값1, 값2)
	
	if 에러 != nil {
		F문자열_출력("common.sV통화.S곱하기2() : %s. 값1 %v, 값2 %v", 
						에러.Error(), 값1, 값2)
						
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.종류 = 통화종류
	s.금액.S곱하기2(값1, 값2)
	
	return s
}
func (s *sV통화) S나누기2(분자, 분모 interface{}) V통화 {
	if 분자 == nil || 분모 == nil {
		F문자열_출력("common.sV통화.S나누기2() : nil 입력값."); return nil
	}

	통화종류, 에러 := F통화_종류(분자, 분모)
	
	if 에러 != nil {
		F문자열_출력("common.sV통화.S나누기2() : %s. 값1 %v, 값2 %v", 
						에러.Error(), 분자, 분모)
						
		return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.종류 = 통화종류
	s.금액.S나누기2(분자, 분모)
	
	return s
}
func (s *sV통화) S반대부호값2(값 I통화) V통화 {
	if 값 == nil || 값.G금액() == nil {
		F문자열_출력("common.sV통화.S절대값2() : nil 입력값."); return nil
	}
	
	s.잠금.Lock(); defer s.잠금.Unlock()
	s.종류 = 값.G종류()
	s.금액.S반대부호값2(값.G금액())
	
	return s
}
func (s *sV통화) String() string {
	// TODO. 예 : KRW 1,000,000 (통화종류, 콤마로 분리된 금액)
	return string(s.종류) + " " + F문자열(s.금액.String())
} 
func (s *sV통화) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	종류 := F임의_통화종류()
	
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()
	
	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}
	
	return reflect.ValueOf(NV통화(종류, big.NewRat(분자, 분모)))		
}

type sC매개변수 struct {
	이름 string
	값 I상수형
}
func (s *sC매개변수) G이름() string { return s.이름 }
func (s *sC매개변수) G값() I상수형 { return s.값 }