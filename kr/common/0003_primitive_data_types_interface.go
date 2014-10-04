package common

import (
	"math/big"
	"time"
)

// 구조체를 공개하면 new()로 직접 생성해서 초기화가 적절하게 되지 않는 경우가 발생함.
// 항상 적절한 초기화가 되도록 구조체 자체는 외부에 숨기고,
// 생성자(N으로 시작됨)를 통해서만 생성할 수 있도록 하여, 생성자에서 적절한 초기화가 이루어지도록 함.
// 구조체를 사용하기 위해서는 외부에 공개된 관련 인터페이스를 사용함.
// 예) SC정수를 사용하기 위해서 NC정수로 생성해서 C정수 인터페이스를 통해서 사용.
// Go언어에는 생성자가 따로 없어서 이런 식으로 해결함.

// 기본 데이터 타입은 Go언어 내장 자료형을 사용하면 되며, 별도의 변수형이 필요없음.
type I정수 interface { G값() int64 }

func NC정수(값 int64) C정수 { return &sC정수64{값} }
type C정수 interface {
	I상수형
	I정수형
	I정수
	G변수형() V정수
}

func NV정수(값 int64) V정수 {
	return &sV정수64{값: 값}
}
	
type V정수 interface {
	I변수형
	I정수형
	I정수
	G상수형() C정수
	
	S값(값 int64) V정수
	S절대값() V정수
	S더하기(값 int64) V정수
	S빼기(값 int64) V정수
	S곱하기(값 int64) V정수
	S나누기(값 int64) V정수
}

type I부호없는_정수 interface { G값() uint64 }

func NC부호없는_정수(값 uint64) C부호없는_정수 { return &sC부호없는_정수64{값} }
type C부호없는_정수 interface {
	I상수형
	I정수형
	I부호없는_정수
	G변수형() V부호없는_정수
}

func NV부호없는_정수(값 uint64) V부호없는_정수 {
	return &sV부호없는_정수64{값: 값}
}

type V부호없는_정수 interface {
	I변수형
	I정수형
	I부호없는_정수
	G상수형() C부호없는_정수
	
	S값(값 uint64) V부호없는_정수
	S더하기(값 uint64) V부호없는_정수
	S빼기(값 uint64) V부호없는_정수
	S곱하기(값 uint64) V부호없는_정수
	S나누기(값 uint64) V부호없는_정수
}

type I실수 interface { G값() float64 }

func NC실수(값 float64) C실수   { return &sC실수64{값} }
type C실수 interface {
	I상수형
	I실수형
	I실수
	G변수형() V실수
}

func NV실수(값 float64) V실수 {
	return &sV실수64{값: 값}
}

type V실수 interface {
	I변수형
	I실수형
	I실수
	G상수형() C실수
	
	S값(값 float64) V실수
	S절대값() V실수
	S더하기(값 float64) V실수
	S빼기(값 float64) V실수
	S곱하기(값 float64) V실수
	S나누기(값 float64) V실수
}

func NC참거짓(값 bool) C참거짓 {
	if 값 { return c참
	} else { return c거짓 }
}

type C참거짓 interface {
	I상수형
	G값() bool
}

func NC문자열(값 string) C문자열 { return &sC문자열{값} }
type C문자열 interface {
	I상수형
	G값() string
}

// 시점
type I시점 interface {
	G값() time.Time
}

func NC시점(값 time.Time) C시점 { return &sC시점{값} }
func NC시점_문자열(값 string) C시점 {
	시점, 에러 := F문자열2시점(값)
	
	if 에러 != nil { return nil }
	
	return NC시점(시점)
}

type C시점 interface {
	I상수형
	I시점
	G변수형() V시점
}

func NV시점(값 time.Time) V시점 {
	return &sV시점{값:값}
}

func NV시점_문자열(값 string) V시점 {
	시점, 에러 := F문자열2시점(값)
	
	if 에러 != nil { return nil }
	
	return NV시점(시점)
}
type V시점 interface {
	I변수형
	I시점
	G상수형() C시점
	S값(값 time.Time) V시점
	S날짜_더하기(연, 월, 일 int) V시점
}

// 정밀수
type I정밀수 interface {
	I실수형
	G값() string
	GRat() *big.Rat
	G같음(값 interface{}) bool
	G비교(값 interface{}) int	// -1 : 더 작음, 0 : 같음, 1 : 더 큼. -2 : 숫자 아님  
}

func NC정밀수(값 interface{}) C정밀수 {
	if 값 == nil { return nil }
	
	var 정밀수 *big.Rat
	
	switch 값.(type) {
	case C정밀수:
		return 값.(C정밀수)	// 상수형은 굳이 새로운 인스턴스를 생성할 필요가 없다.
	case V정밀수:
		정밀수 = 값.(V정밀수).GRat()
	case *big.Rat:
		정밀수 = new(big.Rat).Set(값.(*big.Rat))	
	default:
		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))
		
		if !성공 {
			F문자열_출력("common.NC정밀수() : 입력값이 숫자가 아님. %v", 값)
			return nil
		}
	}

	return &sC정밀수{정밀수}
}
type C정밀수 interface {
	I상수형
	I정밀수
	G변수형() V정밀수
}


func NV정밀수(값 interface{}) V정밀수 {
	if 값 == nil { return nil }
	
	var 정밀수 *big.Rat
	
	switch 값.(type) {
	case *big.Rat:
		정밀수 = new(big.Rat).Set(값.(*big.Rat))
	case I정밀수:
		정밀수 = 값.(I정밀수).GRat()
	default:
		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))
		
		if !성공 {
			F문자열_출력("common.NV정밀수() : 입력값이 숫자가 아님. %v", 값)
			return nil
		}
	}

	return &sV정밀수{값: 정밀수}
}
type V정밀수 interface {
	I변수형
	I정밀수
	G상수형() C정밀수
	
	S값(값 interface{}) V정밀수
	S반올림(소숫점_이하_자릿수 int) V정밀수
	S절대값() V정밀수
	S더하기(값 interface{}) V정밀수
	S빼기(값 interface{}) V정밀수
	S곱하기(값 interface{}) V정밀수
	S나누기(값 interface{}) V정밀수
	S역수() V정밀수
	S반대부호값() V정밀수
	
	S절대값2(값 interface{}) V정밀수
	S더하기2(값1, 값2 interface{}) V정밀수
	S빼기2(값1, 값2 interface{}) V정밀수
	S곱하기2(값1, 값2 interface{}) V정밀수
	S나누기2(값1, 값2 interface{}) V정밀수
	S역수2(값 interface{}) V정밀수
	S반대부호값2(값 interface{}) V정밀수
}

// 통화
type P통화 string

const (
	INVALID_CURRENCY P통화 = "INVALID_CURRENCY"
	KRW  	= "KRW"
	USD     = "USD"
	CNY     = "CNY"
	EUR     = "EUR"
)

type I통화 interface {
	G종류() P통화
	G금액() C정밀수
	G같음(값 I통화) bool
}

func NC원화(금액 interface{}) C통화 { return NC통화(KRW, 금액) }
func NC달러(금액 interface{}) C통화 { return NC통화(USD, 금액) }
func NC위안화(금액 interface{}) C통화 { return NC통화(CNY, 금액) }
func NC유로화(금액 interface{}) C통화 { return NC통화(EUR, 금액) }

func NC통화(종류 P통화, 금액 interface{}) C통화 {
	v금액 := NV정밀수(금액)
	
	if v금액 == nil {
		F문자열_출력("NC통화() : 금액이 숫자가 아님. %v", 금액)
		return nil
	}
	
	c금액 := v금액.S반올림(F통화종류별_정밀도(종류)).G상수형()
	
	return &sC통화{종류, c금액}
}
type C통화 interface {
	I상수형
	I통화
	G변수형() V통화
}

// 변수형 생성자
func NV원화(금액 interface{}) V통화 { return NV통화(KRW, 금액) }
func NV달러(금액 interface{}) V통화 { return NV통화(USD, 금액) }
func NV위안화(금액 interface{}) V통화 { return NV통화(CNY, 금액) }
func NV유로화(금액 interface{}) V통화 { return NV통화(EUR, 금액) }

func NV통화(종류 P통화, 금액 interface{}) V통화 {
	v금액 := NV정밀수(금액)
	
	if v금액 == nil {
		F문자열_출력("NC통화() : 금액이 숫자가 아님. %v", 금액)
		return nil
	}
	
	v금액 = v금액.S반올림(F통화종류별_정밀도(종류))
	
	return &sV통화{종류: 종류, 금액: v금액}
}
type V통화 interface {
	I변수형
	I통화
	G상수형() C통화
	S종류(종류 P통화)
	S금액(금액 I정밀수) V통화
	
	S절대값() V통화
	S더하기(값 interface{}) V통화
	S빼기(값 interface{}) V통화
	S곱하기(값 interface{}) V통화
	S나누기(값 interface{}) V통화
	S반대부호값() V통화
	
	S절대값2(값 I통화) V통화
	S더하기2(값1, 값2 interface{}) V통화
	S빼기2(값1, 값2 interface{}) V통화
	S곱하기2(값1, 값2 interface{}) V통화
	S나누기2(값1, 값2 interface{}) V통화
	S반대부호값2(값 I통화) V통화
}

func NC매개변수(이름 string, 값 interface{}) C매개변수 {
	var 상수형 I상수형 = F상수형(값)
	
	if 상수형 == nil { return nil }
	
	return &sC매개변수{이름, 상수형}
}
type C매개변수 interface {
	G이름() string
	G값() I상수형
}