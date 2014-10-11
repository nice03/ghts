package lib

import (
	"math/big"
	"testing/quick"
	"time"
)

type I테스트용_가상_객체 interface {
	테스트용_가상_객체()
}

type I정수_식별코드 interface {
	G식별코드() uint64
}
type I문자열_식별코드 interface {
	G식별코드() string
}

// 자료형 공통
type I기본_문자열 interface {
	String() string
}

type I임의값_생성 interface {
	quick.Generator
}

type I자료형_공통 interface {
	I기본_문자열
	I임의값_생성
}

type I상수형 interface {
	I자료형_공통
	상수형임()
}

type I변수형 interface {
	I자료형_공통
	변수형임()
}

type I문자열형 interface {
	G문자열() string
}

type I실수형 interface {
	G실수() float64
	G정밀수() C정밀수
}

type I정수형 interface {
	G정수() int64
	G실수() float64
	G정밀수() C정밀수
}

type I통화종류 interface {
	G통화종류() P통화종류
}

// 정수
type I정수 interface {
	I정수형
	G값() int64
}

type C정수 interface {
	I상수형
	I정수
	G변수형() V정수
}

type V정수 interface {
	I변수형
	I정수
	G상수형() C정수

	S값(값 int64) V정수
	S절대값() V정수
	S더하기(값 int64) V정수
	S빼기(값 int64) V정수
	S곱하기(값 int64) V정수
	S나누기(값 int64) V정수
}

// 부호없는 정수
type I부호없는_정수 interface {
	I정수형
	G값() uint64
}

type C부호없는_정수 interface {
	I상수형
	I부호없는_정수
	G변수형() V부호없는_정수
}

type V부호없는_정수 interface {
	I변수형
	I부호없는_정수
	G상수형() C부호없는_정수

	S값(값 uint64) V부호없는_정수
	S더하기(값 uint64) V부호없는_정수
	S빼기(값 uint64) V부호없는_정수
	S곱하기(값 uint64) V부호없는_정수
	S나누기(값 uint64) V부호없는_정수
}

// 실수
type I실수 interface {
	I실수형
	G값() float64
}

type C실수 interface {
	I상수형
	I실수
	G변수형() V실수
}

type V실수 interface {
	I변수형
	I실수
	G상수형() C실수

	S값(값 float64) V실수
	S절대값() V실수
	S더하기(값 float64) V실수
	S빼기(값 float64) V실수
	S곱하기(값 float64) V실수
	S나누기(값 float64) V실수
}

type I참거짓 interface {
	G값() bool
}

type C참거짓 interface {
	I상수형
	I참거짓
	G변수형() V참거짓
}

type V참거짓 interface {
	I변수형
	I참거짓
	G상수형() C참거짓
	S값(값 bool) V참거짓
}

type C문자열 interface {
	I상수형
	G값() string
}

// 시점
type I시점 interface {
	G값() time.Time
}

type C시점 interface {
	I상수형
	I시점
	G변수형() V시점
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
	G비교(값 interface{}) int // -1 : 더 작음, 0 : 같음, 1 : 더 큼. -2 : 숫자 아님
}

type C정밀수 interface {
	I상수형
	I정밀수
	G변수형() V정밀수
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
}

type I통화 interface {
	G종류() P통화종류
	G값() C정밀수
	G같음(값 I통화) bool
}

type C통화 interface {
	I상수형
	I통화
	G변수형() V통화
}

type V통화 interface {
	I변수형
	I통화
	G상수형() C통화
	S종류(종류 P통화종류)
	S값(금액 interface{}) V통화

	S절대값() V통화
	S더하기(값 interface{}) V통화
	S빼기(값 interface{}) V통화
	S곱하기(값 interface{}) V통화
	S나누기(값 interface{}) V통화
	S반대부호값() V통화
}

type C매개변수 interface {
	G이름() string
	G값() I상수형
}

/*
type I환율 interface {
	G원래통화() P통화종류
	G변환통화() P통화종류
	G환율() C정밀수
}

type I종목별_포트폴리오 interface {
	I정수_식별코드
	I통화종류

	G종목() C종목
	G단가() C통화

	G롱포지션_수량() uint64
	G숏포지션_수량() uint64
	G순_수량() uint64
	G총_수량() uint64

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

type C종목별_포트폴리오 interface {
	I상수형
	I종목별_포트폴리오
}

type V종목별_포트폴리오 interface {
	I변수형
	I종목별_포트폴리오

	S종목(종목 C종목)
	S단가(단가 C통화)

	S롱포지션_수량(수량 uint64)
	S숏포지션_수량(수량 uint64)
	S순_수량(수량 uint64)
	S총_수량(수량 uint64)

	S롱포지션_금액(금액 C통화)
	S숏포지션_금액(금액 C통화)
	S순_금액(금액 C통화)
	S총_금액(금액 C통화)
}

type I포트폴리오 interface {
	I정수_식별코드
	I통화종류

	G보유_종목_모음() []C종목
	G종목별_포트폴리오(종목코드 string) C종목별_포트폴리오
	G전종목_포트폴리오() []C종목별_포트폴리오

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

type C포트폴리오 interface {
	I상수형
	I포트폴리오
}

type V포트폴리오 interface {
	I변수형
	I포트폴리오

	S종목별_포트폴리오(종목별_포트폴리오 C종목별_포트폴리오)
	G상수형() C포트폴리오
} */