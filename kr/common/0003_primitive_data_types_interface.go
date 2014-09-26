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
type C정수 interface {
	I상수형
	I정수형
	G값() int64
}

type C부호없는_정수 interface {
	I상수형
	I정수형
	G값() uint64
}

type C실수 interface {
	I상수형
	I실수형
	G값() float64
}

type C문자열 interface {
	I상수형
	G값() string
}

type C참거짓 interface {
	I상수형
	G값() bool
}

// 시점
type I시점 interface {
	G값() time.Time
}

type C시점 interface {
	I상수형
	I시점
}

type V시점 interface {
	I변수형
	I시점
	S값(값 time.Time)
	G상수형() C시점
}

// 큰 정수
type I큰정수 interface {
	I정수형
	G값() *big.Int
}

type C큰정수 interface {
	I상수형
	I큰정수
	G변수형() V큰정수
}

type V큰정수 interface {
	I변수형
	I큰정수
	G상수형() C큰정수
	S값(값 int64)
	S값Big(값 *big.Int)
	
	S절대값(값 I큰정수) V큰정수
	S더하기(값1 I큰정수, 값2 I큰정수) V큰정수
	S빼기(값1 I큰정수, 값2 I큰정수) V큰정수
	S곱하기(값1 I큰정수, 값2 I큰정수) V큰정수
	S나누기(값1 I큰정수, 값2 I큰정수) (V큰정수, error)
	S반대부호값(값 I큰정수) V큰정수
	
	S셀프_절대값() V큰정수
	S셀프_더하기(값 I큰정수) V큰정수
	S셀프_빼기(값 I큰정수) V큰정수
	S셀프_곱하기(값 I큰정수) V큰정수
	S셀프_나누기(값 I큰정수) (V큰정수, error)
	S셀프_반대부호값() V큰정수
}

// 정밀수
type I정밀수 interface {
	I실수형
	G값() *big.Rat
	G반올림_실수(소숫점_이하_자릿수 int) float64
	G반올림_정밀수(소숫점_이하_자릿수 int) *big.Rat
	G반올림_문자열(소숫점_이하_자릿수 int) string
	//G부호() int	// 음수 -1, 제로 0, 양수 1.
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
	S값(값 float64)
	S값Big(값 *big.Rat)
	
	// 연산 함수
	// S로 시작하는 연산 함수는 '메소드 연속 호출'(메소드 체이닝) 기법을 사용하기 위해서 
	//   자기 자신을 결과값으로 설정하고 자기 자신을 반환한다.
	// 예 : v정밀수.S빼기(v1, v2).S셀프_곱하기(v3).S셀프_절대값()
	S절대값(값 I정밀수) V정밀수
	S더하기(값1 I정밀수, 값2 I정밀수) V정밀수
	S빼기(값1 I정밀수, 값2 I정밀수) V정밀수
	S곱하기(값1 I정밀수, 값2 I정밀수) V정밀수
	S나누기(값1 I정밀수, 값2 I정밀수) (V정밀수, error)
	S역수(값 I정밀수) V정밀수
	S반대부호값(값 I정밀수) V정밀수
	
	S셀프_절대값() V정밀수
	S셀프_더하기(값 I정밀수) V정밀수
	S셀프_빼기(값 I정밀수) V정밀수
	S셀프_곱하기(값 I정밀수) V정밀수
	S셀프_나누기(값 I정밀수) (V정밀수, error)
	S셀프_역수() V정밀수
	S셀프_반대부호값() V정밀수
}

// 통화
type P통화 string

const (
	KRW P통화 = "KRW"
	USD     = "USD"
	CNY     = "CNY"
	EUR     = "EUR"
)

type I통화 interface {
	G종류() P통화
	G값() *big.Rat
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
	S종류(종류 P통화)
	S값(금액 float64)
	S값Big(금액 *big.Rat)

	// 연산 함수
	// S로 시작하는 연산 함수는 '메소드 연속 호출'(메소드 체이닝) 기법을 사용하기 위해서 
	//   자기 자신을 결과값으로 설정하고 자기 자신을 반환한다.
	// 예 : v통화.S빼기(v1, v2).S셀프_곱하기(v3).S셀프_절대값()
	S절대값(값 I통화) V통화
	S더하기(값1 I통화, 값2 I통화) V통화
	S빼기(값1 I통화, 값2 I통화) V통화
	S곱하기(값1 I통화, 값2 I통화) V통화
	S나누기(값1 I통화, 값2 I통화) V통화
	S역수(값 I통화) V통화
	S반대부호값(값 I통화) V통화
	
	S셀프_절대값() V통화
	S셀프_더하기(값 I통화) V통화
	S셀프_빼기(값 I통화) V통화
	S셀프_곱하기(값 I통화) V통화
	S셀프_나누기(값 I통화) V통화
	S셀프_역수() V통화
	S셀프_반대부호값() V통화
}

/**************************************************
*                복합 상수형 
***************************************************
*  내부 타입은 변하지만, 값 자체는 상수형임.
*
*  여러가지 타입의 데이터를 주고 받을 때,
*  가변형의 슬라이스나 맵으로 주고 받으면,
*  한꺼번에 주고 받을 수 있어서 관리도 쉽고,
*  내부값을 바꿀 수 없기 때문에 데이터 공유로 인한 문제도 
*  사전에 예방할 수 있다.
***************************************************/
type C복합_상수형 interface {
	I상수형
	G값() I상수형
	G형식명() string
	G참거짓() (bool, error)
	G정수() (int64, error)
	G부호없는_정수() (uint64, error)
	G실수() (float64, error)
	G문자열() (string, error)
	G시점() (time.Time, error)
	G큰정수() (*big.Int, error)
	G정밀수() (*big.Rat, error)
	G통화() (C통화, error)
}