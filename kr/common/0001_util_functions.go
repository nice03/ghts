package common

import (
	"fmt"
	"math/big"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	큰정수_제로 = big.NewInt(0)
	정밀수_제로 = big.NewRat(0, 1)
	
	차이_한도, _ = new(big.Rat).SetString("1/100000000000000000000000000000000")
}

var 큰정수_제로 *big.Int
var 정밀수_제로 *big.Rat






/*
func F공유해도_안전함(값 interface{}) (안전함 bool) {
func F고정값임(값 interface{}) (고정값임 bool) {
	defer func() {
		if r := recover(); r != nil {
			안전함 = false
		}
	}()

	if 값 == nil { return false }

	switch 값.(type) {
	case uint, uint8, uint16, uint32, uint64,
			int, int8, int16, int32, int64,
			float32, float64, bool, string, time.Time:
			return true
	case I변수형, big.Int, *big.Int, big.Rat, *big.Rat:
			return false
	case I상수형:
			// Pass
	default:
			// 'I상수형'을 구현하지 않는 모든 자료형은 안전하지 않은 것으로 판단함.
			return false
	}

	// 'I상수형'을 구현하는 자료형도 실제로 값을 변경하는 메소드가 없는 지 확인해야 함.

	형식 := reflect.TypeOf(값)
	종류 := reflect.TypeOf(값).Kind()
	인터페이스_형식임 := false

	if 종류 == reflect.Interface {
		인터페이스_형식임 = true
	}

	// Interface가  아닌 구조체의 경우, 메소드를 들여다보고 판별함.
	메소드_갯수 := 형식.NumMethod()

	for 인덱스 := 0; 인덱스 < 메소드_갯수; 인덱스++ {
		메소드_이름 := 형식.Method(인덱스).Name
		메소드_입력값_갯수 := 형식.Method(인덱스).Type.NumIn()
		메소드_반환값_갯수 := 형식.Method(인덱스).Type.NumOut()

		if !인터페이스_형식임 {
			// 구조체에 연결된 메소드의 경우,
			// 소스코드 상에 입력 파라메터가 없어도,
			// 내부적으로는 첫 번째 인자로 구조체의 포인터를 가진다.
			// 그래서, 그것을 고려에서 제외한다.
			// 인터페이스의 경우에는 그렇지 않음.
			메소드_입력값_갯수--
		}

		// S로 시작하는 문자열만 값을 변경하는 메소드라고 가정한다.
		// 이 가정이 틀릴 수도 있는 데,
		// 실제로 내부값이 변경되는 지 어떻게 확인할 수 있는 지 모르겠음.
		// reflect로 복사본을 생성한 후 모든 메소드를 실행시켜봐야 하나?
		// 1. 복사본 자동 생성
		// 2. 메소드를 실행.
		// 3. 내부 멤버필드 값이 변했는 지 확인.
		// 가능할까? 1, 2, 3을 하나씩 쪼개서 함수로 만든 후 테스트를 하면서 가능한 지 확인해 봐야 겠다.
		switch {
		case !strings.HasPrefix(메소드_이름, "S"):
			continue
		case 메소드_이름 == "String" &&
				메소드_입력값_갯수 == 0 &&
				메소드_반환값_갯수 == 1:
			메소드_반환값_형식 := 형식.Method(인덱스).Type.Out(0)

			if 메소드_반환값_형식.Kind() == reflect.String &&
				메소드_반환값_형식.String() == "string" {
				// 입력 매개변수가 없고, 반환값이 string형식인 String() 함수는
				// I기본_문자열 인터페이스를 구현하는 메소드이므로,
				// 문제 삼지 않음.
				continue
			}
		default:
			// 이외에 'S'로 시작하는 모든 메소드는 값을 변경하는 메소드로 간주
			return false
		}
	}

	// I상수형을 구현하고,
	// 값을 변경한다고 추정되는 메소드가 없으므로 안전하다고 판정한다.
	return true
} */
func F부호없는_정수2큰정수(값 uint64) *big.Int { return new(big.Int).SetUint64(값) }
func F부호없는_정수2정밀수(값 uint64) *big.Rat {
	return new(big.Rat).SetInt(F부호없는_정수2큰정수(값))
}
func F부호없는_정수2문자열(값 uint64) string    { return strconv.FormatUint(값, 10) }

func F정수2큰정수(값 int64) *big.Int { return big.NewInt(값) }
func F정수2정밀수(값 int64) *big.Rat { return new(big.Rat).SetInt64(값) }
func F정수2문자열(값 int64) string   { return strconv.FormatInt(값, 10) }
func F정수2월(값 int) (time.Month, error) {
	switch 값 {
	case 1:
		return time.January, nil
	case 2:
		return time.February, nil
	case 3:
		return time.March, nil
	case 4:
		return time.April, nil
	case 5:
		return time.May, nil
	case 6:
		return time.June, nil
	case 7:
		return time.July, nil
	case 8:
		return time.August, nil
	case 9:
		return time.September, nil
	case 10:
		return time.October, nil
	case 11:
		return time.November, nil
	case 12:
		return time.December, nil
	default:
		에러 := fmt.Errorf("%scommon.F정수2월() : 예상치 못한 월. 입력값 %v.", F소스코드_위치(2), 값)
		return time.January, 에러
	}
}

func F실수2정밀수(값 float64) *big.Rat {
	값_문자열 := strconv.FormatFloat(값, 'f', -1, 64)
	정밀값, 변환_성공 := new(big.Rat).SetString(값_문자열)

	if !변환_성공 {
		정밀값 = new(big.Rat).SetFloat64(값)
	}

	return 정밀값
}
func F실수2문자열(숫자 float64) string { 
	return strconv.FormatFloat(숫자, 'f', -1, 64)
}

func F큰정수(값 interface{}) (*big.Int, error) {
	if F_nil_존재함(값) {
		에러 := fmt.Errorf("%sF큰정수() 변환 에러. 입력값 %v.\n", F소스코드_위치(2), 값)
		
		return nil, 에러
	}
	
	switch 값.(type) {
	case uint:
		return F부호없는_정수2큰정수(uint64(값.(uint))), nil
	case uint8:
		return F부호없는_정수2큰정수(uint64(값.(uint8))), nil
	case uint16:
		return F부호없는_정수2큰정수(uint64(값.(uint16))), nil
	case uint32:
		return F부호없는_정수2큰정수(uint64(값.(uint32))), nil
	case uint64:
		return F부호없는_정수2큰정수(값.(uint64)), nil
	case int:
		return F정수2큰정수(int64(값.(int))), nil
	case int8:
		return F정수2큰정수(int64(값.(int8))), nil
	case int16:
		return F정수2큰정수(int64(값.(int16))), nil
	case int32:
		return F정수2큰정수(int64(값.(int32))), nil
	case int64:
		return F정수2큰정수(값.(int64)), nil
	case big.Int:
		큰정수 := 값.(big.Int)
		return F큰정수_복사(&큰정수), nil
	case *big.Int:
		return F큰정수_복사(값.(*big.Int)), nil
	case I정수형:
		return 값.(I정수형).G큰정수(), nil
	case string:
		return F문자열2큰정수(값.(string))
	case C문자열:
		return F문자열2큰정수(값.(C문자열).G값())
	default:
		if !strings.Contains(F소스코드_위치(2), "_test.go") {
			fmt.Printf("%sF큰정수() 변환 에러. 입력값 %v.\n", F소스코드_위치(2), 값)
		}
		에러 := fmt.Errorf("%sF큰정수() 변환 에러. 입력값 %v.\n", F소스코드_위치(2), 값)

		return nil, 에러
	}
}

func F큰정수_복사(값 *big.Int) *big.Int {
	return new(big.Int).Set(값)
}

func F큰정수2실수(값 *big.Int) float64 {
	실수, 에러 := F문자열2실수(값.String())
	
	if 에러 != nil {
		fmt.Printf("common.F큰정수2실수() : 변환 에러 발생. %v.", 값)
	}
	
	return 실수		
}

func F큰정수2정밀수(값 *big.Int) *big.Rat {
	return new(big.Rat).SetInt(값)
}

func F큰정수_절대값(값 *big.Int) *big.Int {
	return new(big.Int).Abs(값)
}

func F큰정수_더하기(값1, 값2 *big.Int) *big.Int {
	return new(big.Int).Add(값1, 값2)
}

func F큰정수_빼기(값1, 값2 *big.Int) *big.Int {
	return new(big.Int).Sub(값1, 값2)
}

func F큰정수_곱하기(값1, 값2 *big.Int) *big.Int {
	return new(big.Int).Mul(값1, 값2)
}

func F큰정수_나누기(분자, 분모 *big.Int) (*big.Int, error) {
	if 분모.Cmp(큰정수_제로) == 0 {
		에러 := fmt.Errorf("%scommon.F큰정수_나누기() : 분모가 0입니다.", F소스코드_위치(2))
		return nil, 에러
	}
	
	return new(big.Int).Div(분자, 분모), nil
}

func F큰정수_반대부호값(값 *big.Int) *big.Int {
	return new(big.Int).Neg(값)
}

func F정밀수(값 interface{}) (*big.Rat, error) {
	if F_nil_존재함(값) {
		에러 := fmt.Errorf("%sF정밀수() 변환 에러. 입력값 %v.\n", F소스코드_위치(2), 값)
		
		return nil, 에러
	}
	
	switch 값.(type) {
	case uint:
		return F부호없는_정수2정밀수(uint64(값.(uint))), nil
	case uint8:
		return F부호없는_정수2정밀수(uint64(값.(uint8))), nil
	case uint16:
		return F부호없는_정수2정밀수(uint64(값.(uint16))), nil
	case uint32:
		return F부호없는_정수2정밀수(uint64(값.(uint32))), nil
	case uint64:
		return F부호없는_정수2정밀수(값.(uint64)), nil
	case int:
		return F정수2정밀수(int64(값.(int))), nil
	case int8:
		return F정수2정밀수(int64(값.(int8))), nil
	case int16:
		return F정수2정밀수(int64(값.(int16))), nil
	case int32:
		return F정수2정밀수(int64(값.(int32))), nil
	case int64:
		return F정수2정밀수(값.(int64)), nil
	case float32:
		문자열 := strconv.FormatFloat(float64(값.(float32)), 'f', -1, 32)
		정밀수, _ := F문자열2정밀수(문자열)
		return 정밀수, nil
	case float64:
		return F실수2정밀수(값.(float64)), nil
	case big.Int:
		큰정수 := 값.(big.Int)
		return F큰정수2정밀수(&큰정수), nil
	case *big.Int:
		return F큰정수2정밀수(값.(*big.Int)), nil
	case big.Rat:
		정밀값 := 값.(big.Rat)
		return &정밀값, nil
	case *big.Rat:
		return 값.(*big.Rat), nil
	case I실수형:
		return 값.(I실수형).G정밀수(), nil
	case I통화:
		return 값.(I통화).G값(), nil
	case string:
		return F문자열2정밀수(값.(string))
	case C문자열:
		return F문자열2정밀수(값.(C문자열).G값())
	default:
		fmt.Printf("%sF정밀수() 변환 에러. 입력값 %v.\n", F소스코드_위치(2), 값)
		에러 := fmt.Errorf("%sF정밀수() 변환 에러. 입력값 %v.\n", F소스코드_위치(2), 값)

		return nil, 에러
	}
}

var 차이_한도 *big.Rat

func F정밀수_일치(값1, 값2 *big.Rat) bool {
	if F_nil_존재함(값1, 값2) { return false }
	
	if 값1.Cmp(값2) == 0 { return true }

	차이_절대값 := new(big.Rat)
	차이_절대값.Sub(값1, 값2)
	차이_절대값.Abs(차이_절대값)

	차이_한도, _ = 차이_한도.SetString("1/100000000000000000000000000000000")

	if 차이_절대값.Cmp(차이_한도) == -1 {
		// 차이값이 극히 미세하므로 같다고 봐도 무방하다.
		return true
	}

	return false
}

func F정밀수_복사(값 *big.Rat) *big.Rat {
	return new(big.Rat).Set(값)
}

func F정밀수2실수(값 *big.Rat) float64 {
	// 소숫점 이하 아주 미세한 에러를 줄이기 위해서 문자열로 바꾼 후 변환.
	실수값, 에러 := F문자열2실수(F정밀수2문자열(값))		
	
	if 에러 != nil {
		실수값, _ = 값.Float64()
	}	
	
	return 실수값
}

func F정밀수2문자열(값 *big.Rat) string {
	//실수값, _ := 값.Float64(); return F실수2문자열()	// 예전 방식

	문자열 := F정밀수_반올림_문자열(값, 30)

	if strings.Contains(문자열, ".") {
		for strings.HasSuffix(문자열, "0") {
			문자열 = strings.TrimSuffix(문자열, "0")
		}

		if strings.HasSuffix(문자열, ".") {
			문자열 = 문자열 + "0"
		}
	}

	return 문자열
}

func F정밀수_반올림_문자열(값 *big.Rat, 소숫점_이하_자릿수 int) string {
	return 값.FloatString(소숫점_이하_자릿수)
}

func F정밀수_절대값(값 *big.Rat) *big.Rat {
	return new(big.Rat).Abs(값)
}

func F정밀수_더하기(값1, 값2 *big.Rat) *big.Rat {
	return new(big.Rat).Add(값1, 값2)
}

func F정밀수_빼기(값1, 값2 *big.Rat) *big.Rat {
	return new(big.Rat).Sub(값1, 값2)
}

func F정밀수_곱하기(값1, 값2 *big.Rat) *big.Rat {
	return new(big.Rat).Mul(값1, 값2)
}

func F정밀수_나누기(분자, 분모 *big.Rat) (*big.Rat, error) {
	if 분모.Cmp(정밀수_제로) == 0 {
		에러 := fmt.Errorf("%scommon.F정밀수_나누기() : 분모가 0입니다.", F소스코드_위치(2))
		
		return nil, 에러
	}
	
	return new(big.Rat).Quo(분자, 분모), nil
}

func F정밀수_역수(값 *big.Rat) (*big.Rat, error) {
	if 값.Cmp(정밀수_제로) == 0 {
		에러 := fmt.Errorf("%scommon.F정밀수_역수() : 0의 역수는 무한대 입니다.", F소스코드_위치(2))
		
		return nil, 에러
	}
		
	return new(big.Rat).Inv(값), nil
}

func F정밀수_반대부호값(값 *big.Rat) *big.Rat {
	return new(big.Rat).Neg(값)
}

func F참거짓2문자열(값 bool) string { return strconv.FormatBool(값) }

func F문자열2실수(문자열 string) (float64, error) {
	실수, 에러 := strconv.ParseFloat(strings.Replace(문자열, ",", "", -1), 64)
	if 에러 != nil {
		return 0.0, 에러
	}

	return 실수, nil
}

func F문자열2큰정수(문자열 string) (*big.Int, error) {
	큰정수, 성공 := new(big.Int).SetString(문자열, 10)
	
	if !성공 {
		에러 := fmt.Errorf("%scommon.F문자열2큰정수() : 변환 에러 발생. 문자열 %v.", 
							F소스코드_위치(2), 문자열)
							
		return nil, 에러
	}
	
	return 큰정수, nil
}

func F문자열2정밀수(문자열 string) (*big.Rat, error) {
	정밀수, 성공 := new(big.Rat).SetString(문자열)
	
	if !성공 {
		에러 := fmt.Errorf("%scommon.F문자열2정밀수() : 변환 에러 발생. 문자열 %v.", 
							F소스코드_위치(2), 문자열)
							
		return nil, 에러
	}
	
	return 정밀수, nil
}

func F문자열2일자(일자_문자열 string) (time.Time, error) {
	일자, 에러 := time.Parse("2006-01-02", 일자_문자열)
	if 에러 != nil {
		제로값 := reflect.Zero(reflect.TypeOf(time.Now())).Interface().(time.Time)
		return 제로값, 에러
	}

	return 일자, nil
}

func F일자2문자열(일자 time.Time) string {
	return 일자.Format("2006-01-02")
}

func F시점_복사(값 time.Time) time.Time {
	복사본 := time.Date(값.Year(), 값.Month(), 값.Day(), 
						값.Hour(), 값.Minute(), 값.Second(),
						값.Nanosecond(), 값.Location())
	
	return 복사본
}

var 임시 bool = false

func F값_일치(값1, 값2 interface{}) (값_일치 bool) {
	defer func() {
		에러 := recover()
		
		if 에러 != nil {
			fmt.Println(에러)
		
			값_일치 = false
		}
	}()
	
	if 값1 == nil && 값2 == nil { return true }
	if F_nil_존재함(값1, 값2) { return false }
	
	//i := 0	// 디버깅용 체크포인트 인덱스
	
	// 숫자는 형식이 달라도 비교할 수 있음.
	// 그 외 알려진 형식에 대해서 비교. (현재는 I통화 만 있음)
	
	if !임시 {
		fmt.Printf("TODO : %s F값_일치()\n\n", F소스코드_위치(1))
		임시 = true
	}
	
	switch 값1.(type) {
	/*
	case I통화:
		// 형변환 에러나면 defer문의 recover()에서 false 반환.
		통화1 := 값1.(I통화)
		통화2 := 값2.(I통화)
		
		if 통화1.G종류() == 통화2.G종류() &&
			F정밀수_일치(통화1.G금액(), 통화2.G금액()) {
			return true
		} else {
			return false
		} */
	case I실수형, I정수형,
		uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64,
		float32, float64,
		big.Int, *big.Int,
		big.Rat, *big.Rat:

		정밀값1, 에러1 := F정밀수(값1)
		정밀값2, 에러2 := F정밀수(값2)

		switch {
		case 에러1 != nil, 에러2 != nil:
			값_일치 = false	//; F체크포인트(&i, "F수치형_값_일치() : nil 값.")
		case F정밀수_일치(정밀값1, 정밀값2):
			값_일치 = true	//; F체크포인트(&i, "F수치형_값_일치() : F정밀수_같음().")
		default:
			값_일치 = false	//; F체크포인트(&i, "F수치형_값_일치() : default")
		}

		//F체크포인트(&i,
		//	"F값_일치() : 값_일치 '%v', F정밀수_같음() A. 값1 %v, 값2 %v.",
		//	값_일치, 값1, 값2)

		return 값_일치
	}
	
	// F값_일치() 개선 계획을 구현할 것.
	// 두 값의 내부값 기록본을 서로 비교.
	// 참조형 필드가 가리키는 개체가 달라고 값만 같으면 같은 것으로 판정할 수 있게 될 것임.

	// 마지막으로 reflect.DeepEqual()로 비교.
	값_일치 = reflect.DeepEqual(값1, 값2)

	//F체크포인트(&i, "F값_일치() : 값_일치 '%v'. reflect.DeepEqual(). 값1 %v, 값2 %v.", 값_일치, 값1, 값2)

	return 값_일치
}

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
// 슬라이스 복사하는 builtin.copy()가 존재하지만, 복사본 슬라이스 변수를 생성하는 게 귀찮음.
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

// 디버깅 할 때 각 체크포인트의 위치와 번호를 출력해주는 편의 함수.
func F체크포인트(체크포인트_번호 *int, 포맷_문자열 string, 기타 ...interface{}) {
	fmt.Printf("%s체크포인트 %v "+포맷_문자열+"\n\n", append([]interface{}{F소스코드_위치(2), *체크포인트_번호}, 기타...)...)
	(*체크포인트_번호)++
}

// 소스코드 위치를 나타내는 함수. runtime.Caller()의 한글화 버전임.
// '건너뛰는_단계'값이 커질수록 호출 경로를 거슬러 올라감.
//
// 0 = F소스코드_위치() 자기자신의 위치.
// 1 = F소스코드_위치()를 호출한 메소드의 위치.
// 2 = F소스코드_위치()를 호출한 메소드를 호출한 메소드의 위치
// 3, 4, 5,....n = 계속 거슬러 올라감.
//
// 다른 모듈을 위해서 사용되는 라이브러리 펑션인 경우 2가 적당함.
// 그렇지 않다면, 1이 적당함.
func F소스코드_위치(건너뛰는_단계 int) string {
	_, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)

	파일명 := filepath.Base(파일_경로)
	return 파일명 + ":" + F정수2문자열(int64(행_번호)) + ": "
}

// '참거짓'이 false이면 Fail하는 테스트용 편의 함수.
func F참인지_확인(테스트 testing.TB, 참거짓 bool, 포맷_문자열 string,
	추가_매개변수 ...interface{}) (테스트_통과 bool) {

	if !참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s"+포맷_문자열+"\n\n",
				append([]interface{}{F소스코드_위치(2)}, 추가_매개변수...)...)
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// '참거짓'이 true이면 Fail하는 테스트용 편의 함수.
func F거짓인지_확인(테스트 testing.TB, 참거짓 bool, 포맷_문자열 string,
	추가_매개변수 ...interface{}) (테스트_통과 bool) {
	if 참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s"+포맷_문자열+"\n\n",
				append([]interface{}{F소스코드_위치(2)}, 추가_매개변수...)...)
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
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s에러 발생. : %s \n\n", F소스코드_위치(2), 에러.Error())
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
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s에러 발생. : %s \n\n", F소스코드_위치(2), 에러.Error())
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 다르면 Fail하는 테스트용 편의 함수.
func F같은값_확인(테스트 testing.TB, 값1, 값2 interface{}) (테스트_통과 bool) {
	if !F값_일치(값1, 값2) {//&&
	//if !reflect.DeepEqual(값1, 값2) {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s값 불일치. 값1: %v %v 값2: %v %v.\n\n", 
						F소스코드_위치(2), 
						reflect.TypeOf(값1), 값1, 
						reflect.TypeOf(값2), 값2)
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// 기대값과 실제값이 같으면 Fail하는 테스트용 편의 함수.
func F다른값_확인(테스트 testing.TB, 값1, 값2 interface{}) (테스트_통과 bool) {
	//if F값_일치(기대값, 실제값) ||
	if reflect.DeepEqual(값1, 값2) {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			fmt.Printf("%s값 일치. 값1: %#v 값2: %#v.\n\n", F소스코드_위치(2), 값1, 값2)
		}

		테스트.FailNow()
		//테스트.Fail()

		return false
	}

	return true
}

// nil값이 아니면 Fail하는 테스트용 편의 함수.
func F_nil_확인(테스트 testing.TB, 값 interface{}) (테스트_통과 bool) {

	return F참인지_확인(테스트, 값 == nil, "")

	/*
	fmt.Printf("%s 값: %v %v.\n\n", F소스코드_위치(2), reflect.TypeOf(값), 값)
	
	if 값 == nil {
		return true
	}
	
	switch 테스트.(type) {
	case I테스트용_가상_객체:
			// PASS
	default:
		fmt.Printf("%snil 아님. 값: %v %#v.\n\n", F소스코드_위치(2), reflect.TypeOf(값), 값)
	}

	테스트.FailNow()
	//테스트.Fail()

	return false */
}

func F_nil_존재함(검사대상들 ...interface{}) bool {
	for _, 검사대상 := range 검사대상들 {
		if 검사대상 == nil {
			return true
		}
	}
	
	return false
}