package common

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"testing"
	//"time"
)

/*
type s진짜_상수 struct{}
func (s *s진짜_상수) 상수형임()          {}
func (s *s진짜_상수) G같음(값 interface{}) bool { return false }
func (s *s진짜_상수) String() string { return "" }

type s가짜_상수_1 struct{}

func (s *s가짜_상수_1) G값() int { return 0 }

type s가짜_상수_2 struct{}

func (s *s가짜_상수_2) 상수형임()   {}
func (s *s가짜_상수_2) G값() int { return 0 }
func (s *s가짜_상수_2) G같음(값 interface{}) bool { return false }
func (s *s가짜_상수_2) S값()     {}

type s가짜_상수_3 struct{}

func (s *s가짜_상수_3) 상수형임()   {}
func (s *s가짜_상수_3) G값() int { return 0 }
func (s *s가짜_상수_3) G같음(값 interface{}) bool { return false }
func (s *s가짜_상수_3) String() {}

type s가짜_상수_4 struct{}

func (s *s가짜_상수_4) 상수형임()                    {}
func (s *s가짜_상수_4) G값() int                  { return 0 }
func (s *s가짜_상수_4) G같음(값 interface{}) bool { return false }
func (s *s가짜_상수_4) String(문자열 string) string { return "" }

type s가짜_상수_5 struct{}

func (s *s가짜_상수_5) 상수형임()          {}
func (s *s가짜_상수_5) 변수형임()          {}
func (s *s가짜_상수_5) G값() int        { return 0 }
func (s *s가짜_상수_5) G같음(값 interface{}) bool { return false }
func (s *s가짜_상수_5) String() string { return "" }


func TestF공유해도_안전함(테스트 *testing.T) {
	F참인지_확인(테스트, F공유해도_안전함(uint(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(uint8(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(uint16(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(uint32(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(uint64(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(int(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(int8(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(int16(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(int32(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(int64(0)), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(true), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(false), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함("문자열"), "TestF상수형임() : false negative.")
	F참인지_확인(테스트, F공유해도_안전함(time.Now()), "TestF상수형임() : false negative.")

	F참인지_확인(테스트, F공유해도_안전함(&s진짜_상수{}), "TestF상수형임() : false negative.")

	i := big.NewInt(0)
	r := big.NewRat(1,1)

	F거짓인지_확인(테스트, F공유해도_안전함(*i), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(big.NewInt(0)), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(*r), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(big.NewRat(1,1)), "TestF상수형임() : false positive.")

	F거짓인지_확인(테스트, F공유해도_안전함(&s가짜_상수_1{}), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(&s가짜_상수_2{}), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(&s가짜_상수_3{}), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(&s가짜_상수_4{}), "TestF상수형임() : false positive.")
	F거짓인지_확인(테스트, F공유해도_안전함(&s가짜_상수_5{}), "TestF상수형임() : false positive.")
} */


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

func TestF정수2정밀수(테스트 *testing.T) {
	원래값 := int64(1001)
	var 변환값 *big.Rat = F정수2정밀수(원래값)
	
	참거짓 := 변환값.Cmp(new(big.Rat).SetInt64(원래값)) == 0
	
	F참인지_확인(테스트, 참거짓, "TestF정수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}
func TestF정수2문자열(테스트 *testing.T) {
	F같은값_확인(테스트, F정수2문자열(int64(100)), "100")
}

func TestF실수2정밀수(테스트 *testing.T) {
	원래값 := float64(1001.020023)
	var 변환값 *big.Rat = F실수2정밀수(원래값)
	
	참거짓 := 변환값.FloatString(6) == "1001.020023"
	
	F참인지_확인(테스트, 참거짓, "TestF실수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
}

func TestF큰정수2정밀수(테스트 *testing.T) {
	원래값 := big.NewInt(1001)
	var 변환값 *big.Rat = F큰정수2정밀수(원래값)
	
	참거짓 := 변환값.Cmp(big.NewRat(원래값.Int64(), 1)) == 0
	
	F참인지_확인(테스트, 참거짓, "TestF큰정수2정밀수() : 원래값  %v, 변환값 %v", 원래값, 변환값)
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
	파일명, 행_번호 :=  소스코드_위치[0], 소스코드_위치[1]
	
	fmt.Printf("파일명 '%v', 행_번호 '%v'. \n\n", 파일명, 행_번호)
	fmt.Printf("파일명 '%v', 행_번호 '%v', 나머지 '%v'. \n\n", 소스코드_위치[0], 소스코드_위치[1], 소스코드_위치[2])
	
	F참인지_확인(테스트, strings.HasPrefix(파일명, "0001_"),
		"TestF소스코드_위치() : F소스코드_위치() 파일명_에러. 값 %v", 파일명)

	F참인지_확인(테스트, strings.HasSuffix(파일명, ".go"),
		"TestF소스코드_위치() : F소스코드_위치() 파일명_에러. 값 %v", 파일명)

	소스코드_위치 = strings.Split(F소스코드_위치(1), ":")
	파일명, 행_번호 =  소스코드_위치[0], 소스코드_위치[1]
	_, _, 행_번호_예상값, _ := runtime.Caller(0)
	
	fmt.Printf("파일명 '%v', 행_번호 '%v'. \n\n", 파일명, 행_번호)
	fmt.Printf("파일명 '%v', 행_번호 '%v', 나머지 '%v'. \n\n", 소스코드_위치[0], 소스코드_위치[1], 소스코드_위치[2])

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
	테스트_결과_반환값  = F참인지_확인(가상_테스트, false, "")
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
	테스트_결과_반환값  = F거짓인지_확인(가상_테스트, false, "")
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
	if !테스트_통과 || !테스트_결과_반환값{
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
