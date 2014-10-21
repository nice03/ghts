// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"fmt"
	"math/big"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func F엄격한_매개변수_안전성_검사(값_모음 ...I가변형) bool {
	if !F매개변수_안전성_검사(값_모음...) {
		return false
	}

	for _, 값 := range 값_모음 {
		종류 := reflect.TypeOf(값).Kind()

		if 종류 == reflect.Chan || 종류 == reflect.Func {
			F문자열_출력("안전하지 않은 매개변수 형식 : %v", F값_확인_문자열(값))

			에러 := F에러_생성("%s안전하지 않은 매개변수 형식 : %s\n%s\n%s\n%s\n%s",
				F소스코드_위치(1), F값_확인_문자열(값),
				F소스코드_위치(2), F소스코드_위치(3), F소스코드_위치(4), F소스코드_위치(5))

			panic(에러)

			return false
		}
	}

	return true
}

// 매개변수가 data race를 일으킬 위험이 있는 지 검사.
// 현재는 알려진 몇몇 형식에 대해서만 제대로 작동함.
// 이후 추가하거나, 근본적인 자동 검사가 가능하도록 개선할 것.
func F매개변수_안전성_검사(값_모음 ...I가변형) bool {
	if P매개변수_안전성_검사_건너뛰기 {
		return true
	}

	값_모음 = F중첩된_외부_슬라이스_제거(값_모음)

	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case I변수형:
			// mutable하고 자동으로 복사도 안 되는 형식은 위험하다고 판단.
			// continue 하지 않고, panic()을 향해 감.
			// proceed to panic()
		case nil, uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64,
			float32, float64, complex64, complex128,
			bool, string, time.Time:
			// CallByValue에 의해서 자동으로 복사본이 생성되는 형식.
			continue
		case *sC부호없는_정수64, *sC정수64, *sC실수64, *sC정밀수,
			*sC참거짓, *sC문자열, *sC시점, *sC통화, *sC매개변수,
			*sC안전한_가변형, *s안전한_배열, *s안전한_맵:
			// Immutable 하므로 race condition이 발생하지 않는 형식.
			// 앞으로 여기에 검증된 상수형을 더 추가해야 됨.
			continue
		case *sV부호없는_정수64, *sV정수64, *sV실수64, *sV정밀수,
			*sV참거짓, *sV시점, *sV통화:
			// Atomic 하거나 Mutex로 보호되어 있으면서도,
			// Deadlock을 발생시키지 않는 자료형들.
			continue
		}

		종류 := reflect.TypeOf(값).Kind()

		switch 종류 {
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
			reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128, reflect.String:
			// CallByValue에 의해서 자동으로 복사본이 생성되는 커스텀 형식.
			continue
		case reflect.Chan, reflect.Func:
			// 문제의 소지가 있지만, 상수형으로 변경할 수 없는 형식
			continue
		}

		F문자열_출력("안전하지 않은 매개변수 형식 : %v", F값_확인_문자열(값))

		에러 := F에러_생성("%s안전하지 않은 매개변수 형식 : %s\n%s\n%s\n%s\n%s",
			F소스코드_위치(1), F값_확인_문자열(값),
			F소스코드_위치(2), F소스코드_위치(3), F소스코드_위치(4), F소스코드_위치(5))

		panic(에러)

		return false
	}

	return true
}

// 알려진 몇몇 자료형을 안전한 형식으로 변환.
// 매개변수 안전성 검사를 의도적으로 생략함.
func F안전한_형식으로_변환_시도(값_모음 ...I가변형) []I가변형 {
	if F_nil값임(값_모음) {
		return nil
	}

	값_모음 = F중첩된_외부_슬라이스_제거(값_모음)

	반환값 := make([]I가변형, len(값_모음))

	for 인덱스, 값 := range 값_모음 {
		if F_nil값임(값) {
			반환값[인덱스] = nil
			continue
		}

		switch 값.(type) {
		case nil, uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64, float32, float64,
			bool, string, time.Time, 
			*sC부호없는_정수64, *sC정수64, *sC실수64, *sC정밀수,
			*sC참거짓, *sC문자열, *sC시점, *sC통화, *sC매개변수,
			*sC안전한_가변형, *s안전한_배열, *s안전한_맵,
			*sV부호없는_정수64, *sV정수64, *sV실수64, *sV정밀수,
			*sV참거짓, *sV시점, *sV통화:
			// 이미 안전한 형식이므로 변환이 필요없음.
			반환값[인덱스] = 값
			continue
		case *big.Int, *big.Rat:
			반환값[인덱스] = NC정밀수(값)
		case error, reflect.Value, reflect.Type:
			// 디버깅과 테스트에 자주 쓰이지만,
			// 어떻게 해야 될 지 모르겠으니 일단 그냥 넘어가자.
			continue
		case I변수형:
			상수형 := F상수형(값)

			if 상수형 != nil {
				// I상수형 변환 성공
				반환값[인덱스] = 상수형
				continue
			} else {
				F문자열_출력("예상하지 못한 변수형. %v번째 입력값 %v %v.",
					인덱스, F값_확인_문자열(값))
				반환값[인덱스] = 값 // I상수형 변환 실패
				continue
			}
		}

		종류 := reflect.TypeOf(값).Kind()

		switch 종류 {
		case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
			reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.Chan, reflect.Func, reflect.String:
			반환값[인덱스] = 값
			continue
		}

		F문자열_출력("예상하지 못한 형식. %v번째 입력값 %v", 인덱스, F값_확인_문자열(값))

		반환값[인덱스] = 값
	}

	return 반환값
}

// 몇몇 기본 자료형을 'I상수형'으로 변환.
// 매개변수 안전성 검사를 의도적으로 생략함.
func F상수형(값 I가변형) I상수형 {
	if F_nil값임(값) {
		return nil
	}

	// 지금 알려진 것만 우선 포함 시킴.
	switch 값.(type) {
	case nil:
		return nil
	case *sC정수64, *sC부호없는_정수64, *sC실수64, *sC정밀수, *sC통화, *sC시점,
		*sC문자열, *sC참거짓:
		return 값.(I상수형)
	case uint:
		return NC부호없는_정수(uint64(값.(uint)))
	case uint8:
		return NC부호없는_정수(uint64(값.(uint8)))
	case uint16:
		return NC부호없는_정수(uint64(값.(uint16)))
	case uint32:
		return NC부호없는_정수(uint64(값.(uint32)))
	case uint64:
		return NC부호없는_정수(값.(uint64))
	case int:
		return NC정수(int64(값.(int)))
	case int8:
		return NC정수(int64(값.(int8)))
	case int16:
		return NC정수(int64(값.(int16)))
	case int32:
		return NC정수(int64(값.(int32)))
	case int64:
		return NC정수(값.(int64))
	case float32:
		// float32는 정밀도가 떨어짐.
		// 바로 float64로 변환할 경우 너무 심하게 느껴질 정도임.
		// 문자열로 변환을 거쳐서 float64로 변환할 경우 그나마 조금 낫지만,
		// 여전히 정밀도가 부족함.
		if !strings.Contains(F소스코드_위치(1), strings.Split(F소스코드_위치(0), ":")[0]) &&
			!strings.Contains(F소스코드_위치(1), "_test.go") {
			F문자열_출력("float32는 아주 부정확함. float64나 I정밀수 자료형을 권장함.")
		}
		문자열 := strconv.FormatFloat(float64(값.(float32)), 'f', -1, 32)
		실수, _ := F문자열2실수(문자열)

		return NC실수(실수)
	case float64:
		return NC실수(값.(float64))
	case bool:
		return NC참거짓(값.(bool))
	case string:
		return NC문자열(값.(string))
	case time.Time:
		return NC시점(값.(time.Time))
	case *big.Int:
		return NC정밀수(값.(*big.Int).String())
	case *big.Rat:
		return NC정밀수(값.(*big.Rat).String())
	case *sV부호없는_정수64:
		return 값.(V부호없는_정수).G상수형()
	case *sV정수64:
		return 값.(V정수).G상수형()
	case *sV실수64:
		return 값.(V실수).G상수형()
	case *sV정밀수:
		return 값.(V정밀수).G상수형()
	case *sV통화:
		return 값.(V통화).G상수형()
	case *sV참거짓:
		return 값.(V참거짓).G상수형()
	case *sV시점:
		return 값.(V시점).G상수형()
	default:
		F문자열_출력("알려진 상수형이 아님. %v", F값_확인_문자열(값))

		return nil
	}
}

// 문자열로 변환 시도.
// 매개변수 안전성 검사를 의도적으로 생략함.
func F문자열(값 I가변형) string {
	if 값 == nil {
		return "nil"
	}
	if F_nil값임(값) {
		return "<nil>"
	}

	switch 값.(type) {
	case string:
		return 값.(string)
	case *big.Rat:
		return F마지막_0_제거(값.(*big.Rat).FloatString(100))
	case float64:
		return strconv.FormatFloat(값.(float64), 'f', -1, 64)
	case time.Time:
		return 값.(time.Time).Format(P시점_형식)
	case I기본_문자열:
		return 값.(I기본_문자열).String()
	}

	return F포맷된_문자열("%v", 값)
}

// 문자열로 변환 시도.
// 매개변수 안전성 검사를 의도적으로 생략함.
func F포맷된_문자열(포맷_문자열 string, 추가_내용 ...I가변형) string {
	에러 := F에러_생성(포맷_문자열, 추가_내용...)

	return 에러.Error()
}

func F금액_문자열(금액_문자열 string) string {
	문자_슬라이스 := strings.Split(금액_문자열, "")
	소숫점_인덱스 := strings.Index(금액_문자열, ".")
	역순_버퍼 := new(bytes.Buffer)
	소숫점_내지_콤마로부터_거리 := 0

	if 소숫점_인덱스 == -1 {
		소숫점_인덱스 = len(금액_문자열) // 정수라서 소숫점이 없는 경우.
	}

	for 인덱스 := len(문자_슬라이스) - 1; 인덱스 >= 0; 인덱스-- {
		switch {
		case 인덱스 > 소숫점_인덱스:
			역순_버퍼.WriteString(문자_슬라이스[인덱스])
		case 인덱스 == 소숫점_인덱스:
			역순_버퍼.WriteString(문자_슬라이스[인덱스])
			소숫점_내지_콤마로부터_거리 = 0
		case 소숫점_내지_콤마로부터_거리 == 3:
			역순_버퍼.WriteString(",")
			역순_버퍼.WriteString(문자_슬라이스[인덱스])
			소숫점_내지_콤마로부터_거리 = 1
		default:
			역순_버퍼.WriteString(문자_슬라이스[인덱스])
			소숫점_내지_콤마로부터_거리++
		}
	}

	역순_금액_문자열 := 역순_버퍼.String()
	문자_슬라이스 = strings.Split(역순_금액_문자열, "")
	버퍼 := new(bytes.Buffer)

	for 인덱스 := len(문자_슬라이스) - 1; 인덱스 >= 0; 인덱스-- {
		버퍼.WriteString(문자_슬라이스[인덱스])
	}

	return 버퍼.String()
}

func F마지막_0_제거(문자열 string) string {
	if !strings.Contains(문자열, ".") {
		return 문자열 // 정수임.
	}

	const asc코드_0 uint8 = uint8(48)
	const asc코드_소숫점 uint8 = uint8(46)

	종료_지점 := len(문자열) - 1

	for 인덱스 := len(문자열) - 1; 인덱스 >= 0; 인덱스-- {
		switch {
		case 문자열[인덱스] == asc코드_0:
			continue
		case 문자열[인덱스] == asc코드_소숫점:
			종료_지점 = 인덱스
		default:
			종료_지점 = 인덱스 + 1
		}

		break
	}

	return 문자열[:종료_지점]
}

func F반올림(값 I가변형, 소숫점_이하_자릿수 int) C정밀수 {
	F매개변수_안전성_검사(값)

	if !F숫자형식임(값) {
		F문자열_출력("숫자형식이 아님. %v", F값_확인_문자열(값))
		return nil
	}

	정밀수 := NV정밀수(값)

	if 정밀수 == nil {
		F문자열_출력("예상치 못한 경우. %v", F값_확인_문자열(값))
		return nil
	}

	return 정밀수.S반올림(소숫점_이하_자릿수).G상수형()
}

func F문자열2실수(값 string) (float64, error) {
	실수, 에러 := strconv.ParseFloat(strings.Replace(값, ",", "", -1), 64)

	if 에러 != nil {
		F문자열_출력("실수 변환 에러. %v", 값)
		return 0.0, 에러
	}

	return 실수, nil
}

func F문자열2시점(값 string) (time.Time, error) {
	시점, 에러 := time.Parse(P시점_형식, 값)

	if 에러 == nil {
		return 시점, nil
	}

	시점, 에러 = time.Parse(P일자_형식, 값)

	if 에러 == nil {
		return 시점, nil
	}

	F문자열_출력("시점 변환 에러. %v", 값)

	return time.Time{}, 에러
}

func F시점_문자열(시점 time.Time) string {
	return 시점.Format(P시점_형식)
}

func F일자_문자열(일자 time.Time) string {
	return 일자.Format(P일자_형식)
}

func F시점_복사(값 time.Time) time.Time {
	복사본 := time.Date(값.Year(), 값.Month(), 값.Day(),
		값.Hour(), 값.Minute(), 값.Second(),
		값.Nanosecond(), 값.Location())

	return 복사본
}

func F임의_통화종류() P통화종류 {
	return P통화종류(int(rand.Int31n(int32(len(통화종류_문자열_모음)))))
}

func F통화종류별_정밀도(통화 P통화종류) int {
	switch 통화 {
	case KRW:
		return 0
	case USD, CNY, EUR:
		return 2
	default:
		return 2
	}
}

func F통화종류(값_모음 ...I가변형) P통화종류 {
	F매개변수_안전성_검사(값_모음...)

	통화종류_맵 := make(map[P통화종류]P통화종류)

	for _, 값 := range 값_모음 {
		if F통화형식임(값) {
			통화종류 := 값.(I통화).G종류()

			// 중복없도록 하기 위해서 맵을 사용함.
			// 대입하는 값은 아무 의미 없음.
			통화종류_맵[통화종류] = 통화종류
		}
	}

	if len(통화종류_맵) == 1 {
		for 통화종류, _ := range 통화종류_맵 {
			return 통화종류
		}
	}

	return INVALID_CURRENCY_TYPE
}

func F통화형식임(값_모음 ...I가변형) bool {
	//F매개변수_안전성_검사(값_모음...)

	if F_nil값_존재함(값_모음...) {
		return false
	}

	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case *sC통화, *sV통화:
			continue
		default:
			return false
		}
	}

	return true
}

func F통화_같음(값1, 값2 I가변형) bool {
	if !F통화형식임(값1, 값2) {
		return false
	}

	통화1 := 값1.(I통화)
	통화2 := 값2.(I통화)

	if 통화1 == nil || 통화2 == nil {
		F문자열_출력("예상치 못한 경우. %v", F값_확인_문자열(값1, 값2))
		return false
	}

	return 통화1.G같음(통화2)
}

func FbigRat형식임(값 I가변형) bool {
	_, bigRat형식임 := 값.(*big.Rat)

	return bigRat형식임
}

func F숫자형식임(값_모음 ...I가변형) bool {
	//F매개변수_안전성_검사(값_모음...)

	if F_nil값_존재함(값_모음...) {
		return false
	}

	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64,
			float32, float64, complex64, complex128,
			*sC정수64, *sC부호없는_정수64, *sC실수64, *sC정밀수,
			*big.Int, *big.Rat, *sV정수64, *sV부호없는_정수64, *sV실수64, *sV정밀수:
			continue
		default:
			문자열 := F포맷된_문자열("%v", 값)
			if _, 에러 := strconv.ParseFloat(문자열, 64); 에러 != nil {
				return false
			}
		}
	}

	return true
}

func F숫자_같음(값1, 값2 I가변형) bool {
	if !F숫자형식임(값1, 값2) {
		return false
	}

	정밀수1 := NC정밀수(값1)
	정밀수2 := NC정밀수(값2)

	if 정밀수1 == nil || 정밀수2 == nil {
		F문자열_출력("예상치 못한 경우. %v", F값_확인_문자열(값1, 값2))

		return false
	}

	return 정밀수1.G같음(정밀수2)
}

func F참거짓형식임(값_모음 ...I가변형) bool {
	F매개변수_안전성_검사(값_모음...)

	if F_nil값_존재함(값_모음...) {
		return false
	}

	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case bool, *sC참거짓, *sV참거짓:
			continue
		default:
			return false
		}
	}

	return true
}

func F참거짓_같음(값1, 값2 I가변형) bool {
	if !F참거짓형식임(값1, 값2) {
		return false
	}

	var 참거짓1, 참거짓2 bool

	switch 값1.(type) {
	case bool:
		참거짓1 = 값1.(bool)
	case *sC참거짓, *sV참거짓:
		참거짓1 = 값1.(I참거짓).G값()
	default:
		return false
	}

	switch 값2.(type) {
	case bool:
		참거짓2 = 값2.(bool)
	case *sC참거짓, *sV참거짓:
		참거짓2 = 값2.(I참거짓).G값()
	default:
		return false
	}

	return 참거짓1 == 참거짓2
}

func F문자열형식임(값_모음 ...I가변형) bool {
	F매개변수_안전성_검사(값_모음...)

	if F_nil값_존재함(값_모음...) {
		return false
	}

	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case string, *sC문자열:
			continue
		default:
			return false
		}
	}

	return true
}

func F문자열_같음(값1, 값2 I가변형) bool {
	if !F문자열형식임(값1, 값2) {
		return false
	}

	var 문자열1, 문자열2 string

	switch 값1.(type) {
	case string:
		문자열1 = 값1.(string)
	case *sC문자열:
		문자열1 = 값1.(*sC문자열).G값()
	default:
		return false
	}

	switch 값2.(type) {
	case string:
		문자열2 = 값2.(string)
	case *sC문자열:
		문자열2 = 값2.(*sC문자열).G값()
	default:
		return false
	}

	return 문자열1 == 문자열2
}

func F시점형식임(값_모음 ...I가변형) bool {
	F매개변수_안전성_검사(값_모음...)

	if F_nil값_존재함(값_모음...) {
		return false
	}

	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case time.Time, *sC시점, *sV시점:
			continue
		default:
			return false
		}
	}

	return true
}

func F시점_같음(값1, 값2 I가변형) bool {
	if !F시점형식임(값1, 값2) {
		return false
	}

	var 시점1, 시점2 time.Time

	switch 값1.(type) {
	case *sC시점, *sV시점:
		시점1 = 값1.(I시점).G값()
	case time.Time:
		시점1 = 값1.(time.Time)
	default:
		return false
	}

	switch 값2.(type) {
	case *sC시점, *sV시점:
		시점2 = 값2.(I시점).G값()
	case time.Time:
		시점2 = 값2.(time.Time)
	default:
		return false
	}

	return 시점1.Equal(시점2)
}

func F값_같음(값1, 값2 I가변형) (값_같음 bool) {
	defer func() {
		if 에러 := recover(); 에러 != nil {
			F문자열_출력("%v %v", 에러, F값_확인_문자열(값1, 값2))
			값_같음 = false
		}
	}()

	F매개변수_안전성_검사(값1, 값2)

	switch {
	case F_nil값_존재함(값1, 값2):
		return F_nil값임(값1) && F_nil값임(값2)
	case F통화_같음(값1, 값2), F숫자_같음(값1, 값2), F참거짓_같음(값1, 값2),
		F시점_같음(값1, 값2), F문자열_같음(값1, 값2), reflect.DeepEqual(값1, 값2):
		return true
	}

	return false
}

// Exponential Back-off
// 실패한 후 반복하는 횟수에 따라 기다리는 최대시간이 기하급수적으로 커짐.
func F잠시_대기(반복횟수 int) {
	len_대기시간_한도 := len(대기시간_한도) - 1
	var 임의_대기시간 int64

	if 반복횟수 > len_대기시간_한도 {
		임의_대기시간 = rand.Int63n(대기시간_한도[len_대기시간_한도])
		
		if 임의_대기시간 > P최대_대기시간 {
			F문자열_출력("'임의_대기시간'이 'P최대_대기시간'보다 큼. %v", 임의_대기시간)
			임의_대기시간 = rand.Int63n(대기시간_한도[len_대기시간_한도])
		}
	} else {
		임의_대기시간 = rand.Int63n(P최대_대기시간)
	}
	
	
	
	time.Sleep(time.Duration(P최소_대기시간 + 임의_대기시간))
}

// 호출할 때 매개변수_모음 뒤에 "..."를 빼먹은 경우에 중첩된 슬라이스를
// 정상 슬라이스로 변환하기 위함.
func F중첩된_외부_슬라이스_제거(슬라이스 []I가변형) []I가변형 {
	반복횟수 := 0

	for len(슬라이스) == 1 {
		if _, ok := 슬라이스[0].([]I가변형); !ok {
			break
		}

		슬라이스 = 슬라이스[0].([]I가변형)
		반복횟수++
	}

	if 반복횟수 > 0 {
		F문자열_출력("중첩된 슬라이스 발견. 중첩 깊이 : %v", 반복횟수)
	}

	return 슬라이스
}

func F가변형2interface(값_모음 []I가변형) []interface{} {
	if 값_모음 == nil {
		return nil
	}

	반환값 := make([]interface{}, len(값_모음))

	for 인덱스, 값 := range 값_모음 {
		반환값[인덱스] = 값
	}

	return 반환값
}

func F_nil값_존재함(값_모음 ...I가변형) bool {
	for _, 값 := range 값_모음 {
		if F_nil값임(값) {
			return true
		}
	}

	return false
}

func F_nil값임(값 I가변형) bool {
	if 값 == nil {
		return true
	}

	종류 := reflect.TypeOf(값).Kind()

	// chan, func, interface, map, pointer, or slice
	switch 종류 {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice:
		if reflect.ValueOf(값).IsNil() {
			return true
		}
	}

	switch 값.(type) {
	case I정밀수:
		return 값.(I정밀수).GRat() == nil
	case I통화:
		return 값.(I통화).G값() == nil
	}

	return false
}

// reflect 편의 함수

// 슬라이스 복사하는 편의 함수.
// 는 복사본 슬라이스 변수를 생성해 줘야 함.
// 슬라이스를 주고 받는 것은 안전성에 문제의 소지가 있다.
// builtin.copy()를 쓰도록 하자.
/*
func F슬라이스_복사(원본slice I가변형) I가변형 {
	// 슬라이스를 주고 받는 것은 안전성에 문제의 소지가 있다.
	//if !F매개변수_안전성_검사(원본slice) { return nil }

	원본값 := reflect.ValueOf(원본slice)

	// 원본값 검사.
	if 원본값.IsNil() {
		panic("F슬라이스_복사 : 원본이 nil임.")
	}
	if !원본값.IsValid() {
		panic("F슬라이스_복사 : 원본값이 유효하지 않은 zero값임.")
	}
	if 원본값.Kind() != reflect.Slice {
		panic("F슬라이스_복사 : 원본이 slice가 아님.")
	}
	if 원본값.Len() == 0 {
		panic("F슬라이스_복사 : 원본 슬라이스 길이가 0임.")
	}

	구성요소형식 := 원본값.Index(0).Type()
	슬라이스형식 := reflect.SliceOf(구성요소형식)
	복사본 := reflect.MakeSlice(슬라이스형식, 원본값.Len(), 원본값.Cap())
	reflect.Copy(복사본, 원본값)

	return 복사본.Interface()
} */

// 테스트 편의 함수
// 테스트 함수에서는 매개변수 안전성 검사를 하지 않도록 하자.

func F테스트_모드() bool { return 테스트_모드.G값() }
func F테스트_모드_시작()   { 테스트_모드.S값(true) }
func F테스트_모드_종료()   { 테스트_모드.S값(false) }

func F문자열_출력_일시정지_모드() bool { return 문자열_출력_일시정시.G값() }
func F문자열_출력_일시정지_시작()      { 문자열_출력_일시정시.S값(true) }
func F문자열_출력_일시정지_종료()      { 문자열_출력_일시정시.S값(false) }

// '참거짓'이 false이면 Fail하는 테스트용 편의 함수.
func F참인지_확인(테스트 testing.TB, 참거짓 bool, 추가_매개변수 ...I가변형) (테스트_통과 bool) {
	if !참거짓 {
		switch 테스트.(type) {
		case i테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()

			if 추가_매개변수 == nil || len(추가_매개변수) == 0 {
				F문자열_출력2(1, "주어진 조건이 false임.")
			} else {
				switch 추가_매개변수[0].(type) {
				case string:
					포맷_문자열 := 추가_매개변수[0].(string)
					F문자열_출력2(1, 포맷_문자열, 추가_매개변수[1:]...)
				default:
					포맷_문자열 := "주어진 조건이 false임.\n"

					for 반복횟수 := 0; 반복횟수 < len(추가_매개변수); 반복횟수++ {
						포맷_문자열 = 포맷_문자열 + " %v\n"
					}
					포맷_문자열 = 포맷_문자열 + ".\n"

					F문자열_출력2(1, 포맷_문자열, 추가_매개변수...)
				}
			}
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// '참거짓'이 true이면 Fail하는 테스트용 편의 함수.
func F거짓인지_확인(테스트 testing.TB, 참거짓 bool, 추가_매개변수 ...I가변형) (테스트_통과 bool) {
	if 참거짓 {
		switch 테스트.(type) {
		case i테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()

			if 추가_매개변수 == nil || len(추가_매개변수) == 0 {
				F문자열_출력2(1, "주어진 조건이 true임.")
			} else {
				switch 추가_매개변수[0].(type) {
				case string:
					포맷_문자열 := 추가_매개변수[0].(string)
					F문자열_출력2(1, 포맷_문자열, 추가_매개변수[1:]...)
				default:
					포맷_문자열 := "주어진 조건이 true임.\n"

					for 반복횟수 := 0; 반복횟수 < len(추가_매개변수); 반복횟수++ {
						포맷_문자열 = 포맷_문자열 + " %v\n"
					}
					포맷_문자열 = 포맷_문자열 + ".\n"

					F문자열_출력2(1, 포맷_문자열, 추가_매개변수...)
				}
			}
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// 에러가 발생하면 Fail하는 테스트용 편의 함수.
func F에러없음_확인(테스트 testing.TB, 에러 error) (테스트_통과 bool) {
	if 에러 != nil {
		switch 테스트.(type) {
		case i테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()
			F문자열_출력2(1, 에러.Error())
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// 에러가 발생하지 않으면 Fail하는 테스트용 편의 함수.
func F에러발생_확인(테스트 testing.TB, 에러 error) (테스트_통과 bool) {
	if 에러 == nil {
		switch 테스트.(type) {
		case i테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()

			F문자열_출력2(1, 에러.Error())
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 다르면 Fail하는 테스트용 편의 함수.
func F같은값_확인(테스트 testing.TB, 값1, 값2 I가변형) (테스트_통과 bool) {
	값_모음 := F안전한_형식으로_변환_시도(값1, 값2)
	값1 = 값_모음[0]
	값2 = 값_모음[1]

	if !F값_같음(값1, 값2) {
		switch 테스트.(type) {
		case i테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()
			F문자열_출력2(1, "서로 다름. %v", F값_확인_문자열(값1, 값2))
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 같으면 Fail하는 테스트용 편의 함수.
func F다른값_확인(테스트 testing.TB, 값1, 값2 I가변형) (테스트_통과 bool) {
	값_모음 := F안전한_형식으로_변환_시도(값1, 값2)
	값1 = 값_모음[0]
	값2 = 값_모음[1]

	if F값_같음(값1, 값2) {
		switch 테스트.(type) {
		case i테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()
			F문자열_출력2(1, "서로 같음. %v.", F값_확인_문자열(값1, 값2))
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

func F패닉발생_확인(테스트 testing.TB, 함수 I가변형, 매개변수 ...I가변형) (테스트_통과 bool) {
	defer func() {
		에러 := recover()

		if 에러 != nil {
			테스트_통과 = true
		} else {
			switch 테스트.(type) {
			case i테스트용_가상_객체:
				// PASS
			default:
				F문자열_출력_일시정지_종료()
				F문자열_출력2(1, "패닉이 발생하지 않음. 함수 %v %v",
					F값_확인_문자열(함수), F값_확인_문자열(매개변수...))
			}

			테스트_통과 = false

			테스트.FailNow()
		}
	}()

	입력값 := make([]reflect.Value, len(매개변수))

	for 인덱스, 값 := range 매개변수 {
		입력값[인덱스] = reflect.ValueOf(값)
	}

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_종료()

	reflect.ValueOf(함수).Call(입력값)

	return false
}

// 소스코드 위치를 나타내는 함수. runtime.Caller()의 한글화 버전임.
// '건너뛰는_단계'값이 커질수록 호출 경로를 거슬러 올라감.
//
// -1 = F소스코드_위치() 자기자신의 위치.
// 0 = F소스코드_위치()를 호출한 메소드의 위치.
// 1 = F소스코드_위치()를 호출한 메소드를 호출한 메소드의 위치
// 2, 3, 4,....n = 계속 거슬러 올라감.
//
// 다른 모듈을 위해서 사용되는 라이브러리 펑션인 경우 1가 적당함.
// 그렇지 않다면, 0이 적당함.
func F소스코드_위치(건너뛰는_단계 int) string {
	건너뛰는_단계 = 건너뛰는_단계 + 1 // 이 메소드를 호출한 함수를 기준으로 0이 되게 하기 위함.
	pc, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)
	함수_이름 := runtime.FuncForPC(pc).Name()

	함수_이름 = strings.Replace(함수_이름, "github.com/gh-system/", "", -1)

	파일명 := filepath.Base(파일_경로)
	return 파일명 + ":" + F문자열(행_번호) + ":" + 함수_이름 + "() "
}

// 에러 처리 편의 함수.
func F에러_생성(문자열 string, 추가_내용 ...I가변형) error {
	for strings.HasSuffix(문자열, "\n") {
		문자열 += "\n"
	}

	추가_내용_ := make([]interface{}, len(추가_내용))

	for 인덱스, 내용 := range 추가_내용 {
		추가_내용_[인덱스] = 내용
	}

	return fmt.Errorf(문자열, 추가_내용_...)
}

func F문자열_출력(문자열 string, 추가_매개변수 ...I가변형) {
	F문자열_출력2(1, 문자열, 추가_매개변수...)
}

func F문자열_출력2(추가적인_건너뛰기_단계 int, 문자열 string, 추가_매개변수 ...I가변형) {
	if F문자열_출력_일시정지_모드() {
		return
	}

	i := 추가적인_건너뛰기_단계

	fmt.Println("")
	fmt.Printf(F소스코드_위치(1+i)+": "+문자열+"\n", F가변형2interface(추가_매개변수)...)
	fmt.Println(F소스코드_위치(2 + i))
	fmt.Println(F소스코드_위치(3 + i))
	fmt.Println(F소스코드_위치(4 + i))
	fmt.Println(F소스코드_위치(5 + i))
	fmt.Println(F소스코드_위치(6 + i))
	fmt.Println(F소스코드_위치(7 + i))
	fmt.Println(F소스코드_위치(8 + i))
	fmt.Println(F소스코드_위치(9 + i))
	fmt.Println(F소스코드_위치(10 + i))
}

// 디버깅 편의 함수.
// 디버깅 편의 함수는 일시적으로 사용하며,
// 실제 production 환경에서는 사용되지 않는다고 보고,
// 매개변수의 안전성을 검사하지 않는다.
func F체크포인트(체크포인트_번호 *int, 추가_매개변수 ...I가변형) {
	추가_매개변수 = append([]I가변형{F소스코드_위치(1)}, 추가_매개변수...)
	문자열 := F포맷된_문자열("%s체크포인트 %v ", F소스코드_위치(1), *체크포인트_번호)

	fmt.Println(append([]interface{}{문자열}, F가변형2interface(추가_매개변수)...)...)
	(*체크포인트_번호)++
}

func F값_확인(값_모음 ...I가변형) {
	fmt.Println(F소스코드_위치(1), "값_확인 :", F값_확인_문자열(값_모음...))
}

func F값_확인_문자열(값_모음 ...I가변형) string {
	버퍼 := new(bytes.Buffer)

	for 인덱스, 값 := range 값_모음 {
		if 인덱스 == 0 {
			버퍼.WriteString(" ")
		} else {
			버퍼.WriteString(", ")
		}

		if len(값_모음) == 1 {
			버퍼.WriteString(
				F포맷된_문자열("형식 : %v, 값 : %v", reflect.TypeOf(값), 값))
		} else {
			버퍼.WriteString(
				F포맷된_문자열("형식%v : %v, 값%v : %v",
					인덱스+1, reflect.TypeOf(값), 인덱스+1, 값))
		}
	}
	//버퍼.WriteString("\n")

	return 버퍼.String()
}

// 메모 편의 함수.
var 이미_출력한_TODO_모음 []string = make([]string, 0)

// 해야할 일을 소스코드 위치와 함께 표기해 주는 메소드.
func F메모(문자열 string) {
	for _, 이미_출력한_TODO := range 이미_출력한_TODO_모음 {
		if 문자열 == 이미_출력한_TODO {
			return
		}
	}

	이미_출력한_TODO_모음 = append(이미_출력한_TODO_모음, 문자열)

	fmt.Printf("TODO : %s %s\n\n", F소스코드_위치(1), 문자열)
}
