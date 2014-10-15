package lib

import (
	"bytes"
	"github.com/gh-system/ghts/dep/ps"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type s반환값 struct {
	값 I가변형
	에러 error
}

func (s *s반환값) G값() I가변형 { return s.값 }
func (s *s반환값) G에러() error { return s.에러 }

type s맵_반환값 struct {
	값 I가변형
	찾았음 bool
}

func (s *s맵_반환값) G값() I가변형 { return s.값 }
func (s *s맵_반환값) G찾았음() bool { return s.찾았음 }


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

func (s *s정수64) G값() int64      { return s.값 }
func (s *s정수64) G정수() int64     { return s.값 }
func (s *s정수64) G실수() float64   { return float64(s.값) }
func (s *s정수64) G정밀수() C정밀수     { return NC정밀수(s.값) }
func (s *s정수64) String() string { return strconv.FormatInt(s.값, 10) }
func (s *s정수64) generate(임의값_생성기 *rand.Rand) int64 {
	값 := 임의값_생성기.Int63() // 0 혹은 양의 정수.

	if 임의값_생성기.Int31n(2) == 0 {
		값 = 값 * -1
	}

	return 값
}

type sC정수64 struct{ *s정수64 }

func (s *sC정수64) 상수형임()     {}
func (s *sC정수64) G변수형() V정수 { return NV정수(s.G값()) }
func (s *sC정수64) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC정수(s.s정수64.generate(임의값_생성기)))
}

type sV정수64 struct {
	잠금 sync.RWMutex
	*s정수64
}

func (s *sV정수64) 변수형임() {}

func (s *sV정수64) G값() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정수64.G값()
}
func (s *sV정수64) G정수() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정수64.G정수()
}
func (s *sV정수64) G실수() float64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정수64.G실수()
}
func (s *sV정수64) G정밀수() C정밀수 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정수64.G정밀수()
}
func (s *sV정수64) G상수형() C정수 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return NC정수(s.G값())
}
func (s *sV정수64) S값(값 int64) V정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.s정수64.값 = 값
	return s
}

// 주어진 값에 대해서 CAS(Compare And Set).
func (s *sV정수64) S_CAS(원래값, 새로운값 int64) bool {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	if 원래값 == s.s정수64.값 {
		s.s정수64.값 = 새로운값
		return true
	}

	// 다른 goroutine에서 값을 이미 변경했음.
	return false
}

func (s *sV정수64) S_CAS_함수(
	함수 func(*sV정수64, ...I가변형) int64,
	매개변수 ...I가변형) (반환값 V정수) {
	// 에러발생 하면 nil 반환
	defer func() {
		if 에러 := recover(); 에러 != nil {
			s.잠금.Lock()
			defer s.잠금.Unlock()
			s.s정수64.값 = 0
			반환값 = nil
		}
	}()

	var 원래값, 새로운값 int64
	var 반복횟수 int = 0

	for {
		원래값 = s.G값()
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.S_CAS(원래값, 새로운값) {
			return s
		}

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.S_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
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

	return s.S_CAS_함수(함수)
}
func (s *sV정수64) S더하기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()

			정수 = 정수 + 값

			return 정수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정수64) S빼기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()

			정수 = 정수 - 값

			return 정수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정수64) S곱하기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()

			정수 = 정수 * 값

			return 정수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정수64) S나누기(값 int64) V정수 {
	함수 :=
		func(구조체 *sV정수64, 매개변수 ...I가변형) int64 {
			정수 := 구조체.G값()

			if 값 == 0 {
				F문자열_출력("0으로 나눌 수 없음.")
			}

			정수 = 정수 / 값

			return 정수
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV정수64) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정수64.String()
}
func (s *sV정수64) Generate(
	임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	return reflect.ValueOf(NV정수(s.s정수64.generate(임의값_생성기)))
}

// 부호없는_정수

type s부호없는_정수64 struct{ 값 uint64 }

func (s *s부호없는_정수64) G값() uint64     { return s.값 }
func (s *s부호없는_정수64) G정수() int64     { return int64(s.값) }
func (s *s부호없는_정수64) G실수() float64   { return float64(s.값) }
func (s *s부호없는_정수64) G정밀수() C정밀수     { return NC정밀수(s.값) }
func (s *s부호없는_정수64) String() string { return strconv.FormatUint(s.값, 10) }

type sC부호없는_정수64 struct{ *s부호없는_정수64 }

func (s *sC부호없는_정수64) 상수형임() {}
func (s *sC부호없는_정수64) G변수형() V부호없는_정수 {
	return NV부호없는_정수(s.s부호없는_정수64.값)
}
func (s *sC부호없는_정수64) Generate(
	임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	return reflect.ValueOf(NC부호없는_정수(uint64(임의값_생성기.Int63())))
}

type sV부호없는_정수64 struct {
	잠금 sync.RWMutex
	*s부호없는_정수64
}

func (s *sV부호없는_정수64) 변수형임() {}
func (s *sV부호없는_정수64) G값() uint64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s부호없는_정수64.G값()
}
func (s *sV부호없는_정수64) G정수() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s부호없는_정수64.G정수()
}
func (s *sV부호없는_정수64) G실수() float64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s부호없는_정수64.G실수()
}
func (s *sV부호없는_정수64) G정밀수() C정밀수 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s부호없는_정수64.G정밀수()
}
func (s *sV부호없는_정수64) G상수형() C부호없는_정수 {
	s.잠금.RLock()
	값 := s.s부호없는_정수64.값
	s.잠금.RUnlock()
	return NC부호없는_정수(값)
}
func (s *sV부호없는_정수64) S값(값 uint64) V부호없는_정수 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.s부호없는_정수64.값 = 값
	return s
}

// 주어진 값에 대해서 CAS(Compare And Set).
func (s *sV부호없는_정수64) S_CAS(원래값, 새로운값 uint64) bool {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	if 원래값 == s.s부호없는_정수64.값 {
		s.s부호없는_정수64.값 = 새로운값
		return true
	}

	// 다른 goroutine에서 값을 이미 변경했음.
	return false
}

func (s *sV부호없는_정수64) S_CAS_함수(
	함수 func(*sV부호없는_정수64, ...I가변형) uint64,
	매개변수 ...I가변형) (반환값 V부호없는_정수) {
	// 에러발생 하면 nil 반환
	defer func() {
		if 에러 := recover(); 에러 != nil {
			s.잠금.Lock()
			defer s.잠금.Unlock()
			s.s부호없는_정수64.값 = 0
			반환값 = nil
		}
	}()

	var 원래값, 새로운값 uint64
	var 반복횟수 int = 0

	for {
		원래값 = s.G값()
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.S_CAS(원래값, 새로운값) {
			return s
		}

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.S_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
		반복횟수++
		F잠시_대기(반복횟수)
	}
}
func (s *sV부호없는_정수64) S더하기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()

			부호없는_정수 = 부호없는_정수 + 값

			return 부호없는_정수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV부호없는_정수64) S빼기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()

			부호없는_정수 = 부호없는_정수 - 값

			return 부호없는_정수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV부호없는_정수64) S곱하기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()

			부호없는_정수 = 부호없는_정수 * 값

			return 부호없는_정수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV부호없는_정수64) S나누기(값 uint64) V부호없는_정수 {
	함수 :=
		func(구조체 *sV부호없는_정수64, 매개변수 ...I가변형) uint64 {
			부호없는_정수 := 구조체.G값()

			if 값 == 0 {
				F문자열_출력("0으로 나눌 수 없음.")
			}

			부호없는_정수 = 부호없는_정수 / 값

			return 부호없는_정수
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV부호없는_정수64) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s부호없는_정수64.String()
}
func (s *sV부호없는_정수64) Generate(
	임의값_생성기 *rand.Rand,
	크기 int) reflect.Value {
	return reflect.ValueOf(NV부호없는_정수(uint64(임의값_생성기.Int63())))
}

// 실수

type s실수64 struct{ 값 float64 }

func (s *s실수64) G값() float64  { return s.값 }
func (s *s실수64) G실수() float64 { return s.값 }
func (s *s실수64) G정밀수() C정밀수   { return NC정밀수(s.값) }
func (s *s실수64) String() string {
	return strconv.FormatFloat(s.값, 'f', -1, 64)
}
func (s *s실수64) generate(
	임의값_생성기 *rand.Rand) float64 {
	값 := 임의값_생성기.NormFloat64()

	if 임의값_생성기.Int31n(2) == 0 {
		값 = 값 * -1.0
	}

	return 값
}

type sC실수64 struct{ *s실수64 }

func (s *sC실수64) 상수형임()     {}
func (s *sC실수64) G변수형() V실수 { return NV실수(s.s실수64.값) }
func (s *sC실수64) Generate(
	임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC실수(s.s실수64.generate(임의값_생성기)))
}

type sV실수64 struct {
	잠금 sync.RWMutex
	*s실수64
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
	s.s실수64.값 = 값
	return s
}

// 주어진 값에 대해서 CAS(Compare And Set).
func (s *sV실수64) S_CAS(원래값, 새로운값 float64) bool {
	s.잠금.Lock()
	defer s.잠금.Unlock()

	if 원래값 == s.s실수64.값 {
		s.s실수64.값 = 새로운값
		return true
	}

	// 다른 goroutine에서 값을 이미 변경했음.
	return false
}

func (s *sV실수64) S_CAS_함수(
	함수 func(*sV실수64, ...I가변형) float64,
	매개변수 ...I가변형) (반환값 V실수) {
	// 에러발생 하면 nil 반환
	defer func() {
		if 에러 := recover(); 에러 != nil {
			s.잠금.Lock()
			defer s.잠금.Unlock()
			s.s실수64.값 = 0
			반환값 = nil
		}
	}()

	var 원래값, 새로운값 float64
	var 반복횟수 int = 0

	for {
		원래값 = s.G값()
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.S_CAS(원래값, 새로운값) {
			return s
		}

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.S_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
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

	return s.S_CAS_함수(함수)
}
func (s *sV실수64) S더하기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 + 값

			return 실수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV실수64) S빼기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 - 값

			return 실수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV실수64) S곱하기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			실수 = 실수 * 값

			return 실수
		}

	return s.S_CAS_함수(함수)
}
func (s *sV실수64) S나누기(값 float64) V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			if 값 == 0 {
				F문자열_출력("0으로 나눌 수 없음.")
			}

			실수 = 실수 / 값

			return 실수
		}

	return s.S_CAS_함수(함수, 값)
}
func (s *sV실수64) S역수() V실수 {
	함수 :=
		func(구조체 *sV실수64, 매개변수 ...I가변형) float64 {
			실수 := 구조체.G값()

			if 실수 == 0 {
				F문자열_출력("0의 역수를 구할 수 없음.")
			}

			실수 = 1 / 실수

			return 실수
		}

	return s.S_CAS_함수(함수)
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

func (s *sC참거짓) 상수형임()      {}
func (s *sC참거짓) G변수형() V참거짓 { return NV참거짓(s.s참거짓.값) }
func (s *sC참거짓) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC참거짓(s.s참거짓.generate(임의값_생성기)))
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
	return reflect.ValueOf(NV참거짓(s.s참거짓.generate(임의값_생성기)))
}

type sC문자열 struct{ 값 string }

func (s *sC문자열) 상수형임()          {}
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

func (s *sC시점) 상수형임()     {}
func (s *sC시점) G변수형() V시점 { return NV시점(s.s시점.값) }
func (s *sC시점) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(NC시점(s.s시점.generate(임의값_생성기)))
}

type sV시점 struct {
	잠금 sync.RWMutex
	*s시점
}

func (s *sV시점) 변수형임() {}
func (s *sV시점) G값() time.Time {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s시점.G값()
}
func (s *sV시점) G상수형() C시점 {
	return NC시점(s.G값())
}
func (s *sV시점) S값(값 time.Time) V시점 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.s시점.값 = 값 // time.Time은 매개변수로 넘어오면서 자동으로 복사됨.
	return s
}
func (s *sV시점) S일자_더하기(연, 월, 일 int) V시점 {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.s시점.값 = s.s시점.값.AddDate(연, 월, 일)
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
	실수, 에러 := F문자열2실수(s.String())

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

	차이_절대값 := new(big.Rat).Abs(new(big.Rat).Sub(s.GRat(), 정밀수.GRat()))
	차이_한도, _ := new(big.Rat).SetString("1/1000000000000000000000000000000000")

	if 차이_절대값.Cmp(차이_한도) == -1 {
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
	if F_nil값_존재함(s, s.GRat()) {
		return "<nil"
	}

	return F마지막_0_제거(s.GRat().FloatString(100))
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
func (s *sV정밀수) G같음(값 I가변형) bool {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.G같음(값)
}
func (s *sV정밀수) G비교(값 I가변형) int {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	return s.s정밀수.G비교(값)
}
func (s *sV정밀수) S값(값 I가변형) V정밀수 {
	if F_nil값임(값) ||
		(!F숫자형식임(값) && !F문자열형식임(값)) {
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

// 주어진 값에 대해서 CAS(Compare And Set).
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
	함수 func(*sV정밀수, ...I가변형) *big.Rat,
	매개변수 ...I가변형) (반환값 V정밀수) {
	// 에러발생 하면 nil 반환
	defer func() {
		if 에러 := recover(); 에러 != nil {
			s.잠금.Lock()
			defer s.잠금.Unlock()
			s.s정밀수.값 = nil
			반환값 = nil
		}
	}()

	var 원래값, 새로운값 *big.Rat
	var 반복횟수 int = 0

	for {
		원래값 = s.GRat()
		새로운값 = 함수(s, 매개변수...)

		// 내부값이 변하지 않았을 경우에만 새로운 값으로 설정.
		if s.S_CAS(원래값, 새로운값) {
			return s
		}

		// 함수()가 실행되는 동안 다른 goroutine이 내부값을 변경한 경우에는
		// s.S_CAS()가 실패하며, 그럴 경우 처음부터 새로 계산함.
		반복횟수++
		F잠시_대기(반복횟수)
	}
}

func (s *sV정밀수) S반올림(소숫점_이하_자릿수 int) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
			정밀수 := 구조체.GRat()

			if 정밀수 == nil {
				return nil
			}

			return 정밀수.Abs(정밀수)
		}

	return s.S_CAS_함수(함수)
}
func (s *sV정밀수) S더하기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
func (s *sV정밀수) S빼기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
func (s *sV정밀수) S곱하기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
func (s *sV정밀수) S나누기(값 I가변형) V정밀수 {
	함수 :=
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
		func(구조체 *sV정밀수, 매개변수 ...I가변형) *big.Rat {
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
func (s *sC통화) G비교(값 I통화) int {
	if F_nil값임(s) && F_nil값임(값) {
		if s.G종류() == 값.G종류() {
			return 0
		} else {
			return -2
		}
	}
	if F_nil값임(s) || F_nil값임(값) {
		return -2
	}
	if s.G종류() != 값.G종류() {
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

	return s.금액.G상수형()
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
func (s *sV통화) G비교(값 I통화) int {
	if F_nil값임(s) && F_nil값임(값) {
		if s.G종류() == 값.G종류() {
			return 0
		} else {
			return -2
		}
	}
	if F_nil값임(s) || F_nil값임(값) {
		return -2
	}
	if s.G종류() != 값.G종류() {
		return -2
	}

	return s.G값().G비교(값.G값())
}
func (s *sV통화) G상수형() C통화 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()

	return NC통화(s.종류, s.금액.G상수형())
}
func (s *sV통화) S종류(종류 P통화종류) {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	s.종류 = 종류
}
func (s *sV통화) S값(금액 I가변형) V통화 {
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
func (s *sV통화) S더하기(값 I가변형) V통화 {
	F매개변수_안전성_검사(값)

	if F_nil값_존재함(값) ||
		(!F통화형식임(값) && !F숫자형식임(값)) ||
		F통화종류(s.G상수형(), 값) == INVALID_CURRENCY_TYPE {
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
func (s *sV통화) S빼기(값 I가변형) V통화 {
	F매개변수_안전성_검사(값)

	if F_nil값_존재함(값) ||
		(!F통화형식임(값) && !F숫자형식임(값)) ||
		F통화종류(s.G상수형(), 값) == INVALID_CURRENCY_TYPE {
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
func (s *sV통화) S곱하기(값 I가변형) V통화 {
	F매개변수_안전성_검사(값)

	if F_nil값_존재함(값) ||
		(!F통화형식임(값) && !F숫자형식임(값)) ||
		F통화종류(s.G상수형(), 값) == INVALID_CURRENCY_TYPE {
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
func (s *sV통화) S나누기(값 I가변형) V통화 {
	F매개변수_안전성_검사(값)

	if F_nil값_존재함(값) ||
		(!F통화형식임(값) && !F숫자형식임(값)) ||
		F통화종류(s.G상수형(), 값) == INVALID_CURRENCY_TYPE {
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

func (s *sC매개변수) 상수형임() {}
func (s *sC매개변수) G이름() string     { return s.이름 }
func (s *sC매개변수) G값() I가변형 { return s.값 }
func (s *sC매개변수) G숫자형식임() bool    { return F숫자형식임(s.값) }
func (s *sC매개변수) G문자열형식임() bool   { return F문자열형식임(s.값) }
func (s *sC매개변수) G시점형식임() bool    { return F시점형식임(s.값) }
func (s *sC매개변수) G참거짓형식임() bool   { return F참거짓형식임(s.값) }
func (s *sC매개변수) String() string {
	return s.이름 + " " + reflect.TypeOf(s.값).String() + F문자열(s.값)
}
func (s *sC매개변수) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	임의_숫자값 := 임의값_생성기.Int31()
	임의_문자열 := NC문자열("").Generate(임의값_생성기, 크기).Interface().(string)
	
	값_모음 := []I가변형{ uint(임의_숫자값), int(임의_숫자값), float32(임의_숫자값),
					float64(임의_숫자값), time.Now(), true, false, 임의_문자열,
					NC정수(int64(임의_숫자값)), NC부호없는_정수(uint64(임의_숫자값)), 
					NC실수(float64(임의_숫자값)), NC정밀수(임의_숫자값), 
					NC통화(F임의_통화종류(), 임의_숫자값), NC시점(time.Now()), 
					NC참거짓(true), NC문자열(임의_문자열)}
	
	이름 := 임의_문자열
	값 := 값_모음[임의값_생성기.Int31n(int32(len(값_모음)))]

	return reflect.ValueOf(NC매개변수(이름, 값))
}

// 안전한 배열
type s안전한_배열 struct { 값 ps.List }

func (s *s안전한_배열) G비어있음() bool { return s.값.IsNil() }
func (s *s안전한_배열) G길이() int { return s.값.Size() }
func (s *s안전한_배열) S추가(값 I가변형) I안전한_배열 {
	F매개변수_안전성_검사(값)
	
	return &s안전한_배열{s.값.Cons(값)}
}
func (s *s안전한_배열) G슬라이스() []I가변형 {
	반환값 := make([]I가변형, s.G길이())

	if s.G비어있음() { return 반환값 }
	
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
type s안전한_맵 struct { 값 ps.Map }

func (s *s안전한_맵) G비어있음() bool { return s.값.IsNil() }
func (s *s안전한_맵) G길이() int { return s.값.Size() }
func (s *s안전한_맵) G키_모음() []string { return s.값.Keys() }
func (s *s안전한_맵) G값(키 string) I맵_반환값 { return N맵_반환값(s.값.Lookup(키)) }
func (s *s안전한_맵) S값(키 string, 값 I가변형) I안전한_맵 {
	if !F매개변수_안전성_검사(값) { return nil }
	
	return &s안전한_맵{s.값.Set(키, 값)}
}

func (s *s안전한_맵) S삭제(키 string) I안전한_맵 {
	return &s안전한_맵{s.값.Delete(키)}
}

func (s *s안전한_맵) String() string { return s.값.String() }