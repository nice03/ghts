// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/gh-system/ghts/dep/ps"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type s반환값 struct {
	값  I가변형
	에러 error
}

func (s *s반환값) G값() I가변형   { return s.값 }
func (s *s반환값) G에러() error { return s.에러 }

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

// 정수

type s정수64 struct{ 값 int64 }

func (s *s정수64) G값() int64      { return atomic.LoadInt64(&(s.값)) }
func (s *s정수64) G정수() int64     { return s.G값() }
func (s *s정수64) G실수() float64   { return float64(s.G값()) }
func (s *s정수64) G정밀수() C정밀수     { return NC정밀수(s.G값()) }
func (s *s정수64) String() string { return strconv.FormatInt(s.G값(), 10) }
func (s *s정수64) generate(임의값_생성기 *rand.Rand) int64 {
	값 := 임의값_생성기.Int63() // 0 혹은 양의 정수.

	if 임의값_생성기.Int31n(2) == 0 {
		값 = 값 * -1
	}

	return 값
}

type sC정수64 struct{ *s정수64 }

func (s *sC정수64) G상수형임()     {}
func (s *sC정수64) G변수형() V정수 { return NV정수(s.G값()) }
func (s *sC정수64) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC정수(s.s정수64.generate(임의값_생성기)))
}

type sV정수64 struct {
	*s정수64
}

func (s *sV정수64) G변수형임() {}
func (s *sV정수64) G상수형() C정수 {
	return NC정수(s.G값())
}
func (s *sV정수64) S값(값 int64) V정수 {
	atomic.StoreInt64(&(s.s정수64.값), 값)
	return s
}

// CAS(Compare And Swap)
func (s *sV정수64) s_CAS(원래값, 새로운값 int64) bool {
	return atomic.CompareAndSwapInt64(&(s.s정수64.값), 원래값, 새로운값)
}

// CAS(Compare And Swap)
func (s *sV정수64) s_CAS_함수(함수 func(*sV정수64, ...I가변형) int64,
	매개변수 ...I가변형) V정수 {
	var 원래값, 새로운값 int64
	var 반복횟수 int = 0

	for {
		원래값 = s.G값()
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.s_CAS(원래값, 새로운값) {
			return s
		}

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.s_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
		반복횟수++
		F잠시_대기(반복횟수)
	}
}
func (s *sV정수64) S절대값() V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()

			if 정수 < 0 {
				정수 = 정수 * -1
			}

			return 정수
		}

	return s.s_CAS_함수(함수)
}
func (s *sV정수64) S더하기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()
			값 := 매개변수[0].(int64)

			정수 = 정수 + 값

			return 정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정수64) S빼기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()
			값 := 매개변수[0].(int64)

			정수 = 정수 - 값

			return 정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정수64) S곱하기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()
			값 := 매개변수[0].(int64)

			정수 = 정수 * 값

			return 정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정수64) S나누기(값 int64) V정수 {
	if 값 == 0 {
		F문자열_출력("0으로 나눌 수 없음.")
		s.s정수64 = nil
		
		panic(F에러_생성("0으로 나눌 수 없음."))
		
		return nil
	}

	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()
			값 := 매개변수[0].(int64)

			정수 = 정수 / 값
			
			return 정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정수64) String() string {
	return s.s정수64.String()
}
func (s *sV정수64) Generate(
	임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	return reflect.ValueOf(NV정수(s.s정수64.generate(임의값_생성기)))
}

// 부호없는_정수

type s부호없는_정수64 struct{ 값 uint64 }

func (s *s부호없는_정수64) G값() uint64     { return atomic.LoadUint64(&(s.값)) }
func (s *s부호없는_정수64) G정수() int64     { return int64(s.G값()) }
func (s *s부호없는_정수64) G실수() float64   { return float64(s.G값()) }
func (s *s부호없는_정수64) G정밀수() C정밀수     { return NC정밀수(s.G값()) }
func (s *s부호없는_정수64) String() string { return strconv.FormatUint(s.G값(), 10) }

type sC부호없는_정수64 struct{ *s부호없는_정수64 }

func (s *sC부호없는_정수64) G상수형임() {}
func (s *sC부호없는_정수64) G변수형() V부호없는_정수 {
	return NV부호없는_정수(s.G값())
}
func (s *sC부호없는_정수64) Generate(
	임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC부호없는_정수(uint64(임의값_생성기.Int63())))
}

type sV부호없는_정수64 struct {
	*s부호없는_정수64
}

func (s *sV부호없는_정수64) G변수형임() {}
func (s *sV부호없는_정수64) G상수형() C부호없는_정수 {
	return NC부호없는_정수(s.G값())
}
func (s *sV부호없는_정수64) S값(값 uint64) V부호없는_정수 {
	atomic.StoreUint64(&(s.s부호없는_정수64.값), 값)
	return s
}

// CAS(Compare And Swap)
func (s *sV부호없는_정수64) s_CAS(원래값, 새로운값 uint64) bool {
	return atomic.CompareAndSwapUint64(&(s.s부호없는_정수64.값), 원래값, 새로운값)
}

// CAS(Compare And Swap)
func (s *sV부호없는_정수64) s_CAS_함수(
	함수 func(*sV부호없는_정수64, ...I가변형) uint64,
	매개변수 ...I가변형) V부호없는_정수 {
	var 원래값, 새로운값 uint64
	var 반복횟수 int = 0

	for {
		원래값 = s.G값()
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.s_CAS(원래값, 새로운값) {
			return s
		}

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.s_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
		반복횟수++
		F잠시_대기(반복횟수)
	}
}
func (s *sV부호없는_정수64) S더하기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()
			값 := 매개변수[0].(uint64)

			부호없는_정수 = 부호없는_정수 + 값

			return 부호없는_정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV부호없는_정수64) S빼기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()
			값 := 매개변수[0].(uint64)

			부호없는_정수 = 부호없는_정수 - 값

			return 부호없는_정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV부호없는_정수64) S곱하기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()
			값 := 매개변수[0].(uint64)

			부호없는_정수 = 부호없는_정수 * 값

			return 부호없는_정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV부호없는_정수64) S나누기(값 uint64) V부호없는_정수 {
	if 값 == 0 {
		F문자열_출력("0으로 나눌 수 없음.")
		s.s부호없는_정수64 = nil
		
		panic(F에러_생성("0으로 나눌 수 없음."))
		
		return nil
	}

	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()
			값 := 매개변수[0].(uint64)

			부호없는_정수 = 부호없는_정수 / 값

			return 부호없는_정수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV부호없는_정수64) Generate(
	임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	return reflect.ValueOf(NV부호없는_정수(uint64(임의값_생성기.Int63())))
}

// 실수
type s실수64 struct{ 값 float64 }

func (s *s실수64) G값() float64  { return s.값 }
func (s *s실수64) G실수() float64 { return s.G값() }
func (s *s실수64) G정밀수() C정밀수   { return NC정밀수(s.G값()) }
func (s *s실수64) String() string {
	return strconv.FormatFloat(s.G값(), 'f', -1, 64)
}
func (s *s실수64) generate(임의값_생성기 *rand.Rand) float64 {
	값 := 임의값_생성기.Float64() * float64(임의값_생성기.Int63())

	if 임의값_생성기.Int31n(2) == 0 {
		값 = 값 * -1.0
	}

	return 값
}

type sC실수64 struct{ *s실수64 }

func (s *sC실수64) G상수형임()     {}
func (s *sC실수64) G변수형() V실수 { return NV실수(s.s실수64.값) }
func (s *sC실수64) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC실수(s.s실수64.generate(임의값_생성기)))
}

type sV실수64 struct {
	잠금 sync.RWMutex
	*s실수64
}

func (s *sV실수64) G변수형임() {}
func (s *sV실수64) G값() float64 {
	var 값 float64
	s.잠금.RLock()
	값 = s.s실수64.값
	s.잠금.RUnlock()

	return 값
}
func (s *sV실수64) G실수() float64 { return s.G값() }
func (s *sV실수64) G정밀수() C정밀수   { return NC정밀수(s.G값()) }
func (s *sV실수64) G상수형() C실수    { return NC실수(s.G값()) }
func (s *sV실수64) S값(값 float64) V실수 {
	s.잠금.Lock()
	s.s실수64.값 = 값
	s.잠금.Unlock()

	return s
}
// CAS(Compare And Set). float64에 대해서는 내장된 atomic 함수가 없음.
func (s *sV실수64) s_CAS_함수(
	함수 func(*sV실수64, ...I가변형) float64,
	매개변수 ...I가변형) V실수 {
	var 원래값, 새로운값 float64
	var 반복횟수 int = 0

	for {
		원래값 = s.G값()
		새로운값 = 함수(s, 매개변수...)
		
		s.잠금.Lock()

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if 원래값 == s.s실수64.값 {
			s.s실수64.값 = 새로운값
			s.잠금.Unlock()
			
			return s
		}
		
		s.잠금.Unlock()

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.s_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
		반복횟수++
		F잠시_대기(반복횟수)
	}
}
func (s *sV실수64) S절대값() V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			if 실수 < 0 {
				실수 = 실수 * -1.0
			}

			return 실수
		}

	return s.s_CAS_함수(함수)
}
func (s *sV실수64) S더하기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 + 값

			return 실수
		}

	return s.s_CAS_함수(함수)
}
func (s *sV실수64) S빼기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 - 값

			return 실수
		}

	return s.s_CAS_함수(함수)
}
func (s *sV실수64) S곱하기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 * 값

			return 실수
		}

	return s.s_CAS_함수(함수)
}
func (s *sV실수64) S나누기(값 float64) V실수 {
	if 값 == 0.0 {
		F문자열_출력("0으로 나눌 수 없음.")
		
		if s.s실수64.값 * 값 >= 0 {
			s.s실수64.값 = F양의_무한()
		} else {
			s.s실수64.값 = F음의_무한()
		}

		return s
	}

	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 / 값

			return 실수
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV실수64) S역수() V실수 {
	if s.G값() == 0.0 {
		F문자열_출력("0의 역수를 구할 수 없음.")
		
		s.s실수64.값 = F양의_무한()
	
		return s
	}

	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 1 / 실수

			return 실수
		}

	return s.s_CAS_함수(함수)
}
func (s *sV실수64) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s실수64.String()
}
func (s *sV실수64) Generate(
	임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NV실수(s.s실수64.generate(임의값_생성기)))
}

// 참거짓
type s참거짓 struct{ 값 bool }

func (s *s참거짓) G값() bool       { return s.값 }
func (s *s참거짓) String() string { return strconv.FormatBool(s.값) }
func (s *s참거짓) generate(임의값_생성기 *rand.Rand) bool {
	if 임의값_생성기.Int31n(2) == 0 {
		return false
	}

	return true
}

type sC참거짓 struct{ *s참거짓 }

func (s *sC참거짓) G상수형임()      {}
func (s *sC참거짓) G변수형() V참거짓 { return NV참거짓(s.s참거짓.값) }
func (s *sC참거짓) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC참거짓(s.s참거짓.generate(임의값_생성기)))
}

type sV참거짓 struct {
	잠금 sync.RWMutex
	*s참거짓
}

func (s *sV참거짓) G변수형임() {}
func (s *sV참거짓) G값() bool {
	var 값 bool
	s.잠금.RLock()
	값 = s.s참거짓.값
	s.잠금.RUnlock()

	return 값
}
func (s *sV참거짓) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	return s.s참거짓.String()
}
func (s *sV참거짓) G상수형() C참거짓 {
	return NC참거짓(s.G값())
}
func (s *sV참거짓) S값(값 bool) V참거짓 {
	s.잠금.Lock()
	s.s참거짓.값 = 값
	s.잠금.Unlock()

	return s
}
func (s *sV참거짓) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NV참거짓(s.s참거짓.generate(임의값_생성기)))
}

// 문자열
type sC문자열 struct{ 값 string }

func (s *sC문자열) G상수형임()          {}
func (s *sC문자열) G값() string     { return s.값 }
func (s *sC문자열) String() string { return s.값 }
func (s *sC문자열) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	반복횟수_최대값 := int(임의값_생성기.Int31n(20))
	길이 := int32(len(문자열_후보값_모음))
	버퍼 := new(bytes.Buffer)

	for 반복횟수 := 0; 반복횟수 < 반복횟수_최대값; 반복횟수++ {
		버퍼.WriteString(
			문자열_후보값_모음[int(임의값_생성기.Int31n(길이))])
	}

	return reflect.ValueOf(NC문자열(버퍼.String()))
}

// 시점 (time.Time)
type s시점 struct{ 값 time.Time }

func (s *s시점) G값() time.Time  { return s.값 }
func (s *s시점) String() string { return s.값.Format(P시점_형식) }
func (s *s시점) generate(임의값_생성기 *rand.Rand) time.Time {
	연도_차이 := int(임의값_생성기.Int31n(100))

	if 임의값_생성기.Int31n(2) == 0 {
		연도_차이 = 연도_차이 * -1
	}

	연도 := time.Now().Year() + 연도_차이
	월 := time.Month(int(1 + 임의값_생성기.Int31n(12)))
	일 := int(1 + 임의값_생성기.Int31n(31))
	시 := int(임의값_생성기.Int31n(24))
	분 := int(임의값_생성기.Int31n(60))
	초 := int(임의값_생성기.Int31n(60))
	나노초 := time.Now().Nanosecond() + int(임의값_생성기.Int31())

	return time.Date(연도, 월, 일, 시, 분, 초, 나노초, time.Now().Location())
}

type sC시점 struct{ *s시점 }

func (s *sC시점) G상수형임()     {}
func (s *sC시점) G변수형() V시점 { return NV시점(s.s시점.값) }
func (s *sC시점) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC시점(s.s시점.generate(임의값_생성기)))
}

type sV시점 struct {
	잠금 sync.RWMutex
	*s시점
}

func (s *sV시점) G변수형임() {}
func (s *sV시점) G값() time.Time {
	var 값 time.Time
	s.잠금.RLock()
	값 = s.s시점.값
	s.잠금.RUnlock()

	return 값
}
func (s *sV시점) G상수형() C시점 { return NC시점(s.G값()) }
func (s *sV시점) S값(값 time.Time) V시점 {
	s.잠금.Lock()
	s.s시점.값 = 값 // time.Time은 매개변수로 넘어오면서 자동으로 복사됨.
	s.잠금.Unlock()

	return s
}
func (s *sV시점) S일자_더하기(연, 월, 일 int) V시점 {
	s.잠금.Lock()
	s.s시점.값 = s.s시점.값.AddDate(연, 월, 일)
	s.잠금.Unlock()

	return s
}
func (s *sV시점) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s시점.String()
}
func (s *sV시점) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NV시점(s.s시점.generate(임의값_생성기)))
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
	실수, 에러 := F문자열2실수(s.G값())

	if 에러 != nil {
		실수, _ = s.값.Float64()
	}

	return 실수
}
func (s *s정밀수) G같음(값 I가변형) bool {
	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		return false
	}

	차이_절대값 := new(big.Rat)
	차이_절대값 = 차이_절대값.Sub(s.GRat(), 정밀수.GRat())
	차이_절대값 = 차이_절대값.Abs(차이_절대값)

	// '차이_한도'는 init, const, 공통 변수 모아놓은 곳에 정의되어 있음.
	if 차이_절대값.Cmp(차이_한도.GRat()) == -1 {
		return true
	}

	return false
}

func (s *s정밀수) G비교(값 I가변형) int {
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
	값 := s.GRat()

	if F_nil값임(값) {
		return "<nil>"
	}

	return F마지막_0_제거(값.FloatString(100))
}
func (s *s정밀수) Generate도우미(임의값_생성기 *rand.Rand) s정밀수 {
	분자 := 임의값_생성기.Int63()
	분모 := 임의값_생성기.Int63()

	// 0.0 ~ 1.0 임의값 생성 후 0.5 기준으로 부호 결정.
	if 임의값_생성기.Int31n(2) == 0 {
		분자 = 분자 * -1
	}

	return s정밀수{big.NewRat(분자, 분모)}
}

type sC정밀수 struct{ *s정밀수 }

func (s *sC정밀수) G상수형임()      {}
func (s *sC정밀수) G정밀수() C정밀수 { return s }
func (s *sC정밀수) G변수형() V정밀수 { return NV정밀수(s) }
func (s *sC정밀수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	s내부값 := s.s정밀수.Generate도우미(임의값_생성기)
	return reflect.ValueOf(&sC정밀수{&s내부값})
}

type sV정밀수 struct {
	잠금 sync.RWMutex
	*s정밀수
}

func (s *sV정밀수) G변수형임() {}
func (s *sV정밀수) G값() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.String()
}
func (s *sV정밀수) GRat() *big.Rat {
	var 값 *big.Rat
	s.잠금.RLock()
	값 = s.s정밀수.GRat()
	s.잠금.RUnlock()

	return 값
}
func (s *sV정밀수) G실수() float64 {
	var 값 float64
	s.잠금.RLock()
	값 = s.s정밀수.G실수()
	s.잠금.RUnlock()

	return 값
}
func (s *sV정밀수) G정밀수() C정밀수 { return NC정밀수(s.GRat()) }
func (s *sV정밀수) G상수형() C정밀수 { return NC정밀수(s.GRat()) }
func (s *sV정밀수) G같음(값 I가변형) bool {
	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		return false
	}

	var 내부값 *big.Rat
	s.잠금.RLock()
	내부값 = s.GRat()
	s.잠금.RUnlock()

	return 정밀수.G같음(내부값)
}
func (s *sV정밀수) G비교(값 I가변형) int {
	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		return -2
	}

	var 내부값 *big.Rat
	s.잠금.RLock()
	내부값 = s.GRat()
	s.잠금.RUnlock()

	// 비교 순서가 거꾸로 이므로, 부호를 바꿔준다.
	return (정밀수.G비교(내부값) * -1)
}
func (s *sV정밀수) S값(값 I가변형) V정밀수 {
	정밀수 := NC정밀수(값)

	if 정밀수 == nil {
		return nil
	}

	bigRat값 := 정밀수.GRat()

	if bigRat값 == nil {
		return nil
	}

	s.잠금.Lock()
	s.s정밀수.값.Set(bigRat값)
	s.잠금.Unlock()

	return s
}

// 주어진 값에 대해서 CAS(Compare And Set).
func (s *sV정밀수) s_CAS(원래값, 새로운값 *big.Rat) bool {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	if s.s정밀수.값 == 원래값 {
		s.s정밀수.값.Set(새로운값)
		return true
	}

	return false
}

func (s *sV정밀수) s_CAS_함수(함수 func(*sV정밀수, ...I가변형) *big.Rat,
	매개변수 ...I가변형) (반환값 V정밀수) {
	var 쓰기_잠금 bool = false
	
	// 에러발생 하면 nil 반환
	defer func() {
		if 에러 := recover(); 에러 != nil {
			if !쓰기_잠금 { s.잠금.Lock() }
			defer s.잠금.Unlock()
			s.s정밀수.값 = nil
			반환값 = nil
		}
	}()

	var 원래값, 새로운값 *big.Rat
	var 반복횟수 int = 0

	for {
		s.잠금.RLock()
		원래값 = s.s정밀수.값
		s.잠금.RUnlock()

		새로운값 = 함수(s, 매개변수...)

		if 새로운값 != nil {
			// 복사본을 형성해서 독립성 보장.
			새로운값 = new(big.Rat).Set(새로운값)
		}

		s.잠금.Lock()
		쓰기_잠금 = true

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.s정밀수.값.Cmp(원래값) == 0 {
			s.s정밀수.값.Set(새로운값)
			s.잠금.Unlock()
			쓰기_잠금 = false

			return s
		}

		s.잠금.Unlock()
		쓰기_잠금 = false

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는 처음부터 새로 계산함.
		// 이렇게 약간의 비효율성이 발생하지만, 대신에 골치아픈 데드락이 원천적으로 발생하지 않음.
		반복횟수++
		F잠시_대기(반복횟수)
	}
}

func (s *sV정밀수) S반올림(소숫점_이하_자릿수 int) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			정밀수 := 구조체.GRat()
			소숫점_이하_자릿수 := 매개변수[0].(int)

			if 정밀수 == nil {
				return nil
			}

			문자열 := 정밀수.FloatString(소숫점_이하_자릿수)

			return NC정밀수(문자열).GRat()
		}

	return s.s_CAS_함수(함수, 소숫점_이하_자릿수)
}

func (s *sV정밀수) S절대값() V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil {
				return nil
			}

			return 정밀수.Abs(정밀수)
		}

	return s.s_CAS_함수(함수)
}
func (s *sV정밀수) S더하기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			값 := NC정밀수(매개변수[0])

			if 값 == nil {
				return nil
			}

			정밀수 := 값.GRat()
			원래값 := 구조체.GRat()

			if 정밀수 == nil || 원래값 == nil {
				return nil
			}

			return 정밀수.Add(원래값, 정밀수)
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정밀수) S빼기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			값 := NC정밀수(매개변수[0])

			if 값 == nil {
				return nil
			}

			정밀수 := 값.GRat()
			원래값 := 구조체.GRat()

			if 정밀수 == nil || 원래값 == nil {
				return nil
			}

			return 정밀수.Sub(원래값, 정밀수)
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정밀수) S곱하기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			값 := NC정밀수(매개변수[0])

			if 값 == nil {
				return nil
			}

			정밀수 := 값.GRat()
			원래값 := 구조체.GRat()

			if 정밀수 == nil || 원래값 == nil {
				return nil
			}

			return 정밀수.Mul(원래값, 정밀수)
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정밀수) S나누기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			값 := NC정밀수(매개변수[0])

			if 값 == nil {
				return nil
			}

			정밀수 := 값.GRat()
			원래값 := 구조체.GRat()

			if 정밀수 == nil || 원래값 == nil {
				return nil
			}

			if 정밀수.Cmp(big.NewRat(0, 1)) == 0 {
				return nil
			}

			return 정밀수.Quo(원래값, 정밀수)
		}

	return s.s_CAS_함수(함수, 값)
}
func (s *sV정밀수) S역수() V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil ||
				정밀수.Cmp(big.NewRat(0, 1)) == 0 {
				return nil
			}

			return 정밀수.Inv(정밀수)
		}

	return s.s_CAS_함수(함수)
}
func (s *sV정밀수) S반대부호값() V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil {
				return nil
			}

			return 정밀수.Neg(정밀수)
		}

	return s.s_CAS_함수(함수)
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
type P통화종류 int

var 통화종류_문자열_모음 = [...]string{"KRW", "USD", "CNY", "EUR"}

func (p P통화종류) String() string {
	if int(p) == -1 {
		return "INVALID_CURRENCY_TYPE"
	}

	return 통화종류_문자열_모음[p]
}

type sC통화 struct {
	종류 P통화종류
	금액 C정밀수
}

func (s *sC통화) G상수형임()      {}
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
func (s *sC통화) G비교(값 I통화) int {
	switch {
	case F_nil값임(s) && F_nil값임(값):
		if s.G종류() == 값.G종류() {
			return 0
		} else {
			return -2
		}
	case F_nil값임(s) || F_nil값임(값):
		return -2
	case s.G종류() != 값.G종류():
		return -2
	}

	return s.G값().G비교(값.G값())
}
func (s *sC통화) G변수형() V통화 { return NV통화(s.종류, s.금액) }
func (s *sC통화) String() string {
	return s.종류.String() + " " + F금액_문자열(s.금액.String())
}
func (s *sC통화) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	종류 := F임의_통화종류()
	금액 := 0.0

	for math.Abs(금액) < 1 {
		분자 := 임의값_생성기.Int63()

		if 임의값_생성기.Int31n(2) == 0 {
			분자 = 분자 * -1
		}

		금액, _ = big.NewRat(분자, 1000).Float64()
	}

	return reflect.ValueOf(NC통화(종류, 금액))
}

// V정밀수에 이미 내부적으로 Mutex가 구현되어 있으므로,
// V통화에서 굳이 다시 Mutex를 사용할 필요가 없음.
// 통화종류는 자주 바뀌는 항목이 아니니, 일관성에 크게 신경쓸 필요없을 듯.
type sV통화 struct {
	종류 P통화종류
	금액 V정밀수
}

func (s *sV통화) G변수형임()      {}
func (s *sV통화) G종류() P통화종류 { return s.종류 }
func (s *sV통화) G값() C정밀수   { return s.금액.G상수형() }
func (s *sV통화) G같음(값 I통화) bool {
	if 값 == nil ||
		s.종류 != 값.G종류() ||
		!s.금액.G같음(값.G값()) {
		return false
	}

	return true
}
func (s *sV통화) G비교(값 I통화) int {
	switch {
	case F_nil값임(s) && F_nil값임(값):
		if s.G종류() == 값.G종류() {
			return 0
		} else {
			return -2
		}
	case F_nil값임(s) || F_nil값임(값):
		return -2
	case s.G종류() != 값.G종류():
		return -2
	}

	return s.G값().G비교(값.G값())
}
func (s *sV통화) G상수형() C통화 {
	return NC통화(s.종류, s.금액.G상수형())
}
func (s *sV통화) S값(금액 I가변형) V통화 {
	s.금액.S값(금액)

	return s
}

func (s *sV통화) S절대값() V통화 {
	s.금액.S절대값()
	return s
}
func (s *sV통화) g금액(값 I가변형) I가변형 {
	switch {
	case F_nil값임(값):
		return nil
	case F통화형식임(값):
		if s.종류 != 값.(I통화).G종류() {
			return nil
		}

		return 값.(I통화).G값()
	case F숫자형식임(값):
		return 값
	}

	return nil
}
func (s *sV통화) S더하기(값 I가변형) V통화 {
	금액 := s.g금액(값)

	if 금액 == nil {
		return nil
	}

	s.금액.S더하기(금액)

	return s
}
func (s *sV통화) S빼기(값 I가변형) V통화 {
	금액 := s.g금액(값)

	if 금액 == nil {
		return nil
	}

	s.금액.S빼기(금액)

	return s
}
func (s *sV통화) S곱하기(값 I가변형) V통화 {
	금액 := s.g금액(값)

	if 금액 == nil {
		return nil
	}

	s.금액.S곱하기(금액)

	return s
}
func (s *sV통화) S나누기(값 I가변형) V통화 {
	금액 := s.g금액(값)

	if 금액 == nil {
		return nil
	}

	s.금액.S나누기(금액)

	return s
}
func (s *sV통화) S반대부호값() V통화 {
	s.금액.S반대부호값()
	return s
}
func (s *sV통화) String() string {
	return s.종류.String() + " " + F금액_문자열(s.금액.String())
}
func (s *sV통화) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	종류 := F임의_통화종류()
	금액 := 0.0

	for math.Abs(금액) < 1 {
		분자 := 임의값_생성기.Int63()

		if 임의값_생성기.Int31n(2) == 0 {
			분자 = 분자 * -1
		}

		금액, _ = big.NewRat(분자, 1000).Float64()
	}

	return reflect.ValueOf(NV통화(종류, 금액))
}

// 생성자에서 매개변수의 안전성을 검사하도록 할 것.
type sC매개변수 struct {
	이름 string
	값  I가변형
}

func (s *sC매개변수) G상수형임()         {}
func (s *sC매개변수) G이름() string   { return s.이름 }
func (s *sC매개변수) G값() I가변형      { return s.값 }
func (s *sC매개변수) G숫자형식임() bool  { return F숫자형식임(s.값) }
func (s *sC매개변수) G문자열형식임() bool { return F문자열형식임(s.값) }
func (s *sC매개변수) G시점형식임() bool  { return F시점형식임(s.값) }
func (s *sC매개변수) G참거짓형식임() bool { return F참거짓형식임(s.값) }
func (s *sC매개변수) String() string {
	return s.이름 + " " + reflect.TypeOf(s.값).String() + F문자열(s.값)
}
func (s *sC매개변수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	임의_숫자값 := 임의값_생성기.Int31()
	임의_문자열 := NC문자열("").(*sC문자열).Generate(임의값_생성기, 크기).Interface().(string)

	값_모음 := []I가변형{uint(임의_숫자값), int(임의_숫자값), float32(임의_숫자값),
		float64(임의_숫자값), time.Now(), true, false, 임의_문자열,
		NC정수(int64(임의_숫자값)), NC부호없는_정수(uint64(임의_숫자값)),
		NC실수(float64(임의_숫자값)), NC정밀수(임의_숫자값),
		NC통화(F임의_통화종류(), 임의_숫자값), NC시점(time.Now()),
		NC참거짓(true), NC문자열(임의_문자열)}

	이름 := 임의_문자열
	값 := 값_모음[임의값_생성기.Int31n(int32(len(값_모음)))]

	return reflect.ValueOf(NC매개변수(이름, 값))
}

// 안전한 가변형
type sC안전한_가변형 struct{ 값 I가변형 }

func (s *sC안전한_가변형) G상수형임()    {}
func (s *sC안전한_가변형) G값() I가변형 { return s.값 }
func (s *sC안전한_가변형) String() string {
	if _, ok := s.값.(I기본_문자열); ok {
		return s.값.(I기본_문자열).String()
	}
	
	return F포맷된_문자열("%v", s.값)
}

// 안전한 배열
type s안전한_배열 struct{ 값 ps.List }

func (s *s안전한_배열) G비어있음() bool { return s.값.IsNil() }
func (s *s안전한_배열) G길이() int    { return s.값.Size() }
func (s *s안전한_배열) S추가(값 I가변형) I안전한_배열 {
	F매개변수_안전성_검사(값)

	return &s안전한_배열{s.값.Cons(값)}
}
func (s *s안전한_배열) G슬라이스() []I가변형 {
	반환값 := make([]I가변형, s.G길이())

	if s.G비어있음() {
		return 반환값
	}

	현재_항목 := s.값
	var 인덱스 int

	for !현재_항목.IsNil() {
		인덱스 = 현재_항목.Size() - 1
		반환값[인덱스] = 현재_항목.Head()
		현재_항목 = 현재_항목.Tail()
	}

	return 반환값
}

// 안전한 맵
type s안전한_맵 struct{ 값 ps.Map }

func (s *s안전한_맵) G비어있음() bool     { return s.값.IsNil() }
func (s *s안전한_맵) G길이() int        { return s.값.Size() }
func (s *s안전한_맵) G키_모음() []string { return s.값.Keys() }
func (s *s안전한_맵) G존재함(키 string) bool {
	_, ok := s.값.Lookup(키)

	return ok
}
func (s *s안전한_맵) G값(키 string) I가변형 {
	값, ok := s.값.Lookup(키)

	if !ok {
		return nil
	}

	return 값
}
func (s *s안전한_맵) G맵() map[string]I가변형 {
	맵 := make(map[string]I가변형)

	키_모음 := s.G키_모음()

	for _, 키 := range 키_모음 {
		맵[키] = s.G값(키)
	}

	return 맵
}
func (s *s안전한_맵) S값(키 string, 값 I가변형) I안전한_맵 {
	if !F매개변수_안전성_검사(값) {
		return nil
	}

	return &s안전한_맵{s.값.Set(키, 값)}
}

func (s *s안전한_맵) S삭제(키 string) I안전한_맵 {
	return &s안전한_맵{s.값.Delete(키)}
}

func (s *s안전한_맵) String() string { return s.값.String() }

// NV문자열키_맵 의 구현체 및 그 부속 구조체들
type sC키_값_string_I가변형 struct {
	키 string
	값 I가변형
}

func (s *sC키_값_string_I가변형) G키() string { return s.키 }
func (s *sC키_값_string_I가변형) G값() I가변형   { return s.값 }

// 'V문자열키_맵'의 기본 구현체
type sV기본_문자열키_맵 struct {
	잠금  sync.RWMutex
	저장소 map[string]I가변형
}

func (s *sV기본_문자열키_맵) G변수형임() {}

func (s *sV기본_문자열키_맵) G수량() int {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return len(s.저장소)
}

func (s *sV기본_문자열키_맵) G키_모음() []string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	키_모음 := make([]string, len(s.저장소))
	인덱스 := 0

	for 키, _ := range s.저장소 {
		키_모음[인덱스] = 키
		인덱스++
	}

	return 키_모음
}

func (s *sV기본_문자열키_맵) G값(키 string) (I가변형, bool) {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	값, 존재함 := s.저장소[키]

	return 값, 존재함
}

func (s *sV기본_문자열키_맵) G값_모음() []I가변형 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	값_모음 := make([]I가변형, len(s.저장소))
	인덱스 := 0
	
	for _, 값 := range s.저장소 {
		값_모음[인덱스] = 값
		인덱스++
	}

	return 값_모음
}

func (s *sV기본_문자열키_맵) G키_값_모음() []C키_값_string_I가변형 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	키_값_모음 := make([]C키_값_string_I가변형, len(s.저장소))
	인덱스 := 0
	
	for 키, 값 := range s.저장소 {
		키_값_모음[인덱스] = &sC키_값_string_I가변형{키: 키, 값: 값}
		인덱스++
	}

	return 키_값_모음
}

func (s *sV기본_문자열키_맵) G임의값() I가변형 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	for _, 값 := range s.저장소 {
		return 값
	}
	
	return nil
}

func (s *sV기본_문자열키_맵) S값(키 string, 값 I가변형) {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	s.저장소[키] = 값
}

func (s *sV기본_문자열키_맵) S없으면_추가(키 string, 값 I가변형) {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	// 이미 존재하는 지 확인.
	if _, 이미_존재함 := s.저장소[키]; 이미_존재함 {
		return
	}

	s.저장소[키] = 값
}

func (s *sV기본_문자열키_맵) String() string { return F포맷된_문자열("%v\n", s) }

// 'V문자열키_맵'의 고성능 구현체.
// 뮤텍스의 쓰기 lock 병목 현상을 분산시켜 준다.
// COPIED & MODIFIED from 
// http://openmymind.net/Shard-Your-Hash-table-to-reduce-write-locks/
type sV고성능_문자열키_맵 struct {
	중앙_저장소 map[string]*sV기본_문자열키_맵
}

func (s *sV고성능_문자열키_맵) G변수형임() {}

func (s *sV고성능_문자열키_맵) G수량() int {
	수량 := 0

	동시처리_수량 := int(math.Max(float64(runtime.NumCPU() * 3), 10.0))
	입력_채널 := make(chan *sV기본_문자열키_맵, 동시처리_수량)
	출력_채널 := make(chan int, 동시처리_수량 * 2)
	defer close(입력_채널)

	for 반복횟수 := 0; 반복횟수 < 동시처리_수량; 반복횟수++ {
		go func() {
			for 맵_조각 := range 입력_채널 {
				출력_채널 <- 맵_조각.G수량()
			}
		}()
	}
	
	go func() {
		for _, 맵_조각 := range s.중앙_저장소 {
			입력_채널 <- 맵_조각
		}
	}()

	for 반복횟수 := 0; 반복횟수 < len(s.중앙_저장소); 반복횟수++ {
		조각별_수량 := <-출력_채널

		수량 += 조각별_수량
	}

	return 수량
}

func (s *sV고성능_문자열키_맵) G키_모음() []string {
	키_모음 := make([]string, 0)

	동시처리_수량 := int(math.Max(float64(runtime.NumCPU() * 3), 10.0))
	입력_채널 := make(chan *sV기본_문자열키_맵, 동시처리_수량)
	출력_채널 := make(chan []string, 동시처리_수량 * 2)
	defer close(입력_채널)

	for 반복횟수 := 0; 반복횟수 < 동시처리_수량; 반복횟수++ {
		go func() {
			for 맵_조각 := range 입력_채널 {
				출력_채널 <- 맵_조각.G키_모음()
			}
		}()
	}
	
	go func() {
		for _, 맵_조각 := range s.중앙_저장소 {
			입력_채널 <- 맵_조각
		}
	}()

	for 반복횟수 := 0; 반복횟수 < len(s.중앙_저장소); 반복횟수++ {
		조각별_키_모음 := <-출력_채널

		키_모음 = append(키_모음, 조각별_키_모음...)
	}

	return 키_모음
}

func (s *sV고성능_문자열키_맵) G값(키 string) (I가변형, bool) {
	값, 존재함 := s.g맵_조각(키).G값(키)

	return 값, 존재함
}

func (s *sV고성능_문자열키_맵) G값_모음() []I가변형 {
	값_모음 := make([]I가변형, 0)
	
	동시처리_수량 := int(math.Max(float64(runtime.NumCPU() * 3), 10.0))
	입력_채널 := make(chan *sV기본_문자열키_맵, 동시처리_수량)
	출력_채널 := make(chan []I가변형, 동시처리_수량)
	defer close(입력_채널)

	for 반복횟수 := 0; 반복횟수 < 동시처리_수량; 반복횟수++ {
		go func() {
			for 맵_조각 := range 입력_채널 {
				출력_채널 <- 맵_조각.G값_모음()
			}
		}()
	}
	
	go func() {
		for _, 맵_조각 := range s.중앙_저장소 {
			입력_채널 <- 맵_조각
		}
	}()

	for 반복횟수 := 0; 반복횟수 < len(s.중앙_저장소); 반복횟수++ {
		조각별_값_모음 := <-출력_채널

		값_모음 = append(값_모음, 조각별_값_모음...)
	}

	return 값_모음
}

func (s *sV고성능_문자열키_맵) G키_값_모음() []C키_값_string_I가변형 {
	키_값_모음 := make([]C키_값_string_I가변형, 0)

	동시처리_수량 := int(math.Max(float64(runtime.NumCPU() * 3), 10.0))
	입력_채널 := make(chan *sV기본_문자열키_맵, 동시처리_수량)
	출력_채널 := make(chan []C키_값_string_I가변형, 동시처리_수량)
	defer close(입력_채널)

	for 반복횟수 := 0; 반복횟수 < 동시처리_수량; 반복횟수++ {
		go func() {
			for 맵_조각 := range 입력_채널 {
				출력_채널 <- 맵_조각.G키_값_모음()
			}
		}()
	}
	
	go func() {
		for _, 맵_조각 := range s.중앙_저장소 {
			입력_채널 <- 맵_조각
		}
	}()

	for 반복횟수 := 0; 반복횟수 < len(s.중앙_저장소); 반복횟수++ {
		조각별_키_값_모음 := <-출력_채널

		키_값_모음 = append(키_값_모음, 조각별_키_값_모음...)
	}

	return 키_값_모음
}

func (s *sV고성능_문자열키_맵) G임의값() I가변형 {
	for _, 조각_맵 := range s.중앙_저장소 {
		값 := 조각_맵.G임의값()
		
		if 값 != nil { return 값 }
	}
	
	return nil
}

func (s *sV고성능_문자열키_맵) S값(키 string, 값 I가변형) {
	s.g맵_조각(키).S값(키, 값)
}

func (s *sV고성능_문자열키_맵) S없으면_추가(키 string, 값 I가변형) {
	s.g맵_조각(키).S없으면_추가(키, 값)
}

func (s *sV고성능_문자열키_맵) g맵_조각(키 string) *sV기본_문자열키_맵 {
	// 균일하게 각 조각 맵에 퍼뜨리기 위해서 SHA1 해쉬함수를 이용한다.
	hasher := sha1.New()
	hasher.Write([]byte(키))
	맵_조각_키 := fmt.Sprintf("%x", hasher.Sum(nil))[0:2]

	return s.중앙_저장소[맵_조각_키]
}

func (s *sV고성능_문자열키_맵) String() string { return F포맷된_문자열("%v\n", s) }