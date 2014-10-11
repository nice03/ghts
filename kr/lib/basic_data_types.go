package lib

import (
	"bytes"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	//"sync/atomic"
	"time"
	//"unsafe"
)

func NC정수(값 int64) C정수 { return &sC정수64{값} }
func NV정수(값 int64) V정수 { return &sV정수64{값: 값} }

func NC부호없는_정수(값 uint64) C부호없는_정수 { return &sC부호없는_정수64{값} }
func NV부호없는_정수(값 uint64) V부호없는_정수 { return &sV부호없는_정수64{값: 값} }

func NC실수(값 float64) C실수 { return &sC실수64{값} }
func NV실수(값 float64) V실수 { return &sV실수64{값: 값} }

func NC참거짓(값 bool) C참거짓 {
	if 값 {
		return c참
	} else {
		return c거짓
	}
}
func NV참거짓(값 bool) V참거짓 { return &sV참거짓{s참거짓: &s참거짓{값}} }

func NC문자열(값 string) C문자열 { return &sC문자열{값} }

func NC시점(값 time.Time) C시점 { return &sC시점{값} }
func NC시점_문자열(값 string) C시점 {
	시점, 에러 := F문자열2시점(값)

	if 에러 != nil {
		return nil
	}

	return NC시점(시점)
}

func NV시점(값 time.Time) V시점 {
	return &sV시점{값: 값}
}

func NV시점_문자열(값 string) V시점 {
	시점, 에러 := F문자열2시점(값)

	if 에러 != nil {
		return nil
	}

	return NV시점(시점)
}

func NC정밀수(값 interface{}) C정밀수 {
	if 값 == nil {
		return nil
	}

	var 정밀수 *big.Rat

	switch 값.(type) {
	case *sC정밀수:
		return 값.(*sC정밀수) // 상수형은 굳이 새로운 인스턴스를 생성할 필요가 없다.
	case *sV정밀수:
		정밀수 = 값.(*sV정밀수).GRat()
	case *big.Rat:
		정밀수 = new(big.Rat).Set(값.(*big.Rat))
	default:
		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))

		if !성공 {
			return nil
		}
	}

	return &sC정밀수{&s정밀수{정밀수}}
}

func NV정밀수(값 interface{}) V정밀수 {
	if 값 == nil {
		return nil
	}

	var 정밀수 *big.Rat

	switch 값.(type) {
	case *big.Rat:
		정밀수 = new(big.Rat).Set(값.(*big.Rat))
	case I정밀수:
		정밀수 = 값.(I정밀수).GRat()
	default:
		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))

		if !성공 {
			F문자열_출력("common.NV정밀수() : 입력값이 숫자가 아님. %v", 값)

			return nil
		}
	}

	return &sV정밀수{s정밀수: &s정밀수{정밀수}}
}

// 통화
type P통화종류 int

var 통화종류_문자열_모음 = [...]string{"KRW", "USD", "CNY", "EUR"}

func (p P통화종류) String() string {
	if int(p) == -1 {
		return "INVALID"
	}

	return 통화종류_문자열_모음[p]
}

func NC원화(금액 interface{}) C통화  { return NC통화(KRW, 금액) }
func NC달러(금액 interface{}) C통화  { return NC통화(USD, 금액) }
func NC위안화(금액 interface{}) C통화 { return NC통화(CNY, 금액) }
func NC유로화(금액 interface{}) C통화 { return NC통화(EUR, 금액) }

func NC통화(종류 P통화종류, 금액 interface{}) C통화 {
	v금액 := NV정밀수(금액)

	if v금액 == nil {
		F문자열_출력("NC통화() : 금액이 숫자가 아님. %v", 금액)
		return nil
	}

	c금액 := v금액.S반올림(F통화종류별_정밀도(종류)).G상수형()

	return &sC통화{종류, c금액}
}

// 변수형 생성자
func NV원화(금액 interface{}) V통화  { return NV통화(KRW, 금액) }
func NV달러(금액 interface{}) V통화  { return NV통화(USD, 금액) }
func NV위안화(금액 interface{}) V통화 { return NV통화(CNY, 금액) }
func NV유로화(금액 interface{}) V통화 { return NV통화(EUR, 금액) }

func NV통화(종류 P통화종류, 금액 interface{}) V통화 {
	v금액 := NV정밀수(금액)

	if v금액 == nil {
		F문자열_출력("NC통화() : 금액이 숫자가 아님. %v", 금액)
		return nil
	}

	v금액 = v금액.S반올림(F통화종류별_정밀도(종류))

	return &sV통화{종류: 종류, 금액: v금액}
}

func NC매개변수(이름 string, 값 interface{}) C매개변수 {
	var 상수형 I상수형 = F상수형(값)

	if 상수형 == nil {
		return nil
	}

	return &sC매개변수{이름, 상수형}
}

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

func (s *sC정수64) 상수형임()          {}
func (s *sC정수64) G값() int64      { return s.값 }
func (s *sC정수64) G정수() int64     { return s.값 }
func (s *sC정수64) G실수() float64   { return float64(s.값) }
func (s *sC정수64) G정밀수() C정밀수     { return NC정밀수(s.값) }
func (s *sC정수64) G변수형() V정수      { return NV정수(s.값) }
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
	값  int64
}

func (s *sV정수64) 변수형임() {}
func (s *sV정수64) G값() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값
}
func (s *sV정수64) G정수() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값
}
func (s *sV정수64) G실수() float64 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return float64(값)
}
func (s *sV정수64) G정밀수() C정밀수 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return NC정밀수(값)
}
func (s *sV정수64) G상수형() C정수 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return NC정수(값)
}
func (s *sV정수64) S값(값 int64) V정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = 값
	return s
}
func (s *sV정수64) S절대값() V정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	if s.값 < 0 {
		s.값 = s.값 * -1
	}
	return s
}
func (s *sV정수64) S더하기(값 int64) V정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 + 값
	return s
}
func (s *sV정수64) S빼기(값 int64) V정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 - 값
	return s
}
func (s *sV정수64) S곱하기(값 int64) V정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 * 값
	return s
}
func (s *sV정수64) S나누기(값 int64) V정수 {
	if 값 == 0 {
		F문자열_출력("sV정수.S나누기() : 0으로 나눌 수 없음.")
		return nil
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 / 값
	return s
}
func (s *sV정수64) String() string {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
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

type sC부호없는_정수64 struct{ 값 uint64 }

func (s *sC부호없는_정수64) 상수형임()        {}
func (s *sC부호없는_정수64) G값() uint64   { return s.값 }
func (s *sC부호없는_정수64) G정수() int64   { return int64(s.값) }
func (s *sC부호없는_정수64) G실수() float64 { return float64(s.값) }
func (s *sC부호없는_정수64) G정밀수() C정밀수   { return NC정밀수(s.값) }
func (s *sC부호없는_정수64) G변수형() V부호없는_정수 {
	return NV부호없는_정수(s.값)
}
func (s *sC부호없는_정수64) String() string { return F문자열(s.값) }
func (s *sC부호없는_정수64) Generate(
	임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	return reflect.ValueOf(NC부호없는_정수(uint64(임의값_생성기.Uint32())))
}

type sV부호없는_정수64 struct {
	잠금 sync.RWMutex
	값  uint64
}

func (s *sV부호없는_정수64) 변수형임() {}
func (s *sV부호없는_정수64) G값() uint64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값
}
func (s *sV부호없는_정수64) G정수() int64 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return new(big.Int).SetUint64(값).Int64()
}
func (s *sV부호없는_정수64) G실수() float64 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return float64(값)
}
func (s *sV부호없는_정수64) G정밀수() C정밀수 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return NC정밀수(값)
}
func (s *sV부호없는_정수64) G상수형() C부호없는_정수 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return NC부호없는_정수(값)
}
func (s *sV부호없는_정수64) S값(값 uint64) V부호없는_정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값 = 값
	return s
}
func (s *sV부호없는_정수64) S더하기(값 uint64) V부호없는_정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값 = s.값 + 값
	return s
}
func (s *sV부호없는_정수64) S빼기(값 uint64) V부호없는_정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값 = s.값 - 값
	return s
}
func (s *sV부호없는_정수64) S곱하기(값 uint64) V부호없는_정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값 = s.값 * 값
	return s
}
func (s *sV부호없는_정수64) S나누기(값 uint64) V부호없는_정수 {
	if 값 == 0 {
		F문자열_출력("sV부호없는_정수.S나누기() : 0으로 나눌 수 없음.")
		return nil
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값 = s.값 / 값
	return s
}
func (s *sV부호없는_정수64) String() string {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
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
func (s *sC실수64) G실수() float64   { return s.값 }
func (s *sC실수64) G정밀수() C정밀수     { return NC정밀수(s.값) }
func (s *sC실수64) G변수형() V실수      { return NV실수(s.값) }
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
	값  float64
}

func (s *sV실수64) 변수형임() {}
func (s *sV실수64) G값() float64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값
}
func (s *sV실수64) G실수() float64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값
}
func (s *sV실수64) G정밀수() C정밀수 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return NC정밀수(값)
}
func (s *sV실수64) G상수형() C실수 {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
	return NC실수(값)
}
func (s *sV실수64) S값(값 float64) V실수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = 값
	return s
}
func (s *sV실수64) S절대값() V실수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	if s.값 < 0 {
		s.값 = s.값 * -1
	}
	return s
}
func (s *sV실수64) S더하기(값 float64) V실수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 + 값
	return s
}
func (s *sV실수64) S빼기(값 float64) V실수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 - 값
	return s
}
func (s *sV실수64) S곱하기(값 float64) V실수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 * 값
	return s
}
func (s *sV실수64) S나누기(값 float64) V실수 {
	if 값 == 0 {
		F문자열_출력("sV실수.S나누기() : 0으로 나눌 수 없음.")
		return nil
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.값 = s.값 / 값
	return s
}
func (s *sV실수64) String() string {
	s.잠금.RLock()
	값 := s.값
	s.잠금.RUnlock()
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

type s참거짓 struct{ 값 bool }

func (s *s참거짓) G값() bool       { return s.값 }
func (s *s참거짓) String() string { return strconv.FormatBool(s.값) }
func (s *s참거짓) generate_도우미(임의값_생성기 *rand.Rand) bool {
	정수 := int(임의값_생성기.Int31n(1))

	if 정수 == 0 {
		return true
	} else {
		return false
	}
}

type sC참거짓 struct{ *s참거짓 }

func (s *sC참거짓) 상수형임()      {}
func (s *sC참거짓) G변수형() V참거짓 { return NV참거짓(s.s참거짓.값) }
func (s *sC참거짓) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC참거짓(s.s참거짓.generate_도우미(임의값_생성기)))
}

type sV참거짓 struct {
	잠금 sync.RWMutex
	*s참거짓
}

func (s *sV참거짓) 변수형임() {}
func (s *sV참거짓) G값() bool {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s참거짓.값
}
func (s *sV참거짓) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s참거짓.String()
}
func (s *sV참거짓) G상수형() C참거짓 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return NC참거짓(s.s참거짓.값)
}
func (s *sV참거짓) S값(값 bool) V참거짓 {
	s.잠금.Lock()
	s.잠금.Unlock()
	s.s참거짓.값 = 값
	return s
}
func (s *sV참거짓) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NV참거짓(s.s참거짓.generate_도우미(임의값_생성기)))
}

type sC문자열 struct{ 값 string }

func (s *sC문자열) 상수형임()          {}
func (s *sC문자열) G값() string     { return s.값 }
func (s *sC문자열) String() string { return s.값 }
func (s *sC문자열) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	반복횟수_최대값 := int(임의값_생성기.Int31n(20))

	버퍼 := new(bytes.Buffer)
	var 길이 = int32(len(문자열_후보값_모음) - 1)

	for 반복횟수 := 0; 반복횟수 < 반복횟수_최대값; 반복횟수++ {
		버퍼.WriteString(문자열_후보값_모음[int(임의값_생성기.Int31n(길이))])
	}

	return reflect.ValueOf(NC문자열(버퍼.String()))
}

// 시점 (time.Time)
type sC시점 struct{ 값 time.Time }

func (s *sC시점) 상수형임()          {}
func (s *sC시점) G값() time.Time  { return s.값 }
func (s *sC시점) G변수형() V시점      { return NV시점(s.값) }
func (s *sC시점) String() string { return F문자열(s.값) }
func (s *sC시점) Generate(임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	연도 := int(1900 + 임의값_생성기.Int31n(300))
	월 := time.Month(int(1 + 임의값_생성기.Int31n(11)))
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
	값  time.Time
}

func (s *sV시점) 변수형임() {}
func (s *sV시점) G값() time.Time {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값
}
func (s *sV시점) G상수형() C시점 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return NC시점(s.값)
}
func (s *sV시점) S값(값 time.Time) V시점 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값 = 값
	return s
}
func (s *sV시점) S날짜_더하기(연, 월, 일 int) V시점 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.값.AddDate(연, 월, 일)
	return s
}
func (s *sV시점) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.값.Format(P시점_포맷)
}
func (s *sV시점) Generate(임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	연도 := int(1900 + 임의값_생성기.Int31n(300))
	월 := time.Month(int(1 + 임의값_생성기.Int31n(11)))
	일 := int(1 + 임의값_생성기.Int31n(30))
	시 := int(임의값_생성기.Int31n(24))
	분 := int(임의값_생성기.Int31n(59))
	초 := int(임의값_생성기.Int31n(59))
	나노초 := 임의값_생성기.Int()

	값 := time.Date(연도, 월, 일, 시, 분, 초, 나노초, time.Now().Location())

	return reflect.ValueOf(NV시점(값))
}

// 정밀수
type s정밀수 struct{ 값 *big.Rat }

func (s *s정밀수) G값() string { return s.String() }
func (s *s정밀수) GRat() *big.Rat {
	if s.값 == nil {
		return nil
	}

	return new(big.Rat).Set(s.값)
}
func (s *s정밀수) G실수() float64 {
	실수, 에러 := F문자열2실수(s.String())

	if 에러 != nil {
		실수, _ = s.값.Float64()
	}

	return 실수
}
func (s *s정밀수) G같음(값 interface{}) bool {
	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		return false
	}

	const 차이_한도 string = "1/1000000000000000000000000000000000000"

	차이_절대값 := new(big.Rat).Abs(new(big.Rat).Sub(s.GRat(), 정밀수.GRat()))

	if 차이_절대값.Cmp(NC정밀수(차이_한도).GRat()) == -1 {
		return true
	}

	return false
}

func (s *s정밀수) G비교(값 interface{}) int {
	if s.G같음(값) {
		return 0
	}

	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		return -2
	}

	return s.값.Cmp(정밀수.GRat())
}
func (s *s정밀수) String() string {
	return F마지막_0_제거(s.GRat().FloatString(100))
}
func (s *s정밀수) Generate도우미(임의값_생성기 *rand.Rand) s정밀수 {
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()

	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Float32() < 0.5 {
		분자 = 분자 * -1
	}

	return s정밀수{big.NewRat(분자, 분모)}
}

type sC정밀수 struct{ *s정밀수 }

func (s *sC정밀수) 상수형임()      {}
func (s *sC정밀수) G정밀수() C정밀수 { return s }
func (s *sC정밀수) G변수형() V정밀수 { return NV정밀수(s) }
func (s *sC정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	s내부 := s.s정밀수.Generate도우미(임의값_생성기)
	return reflect.ValueOf(&sC정밀수{&s내부})
}

type sV정밀수 struct {
	잠금 sync.RWMutex
	*s정밀수
}

func (s *sV정밀수) 변수형임() {}
func (s *sV정밀수) G값() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.String()
}
func (s *sV정밀수) GRat() *big.Rat {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.GRat()
}
func (s *sV정밀수) G실수() float64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.G실수()
}
func (s *sV정밀수) G정밀수() C정밀수 { return NC정밀수(s.G값()) }
func (s *sV정밀수) G상수형() C정밀수 { return NC정밀수(s.G값()) }
func (s *sV정밀수) G같음(값 interface{}) bool {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.G같음(값)
}
func (s *sV정밀수) G비교(값 interface{}) int {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.G비교(값)
}
func (s *sV정밀수) S값(값 interface{}) V정밀수 {
	if 값 == nil {
		s.잠금.Lock()
		defer s.잠금.Unlock()
		s.s정밀수.값 = nil
		return s
	}

	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		s.잠금.Lock()
		defer s.잠금.Unlock()
		s.s정밀수.값 = nil
		return nil
	}

	bigRat := 정밀수.GRat()

	s.잠금.Lock()
	defer s.잠금.Unlock()
	if s.s정밀수.값 == nil {
		s.s정밀수.값 = new(big.Rat)
	}

	s.s정밀수.값.Set(bigRat)

	return s
}

// 주어진 함수의 결과값에 대해서 CAS(Compare And Set).
func (s *sV정밀수) S_CAS(원래값, 새로운값 *big.Rat) bool {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	switch {
	case s.s정밀수.값 == nil && 원래값 == nil:
		s.s정밀수.값 = new(big.Rat).Set(새로운값)
		return true
	case s.s정밀수.값 != nil && s.s정밀수.값.Cmp(원래값) == 0:
		if 새로운값 == nil {
			s.s정밀수.값 = nil
		} else {
			s.s정밀수.값.Set(새로운값)
		}

		return true
	default:
		return false
	}
}

func (s *sV정밀수) S_CAS_함수(
	함수 func(*sV정밀수, ...interface{}) *big.Rat,
	매개변수 ...interface{}) V정밀수 {
	var 현재값, 새로운값 *big.Rat
	var 반복횟수 int = 0

	for {
		현재값 = s.GRat()

		// 넘겨받은 함수를 실행.
		// 실행이 마치기 전에 다른 goroutine에서 내부값을 변경할 수 있음.
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.S_CAS(현재값, 새로운값) {
			return s
		}

		반복횟수++

		// Exponential Back-off
		F잠시_대기(반복횟수)
	}
}

func (s *sV정밀수) S반올림(소숫점_이하_자릿수 int) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil {
				return nil
			}

			소숫점_이하_자릿수 := 매개변수[0].(int)
			문자열 := 정밀수.FloatString(소숫점_이하_자릿수)
			return NC정밀수(문자열).GRat()
		}

	return s.S_CAS_함수(함수, 소숫점_이하_자릿수)
}

func (s *sV정밀수) S절대값() V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil {
				return nil
			}

			return 정밀수.Abs(정밀수)
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정밀수) S더하기(값 interface{}) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			값 := 매개변수[0]

			if 값 == nil {
				return nil
			}

			c정밀수 := NC정밀수(값)
			원래값 := 구조체.GRat()

			if c정밀수 == nil ||
				c정밀수.GRat() == nil ||
				원래값 == nil {
				return nil
			}

			정밀수 := c정밀수.GRat()

			return 정밀수.Add(원래값, 정밀수)
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV정밀수) S빼기(값 interface{}) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			값 := 매개변수[0]

			if 값 == nil {
				return nil
			}

			c정밀수 := NC정밀수(값)
			원래값 := 구조체.GRat()

			if c정밀수 == nil ||
				c정밀수.GRat() == nil ||
				원래값 == nil {
				return nil
			}

			정밀수 := c정밀수.GRat()

			return 정밀수.Sub(원래값, 정밀수)
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV정밀수) S곱하기(값 interface{}) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			값 := 매개변수[0]

			if 값 == nil {
				return nil
			}

			c정밀수 := NC정밀수(값)
			원래값 := 구조체.GRat()

			if c정밀수 == nil ||
				c정밀수.GRat() == nil ||
				원래값 == nil {
				return nil
			}

			정밀수 := c정밀수.GRat()

			return 정밀수.Mul(원래값, 정밀수)
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV정밀수) S나누기(값 interface{}) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			값 := 매개변수[0]

			if 값 == nil {
				return nil
			}

			c정밀수 := NC정밀수(값)
			원래값 := 구조체.GRat()

			if c정밀수 == nil ||
				c정밀수.GRat() == nil ||
				원래값 == nil {
				return nil
			}

			정밀수 := c정밀수.GRat()

			if 정밀수.Cmp(big.NewRat(0, 1)) == 0 {
				return nil
			}

			return 정밀수.Quo(원래값, 정밀수)
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV정밀수) S역수() V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil ||
				정밀수.Cmp(big.NewRat(0, 1)) == 0 {
				return nil
			}

			return 정밀수.Inv(정밀수)
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정밀수) S반대부호값() V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...interface{}) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil {
				return nil
			}

			return 정밀수.Neg(정밀수)
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정밀수) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.String()
}
func (s *sV정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	s내부값 := s.s정밀수.Generate도우미(임의값_생성기)
	return reflect.ValueOf(&sV정밀수{s정밀수: &s내부값})
}

// 통화
type sC통화 struct {
	종류 P통화종류
	금액 C정밀수
}

func (s *sC통화) 상수형임()      {}
func (s *sC통화) G종류() P통화종류 { return s.종류 }
func (s *sC통화) G값() C정밀수   { return s.금액 }
func (s *sC통화) G같음(값 I통화) bool {
	if 값 == nil {
		return false
	}

	if s.종류 == 값.G종류() &&
		s.금액.G같음(값.G값()) {
		return true
	}

	return false
}
func (s *sC통화) G변수형() V통화 { return NV통화(s.종류, s.금액) }
func (s *sC통화) String() string {
	return s.종류.String() + " " + F금액_문자열(s.금액.String())
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
	종류 P통화종류
	금액 V정밀수
}

func (s *sV통화) 변수형임() {}
func (s *sV통화) G종류() P통화종류 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.종류
}
func (s *sV통화) G값() C정밀수 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	return NC정밀수(s.금액)
}
func (s *sV통화) G같음(값 I통화) bool {
	if 값 == nil {
		return false
	}

	통화종류 := 값.G종류()
	금액 := 값.G값()

	s.잠금.RLock()
	defer s.잠금.RUnlock()
	if s.종류 == 통화종류 &&
		s.금액.G같음(금액) {
		return true
	}

	return false
}
func (s *sV통화) G상수형() C통화 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	return NC통화(s.종류, s.금액)
}
func (s *sV통화) S종류(종류 P통화종류) {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.종류 = 종류
}
func (s *sV통화) S값(금액 interface{}) V통화 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S값(금액)
	return s
}

func (s *sV통화) S절대값() V통화 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S절대값()
	return s
}
func (s *sV통화) S더하기(값 interface{}) V통화 {
	if 종류 := F통화종류(s, 값); 종류 == INVALID_CURRENCY_TYPE {
		s.잠금.Lock()
		defer s.잠금.Unlock()
		s.금액.S값(nil)
		return s
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S더하기(값)
	return s
}
func (s *sV통화) S빼기(값 interface{}) V통화 {
	if 종류 := F통화종류(s, 값); 종류 == INVALID_CURRENCY_TYPE {
		s.잠금.Lock()
		defer s.잠금.Unlock()
		s.금액.S값(nil)
		return s
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S빼기(값)
	return s
}
func (s *sV통화) S곱하기(값 interface{}) V통화 {
	if 종류 := F통화종류(s, 값); 종류 == INVALID_CURRENCY_TYPE {
		s.잠금.Lock()
		defer s.잠금.Unlock()
		s.금액.S값(nil)
		return s
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S곱하기(값)
	return s
}
func (s *sV통화) S나누기(값 interface{}) V통화 {
	if 종류 := F통화종류(s, 값); 종류 == INVALID_CURRENCY_TYPE {
		s.잠금.Lock()
		defer s.잠금.Unlock()
		s.금액.S값(nil)
		return s
	}

	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S나누기(값)
	return s
}
func (s *sV통화) S반대부호값() V통화 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.금액.S반대부호값()
	return s
}
func (s *sV통화) String() string {
	var 금액 *big.Rat

	s.잠금.RLock()
	금액 = s.금액.GRat()
	s.잠금.RUnlock()

	return s.종류.String() + " " + F금액_문자열(F문자열(금액))
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
	값  I상수형
}

func (s *sC매개변수) G이름() string { return s.이름 }
func (s *sC매개변수) G값() I상수형    { return s.값 }
