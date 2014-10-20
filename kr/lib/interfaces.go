// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"math/big"
	"testing/quick"
	"time"
)

type I가변형 interface{}

type I반환값 interface {
	G값() I가변형
	G에러() error
}

type i테스트용_가상_객체 interface {
	테스트용_가상_객체()
}

type I식별코드 interface {
	G식별코드() uint64
}
type I종목코드 interface {
	G종목코드() string
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
	S역수() V실수
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
	S일자_더하기(연, 월, 일 int) V시점
}

// 정밀수
type I정밀수 interface {
	I실수형
	G값() string
	GRat() *big.Rat
	G같음(값 I가변형) bool
	G비교(값 I가변형) int // -1 : 더 작음, 0 : 같음, 1 : 더 큼. -2 : 숫자 아님
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

	S값(값 I가변형) V정밀수
	S반올림(소숫점_이하_자릿수 int) V정밀수
	S절대값() V정밀수
	S더하기(값 I가변형) V정밀수
	S빼기(값 I가변형) V정밀수
	S곱하기(값 I가변형) V정밀수
	S나누기(값 I가변형) V정밀수
	S역수() V정밀수
	S반대부호값() V정밀수
}

type I통화 interface {
	G종류() P통화종류
	G값() C정밀수
	G같음(값 I통화) bool
	G비교(값 I통화) int // -1 : 더 작음, 0 : 같음, 1 : 더 큼, -2 : 비교 불가
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
	//S종류(종류 P통화종류)
	S값(금액 I가변형) V통화

	S절대값() V통화
	S더하기(값 I가변형) V통화
	S빼기(값 I가변형) V통화
	S곱하기(값 I가변형) V통화
	S나누기(값 I가변형) V통화
	S반대부호값() V통화
}

type I환율 interface {
	G원래통화() P통화종류
	G목표통화() P통화종류
	G환율() C정밀수
}

type C매개변수 interface {
	I상수형
	G이름() string
	G값() I가변형
	G숫자형식임() bool
	G문자열형식임() bool
	G시점형식임() bool
	G참거짓형식임() bool
}

// 생성자에서 공유해도 안전한 타입인지 검사하는 가변형
type C안전한_가변형 interface {
	I상수형
	G값() I가변형
}

// Persistent 리스트.
// 순서가 반대인 링크드 리스트. 데이터를 공유로 인한 문제가 없음.
// 원소를 추가하면 새로운 항목이 배열의 첫번째 항목이 됨.
// 추가만 가능하고, 삭제나 변경이 안 되므로,
// 배열의 첫번째 항목이 새로 생성되는 것은 새로운 배열이 생성되는 것과 비슷한 효과를 냄.
// ps패키지(github.com/mndrix/ps) 의 List를 한글화 했음.
type I안전한_배열 interface {
	G비어있음() bool
	G길이() int

	// 반환값을 변수에 저장하지 않으면 추가한 항목은 사라짐.
	// 메소드 체이닝은 가능함. s = S추가(1).S추가(2).S추가(3).....
	S추가(값 I가변형) I안전한_배열
	G슬라이스() []I가변형
}

// Persistent 맵
// 매번 추가, 삭제할 때마다 새로운 맵이 생성 및 할당되므로.
// mutable 데이터를 공유하면서 생기는 문제가 없음.
// ps패키지(github.com/mndrix/ps) 의 Map을 한글화 했음.
type I안전한_맵 interface {
	G비어있음() bool
	G길이() int
	G키_모음() []string
	G존재함(키 string) bool
	G값(키 string) I가변형
	G맵() map[string]I가변형

	// 반환값을 변수에 저장하지 않으면 추가한 항목은 사라짐.
	S값(이름 string, 값 I가변형) I안전한_맵

	// 반환값을 변수에 저장하지 않으면 추가한 항목은 사라짐.
	S삭제(이름 string) I안전한_맵
	I기본_문자열
}

type C종목 interface {
	I상수형
	G코드() string
	G명칭() string
}

//
//	포트폴리오 구성 계획.
//
// 시스템 트레이딩의 목표가 적절한 리스크 내에서 자산증식을 목표로 한다면,
// 트레이딩 시스템은 현실에서의 가격변동을 확률적으로 계산해서,
// 적절한 리스크 범위 이내에서, 자산을 최대한 증식하여서 자산의 스냅샷인 포트폴리오를 키우는
// 프로그램이 된다.
// 즉, 트레이딩 시스템의 궁극적인 목표인 자산증식의 기본 관리 대상이 포트폴리오이다.
// 알파 모듈, 리스크 관리 모듈, 주문모듈등은 포트폴리오 를 키우기 위한 확률 계산 로직일 뿐이다.
// 포트폴리오를 얼마나 정확하고, 빠르게 관리할 수 있는 지가 트레이딩 시스템의 기본 토대가 된다.
//
// 포트폴리오는 대표적인 mutable 데이터.
// 그리고, 여러가지 기준으로 합산이 필요한 데이터.
//
// PLAN A - DB를 이용.
// : RDBMS를 이용하면, 여러가지 기준으로 데이터를 집계하고, 동기화 문제도 해결된다.
//   일정 기간이 지나면 변동내역의 숫자가 많아지면서, 집계 해야할 데이터가 많아서
//   속도가 느려질 수 있으니, 일정 주기(매년? 매달? 매주? 매일?)로 스냅샷을 만들어서
//   변동사항의 숫자를 일정규모 이하로 유지하도록 하자.
//   이것은 재고관리 프로그램을 만들 때 많이 해 봤던 작업이다.
//
// PLAN B - 메모리 상에 각 기준별  스냅샷 유지.
// : DB는 변동사항을 디스크에 기록하고, 시스템 시작할 때 스냅샷을 재구성 하는 용도로만 사용하고,
//   실제 작업은 메모리에 있는 포트폴리오 스냅샷을 이용.
//	 포트폴리오 변동사항은 DB에 기록함과 동시에 메모리 상의 스냅샷을 업데이트 하는 작업을 동시에 해야 함.
//   장점 : 더 빠른 성능.
//	 단점 : 더 많은 개발노력 필요. 동기화 문제. 자료 구분 문제. 메모리 낭비 문제
// 	 이것은 탈정규화를 통해서 성능향상을 추구하는 방식임.
//	 캐쉬 메모리를 효율적으로 사용할려면 RDBMS의 정규화가 답임.
//
// 결론 : 우선은 PLAN A 대로 진행한 후, 성능 상 병목현상이 발생하면 PLAN B를 진행하자.
// PLAN B를 실행하더라도 시스템을 새로 시작할 때마다, DB 기록을 바탕으로 스냅샷을 재구성할려면,
// PLAN A에서 만들어 놓은 로직이 필요하므로, 개발 노력의 낭비가 아니다.
//

/*
type C단순_가격정보 interface {
	I상수형
	G종목코드() C문자열
	G가격() C통화
}

type C가격정보 interface {
	C단순_가격정보
	G시점() C시점
}

type I최신_가격정보_모음 interface {
	G종목_모음() []C종목
	G가격정보(종목 string) C단순_가격정보
	G가격정보_모음() []C단순_가격정보
}


// 포트폴리오 변동사항을 DB에 기록할 포맷.
type C포트폴리오_변동내역 interface {
	G전략코드() C문자열
	G종목코드() C문자열
	G롱포지션_변동수량() C정수
	G숏포지션_변동수량() C정수
	G시점() C시점
}

// 'C포트폴리오_변동내역'의 종목별 집계
type I종목별_포트폴리오 interface {
	G전략코드() C문자열
	G종목코드() C문자열

	G롱포지션_수량() C정수
	G숏포지션_수량() C정수
	G순_수량() C정수
	G총_수량() C정수

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

// 'C포트폴리오_변동내역'의 전략별 집계
type I전략별_포트폴리오 interface {
	G전략코드() C문자열

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

// 'C포트폴리오_변동내역'의 산업군별 집계
type I업종별_포트폴리오 interface {
	G업종코드() C문자열

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화
}

type I전체_포트폴리오 interface {
	G종목_모음() []C종목
	G종목별_포트폴리오_모음() []C종목별_포트폴리오
	G종목별_포트폴리오(종목코드 string) C종목별_포트폴리오

	G전략_모음() []C전략
	G전략별_포트폴리오_모음() []C전략별_포트폴리오
	G전략별_포트폴리오(전략코드 string) C전략별_포트폴리오

	G업종_모음() []C업종
	G업종별_포트폴리오_모음() []C업종별_포트폴리오
	G업종별_포트폴리오(업종코드 string) C업종별_포트폴리오

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화

	G자본() C통화
} */
