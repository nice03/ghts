package common

import (
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

func F안전한_매개변수(값_모음 ...interface{}) bool {
	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64,
			float32, float64, bool, string, time.Time:
			// CallByValue에 의해서 자동으로 복사본이 생성되는 형식.
			//OK to PASS
		case *sC부호없는_정수64, *sC정수64, *sC실수64,
			*sC참거짓, *sC문자열, *sC시점, *sC정밀수, *sC통화:
			// Immutable 하므로 race condition이 발생하지 않는 형식.
			// 앞으로 여기에 검증된 상수형을 더 추가해야 됨.
			// OK to PASS
		default:
			// 알려진 상수형이 아닌 경우에는 안전하지 않다고 판단.		
			return false
		}
	}

	return true
}

func F상수형(값 interface{}) I상수형 {
	if 값 == nil {
		return nil
	}
	
	/*
	switch 형식 {
	case "*common.sC정수64", "*common.sC부호없는_정수64", "*common.sC실수64", 
			"*common.sC정밀수", "*common.sC통화", "*common.sC참거짓", 
			"*common.sC문자열", "*common.sC시점":
		return 값.(I상수형)
	} */

	// 지금 알려진 것만 우선 포함 시킴.
	switch 값.(type) {
	case *sC정수64, *sC부호없는_정수64, *sC실수64, *sC정밀수, *sC통화,
		*sC참거짓, *sC문자열, *sC시점:
		return 값.(I상수형)
	case bool:
		return NC참거짓(값.(bool))
	case string:
		return NC문자열(값.(string))
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
		실수, 에러 := F문자열2실수(문자열)

		if 에러 != nil {
			return nil
		}

		return NC실수(실수)
	case float64:
		return NC실수(값.(float64))
	case time.Time:
		return NC시점(값.(time.Time))
	case *big.Int:
		return NC정밀수(값.(*big.Int).String())
	case *big.Rat:
		return NC정밀수(값.(*big.Rat).String())
	case V부호없는_정수:
		return 값.(V부호없는_정수).G상수형()
	case V정수:
		return 값.(V정수).G상수형()
	case V실수:
		return 값.(V실수).G상수형()
	case V정밀수:
		return 값.(V정밀수).G상수형()
	case V통화:
		return 값.(V통화).G상수형()
	case V참거짓:
		return 값.(V참거짓).G상수형()
	case V시점:
		return 값.(V시점).G상수형()
	default:
		F문자열_출력("알려진 상수형이 아님. 입력값 %v %v.", reflect.TypeOf(값), 값)

		return nil
	}
}

func F문자열(값 interface{}) string {
	switch 값.(type) {
	case string:
		return 값.(string)
	case C문자열:
		return 값.(C문자열).G값()
	case *big.Rat:
		F마지막_0_제거(값.(*big.Rat).FloatString(100))
	case float64:
		return strconv.FormatFloat(값.(float64), 'f', -1, 64)
	case time.Time:
		return 값.(time.Time).Format(P시점_포맷)
	case float32: // 잘못 변환하면 너무 부정확해 짐. 비트를 32로 지정해야 그나마 나아짐.
		return strconv.FormatFloat(float64(값.(float32)), 'f', -1, 32)
	case I기본_문자열:
		return 값.(I기본_문자열).String()
	}

	return F포맷된_문자열_생성("%v", 값)
}


//	case uint:
//		return strconv.FormatUint(uint64(값.(uint)), 10)
//	case uint8:
//		return strconv.FormatUint(uint64(값.(uint8)), 10)
//	case uint16:
//		return strconv.FormatUint(uint64(값.(uint16)), 10)
//	case uint32:
//		return strconv.FormatUint(uint64(값.(uint32)), 10)
//	case uint64:
//		return strconv.FormatUint(값.(uint64), 10)
//	case int:
//		return strconv.FormatInt(int64(값.(int)), 10)
//	case int8:
//		return strconv.FormatInt(int64(값.(int8)), 10)
//	case int16:
//		return strconv.FormatInt(int64(값.(int16)), 10)
//	case int32:
//		return strconv.FormatInt(int64(값.(int32)), 10)
//	case int64:
//		return strconv.FormatInt(값.(int64), 10)
//	case float64:
//		return strconv.FormatFloat(값.(float64), 'f', -1, 64)
//	case big.Int:
//		정밀수 := 값.(big.Int); return (&정밀수).String()
//	case *big.Int:
//		return 값.(*big.Int).String()
//	case string:
//		return 값.(string)
//	case C문자열:
//		return 값.(C문자열).G값()

func F포맷된_문자열_생성(포맷_문자열 string, 추가_내용 ...interface{}) string {
	에러 := F에러_생성(포맷_문자열, 추가_내용...)

	return 에러.Error()
}

func F마지막_0_제거(문자열 string) string {
	if !strings.Contains(문자열, ".") {
		fmt.Println("정수인듯 함.", 문자열)
		return 문자열
	}

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

func F반올림(값 interface{}, 소숫점_이하_자릿수 int) C정밀수 {
	정밀수 := NV정밀수(값)
	
	if 정밀수 == nil { return nil }
	
	return 정밀수.S반올림(소숫점_이하_자릿수).G상수형()
}

func F문자열2실수(값 string) (float64, error) {
	실수, 에러 := strconv.ParseFloat(strings.Replace(값, ",", "", -1), 64)

	if 에러 != nil {
		return    0.0, 에러
	}

	return 실수, nil
}

func F문자열2시점(값 string) (time.Time, error) {
	시점, 에러 := time.Parse(P시점_포맷, 값)

	if 에러 == nil {
		return 시점, nil
	}

	시점, 에러 = time.Parse(P일자_포맷, 값)

	if 에러 == nil {
		return 시점, nil
	}

	return time.Time{}, 에러
}

func F시점2문자열(시점 time.Time) string {
	return 시점.Format(P시점_포맷)
}

func F일자2문자열(일자 time.Time) string {
	return 일자.Format(P일자_포맷)
}

func F시점_복사(값 time.Time) time.Time {
	복사본 := time.Date(값.Year(), 값.Month(), 값.Day(),
		값.Hour(), 값.Minute(), 값.Second(),
		값.Nanosecond(), 값.Location())

	return 복사본
}

func F임의_통화종류() P통화 {
	return P통화(int(rand.Int31n(int32(len(통화종류_문자열_모음)))))
}

func F통화종류별_정밀도(통화 P통화) int {
	switch 통화 {
	case KRW:
		return 0
	case USD, CNY, EUR:
		return 2
	default:
		return 2
	}
}

func F통화형식임(값 interface{}) bool {
	if _, ok := 값.(I통화); ok {
		return true
	}

	return false

	/*
		switch 값.(type) {
		case I통화:
			return true
		default:
			return false
		} */
}

func F통화_종류(값1, 값2 interface{}) (P통화, error) {
	if F_nil값_존재함(값1, 값2) {
		에러 := F에러_생성_출력("common.F통화_종류() : nil 입력값. %v, %v.", 값1, 값2)
		return P통화(0), 에러
	}

	통화형식임1 := F통화형식임(값1)
	통화형식임2 := F통화형식임(값2)

	if !통화형식임1 && !통화형식임2 {
		에러 := F에러_생성_출력("common.F통화_종류() : 두 입력값 모두 통화형식이 아님. "+
			"값1 %v %v, 값2 %v %v.",
			reflect.TypeOf(값1), 값1,
			reflect.TypeOf(값2), 값2)
		return P통화(0), 에러
	}

	if 통화형식임1 && 통화형식임2 &&
		값1.(I통화).G종류() != 값2.(I통화).G종류() {
		에러 := F에러_생성_출력("common.F통화_종류() : 통화 종류 불일치. %v, %v.", 값1, 값2)
		return P통화(0), 에러
	}

	if 통화형식임1 {
		return 값1.(I통화).G종류(), nil
	} else {
		return 값2.(I통화).G종류(), nil
	}
}

func F통화_복사(값 I통화) I통화 {
	if F_nil입력값_에러_출력(값) {
		return nil
	}

	switch 값.(type) {
	case C통화:
		// 상수형은 굳이 새로운 인스턴스를 만들 필요가 없음.
		return 값.(C통화)
	case V통화:
		return NV통화(값.G종류(), 값.G값())
	default:
		F문자열_출력("common.F통화_복사() : 예상치 못한 자료형.", reflect.TypeOf(값))
		return NC통화(값.G종류(), 값.G값())
	}
}

func F숫자_같음(값1, 값2 interface{}) bool {
	if 값1 == nil && 값2 == nil {
		return true
	}
	if 값1 == nil || 값2 == nil {
		return false
	}

	정밀수1 := NC정밀수(값1)
	정밀수2 := NC정밀수(값2)

	if 정밀수1 == nil || 정밀수2 == nil {
		//F문자열_출력("변환 에러. %v %v", 값1, 값2)
		return false
	}

	if 정밀수1.G비교(정밀수2) == 0 {
		return true
	}

	return false
}

func F값_같음(값1, 값2 interface{}) (값_같음 bool) {
	defer func() {
		if 에러 := recover(); 에러 != nil {
			F문자열_출력("common.F값_같음() : %v", 에러)
			값_같음 = false
		}
	}()

	switch {
	case 값1 == nil && 값2 == nil:
		return true
	case 값1 == nil, 값2 == nil:
		return false
	case F통화형식임(값1) && F통화형식임(값2):
		return 값1.(I통화).G같음(값2.(I통화))
	case F숫자_같음(값1, 값2):
		return true
	case reflect.DeepEqual(값1, 값2):
		return true
	}

	F_TODO("Value DeepEqual 기능 추가.")

	return false
}

// reflect 편의 함수

// 인터페이스 Type 구하는 법 : 타입 := reflect.TypeOf((*인터페이스)(nil)).Elem()
func F인터페이스를_구현함(값 interface{}, 인터페이스_형식 reflect.Type) bool {
	if 값 == nil {
		fmt.Println("F인터페이스를_구현함() : 값이 nil임.")

		return false
	}

	if 인터페이스_형식 == nil {
		fmt.Println("F인터페이스를_구현함() : 인터페이스_형식이 nil임.")

		return false
	}

	if 인터페이스_형식.Kind() != reflect.Interface {
		fmt.Println("F인터페이스를_구현함() : '인터페이스_형식'값이 인터페이스가 아님.")

		return false
	}

	return reflect.TypeOf(값).Implements(인터페이스_형식)
}

// 슬라이스 복사하는 편의 함수.
// builtin.copy()는 복사본 슬라이스 변수를 생성해 줘야 함.
func F슬라이스_복사(원본slice interface{}) interface{} {
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
}

// 테스트 편의 함수

var 테스트_모드 = false
func F테스트_모드() bool { return 테스트_모드 }
func F테스트_모드_시작() { 테스트_모드 = true }
func F테스트_모드_종료() { 테스트_모드 = false }

// '참거짓'이 false이면 Fail하는 테스트용 편의 함수.
func F참인지_확인(테스트 testing.TB, 참거짓 bool, 추가_매개변수 ...interface{}) (테스트_통과 bool) {

	if !참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			if 추가_매개변수 == nil || len(추가_매개변수) == 0 {
				fmt.Printf("%s주어진 조건이 false임.\n\n", F소스코드_위치(1))
			} else {
				switch 추가_매개변수[0].(type) {
				case string:
					포맷_문자열 := 추가_매개변수[0].(string)
					fmt.Printf("%s"+포맷_문자열+"\n\n",
						append([]interface{}{F소스코드_위치(1)},
							추가_매개변수[1:]...)...)
				default:
					포맷_문자열 := F소스코드_위치(1) + "%주어진 조건이 false임.\n\n"

					for 반복횟수 := 0; 반복횟수 < len(추가_매개변수); 반복횟수++ {
						포맷_문자열 = 포맷_문자열 + " %v"
					}
					포맷_문자열 = 포맷_문자열 + ".\n\n"

					fmt.Printf(포맷_문자열, 추가_매개변수...)
				}
			}
		}

		//테스트.FailNow()
		테스트.Fail()

		return false
	}

	return true
}

// '참거짓'이 true이면 Fail하는 테스트용 편의 함수.
func F거짓인지_확인(테스트 testing.TB, 참거짓 bool, 추가_매개변수 ...interface{}) (테스트_통과 bool) {
	if 참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			if 추가_매개변수 == nil || len(추가_매개변수) == 0 {
				fmt.Printf("%s주어진 조건이 true임.\n\n", F소스코드_위치(1))
			} else {
				switch 추가_매개변수[0].(type) {
				case string:
					포맷_문자열 := 추가_매개변수[0].(string)
					fmt.Printf("%s"+포맷_문자열+"\n\n",
						append([]interface{}{F소스코드_위치(1)},
							추가_매개변수[1:]...)...)
				default:
					포맷_문자열 := F소스코드_위치(1) + "%주어진 조건이 true임.\n\n"

					for 반복횟수 := 0; 반복횟수 < len(추가_매개변수); 반복횟수++ {
						포맷_문자열 = 포맷_문자열 + " %v"
					}
					포맷_문자열 = 포맷_문자열 + ".\n\n"

					fmt.Printf(포맷_문자열, 추가_매개변수...)
				}
			}
		}

		//테스트.FailNow()
		테스트.Fail()

		return false
	}

	return true
}

// 에러가 발생하면 Fail하는 테스트용 편의 함수.
func F에러없음_확인(테스트 testing.TB, 에러 error) (테스트_통과 bool) {
	if 에러 != nil {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			F에러_출력(에러)
		}

		//테스트.FailNow()
		테스트.Fail()

		return false
	}

	return true
}

// 에러가 발생하지 않으면 Fail하는 테스트용 편의 함수.
func F에러발생_확인(테스트 testing.TB, 에러 error) (테스트_통과 bool) {
	if 에러 == nil {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			F에러_출력(에러)
		}

		//테스트.FailNow()
		테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 다르면 Fail하는 테스트용 편의 함수.
func F같은값_확인(테스트 testing.TB, 값1, 값2 interface{}) (테스트_통과 bool) {
	if !F값_같음(값1, 값2) { //&&
		//if !reflect.DeepEqual(값1, 값2) {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s서로 다름. 값1: %v %v 값2: %v %v.\n\n",
				F소스코드_위치(1),
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}

		//테스트.FailNow()
		테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 같으면 Fail하는 테스트용 편의 함수.
func F다른값_확인(테스트 testing.TB, 값1, 값2 interface{}) (테스트_통과 bool) {
	if F값_같음(값1, 값2) { //||
		//if reflect.DeepEqual(값1, 값2) {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s서로 같음. 값1: %v %v 값2: %v %v.\n\n",
				F소스코드_위치(1),
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}

		//테스트.FailNow()
		테스트.Fail()

		return false
	}

	return true
}

// nil값이 아니면 Fail하는 테스트용 편의 함수.
/*
func F_nil_확인(테스트 testing.TB, 값 interface{}) (테스트_통과 bool) {

	/*
	fmt.Printf("%s 값: %v %v.\n\n", F소스코드_위치(1), reflect.TypeOf(값), 값)

	if 값 == nil {
		return true
	}

	switch 테스트.(type) {
	case I테스트용_가상_객체:
			// PASS
	default:
		fmt.Printf("%snil 아님. 값: %v %#v.\n\n", F소스코드_위치(1), reflect.TypeOf(값), 값)
	}

	//테스트.FailNow()
	테스트.Fail()

	return false
}
*/

// 에러 처리 편의 함수.

func F_nil값_존재함(검사대상_모음 ...interface{}) bool {
	for _, 검사대상 := range 검사대상_모음 {
		if 검사대상 == nil {
			return true
		}
	}

	return false
}

func F_nil입력값_에러_출력(검사대상_모음 ...interface{}) bool {
	if F_nil값_존재함(검사대상_모음...) {
		출력내용 := make([]interface{}, 0)
		출력내용 = append(출력내용, F소스코드_위치(2)+"snil 입력값 에러.")
		출력내용 = append(출력내용, 검사대상_모음...)

		fmt.Println(출력내용)
		return true
	}

	return false
}

func F에러_생성(문자열 string, 추가_내용 ...interface{}) error {
	for strings.HasSuffix(문자열, "\n") {
		문자열 += "\n"
	}

	return fmt.Errorf(문자열, 추가_내용...)
}

func F에러_생성_출력(문자열 string, 추가_내용 ...interface{}) error {
	fmt.Println("")
	fmt.Printf("%s: " + 문자열 + "\n", F소스코드_위치(1), 추가_내용)
	fmt.Println(F소스코드_위치(2))
	fmt.Println(F소스코드_위치(3))

	return F에러_생성(문자열, 추가_내용...)
}

func F에러_출력(에러 error) {
	fmt.Println("")
	fmt.Println(F소스코드_위치(1) + ": " + 에러.Error())
	fmt.Println(F소스코드_위치(2))
	fmt.Println(F소스코드_위치(3))
}

func F문자열_출력(문자열 string, 추가_내용 ...interface{}) {
	fmt.Println("")
	fmt.Printf("%s: " + 문자열 + "\n", append([]interface{}{F소스코드_위치(1)}, 추가_내용...)...)
	fmt.Println(F소스코드_위치(2))
	fmt.Println(F소스코드_위치(3))
	fmt.Println(F소스코드_위치(4))
	fmt.Println(F소스코드_위치(5))
	fmt.Println(F소스코드_위치(6))	
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

// 디버깅 편의 함수.
func F체크포인트(체크포인트_번호 *int, 포맷_문자열 string, 기타 ...interface{}) {
	fmt.Printf("%s체크포인트 %v "+포맷_문자열+"\n\n", append([]interface{}{F소스코드_위치(1), *체크포인트_번호}, 기타...)...)
	(*체크포인트_번호)++
}

func F값_확인(값 ...interface{}) {
	fmt.Println("")
	fmt.Println(append([]interface{}{F소스코드_위치(1), "값_확인 :"}, 값...)...)
}

func F타입(값 interface{}) string {
	return reflect.TypeOf(값).String()
}

// 메모 편의 함수.
var 이미_출력한_TODO_모음 []string = make([]string, 0)

func F_TODO(문자열 string) {
	for _, 이미_출력한_TODO := range 이미_출력한_TODO_모음 {
		if 문자열 == 이미_출력한_TODO {
			return
		}
	}

	이미_출력한_TODO_모음 = append(이미_출력한_TODO_모음, 문자열)

	fmt.Printf("TODO : %s %s\n\n", F소스코드_위치(1), 문자열)
}
