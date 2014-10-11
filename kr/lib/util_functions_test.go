package lib

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestF잠시_대기(테스트 *testing.T) {
	F잠시_대기(1)
	F잠시_대기(100)
}

func TestF안전한_매개변수(테스트 *testing.T) {
	// CallByValue에 의해서 자동으로 복사본이 생성되는 형식.
	검사_결과 := F안전한_매개변수(
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1), float64(1), true, false, "test",
		time.Now())
	F참인지_확인(테스트, 검사_결과)

	// Immutable 하므로 race condition이 발생하지 않는 형식.
	// 앞으로 여기에 검증된 상수형을 더 추가해야 됨.
	검사_결과 = F안전한_매개변수(
		NC부호없는_정수(1), NC정수(1), NC실수(1),
		NC시점(time.Now()), NC정밀수(1), NC통화(KRW, 1))
	F참인지_확인(테스트, 검사_결과)

	// Mutable 한 타입들.
	// 비록 RWMutex로 보호되어 있더라도, 매개변수로 좋지 않음.
	검사_결과 = F안전한_매개변수(
		NV부호없는_정수(1), NV정수(1), NV실수(1),
		NV시점(time.Now()), NV정밀수(1), NV통화(KRW, 1))
	F거짓인지_확인(테스트, 검사_결과)
}

func TestF상수형(테스트 *testing.T) {
	F참인지_확인(테스트, F상수형(nil) == nil)

	F문자열_출력_일시정지_시작()
	F참인지_확인(테스트, F상수형(time.January) == nil)
	F문자열_출력_일시정지_종료()

	값_모음 := []interface{}{
		time.Now(), NC시점(time.Now()), NC시점_문자열("2000-01-01"),
		NV시점(time.Now()), NV시점_문자열("2000-01-01")}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C시점)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{"테스트", NC문자열("테스트")}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C문자열)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{true, false, NC참거짓(true), NC참거짓(false)}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C참거짓)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		NC부호없는_정수(uint64(1)), NV부호없는_정수(uint64(1))}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C부호없는_정수)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{
		int(1), int8(1), int16(1), int32(1), int64(1),
		NC정수(int64(1)), NV정수(int64(1))}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C정수)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{
		float32(1), float64(1), NC실수(1), NV실수(1)}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C실수)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{
		big.NewInt(1), big.NewRat(1, 1), NC정밀수(1), NV정밀수(1)}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C정밀수)
		F참인지_확인(테스트, ok)
	}

	값_모음 = []interface{}{
		NC통화(KRW, 100), NV통화(KRW, 100)}
	for _, 값 := range 값_모음 {
		_, ok := F상수형(값).(C통화)
		F참인지_확인(테스트, ok)
	}
}

func TestF문자열(테스트 *testing.T) {
	값_모음 := []interface{}{true, P참, NC참거짓(true)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F문자열(값), "true")
	}

	값_모음 = []interface{}{"테스트", NC문자열("테스트")}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F문자열(값), "테스트")
	}

	시점, _ := F문자열2시점("2000-01-01")

	값_모음 = []interface{}{
		시점, NC시점(시점), NC시점_문자열("2000-01-01"),
		NV시점(시점), NV시점_문자열("2000-01-01")}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F문자열(값), "2000-01-01 00:00:00 (UTC) Sat +0000")
	}

	값_모음 = []interface{}{
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		NC부호없는_정수(uint64(1)), NV부호없는_정수(uint64(1)),
		int(1), int8(1), int16(1), int32(1), int64(1),
		NC정수(int64(1)), NV정수(int64(1)), big.NewInt(1)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F문자열(값), "1")
	}

	값_모음 = []interface{}{
		float32(1.1), float64(1.1), NC실수(1.1), NV실수(1.1),
		big.NewRat(11, 10), NC정밀수(1.1), NV정밀수(1.1)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F문자열(값), "1.1")
	}

	값_모음 = []interface{}{
		NC통화(KRW, 100), NV통화(KRW, 100)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F문자열(값), "KRW 100")
	}
}

func TestF포맷된_문자열(테스트 *testing.T) {
	값_모음 := []interface{}{true, P참, NC참거짓(true)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F포맷된_문자열("%v", 값), "true")
	}

	값_모음 = []interface{}{"테스트", NC문자열("테스트")}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F포맷된_문자열("%v", 값), "테스트")
	}

	시점, _ := F문자열2시점("2000-01-01")

	값_모음 = []interface{}{
		NC시점(시점), NC시점_문자열("2000-01-01"),
		NV시점(시점), NV시점_문자열("2000-01-01")}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F포맷된_문자열("%v", 값), "2000-01-01 00:00:00 (UTC) Sat +0000")
	}

	값_모음 = []interface{}{
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		NC부호없는_정수(uint64(1)), NV부호없는_정수(uint64(1)),
		int(1), int8(1), int16(1), int32(1), int64(1),
		NC정수(int64(1)), NV정수(int64(1)), big.NewInt(1)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F포맷된_문자열("%v", 값), "1")
	}

	값_모음 = []interface{}{
		float32(1.1), float64(1.1), NC실수(1.1), NV실수(1.1),
		NC정밀수(1.1), NV정밀수(1.1)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F포맷된_문자열("%v", 값), "1.1")
	}

	값_모음 = []interface{}{
		NC통화(KRW, 100), NV통화(KRW, 100)}
	for _, 값 := range 값_모음 {
		F같은값_확인(테스트, F포맷된_문자열("%v", 값), "KRW 100")
	}
}

func TestF금액_문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F금액_문자열("100"), "100")
	F같은값_확인(테스트, F금액_문자열("100.1"), "100.1")
	F같은값_확인(테스트, F금액_문자열("1000"), "1,000")
	F같은값_확인(테스트, F금액_문자열("1000.1"), "1,000.1")
	F같은값_확인(테스트, F금액_문자열("1000000"), "1,000,000")
	F같은값_확인(테스트, F금액_문자열("1000000.1"), "1,000,000.1")
	F같은값_확인(테스트, F금액_문자열("1000000000000"), "1,000,000,000,000")
	F같은값_확인(테스트, F금액_문자열("1000000000000.1"), "1,000,000,000,000.1")
}

func TestF마지막_0_제거(테스트 *testing.T) {
	F같은값_확인(테스트, F마지막_0_제거("100"), "100")
	F같은값_확인(테스트, F마지막_0_제거("100.0"), "100")
	F같은값_확인(테스트, F마지막_0_제거("100.1"), "100.1")
	F같은값_확인(테스트, F마지막_0_제거("100.10"), "100.1")
	F같은값_확인(테스트, F마지막_0_제거("100.1230450000"), "100.123045")
}

func TestF반올림(테스트 *testing.T) {
	F같은값_확인(테스트, F반올림(100.0045, 2).G실수(), 100.0)
	F같은값_확인(테스트, F반올림(100.0045, 3).G실수(), 100.005)

	F문자열_출력_일시정지_시작()
	F같은값_확인(테스트, F반올림("숫자 아닌 형식", 3), nil)
	F문자열_출력_일시정지_종료()
}

func TestF문자열2실수(테스트 *testing.T) {
	실수, 에러 := F문자열2실수("1.1")

	F에러없음_확인(테스트, 에러)
	F같은값_확인(테스트, 실수, 1.1)

	실수, 에러 = F문자열2실수("변환 불가능한 문자열")

	F에러발생_확인(테스트, 에러)
	F같은값_확인(테스트, 실수, 0)
}

func TestF문자열2시점(테스트 *testing.T) {
	시점_원래값 := time.Now()

	시점, 에러 := F문자열2시점(시점_원래값.Format(P시점_포맷))
	F에러없음_확인(테스트, 에러)
	F같은값_확인(테스트, 시점.Format(P시점_포맷), 시점_원래값.Format(P시점_포맷))

	일자, 에러 := F문자열2시점(시점_원래값.Format(P일자_포맷))
	F에러없음_확인(테스트, 에러)
	F같은값_확인(테스트, 일자.Format(P일자_포맷), 시점_원래값.Format(P일자_포맷))
}

func TestF시점_문자열(테스트 *testing.T) {
	시점_원래값 := time.Now()

	F같은값_확인(테스트, F시점_문자열(시점_원래값), 시점_원래값.Format(P시점_포맷))
}

func TestF일자_문자열(테스트 *testing.T) {
	일자 := time.Date(2000, time.Month(1), 1, 0, 0, 0, 0, time.Now().Location())

	F같은값_확인(테스트, F일자_문자열(일자), "2000-01-01")

	시점_원래값 := time.Now()

	F같은값_확인(테스트, F일자_문자열(시점_원래값), 시점_원래값.Format(P일자_포맷))
}

func TestF시점_복사(테스트 *testing.T) {
	시점 := time.Now()
	복사본 := F시점_복사(시점)

	F같은값_확인(테스트, 시점, 복사본)

	시점 = 시점.AddDate(0, 0, 1)

	F다른값_확인(테스트, 시점, 복사본)
}

func TestF임의_통화종류(테스트 *testing.T) {
	통화종류_맵 := make(map[P통화종류]int)

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		통화종류 := F임의_통화종류()
		발생횟수, OK := 통화종류_맵[통화종류]

		if OK {
			발생횟수++
			통화종류_맵[통화종류] = 발생횟수
		} else {
			통화종류_맵[통화종류] = 1
		}
	}

	F같은값_확인(테스트, len(통화종류_맵), len(통화종류_문자열_모음)) // 현재까지 설정한 통화종류

	최소_발생횟수 := int(30.0 / len(통화종류_문자열_모음))

	for 통화종류, 발생횟수 := range 통화종류_맵 {
		if 발생횟수 < 최소_발생횟수 {
			F문자열_출력("발생횟수가 너무 적음 : 통화종류 %v, 발생횟수 %v.",
				통화종류.String(), 발생횟수)
		}
	}
}

func TestF통화종류별_정밀도(테스트 *testing.T) {
	F같은값_확인(테스트, F통화종류별_정밀도(KRW), 0)
	F같은값_확인(테스트, F통화종류별_정밀도(USD), 2)
	F같은값_확인(테스트, F통화종류별_정밀도(CNY), 2)
	F같은값_확인(테스트, F통화종류별_정밀도(EUR), 2)
	F같은값_확인(테스트, F통화종류별_정밀도(P통화종류(-1)), 2)
}

func TestF통화종류(테스트 *testing.T) {
	F같은값_확인(테스트, F통화종류(nil, nil), INVALID_CURRENCY_TYPE)
	F같은값_확인(테스트, F통화종류(NC통화(KRW, 100), nil), KRW)
	F같은값_확인(테스트, F통화종류(nil, NC통화(KRW, 100)), KRW)
	F같은값_확인(테스트, F통화종류(NC통화(KRW, 100), NC통화(KRW, 100)), KRW)
	F같은값_확인(테스트, F통화종류(NC통화(KRW, 1), NC통화(USD, 1)), INVALID_CURRENCY_TYPE)
}

func TestF통화형식임(테스트 *testing.T) {
	값_모음 := []interface{}{NC통화(KRW, 100), NV통화(KRW, 100)}

	F참인지_확인(테스트, F통화형식임(값_모음...))

	for _, 값 := range 값_모음 {
		F참인지_확인(테스트, F통화형식임(값))
	}

	시점_문자열 := "2000-01-01"
	값_모음 = []interface{}{nil, true, P참, NC참거짓(true), "테스트", NC문자열("테스트"),
		NC시점_문자열(시점_문자열).G값(),
		NC시점_문자열(시점_문자열), NV시점_문자열(시점_문자열),
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		NC부호없는_정수(uint64(1)), NV부호없는_정수(uint64(1)),
		int(1), int8(1), int16(1), int32(1), int64(1),
		NC정수(int64(1)), NV정수(int64(1)),
		float32(1.1), float64(1.1), NC실수(1.1), NV실수(1.1),
		big.NewInt(1), big.NewRat(11, 10), NC정밀수(1.1), NV정밀수(1.1)}

	F거짓인지_확인(테스트, F통화형식임(값_모음...))

	for _, 값 := range 값_모음 {
		F거짓인지_확인(테스트, F통화형식임(값))
	}
}

func TestF숫자형식임(테스트 *testing.T) {
	값_모음 := []interface{}{
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1.1), float64(1.1), big.NewInt(1), big.NewRat(11, 10),
		NC부호없는_정수(1), NV부호없는_정수(1), NC정수(1), NV정수(1),
		NC정밀수(1.1), NV정밀수(1.1)}
	F참인지_확인(테스트, F숫자형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F참인지_확인(테스트, F숫자형식임(값))
	}

	값_모음 = []interface{}{
		nil, "문자열", NC문자열("문자열"), time.Now(), NC시점(time.Now()),
		true, false, NC참거짓(true), NC통화(KRW, 100), NV통화(KRW, 100)}
	F거짓인지_확인(테스트, F숫자형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F거짓인지_확인(테스트, F숫자형식임(값))
	}
}

func TestF숫자_같음(테스트 *testing.T) {
	값_모음_1 := []interface{}{
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1), float64(1), big.NewInt(1), big.NewRat(1, 1),
		NC부호없는_정수(1), NV부호없는_정수(1), NC정수(1), NV정수(1),
		NC정밀수(1), NV정밀수(1)}

	값_모음_2 := []interface{}{
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1), float64(1), big.NewInt(1), big.NewRat(1, 1),
		NC부호없는_정수(1), NV부호없는_정수(1), NC정수(1), NV정수(1),
		NC정밀수(1), NV정밀수(1)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F참인지_확인(테스트, F숫자_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	값_모음_2 = []interface{}{
		nil, uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		int(2), int8(2), int16(2), int32(2), int64(2),
		float32(2), float64(2), big.NewInt(2), big.NewRat(2, 1),
		NC부호없는_정수(2), NV부호없는_정수(2), NC정수(2), NV정수(2),
		NC정밀수(2), NV정밀수(2)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F거짓인지_확인(테스트, F숫자_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	값_모음_1 = []interface{}{
		float32(1.1), float64(1.1), big.NewRat(11, 10),
		NC정밀수(1.1), NV정밀수(1.1)}

	값_모음_2 = []interface{}{
		float32(1.1), float64(1.1), big.NewRat(11, 10),
		NC정밀수(1.1), NV정밀수(1.1)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F참인지_확인(테스트, F숫자_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	값_모음_2 = []interface{}{
		nil, float32(2.2), float64(2.2),
		big.NewRat(22, 10), NC정밀수(2.2), NV정밀수(2.2)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F거짓인지_확인(테스트, F숫자_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}
}

func TestF참거짓형식임(테스트 *testing.T) {
	값_모음 := []interface{}{true, false, NC참거짓(true)}
	F참인지_확인(테스트, F참거짓형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F참인지_확인(테스트, F참거짓형식임(값))
	}

	값_모음 = []interface{}{
		nil, uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1.1), float64(1.1), big.NewInt(1), big.NewRat(11, 10),
		NC부호없는_정수(1), NV부호없는_정수(1), NC정수(1), NV정수(1),
		NC정밀수(1.1), NV정밀수(1.1), "문자열", NC문자열("문자열"),
		time.Now(), NC시점(time.Now()), NC통화(KRW, 100), NV통화(KRW, 100)}
	F거짓인지_확인(테스트, F참거짓형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F거짓인지_확인(테스트, F참거짓형식임(값))
	}
}

func TestF참거짓_같음(테스트 *testing.T) {
	값_모음_1 := []interface{}{true, NC참거짓(true)}
	값_모음_2 := []interface{}{true, NC참거짓(true)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F참인지_확인(테스트, F참거짓_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	값_모음_1 = []interface{}{false, NC참거짓(false)}
	값_모음_2 = []interface{}{false, NC참거짓(false)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F참인지_확인(테스트, F참거짓_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	값_모음_1 = []interface{}{true, NC참거짓(true)}
	값_모음_2 = []interface{}{false, NC참거짓(false)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F거짓인지_확인(테스트, F참거짓_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	값_모음_1 = []interface{}{false, NC참거짓(false)}
	값_모음_2 = []interface{}{true, NC참거짓(true)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F거짓인지_확인(테스트, F참거짓_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}
}

func TestF문자열형식임(테스트 *testing.T) {
	값_모음 := []interface{}{"테스트", NC문자열("테스트")}
	F참인지_확인(테스트, F문자열형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F참인지_확인(테스트, F문자열형식임(값))
	}

	값_모음 = []interface{}{
		nil, uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1.1), float64(1.1), big.NewInt(1), big.NewRat(11, 10),
		NC부호없는_정수(1), NV부호없는_정수(1), NC정수(1), NV정수(1),
		NC정밀수(1.1), NV정밀수(1.1), time.Now(), NC시점(time.Now()),
		true, NC참거짓(true), NC통화(KRW, 100), NV통화(KRW, 100)}
	F거짓인지_확인(테스트, F문자열형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F거짓인지_확인(테스트, F문자열형식임(값))
	}
}

func TestF문자열_같음(테스트 *testing.T) {
	문자열1 := "테스트"
	문자열2 := "테스트"

	값_모음_1 := []interface{}{문자열1, NC문자열(문자열1)}
	값_모음_2 := []interface{}{문자열2, NC문자열(문자열2)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F참인지_확인(테스트, F문자열_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	문자열2 = 문자열2 + "2"
	값_모음_2 = []interface{}{문자열2, NC문자열(문자열2)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F거짓인지_확인(테스트, F문자열_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}
}

func TestF시점형식임(테스트 *testing.T) {
	시점 := time.Now()
	값_모음 := []interface{}{시점, NC시점(시점), NV시점(시점)}
	F참인지_확인(테스트, F시점형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F참인지_확인(테스트, F시점형식임(값))
	}

	값_모음 = []interface{}{
		nil, uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1.1), float64(1.1), big.NewInt(1), big.NewRat(11, 10),
		NC부호없는_정수(1), NV부호없는_정수(1), NC정수(1), NV정수(1),
		NC정밀수(1.1), NV정밀수(1.1), "문자열", NC문자열("문자열"),
		true, NC참거짓(true), NC통화(KRW, 100), NV통화(KRW, 100)}
	F거짓인지_확인(테스트, F시점형식임(값_모음...))
	for _, 값 := range 값_모음 {
		F거짓인지_확인(테스트, F시점형식임(값))
	}
}

func TestF시점_같음(테스트 *testing.T) {
	시점1 := time.Now()
	시점2 := F시점_복사(시점1)

	값_모음_1 := []interface{}{시점1, NC시점(시점1), NV시점(시점1)}
	값_모음_2 := []interface{}{시점2, NC시점(시점2), NV시점(시점2)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F참인지_확인(테스트, F시점_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}

	시점2 = 시점2.AddDate(0, 0, 1)
	값_모음_2 = []interface{}{시점2, NC시점(시점2), NV시점(시점2)}

	for _, 값1 := range 값_모음_1 {
		for _, 값2 := range 값_모음_2 {
			F거짓인지_확인(테스트, F시점_같음(값1, 값2), "값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}
	}
}

func TestF값_같음(테스트 *testing.T) {
	// 정수 테스트
	값 := []interface{}{
		uint(100), uint8(100), uint16(100), uint32(100), uint64(100),
		int(100), int8(100), int16(100), int32(100), int64(100),
		float32(100.0), float64(100.0),
		NC정밀수(100), NV정밀수(100)}

	testF값_같음_도우미(테스트, 값)

	// 실수 테스트
	값 = []interface{}{
		float32(100.0345), float64(100.0345),
		NC정밀수(100.0345), NV정밀수(100.0345)}

	testF값_같음_도우미(테스트, 값)

	// 정밀한 실수 테스트
	값 = []interface{}{
		//float32(100.000000345), // float32에서는 에러남.
		float64(100.000000345),
		NC정밀수(100.000000345), NV정밀수(100.000000345)}

	testF값_같음_도우미(테스트, 값)

	// 통화 테스트
	통화종류 := F임의_통화종류()

	값 = []interface{}{
		NC통화(통화종류, 111.1111),
		NC통화(통화종류, big.NewRat(1111111, 10000)),
		NV통화(통화종류, 111.1111),
		NV통화(통화종류, big.NewRat(1111111, 10000))}

	testF값_같음_도우미(테스트, 값)

	// 문자열 테스트
	값 = []interface{}{"테스트", NC문자열("테스트")}
	testF값_같음_도우미(테스트, 값)

	// 참거짓 테스트
	값 = []interface{}{true, P참, NC참거짓(true)}
	testF값_같음_도우미(테스트, 값)

	값 = []interface{}{false, P거짓, NC참거짓(false)}
	testF값_같음_도우미(테스트, 값)

	// 시점 테스트
	시점 := time.Now()
	값 = []interface{}{시점, NC시점(시점), NV시점(시점)}
	testF값_같음_도우미(테스트, 값)
}

func testF값_같음_도우미(테스트 *testing.T, 값 []interface{}) {
	for 인덱스1 := 0; 인덱스1 < (len(값) - 1); 인덱스1++ {
		for 인덱스2 := 인덱스1 + 1; 인덱스2 < len(값); 인덱스2++ {
			F참인지_확인(테스트, F값_같음(값[인덱스1], 값[인덱스2]),
				"common.TestF값_값음() : 값1 %v %v, 값2 %v %v",
				reflect.TypeOf(값[인덱스1]), 값[인덱스1],
				reflect.TypeOf(값[인덱스2]), 값[인덱스2])
		}
	}
}

func TestF슬라이스_복사(테스트 *testing.T) {
	원본_슬라이스 := []string{"1번째", "2번째", "3번째"}
	복사본_슬라이스 := F슬라이스_복사(원본_슬라이스).([]string)

	F같은값_확인(테스트, len(원본_슬라이스), len(복사본_슬라이스))
	F같은값_확인(테스트, 원본_슬라이스[0], 복사본_슬라이스[0])
	F같은값_확인(테스트, 원본_슬라이스[1], 복사본_슬라이스[1])
	F같은값_확인(테스트, 원본_슬라이스[2], 복사본_슬라이스[2])

	// 원본과 복사본의 독립성 확인.
	복사본_슬라이스[0] = "변경된 1번째"
	F다른값_확인(테스트, 원본_슬라이스[0], 복사본_슬라이스[0])
}

// 테스트 편의함수 Fxxx_확인() 테스트용 Mock-Up
// testing.TB 인터페이스를 구현함.
var 테스트_통과 bool = true

type s가상TB struct{ *testing.T }

func (s s가상TB) Error(args ...interface{})                 { 테스트_통과 = false }
func (s s가상TB) Errorf(format string, args ...interface{}) { 테스트_통과 = false }
func (s s가상TB) Fail()                                     { 테스트_통과 = false }
func (s s가상TB) FailNow()                                  { 테스트_통과 = false }
func (s s가상TB) Failed() bool                              { return !테스트_통과 }
func (s s가상TB) Fatal(args ...interface{})                 { 테스트_통과 = false }
func (s s가상TB) Fatalf(format string, args ...interface{}) { 테스트_통과 = false }
func (s s가상TB) Log(args ...interface{})                   {}
func (s s가상TB) Logf(format string, args ...interface{})   {}
func (s s가상TB) Skip(args ...interface{})                  {}
func (s s가상TB) SkipNow()                                  {}
func (s s가상TB) Skipf(format string, args ...interface{})  {}
func (s s가상TB) Skipped() bool                             { return false }
func (s s가상TB) 테스트용_가상_객체()                               {}

func TestS가상TB(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)
	var tb testing.TB = 가상_테스트
	tb.Failed()
	var i테스트용_가상_객체 I테스트용_가상_객체 = 가상_테스트
	i테스트용_가상_객체.테스트용_가상_객체()

	테스트_통과 = true
	가상_테스트.Error()
	if 테스트_통과 || !가상_테스트.Failed() {
		테스트.Errorf("TestS가상TB() : 에러 1")
	}

	테스트_통과 = true
	가상_테스트.Errorf("")
	if 테스트_통과 || !가상_테스트.Failed() {
		테스트.Errorf("TestS가상TB() : 에러 2")
	}

	테스트_통과 = true
	가상_테스트.Fail()
	if 테스트_통과 || !가상_테스트.Failed() {
		테스트.Errorf("TestS가상TB() : 에러 3")
	}

	테스트_통과 = true
	가상_테스트.FailNow()
	if 테스트_통과 || !가상_테스트.Failed() {
		테스트.Errorf("TestS가상TB() : 에러 4")
	}

	테스트_통과 = true
	가상_테스트.Fatal()
	if 테스트_통과 || !가상_테스트.Failed() {
		테스트.Errorf("TestS가상TB() : 에러 5")
	}

	테스트_통과 = true
	가상_테스트.Fatalf("")
	if 테스트_통과 || !가상_테스트.Failed() {
		테스트.Errorf("TestS가상TB() : 에러 6")
	}
}

func TestF참인지_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	// 간단한 형식
	테스트_통과 = true
	테스트_결과_반환값 := F참인지_확인(가상_테스트, true)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F참인지_확인(가상_테스트, false)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	// 포맷 문자열 있는 폼
	테스트_통과 = true
	테스트_결과_반환값 = F참인지_확인(가상_테스트, true, "포맷_문자열")
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F참인지_확인(가상_테스트, false, "포맷_문자열")
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	// 포맷 문자열 없는 폼
	테스트_통과 = true
	테스트_결과_반환값 = F참인지_확인(가상_테스트, true, 1, 2)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F참인지_확인(가상_테스트, false, 1, 2)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}
}

func TestF거짓인지_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	// 간단한 형식
	테스트_통과 = true
	테스트_결과_반환값 := F거짓인지_확인(가상_테스트, true)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F거짓인지_확인(가상_테스트, false)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	// 포맷 문자열 있는 폼
	테스트_통과 = true
	테스트_결과_반환값 = F거짓인지_확인(가상_테스트, true, "포맷_문자열")
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F거짓인지_확인(가상_테스트, false, "포맷_문자열")
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	// 포맷 문자열 없는 폼
	테스트_통과 = true
	테스트_결과_반환값 = F거짓인지_확인(가상_테스트, true, 1, 2)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F거짓인지_확인(가상_테스트, false, 1, 2)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}
}

func TestF에러없음_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F에러없음_확인(가상_테스트, nil)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F에러없음_확인(가상_테스트, fmt.Errorf(""))
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}
}

func TestF에러발생_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F에러발생_확인(가상_테스트, nil)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F에러발생_확인(가상_테스트, fmt.Errorf(""))
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}
}

func TestF같은값_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F같은값_확인(가상_테스트, 1, 1)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F같은값_확인(가상_테스트, 1, 2)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}
}

func TestF다른값_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F다른값_확인(가상_테스트, 1, 1)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Error("%s예상치 못한 테스트 통과.", F소스코드_위치(0))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F다른값_확인(가상_테스트, 1, 2)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Error("%s예상치 못한 테스트 실패.", F소스코드_위치(0))
	}
}

func TestF에러_생성(테스트 *testing.T) {
	F같은값_확인(테스트, F에러_생성("테스트 %v, %v", 1, 2.2).Error(), "테스트 1, 2.2")
}

func TestF소스코드_위치(테스트 *testing.T) {
	소스코드_위치 := strings.Split(F소스코드_위치(-1), ":")
	파일명, 행_번호 := 소스코드_위치[0], 소스코드_위치[1]

	F같은값_확인(테스트, 파일명, "util_functions.go")

	소스코드_위치 = strings.Split(F소스코드_위치(0), ":")
	파일명, 행_번호 = 소스코드_위치[0], 소스코드_위치[1]
	_, _, 행_번호_예상값, _ := runtime.Caller(0)

	F같은값_확인(테스트, 파일명, "util_functions_test.go")
	F같은값_확인(테스트, F문자열(int64(행_번호_예상값-2)), 행_번호)
}