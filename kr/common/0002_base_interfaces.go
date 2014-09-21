package common

import (
	"math/big"
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

// 원본과 복사본은 서로 독립성을 가져야 함.
// 즉, 멤버 필드 중에 참조형이 있으면
// 같은 타입의 새로운 인스턴스를 생성한 후, 값을 복사해야 함.

type I값_복사본 interface {
	G값_복사본() interface{}
}
type I같음 interface {
	G같음(비교값 interface{}) bool
}

type I기본_문자열 interface {
	String() string
}

type I자료형_공통 interface {
	I기본_문자열
	I값_복사본
	I같음
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

/*
type I고정소숫점형 interface {
	I수치형
	G정밀수() *big.Rat
	G고정소숫점(소숫점_이하_자릿수 int) C고정소숫점
}

type I통화형 interface { G통화() C통화 }

type I통화종류 interface { G통화종류() P통화 }

type I환율 interface {
	I통화종류
	G기준통화() P통화
	G환율() C고정소숫점
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
