package common

import (
	"fmt"
	//"math/big"
	"path/filepath"
	"reflect"
	"runtime"
	//"strings"
	"testing"
	//"time"
)

/*
func F값_복사() {}
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

func F값_일치(기대값, 실제값 interface{}) bool {	
	/*
	i := 1
	
	switch 기대값.(type) {
	case I같음:
		값_일치 := 기대값.(I같음).G같음(실제값)
		
		F체크포인트(&i, "F값_일치() : 값_일치 '%v', 2 기대값.G같음(실제값). 기대값 %v, 실제값 %v.", 값_일치, 기대값, 실제값)
		
		return 값_일치
	}

	switch 실제값.(type) {
	case I같음:
		값_일치 := 실제값.(I같음).G같음(기대값)
		
		F체크포인트(&i, "F값_일치() :  값_일치 '%v', 1 실제값.G같음(기대값). 기대값 %v, 실제값 %v.", 값_일치, 기대값, 실제값)
		
		return 값_일치
	}
	
	기대값_형식 := reflect.TypeOf(기대값)
	실제값_형식 := reflect.TypeOf(실제값)
	if 기대값_형식 != 실제값_형식 {
		return false
	}
	
	값_일치 := reflect.DeepEqual(기대값, 실제값)
		
	F체크포인트(&i, "F값_일치() : 값_일치 '%v', 3 reflect.DeepEqual. 기대값 %v, 실제값 %v.", 값_일치, 기대값, 실제값)
	
	return 값_일치 */
	
	return false
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
	파일명, 행_번호 := F소스코드_위치(2)
	fmt.Printf("%s:%d: 체크포인트 %v " + 포맷_문자열 + "\n\n", append([]interface{}{파일명, 행_번호, *체크포인트_번호}, 기타...)...)
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
func F소스코드_위치(건너뛰는_단계 int) (파일명 string, 행_번호 int) {
	_, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)

	파일명 = filepath.Base(파일_경로)
	return
}

// '참거짓'이 false이면 Fail하는 테스트용 편의 함수.
func F참인지_확인(테스트 testing.TB, 참거짓 bool, 포맷_문자열 string,
	추가_매개변수 ...interface{}) {

	if !참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			파일명, 행_번호 := F소스코드_위치(2)
			fmt.Printf("%s:%d: "+포맷_문자열+"\n\n", 
						append([]interface{}{파일명, 행_번호}, 추가_매개변수...)...)
		}

		//테스트.FailNow()
		테스트.Fail()
	}
}

// '참거짓'이 true이면 Fail하는 테스트용 편의 함수.
func F거짓인지_확인(테스트 testing.TB, 참거짓 bool, 포맷_문자열 string,
	추가_매개변수 ...interface{}) {
	if 참거짓 {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			파일명, 행_번호 := F소스코드_위치(2)
			fmt.Printf("%s:%d: "+포맷_문자열+"\n\n", 
						append([]interface{}{파일명, 행_번호}, 추가_매개변수...)...)
		}
		
		//테스트.FailNow()
		테스트.Fail()
	}
}

// 에러가 발생하면 Fail하는 테스트용 편의 함수.
func F에러없음_확인(테스트 testing.TB, 에러 error) {
	if 에러 != nil {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			파일명, 행_번호 := F소스코드_위치(2)
			fmt.Printf("%s:%d: 에러 발생. : %s \n\n", 파일명, 행_번호, 에러.Error())
		}
		
		//테스트.FailNow()
		테스트.Fail()
	}
}

// 에러가 발생하지 않으면 Fail하는 테스트용 편의 함수.
func F에러발생_확인(테스트 testing.TB, 에러 error) {
	if 에러 == nil {
		switch 테스트.(type) {
		case I테스트용_가상_객체:
			// PASS
		default:
			파일명, 행_번호 := F소스코드_위치(2)
			fmt.Printf("%s:%d: 에러 발생. : %s \n\n", 파일명, 행_번호, 에러.Error())
		}
		
		//테스트.FailNow()
		테스트.Fail()
	}
}

// 기대값과 실제값이 다르면 Fail하는 테스트용 편의 함수.
func F같은값_확인(테스트 testing.TB, 기대값, 실제값 interface{}) {
	if !F값_일치(기대값, 실제값) &&
		!reflect.DeepEqual(기대값, 실제값) {
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
	if F값_일치(기대값, 실제값) ||
		reflect.DeepEqual(기대값, 실제값) {
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