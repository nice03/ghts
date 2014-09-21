package temp

import (
	//. "exp/kr/common"
	"fmt"
	"math/big"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func F부호없는_정수2큰정수(값 uint64) *big.Int { return new(big.Int).SetUint64(값) }
func F부호없는_정수2정밀수(값 uint64) *big.Rat {
	return new(big.Rat).SetInt(F부호없는_정수2큰정수(값))
}
func F정수2정밀수(값 int64) *big.Rat { return new(big.Rat).SetInt64(값) }
func F실수2문자열(값 float64) string {
	return strconv.FormatFloat(값, 'f', -1, 64)
}
func F실수2정밀수(값 float64) *big.Rat {
	값_문자열 := strconv.FormatFloat(값, 'f', -1, 64)
	정밀값, 변환_성공 := new(big.Rat).SetString(값_문자열)

	if !변환_성공 {
		fmt.Printf("%s:%d: common.F정밀수() : 문자열을 정밀수로 변환 실패. 값 %v", 값)

		정밀값 = new(big.Rat).SetFloat64(값)
	}

	return 정밀값
}

func F정밀수(값 interface{}) (*big.Rat, error) {
	switch 값.(type) {
	case nil:
		파일명, 행 := F소스코드_위치(2)
		fmt.Printf("%s:%d: F정밀수() nil값. 입력값 %v.\n", 파일명, 행, 값)
		에러 := fmt.Errorf("%s:%d: F정밀수() nil값. 입력값 %v.\n", 파일명, 행, 값)

		return nil, 에러
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
		return F실수2정밀수(float64(값.(float32))), nil
	case float64:
		return F실수2정밀수(값.(float64)), nil
	case big.Rat:
		정밀값 := 값.(big.Rat)
		return &정밀값, nil
	case *big.Rat:
		return 값.(*big.Rat), nil
	case I큰정수형:
		return new(big.Rat).SetInt(값.(I큰정수형).G큰정수()), nil
	case I정밀수형:
		return 값.(I정밀수형).G정밀수(), nil
	case I정밀수:
		return 값.(I정밀수).G값(), nil
	//case I고정소숫점:	// 'I고정소숫점'은 'I정밀수형'임
	//	return 값.(I고정소숫점).G정밀수(), nil
	default:
		파일명, 행 := F소스코드_위치(2)
		fmt.Printf("%s:%d: F정밀수() 변환 에러. 입력값 %v.\n", 파일명, 행, 값)
		에러 := fmt.Errorf("%s:%d: F정밀수() 변환 에러. 입력값 %v.\n", 파일명, 행, 값)

		return nil, 에러
	}
}

func F정밀수_같음(값1, 값2 *big.Rat) bool {
	if 값1.Cmp(값2) == 0 {
		return true
	}

	차이_절대값 := new(big.Rat)
	차이_절대값.Sub(값1, 값2)
	차이_절대값.Abs(차이_절대값)

	차이_한도, _ := new(big.Rat).SetString("1/100000000000000000000000000000000")

	if 차이_절대값.Cmp(차이_한도) == -1 {
		// 차이값이 극히 미세하므로 같다고 봐도 무방하다.
		return true
	}

	return false
}

func F값_일치(값1, 값2 interface{}) bool {
	i := 1

	// 'I같음', 'I수치형', 'I큰정수형', 'I정밀수형', 기타 수치형의 경우
	// 서로 다른 타입끼리도 값을 비교할 수 있으므로 가장 우선적으로 시도함.
	switch 값1.(type) {
	case I수치형, I큰정수형, I정밀수형,
		uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64,
		float32, float64,
		big.Int, *big.Int,
		big.Rat, *big.Rat:
		값_일치 := false

		정밀값1, 에러1 := F정밀수(값1)
		정밀값2, 에러2 := F정밀수(값2)

		switch {
		case 에러1 != nil, 에러2 != nil:
			값_일치 = false
			//F체크포인트(&i, "F값_일치() : nil 값.")
		case F정밀수_같음(정밀값1, 정밀값2):
			값_일치 = true
			//F체크포인트(&i, "F값_일치() : F정밀수_같음().")
		default:
			값_일치 = false
			//F체크포인트(&i, "F값_일치() : default")
		}

		F체크포인트(&i, "F값_일치() : 값_일치 '%v', F정밀수_같음() A. 값1 %v, 값2 %v.",
			값_일치, 값1, 값2)

		return 값_일치
	case I같음:
		값_일치 := 값1.(I같음).G같음(값2)

		F체크포인트(&i, "F값_일치() : 값_일치 '%v', 값1.G같음(값2). 값1 %v, 값2 %v.", 값_일치, 값1, 값2)

		return 값_일치
	}

	switch 값2.(type) {
	case I수치형, I큰정수형, I정밀수형,
		uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64,
		float32, float64,
		big.Int, *big.Int,
		big.Rat, *big.Rat:
		값_일치 := false

		정밀값1, 에러1 := F정밀수(값1)
		정밀값2, 에러2 := F정밀수(값2)

		switch {
		case 에러1 != nil, 에러2 != nil:
			값_일치 = false
			//F체크포인트(&i, "F값_일치() : nil 값.")
		case F정밀수_같음(정밀값1, 정밀값2):
			값_일치 = true
			//F체크포인트(&i, "F값_일치() : F정밀수_같음().")
		default:
			값_일치 = false
			//F체크포인트(&i, "F값_일치() : default")
		}

		F체크포인트(&i, "F값_일치() : 값_일치 '%v', F정밀수_같음() B. 값1 %v, 값2 %v.",
			값_일치, 값1, 값2)

		return 값_일치
	case I같음:
		값_일치 := 값2.(I같음).G같음(값1)

		F체크포인트(&i, "F값_일치() : 값_일치 '%v', 값2.G같음(값1). 값1 %v, 값2 %v.", 값_일치, 값1, 값2)

		return 값_일치
	}

	// 이제 형식이 다르면 같을 수 없다고 판정.
	값1_형식 := reflect.TypeOf(값1)
	값2_형식 := reflect.TypeOf(값2)

	if 값1_형식 != 값2_형식 {
		return false
	}

	// 나머지는 reflect.DeepEqual의 힘을 빌림.
	값_일치 := reflect.DeepEqual(값1, 값2)

	F체크포인트(&i, "F값_일치() : 값_일치 '%v', reflect.DeepEqual. 값1 %v, 값2 %v.", 값_일치, 값1, 값2)

	return 값_일치
}

func F값_복사(값 interface{}) interface{} {
	switch 값.(type) {
	case uint, uint8, uint16, uint32, uint64,
		int, int8, int16, int32, int64,
		float32, float64, bool, string, time.Time:
		return reflect.ValueOf(값).Interface()
	case big.Int:
		값_ := 값.(big.Int)
		값_복사본 := new(big.Int).Set(&값_)
		return *값_복사본
	case *big.Int:
		return new(big.Int).Set(값.(*big.Int))
	case big.Rat:
		값_ := 값.(big.Rat)
		값_복사본 := new(big.Rat).Set(&값_)
		return *값_복사본
	case *big.Rat:
		return new(big.Rat).Set(값.(*big.Rat))
	case I값_복사본:
		return 값.(I값_복사본).G값_복사본()
	default:
		// 잘 모르겠으니 그냥 reflect의 힘을 빌림.
		return reflect.ValueOf(값).Interface()
	}
}

type I수치형 interface {
	수치형임()
}
type I큰정수형 interface {
	I수치형
	G큰정수() *big.Int
}
type I정밀수형 interface {
	I수치형
	G정밀수() *big.Rat
}
type I문자열형 interface {
	G문자열() string
}

type I정밀수 interface {
	I정밀수형
	I문자열형
	G값() *big.Rat
	G실수값() float64
	G반올림값(소숫점_이하_자릿수 int) float64
	G반올림_문자열(소숫점_이하_자릿수 int) string
	G부호() int // 음수 -1, 제로 0, 양수 1.
}

type I값_복사본 interface {
	G값_복사본() interface{}
}
type I같음 interface {
	G같음(비교값 interface{}) bool
}
type I테스트용_가상_객체 interface {
	테스트용_가상_객체()
}

func F소스코드_위치(건너뛰는_단계 int) (파일명 string, 행_번호 int) {
	_, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)

	파일명 = filepath.Base(파일_경로)
	return
}

// 디버깅 할 때 각 체크포인트의 위치와 번호를 출력해주는 편의 함수.
func F체크포인트(체크포인트_번호 *int, 포맷_문자열 string, 기타 ...interface{}) {
	파일명, 행_번호 := F소스코드_위치(2)
	fmt.Printf("%s:%d: 체크포인트 %v "+포맷_문자열+"\n\n", append([]interface{}{파일명, 행_번호, *체크포인트_번호}, 기타...)...)
	(*체크포인트_번호)++
}

// 기대값과 실제값이 다르면 Fail하는 테스트용 편의 함수.
func F같은값_확인(테스트 testing.TB, 기대값, 실제값 interface{}) {
	if !F값_일치(기대값, 실제값) {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			파일명, 행_번호 := F소스코드_위치(2)
			fmt.Printf("%s:%d: 값 불일치. 기대값: %#v 실제값: %#v.\n\n", 파일명, 행_번호, 기대값, 실제값)
		}

		//테스트.FailNow()
		테스트.Fail()
	}
}

// 기대값과 실제값이 같으면 Fail하는 테스트용 편의 함수.
func F다른값_확인(테스트 testing.TB, 기대값, 실제값 interface{}) {
	if F값_일치(기대값, 실제값) {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			파일명, 행_번호 := F소스코드_위치(2)
			fmt.Printf("%s:%d: 값 일치. 기대값: %#v 실제값: %#v.\n\n", 파일명, 행_번호, 기대값, 실제값)
		}

		//테스트.FailNow()
		테스트.Fail()
	}
}

func F나를_위한_문구_dev() {
	fmt.Println("")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("	쉽고 간단하게, 테스트로 검증해 가면서 마음 편하게.")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("")
}

func F메모_dev() {
	F나를_위한_문구_dev()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("TODO : F공유해도_안전함() 에서 S로 시작하는 문자열만 값을 변경하는 메소드라고 가정한다.")
	fmt.Println("		이 가정이 틀릴 수도 있는 데,")
	fmt.Println("		실제로 내부값이 변경되는 지 어떻게 확인할 수 있는 지 모르겠음.")
	fmt.Println("		reflect로 복사본을 생성한 후 모든 메소드를 실행시켜봐야 하나?")
	fmt.Println("		1. 복사본 자동 생성")
	fmt.Println("		2. 모든 메소드를 실행.")
	fmt.Println("		3. 내부 멤버필드 값이 변했는 지 확인.")
	fmt.Println("		1, 2, 3을 하나씩 쪼개서 함수로 만든 후 테스트를 하면서 가능한 지 해 볼 것.")
	fmt.Println("")
	fmt.Println("")
	F나를_위한_문구_dev()
}
