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

// Exponential Back-off
// 실패한 후 반복하는 횟수에 따라 기다리는 최대시간이 기하급수적으로 커짐.
func F잠시_대기(반복횟수 int) {
	const len_대기시간_한도 = len(대기시간_한도) - 1

	if 반복횟수 > len_대기시간_한도 {
		time.Sleep(time.Duration(rand.Int63n(대기시간_한도[len_대기시간_한도])))
	} else {
		time.Sleep(time.Duration(rand.Int63n(대기시간_한도[반복횟수])))
	}
}

func F안전한_매개변수(값_모음 ...interface{}) bool {
	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64,
			float32, float64, bool, string, time.Time:
			// CallByValue에 의해서 자동으로 복사본이 생성되는 형식.
			//OK to PASS
		case *sC부호없는_정수64, *sC정수64, *sC실수64, *sC시점, *sC정밀수, *sC통화:
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

	// 지금 알려진 것만 우선 포함 시킴.
	switch 값.(type) {
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
	case *sV시점:
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
	case *big.Rat:
		return F마지막_0_제거(값.(*big.Rat).FloatString(100))
	case float64:
		return strconv.FormatFloat(값.(float64), 'f', -1, 64)
	case time.Time:
		return 값.(time.Time).Format(P시점_포맷)
	case I기본_문자열:
		return 값.(I기본_문자열).String()
	}

	return F포맷된_문자열("%v", 값)
}

func F포맷된_문자열(포맷_문자열 string, 추가_내용 ...interface{}) string {
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

func F반올림(값 interface{}, 소숫점_이하_자릿수 int) C정밀수 {
	정밀수 := NV정밀수(값)

	if 정밀수 == nil {
		return nil
	}

	return 정밀수.S반올림(소숫점_이하_자릿수).G상수형()
}

func F문자열2실수(값 string) (float64, error) {
	실수, 에러 := strconv.ParseFloat(strings.Replace(값, ",", "", -1), 64)

	if 에러 != nil {
		return 0.0, 에러
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

func F시점_문자열(시점 time.Time) string {
	return 시점.Format(P시점_포맷)
}

func F일자_문자열(일자 time.Time) string {
	return 일자.Format(P일자_포맷)
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

func F통화종류(값_모음 ...interface{}) P통화종류 {
	통화종류_맵 := make(map[P통화종류]P통화종류)

	for _, 값 := range 값_모음 {
		if !F통화형식임(값) {
			continue
		}

		통화종류 := 값.(I통화).G종류()

		// 중복없도록 하기 위해서 맵을 사용함.
		// 대입하는 값은 아무 의미 없음.
		통화종류_맵[통화종류] = 통화종류
	}

	if len(통화종류_맵) == 1 {
		for 통화종류, _ := range 통화종류_맵 {
			return 통화종류
		}
	}

	return INVALID_CURRENCY_TYPE
}

func F통화형식임(값_모음 ...interface{}) bool {
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

func F숫자형식임(값_모음 ...interface{}) bool {
	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64,
			float32, float64, *big.Int, *big.Rat,
			*sC정수64, *sV정수64, *sC부호없는_정수64, *sV부호없는_정수64,
			*sC실수64, *sV실수64, *sC정밀수, *sV정밀수:
			continue
		default:
			_, 성공 := new(big.Rat).SetString(F문자열(값))

			if !성공 {
				return false
			}
		}
	}

	return true
}

func F숫자_같음(값1, 값2 interface{}) bool {
	정밀수1 := NC정밀수(값1)
	정밀수2 := NC정밀수(값2)

	if 정밀수1 == nil || 정밀수2 == nil {
		return false
	}

	return 정밀수1.G같음(정밀수2)
}

func F참거짓형식임(값_모음 ...interface{}) bool {
	for _, 값 := range 값_모음 {
		switch 값.(type) {
		case bool, *sC참거짓:
			continue
		default:
			return false
		}
	}

	return true
}

func F참거짓_같음(값1, 값2 interface{}) bool {
	var 참거짓1, 참거짓2 bool

	switch 값1.(type) {
	case bool:
		참거짓1 = 값1.(bool)
	case *sC참거짓:
		참거짓1 = 값1.(*sC참거짓).G값()
	default:
		return false
	}

	switch 값2.(type) {
	case bool:
		참거짓2 = 값2.(bool)
	case *sC참거짓:
		참거짓2 = 값2.(*sC참거짓).G값()
	default:
		return false
	}

	return 참거짓1 == 참거짓2
}

func F문자열형식임(값_모음 ...interface{}) bool {
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

func F문자열_같음(값1, 값2 interface{}) bool {
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

func F시점형식임(값_모음 ...interface{}) bool {
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

func F시점_같음(값1, 값2 interface{}) bool {
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

func F값_같음(값1, 값2 interface{}) (값_같음 bool) {
	defer func() {
		if 에러 := recover(); 에러 != nil {
			//F문자열_출력("%v", 에러)
			값_같음 = false
		}
	}()

	switch {
	case 값1 == nil, 값2 == nil:
		if 값1 == nil && 값2 == nil {
			return true
		} else if F포맷된_문자열("%v", 값1) == "<nil>" &&
			F포맷된_문자열("%v", 값2) == "<nil>" {
			return true
		}
		return false
	case F통화형식임(값1, 값2):
		return 값1.(I통화).G같음(값2.(I통화))
	case F숫자형식임(값1, 값2):
		return F숫자_같음(값1, 값2)
	case F참거짓형식임(값1, 값2):
		return F참거짓_같음(값1, 값2)
	case F시점형식임(값1, 값2):
		return F시점_같음(값1, 값2)
	case F문자열형식임(값1, 값2):
		return F문자열_같음(값1, 값2)
	case reflect.DeepEqual(값1, 값2):
		return true
	}

	return false
}

// reflect 편의 함수

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

func F테스트_모드() bool { return 테스트_모드.G값() }
func F테스트_모드_시작()   { 테스트_모드.S값(true) }
func F테스트_모드_종료()   { 테스트_모드.S값(false) }

func F문자열_출력_일시정지_모드() bool { return 문자열_출력_일시정시.G값() }
func F문자열_출력_일시정지_시작()      { 문자열_출력_일시정시.S값(true) }
func F문자열_출력_일시정지_종료()      { 문자열_출력_일시정시.S값(false) }

// '참거짓'이 false이면 Fail하는 테스트용 편의 함수.
func F참인지_확인(테스트 testing.TB, 참거짓 bool, 추가_매개변수 ...interface{}) (테스트_통과 bool) {

	if !참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
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

		테스트.Fail()

		return false
	}

	return true
}

// 에러가 발생하면 Fail하는 테스트용 편의 함수.
func F에러없음_확인(테스트 testing.TB, 에러 error) (테스트_통과 bool) {
	if 에러 != nil {
		F문자열_출력_일시정지_종료()

		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력2(1, 에러.Error())
		}

		테스트.Fail()

		return false
	}

	return true
}

// 에러가 발생하지 않으면 Fail하는 테스트용 편의 함수.
func F에러발생_확인(테스트 testing.TB, 에러 error) (테스트_통과 bool) {
	if 에러 == nil {
		F문자열_출력_일시정지_종료()

		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력_일시정지_종료()
			F문자열_출력2(1, 에러.Error())
		}

		테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 다르면 Fail하는 테스트용 편의 함수.
func F같은값_확인(테스트 testing.TB, 값1, 값2 interface{}) (테스트_통과 bool) {
	if !F값_같음(값1, 값2) {
		F문자열_출력_일시정지_종료()

		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력2(1, "서로 다름. 값1: %v %v 값2: %v %v.",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}

		테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 같으면 Fail하는 테스트용 편의 함수.
func F다른값_확인(테스트 testing.TB, 값1, 값2 interface{}) (테스트_통과 bool) {
	if F값_같음(값1, 값2) {
		F문자열_출력_일시정지_종료()

		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			F문자열_출력2(1, "서로 같음. 값1: %v %v 값2: %v %v.",
				reflect.TypeOf(값1), 값1,
				reflect.TypeOf(값2), 값2)
		}

		테스트.Fail()

		return false
	}

	return true
}

// 에러 처리 편의 함수.
func F에러_생성(문자열 string, 추가_내용 ...interface{}) error {
	for strings.HasSuffix(문자열, "\n") {
		문자열 += "\n"
	}

	return fmt.Errorf(문자열, 추가_내용...)
}

func F문자열_출력(문자열 string, 추가_내용 ...interface{}) {
	F문자열_출력2(1, 문자열, 추가_내용...)
}

func F문자열_출력2(추가적인_건너뛰기_단계 int, 문자열 string, 추가_내용 ...interface{}) {
	if F문자열_출력_일시정지_모드() {
		return
	}

	i := 추가적인_건너뛰기_단계

	fmt.Println("")
	fmt.Printf("%s: "+문자열+"\n", append([]interface{}{F소스코드_위치(1 + i)}, 추가_내용...)...)
	fmt.Println(F소스코드_위치(2 + i))
	fmt.Println(F소스코드_위치(3 + i))
	fmt.Println(F소스코드_위치(4 + i))
	fmt.Println(F소스코드_위치(5 + i))
	fmt.Println(F소스코드_위치(6 + i))
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
func F체크포인트(체크포인트_번호 *int, 추가_매개변수 ...interface{}) {
	문자열 := F포맷된_문자열("%s체크포인트 %v ", F소스코드_위치(1), *체크포인트_번호)
	fmt.Println(append([]interface{}{문자열}, 추가_매개변수...)...)
	(*체크포인트_번호)++
}

func F값_확인(값 ...interface{}) {
	fmt.Println("")
	fmt.Println(append([]interface{}{F소스코드_위치(1), "값_확인 :"}, 값...)...)
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
