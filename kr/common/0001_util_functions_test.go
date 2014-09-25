package common

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestF_nil_존재함(테스트 *testing.T) {
	F참인지_확인(테스트, F_nil_존재함(nil), "")
	F참인지_확인(테스트, F_nil_존재함(1, nil, "test", 1.1), "")
	F거짓인지_확인(테스트, F_nil_존재함(1, "test", 1.1), "")
}

func TestF부호없는_정수2큰정수(테스트 *testing.T) {
	원래값 := uint64(1001)
	var 변환값 *big.Int = F부호없는_정수2큰정수(원래값)

	참거짓 := 변환값.Cmp(new(big.Int).SetUint64(원래값)) == 0

	F참인지_확인(테스트, 참거짓, "TestF부호없는_정수2큰정수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}

func TestF부호없는_정수2정밀수(테스트 *testing.T) {
	원래값 := uint64(1001)
	var 변환값 *big.Rat = F부호없는_정수2정밀수(원래값)

	참거짓 := 변환값.Cmp(new(big.Rat).SetInt(F부호없는_정수2큰정수(원래값))) == 0

	F참인지_확인(테스트, 참거짓, "TestF부호없는_정수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}

func TestF부호없는_정수2문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F부호없는_정수2문자열(100), "100")
}

func TestF정수2큰정수(테스트 *testing.T) {
	원래값 := int64(100)
	var 변환값 *big.Int = F정수2큰정수(원래값)
	
	참거짓 := 변환값.Cmp(big.NewInt(원래값)) == 0
	
	F참인지_확인(테스트, 참거짓, "TestF정수2큰정수() : 원래값 %v, 변환값 %v", 원래값, 변환값)
}

func TestF정수2정밀수(테스트 *testing.T) {
	원래값 := int64(1001)
	var 변환값 *big.Rat = F정수2정밀수(원래값)

	참거짓 := 변환값.Cmp(new(big.Rat).SetInt64(원래값)) == 0

	F참인지_확인(테스트, 참거짓, "TestF정수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}

func TestF정수2문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F정수2문자열(int64(100)), "100")
}

func TestF정수2월(테스트 *testing.T) {
	월, 에러 := F정수2월(1)
	F에러없음_확인(테스트, 에러)
	F같은값_확인(테스트, 월, time.January)
	
	월, 에러 = F정수2월(13)
	F에러발생_확인(테스트, 에러)
}

func TestF실수2정밀수(테스트 *testing.T) {
	원래값 := float64(1001.020023)
	var 변환값 *big.Rat = F실수2정밀수(원래값)

	참거짓 := 변환값.FloatString(6) == "1001.020023"

	F참인지_확인(테스트, 참거짓, "TestF실수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}

func TestF실수2문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F실수2문자열(100.25), "100.25")
}

func TestF큰정수2정밀수(테스트 *testing.T) {
	원래값 := big.NewInt(1001)
	var 변환값 *big.Rat = F큰정수2정밀수(원래값)

	참거짓 := 변환값.Cmp(big.NewRat(원래값.Int64(), 1)) == 0

	F참인지_확인(테스트, 참거짓, "TestF큰정수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}

func TestF정밀수_복사(테스트 *testing.T) {
	원래값 := big.NewRat(1001, 10)
	예상값 := big.NewRat(1001, 10)
	복사값 := F정밀수_복사(원래값)
	
	F같은값_확인(테스트, 복사값, 예상값)
	
	원래값.Add(원래값, big.NewRat(100, 1))
	
	F같은값_확인(테스트, 복사값, 예상값)

}

func TestF정밀수2실수(테스트 *testing.T) {
	실수 := F정밀수2실수(big.NewRat(1001, 10))
	
	F같은값_확인(테스트, 실수, 100.1)
}

func TestF정밀수2문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F정밀수2문자열(big.NewRat(100340, 1000)), "100.34")
}

func TestF정밀수_반올림_문자열(테스트 *testing.T) {
	정밀수 := big.NewRat(104594, 10000)
	F같은값_확인(테스트, F정밀수_반올림_문자열(정밀수, 4), "10.4594")
	F같은값_확인(테스트, F정밀수_반올림_문자열(정밀수, 3), "10.459")
	F같은값_확인(테스트, F정밀수_반올림_문자열(정밀수, 2), "10.46")
	F같은값_확인(테스트, F정밀수_반올림_문자열(정밀수, 1), "10.5")
	F같은값_확인(테스트, F정밀수_반올림_문자열(정밀수, 0), "10")
}

func TestF참거짓2문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F참거짓2문자열(true), "true")
	F같은값_확인(테스트, F참거짓2문자열(false), "false")
}

func TestF문자열2실수(테스트 *testing.T) {
	실수, 에러 := F문자열2실수("18.593")

	F에러없음_확인(테스트, 에러)
	F같은값_확인(테스트, 실수, 18.593)
	
	실수, 에러 = F문자열2실수("실수로 변환 불가능한 문자열")
	
	F에러발생_확인(테스트, 에러)
}

func TestF문자열2일자(테스트 *testing.T) {
	일자, 에러 := F문자열2일자("2000-01-01")
	
	F에러없음_확인(테스트, 에러)
	F같은값_확인(테스트, 일자.Format("2006-01-02"), "2000-01-01")

	일자, 에러 = F문자열2일자("변환 불가능한 문자열")
	F에러발생_확인(테스트, 에러)
}

func TestF일자2문자열(테스트 *testing.T) {
	월, _ := F정수2월(1)
	일자 := time.Date(2000, 월, 1, 0, 0, 0, 0, time.Now().Location())

	F같은값_확인(테스트, F일자2문자열(일자), "2000-01-01")
}

type i테스트_인터페이스를_구현함_a interface {
	a()
}
type i테스트_인터페이스를_구현함_b interface {
	b()
}

type s테스트_인터페이스를_구현함_a struct{}

func (s *s테스트_인터페이스를_구현함_a) a() {}

type s테스트_인터페이스를_구현함_b struct{}

func (s *s테스트_인터페이스를_구현함_b) b() {}

func TestF인터페이스를_구현함(테스트 *testing.T) {
	ia := reflect.TypeOf((*i테스트_인터페이스를_구현함_a)(nil)).Elem()
	ib := reflect.TypeOf((*i테스트_인터페이스를_구현함_b)(nil)).Elem()

	sa := new(s테스트_인터페이스를_구현함_a)
	sb := new(s테스트_인터페이스를_구현함_b)

	참거짓 := F인터페이스를_구현함(sa, ia)
	포맷_문자열 := "예상과 다른 결과. 값 %v 인터페이스 %v 구현 여부 %v"
	F참인지_확인(테스트, 참거짓, 포맷_문자열, reflect.TypeOf(sa), reflect.TypeOf(ia), 참거짓)

	참거짓 = F인터페이스를_구현함(sa, ib)
	F거짓인지_확인(테스트, 참거짓, 포맷_문자열, reflect.TypeOf(sa), reflect.TypeOf(ib), 참거짓)

	참거짓 = F인터페이스를_구현함(sb, ia)
	F거짓인지_확인(테스트, 참거짓, 포맷_문자열, reflect.TypeOf(sb), reflect.TypeOf(ia), 참거짓)

	참거짓 = F인터페이스를_구현함(sb, ib)
	F참인지_확인(테스트, 참거짓, 포맷_문자열, reflect.TypeOf(sb), reflect.TypeOf(ib), 참거짓)
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

func TestF체크포인트(테스트 *testing.T) {
	/*
		체크포인트_번호 := 1
		F체크포인트("TestF체크포인트", &체크포인트_번호)
		F체크포인트("TestF체크포인트", &체크포인트_번호)
		F체크포인트("TestF체크포인트", &체크포인트_번호)
	*/
}

func TestF소스코드_위치(테스트 *testing.T) {
	소스코드_위치 := strings.Split(F소스코드_위치(0), ":")
	파일명, 행_번호 := 소스코드_위치[0], 소스코드_위치[1]
	
	F참인지_확인(테스트, strings.HasPrefix(파일명, "0001_"),
		"TestF소스코드_위치() : F소스코드_위치() 파일명_에러. 값 %v", 파일명)

	F참인지_확인(테스트, strings.HasSuffix(파일명, ".go"),
		"TestF소스코드_위치() : F소스코드_위치() 파일명_에러. 값 %v", 파일명)

	소스코드_위치 = strings.Split(F소스코드_위치(1), ":")
	파일명, 행_번호 = 소스코드_위치[0], 소스코드_위치[1]
	_, _, 행_번호_예상값, _ := runtime.Caller(0)

	F참인지_확인(테스트, strings.HasPrefix(파일명, "0001_"),
		"TestF소스코드_위치() : F소스코드_위치() 파일명_에러. 값 %v", 파일명)

	F참인지_확인(테스트, strings.HasSuffix(파일명, "_test.go"),
		"TestF소스코드_위치() : F소스코드_위치() 파일명_에러. 값 %v", 파일명)

	F같은값_확인(테스트, F정수2문자열(int64(행_번호_예상값-2)), 행_번호)
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

	테스트_통과 = true
	테스트_결과_반환값 := F참인지_확인(가상_테스트, true, "")
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(1))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F참인지_확인(가상_테스트, false, "")
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(1))
	}
}

func TestF거짓인지_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F거짓인지_확인(가상_테스트, true, "")
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(1))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F거짓인지_확인(가상_테스트, false, "")
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(1))
	}
}

func TestF에러없음_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F에러없음_확인(가상_테스트, nil)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(1))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F에러없음_확인(가상_테스트, fmt.Errorf(""))
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(1))
	}
}

func TestF에러발생_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F에러발생_확인(가상_테스트, nil)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(1))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F에러발생_확인(가상_테스트, fmt.Errorf(""))
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(1))
	}
}

func TestF같은값_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F같은값_확인(가상_테스트, 1, 1)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 실패.", F소스코드_위치(1))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F같은값_확인(가상_테스트, 1, 2)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Errorf("%s예상치 못한 테스트 통과.", F소스코드_위치(1))
	}
}

func TestF다른값_확인(테스트 *testing.T) {
	가상_테스트 := new(s가상TB)

	테스트_통과 = true
	테스트_결과_반환값 := F다른값_확인(가상_테스트, 1, 1)
	if 테스트_통과 || 테스트_결과_반환값 {
		테스트.Error("%s예상치 못한 테스트 통과.", F소스코드_위치(1))
	}

	테스트_통과 = true
	테스트_결과_반환값 = F다른값_확인(가상_테스트, 1, 2)
	if !테스트_통과 || !테스트_결과_반환값 {
		테스트.Error("%s예상치 못한 테스트 실패.", F소스코드_위치(1))
	}
}
