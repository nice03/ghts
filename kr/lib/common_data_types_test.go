// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"math/rand"
	"testing"
	//"time"
)

func TestC종목(테스트 *testing.T) {
	테스트.Parallel()
	
	종목 := NC종목("코드", "명칭")
	
	_, 상수형임 := 종목.(I상수형)
	F참인지_확인(테스트, 상수형임, "NC종목() 결과값이 상수형이 아님.")

	F같은값_확인(테스트, 종목.G코드(), "코드")
	F같은값_확인(테스트, 종목.G명칭(), "명칭")
	F같은값_확인(테스트, 종목.String(), "코드 명칭")
}

func TestC전략	(테스트 *testing.T) {
	테스트.Parallel()
	
	var c C전략 = NC전략("코드")
	
	_, ok := c.(*sC전략)
	
	F참인지_확인(테스트, ok)
	F같은값_확인(테스트, c.G코드(), "코드")	
}

func TestC포트폴리오_변동내역(테스트 *testing.T) {
	테스트.Parallel()
	
	종목 := NC종목(F임의_문자열(10), F임의_문자열(10))
	전략 := NC전략(F임의_문자열(10))
	롱포지션_변동수량 := rand.Int63()
	숏포지션_변동수량 := rand.Int63()
	
	c := NC포트폴리오_변동내역(종목, 전략, 롱포지션_변동수량, 숏포지션_변동수량)
	
	_, ok := c.(*sC포트폴리오_변동내역)
	
	F참인지_확인(테스트, ok)
	F같은값_확인(테스트, c.G종목().G코드(), 종목.G코드())
	F같은값_확인(테스트, c.G종목().G명칭(), 종목.G명칭())
	F같은값_확인(테스트, c.G전략().G코드(), 전략.G코드())
	F같은값_확인(테스트, c.G롱포지션_변동수량(), 롱포지션_변동수량)
	F같은값_확인(테스트, c.G숏포지션_변동수량(), 숏포지션_변동수량)		
}

func TestN포트폴리오_구성요소(테스트 *testing.T) {
	테스트.Parallel()
	
	F참인지_확인(테스트, n포트폴리오_구성요소(nil, nil, 1, 2) == nil)
	
	종목 := NC종목(F임의_문자열(10), F임의_문자열(10))
	F거짓인지_확인(테스트, n포트폴리오_구성요소(종목, nil, 1, 2) == nil)
	F같은값_확인(테스트, n포트폴리오_구성요소(종목, nil, 1, 2).G키(), 종목.G코드() + "_")
	
	전략 := NC전략(F임의_문자열(10))
	
	s := n포트폴리오_구성요소(종목, 전략, 1, 2)
	
	F같은값_확인(테스트, s.G키(), 종목.G코드() + "_" + 전략.G코드())
	F같은값_확인(테스트, s.종목.G코드(), 종목.G코드())
	F같은값_확인(테스트, s.종목.G명칭(), 종목.G명칭())
	F같은값_확인(테스트, s.전략.G코드(), 전략.G코드())
	F같은값_확인(테스트, s.G롱포지션_수량(), 1)
	F같은값_확인(테스트, s.G숏포지션_수량(), 2)
	F같은값_확인(테스트, s.G순_수량(), -1)
	F같은값_확인(테스트, s.G총_수량(), 3)	
}

func TestC종목별_포트폴리오(테스트 *testing.T) {
	//테스트.Parallel()
	// 단가 도우미를 바꾸기 때문에 다른 테스트에 영향을 미친다.
	
	var 단가_도우미_백업 = 단가_도우미
	
	단가_도우미 = func (종목 C종목) C통화 { return NC통화(KRW, 100) }
	defer func() { 단가_도우미 = 단가_도우미_백업 }()

	종목 := NC종목(F임의_문자열(10), F임의_문자열(10))
	롱포지션_수량 := rand.Int63()
	숏포지션_수량 := rand.Int63()
	
	c := NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
	
	_, ok := c.(I상수형)
	F참인지_확인(테스트, ok)
	
	_, ok = c.(*sC종목별_포트폴리오)
	F참인지_확인(테스트, ok)
	
	F같은값_확인(테스트, c.G종목().G코드(), 종목.G코드())
	F같은값_확인(테스트, c.G롱포지션_수량(), 롱포지션_수량)
	F같은값_확인(테스트, c.G숏포지션_수량(), 숏포지션_수량)
	F같은값_확인(테스트, c.G순_수량(), 롱포지션_수량 - 숏포지션_수량)
	F같은값_확인(테스트, c.G총_수량(), 롱포지션_수량 + 숏포지션_수량)
	F같은값_확인(테스트, c.G롱포지션_금액(), 단가_도우미(종목).G변수형().S곱하기(롱포지션_수량))
	F같은값_확인(테스트, c.G숏포지션_금액(), 단가_도우미(종목).G변수형().S곱하기(숏포지션_수량))
	F같은값_확인(테스트, c.G순_금액(), 단가_도우미(종목).G변수형().S곱하기(c.G순_수량()))
	F같은값_확인(테스트, c.G총_금액(), 단가_도우미(종목).G변수형().S곱하기(c.G총_수량()))
}

func TestC전략별_포트폴리오(테스트 *testing.T) {	
	//테스트.Parallel()
	// 단가 도우미를 바꾸기 때문에 다른 테스트에 영향을 미친다.
	
	var 단가_도우미_백업 = 단가_도우미
	
	단가_도우미 = func (종목 C종목) C통화 { return NC통화(KRW, 100) }
	defer func() { 단가_도우미 = 단가_도우미_백업 }()
	
	전략 := NC전략("전략코드")
	종목별_포트폴리오_모음 := make([]C종목별_포트폴리오, 0)
	
	for 반복횟수 := 0 ; 반복횟수 < 10 ; 반복횟수++ {
		종목 := NC종목(F임의_문자열(10), F임의_문자열(10))
		롱포지션_수량 := rand.Int63()
		숏포지션_수량 := rand.Int63()
		
		종목별_포트폴리오 := NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
		종목별_포트폴리오_모음 = append(종목별_포트폴리오_모음, 종목별_포트폴리오)
	}
	
	c := NC전략별_포트폴리오(전략, 종목별_포트폴리오_모음)
	
	_, ok := c.(I상수형)
	F참인지_확인(테스트, ok)
	
	_, ok = c.(*sC전략별_포트폴리오)
	F참인지_확인(테스트, ok)
	
	F같은값_확인(테스트, c.G전략().G코드(), 전략.G코드())
	
	종목_모음 := c.G종목_모음()
	
	for _, 종목 := range 종목_모음 {
		일치 := false
		
		for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
			if 종목 == nil {
				F문자열_출력("nil 종목")
				continue
			}
			
			if 종목별_포트폴리오 == nil {
				F문자열_출력("nil 종목별_포트폴리오")
				continue
			}
			
			if 종목별_포트폴리오.G종목() == nil {
				F문자열_출력("nil 종목별_포트폴리오.G종목()")
				continue
			}
			
			if 종목.G코드() == 종목별_포트폴리오.G종목().G코드() {
				일치 = true
				break
			}
		}
		
		F참인지_확인(테스트, 일치)
	}
	
	for _, 원본 := range 종목별_포트폴리오_모음 {
		종목별_포트폴리오 := c.G종목별_포트폴리오(원본.G종목())
		
		F같은값_확인(테스트, 종목별_포트폴리오.G종목().G코드(), 원본.G종목().G코드())
		F같은값_확인(테스트, 종목별_포트폴리오.G롱포지션_수량(), 원본.G롱포지션_수량())
		F같은값_확인(테스트, 종목별_포트폴리오.G숏포지션_수량(), 원본.G숏포지션_수량())
	}
	
	종목별_포트폴리오_모음_복사본 := c.G종목별_포트폴리오_모음()
	
	F같은값_확인(테스트, len(종목별_포트폴리오_모음_복사본), len(종목별_포트폴리오_모음))
	
	for _, 원본 := range 종목별_포트폴리오_모음 {
		일치 := false
		
		for _, 복사본 := range 종목별_포트폴리오_모음_복사본 {
			if 원본.G종목().G코드() == 복사본.G종목().G코드() &&
				원본.G롱포지션_수량() == 복사본.G롱포지션_수량() &&
				원본.G숏포지션_수량() == 복사본.G숏포지션_수량() {
				일치 = true
				
				break
			}
		}
		
		F참인지_확인(테스트, 일치)
	}

	롱포지션_금액_원본 := NV통화(KRW, 0.0)
	숏포지션_금액_원본 := NV통화(KRW, 0.0)
	순_금액_원본 := NV통화(KRW, 0.0)
	총_금액_원본 := NV통화(KRW, 0.0)
	
	for _, 원본 := range 종목별_포트폴리오_모음 {
		롱포지션_금액_원본.S더하기(원본.G롱포지션_금액())
		숏포지션_금액_원본.S더하기(원본.G숏포지션_금액())
		순_금액_원본.S더하기(원본.G순_금액())
		총_금액_원본.S더하기(원본.G총_금액())
	}
	
	F같은값_확인(테스트, c.G롱포지션_금액(), 롱포지션_금액_원본)
	F같은값_확인(테스트, c.G숏포지션_금액(), 숏포지션_금액_원본)
	F같은값_확인(테스트, c.G순_금액(), 순_금액_원본)
	F같은값_확인(테스트, c.G총_금액(), 총_금액_원본)
}

func TestV포트폴리오(테스트 *testing.T) {
	//테스트.Parallel()
	// 단가 도우미를 바꾸기 때문에 다른 테스트에 영향을 미친다.
	
	var 단가_도우미_백업 = 단가_도우미
	
	단가_도우미 = func (종목 C종목) C통화 { return NC통화(KRW, 100) }
	defer func() { 단가_도우미 = 단가_도우미_백업 }()
	
	
	
	자본 := NC통화(KRW, 1000000)
	포트폴리오 := NV포트폴리오(자본)
	
	// 간단한 테스트부터
	원래값_모음 := make([](map[string]I가변형), 0)
	수량 := 10
	
	const (
		p종목 string = "종목"
		p전략 string = "전략"
		p롱포지션_수량 string = "롱포지션_수량"
		p숏포지션_수량 string = "숏포지션_수량"
	)
	
	for 인덱스 := 0 ; 인덱스 < 수량 ; 인덱스++ {
		인덱스_문자열 := F문자열(인덱스)
		
		종목 := NC종목("코드"+인덱스_문자열, "명칭"+인덱스_문자열)
		전략 := NC전략("코드"+인덱스_문자열)
		롱포지션_수량 := int64(인덱스)
		숏포지션_수량 := int64(인덱스)
		
		변동내역 := NC포트폴리오_변동내역(종목, 전략, 롱포지션_수량, 숏포지션_수량)
		포트폴리오.S변동(변동내역)
		
		원래값 := make(map[string]I가변형)
		원래값[p종목] = 종목
		원래값[p전략] = 전략
		원래값[p롱포지션_수량] = 롱포지션_수량
		원래값[p숏포지션_수량] = 숏포지션_수량
		원래값_모음 = append(원래값_모음, 원래값)
	}
	
	// G자본() 테스트
	F같은값_확인(테스트, 포트폴리오.G자본(), 자본)

	// G종목_모음() 테스트
	종목_모음 := 포트폴리오.G종목_모음()
	F같은값_확인(테스트, len(종목_모음), 수량)

	for _, 종목 := range 종목_모음 {
		일치 := false
		for _, 원래값 := range 원래값_모음 {
			종목_원래값 := 원래값[p종목]
			
			if 종목.G코드() == 종목_원래값.(C종목).G코드() {
				일치 = true
				
				break
			}
		}
		
		F참인지_확인(테스트, 일치)
	}
	
	// G종목별_포트폴리오() 테스트
	for _, 원래값 := range 원래값_모음 {
		종목 := 원래값[p종목].(C종목)
		롱포지션_수량 := 원래값[p롱포지션_수량].(int64)
		숏포지션_수량 := 원래값[p숏포지션_수량].(int64)
		종목별_포트폴리오_원래값 := NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
	
		종목별_포트폴리오 := 포트폴리오.G종목별_포트폴리오(종목)
		
		F같은값_확인(테스트, 
					종목별_포트폴리오.G롱포지션_수량(), 
					종목별_포트폴리오_원래값.G롱포지션_수량())
					
		F같은값_확인(테스트, 
					종목별_포트폴리오.G숏포지션_수량(), 
					종목별_포트폴리오_원래값.G숏포지션_수량())
	}

	// G종목별_포트폴리오_모음() 테스트
	종목별_포트폴리오_모음 := 포트폴리오.G종목별_포트폴리오_모음()
	
	F같은값_확인(테스트, len(종목별_포트폴리오_모음), len(원래값_모음))
	
	for _, 원래값 := range 원래값_모음 {
		종목 := 원래값[p종목].(C종목)
		롱포지션_수량 := 원래값[p롱포지션_수량].(int64)
		숏포지션_수량 := 원래값[p숏포지션_수량].(int64)
		종목별_포트폴리오_원래값 := NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
		
		찾았음 := false
		for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
			if 종목별_포트폴리오.G종목().G코드() == 종목별_포트폴리오_원래값.G종목().G코드() {
				찾았음 = true
				
				F같은값_확인(테스트, 
					종목별_포트폴리오.G롱포지션_수량(), 
					종목별_포트폴리오_원래값.G롱포지션_수량())
					
				F같은값_확인(테스트, 
					종목별_포트폴리오.G숏포지션_수량(), 
					종목별_포트폴리오_원래값.G숏포지션_수량())
			}
		}
		
		F참인지_확인(테스트, 찾았음)
	}
	
	// G전략_모음() 테스트
	전략_모음 := 포트폴리오.G전략_모음()
	F같은값_확인(테스트, len(전략_모음), 수량)

	for _, 전략 := range 전략_모음 {
		일치 := false
		for _, 원래값 := range 원래값_모음 {
			전략_원래값 := 원래값[p전략].(C전략)
			
			if 전략.G코드() == 전략_원래값.G코드() {
				일치 = true
				
				break
			}			
		}
		
		F참인지_확인(테스트, 일치)
	}
	
	// G전략별_포트폴리오() 테스트
	for _, 원래값 := range 원래값_모음 {
		종목 := 원래값[p종목].(C종목)
		롱포지션_수량 := 원래값[p롱포지션_수량].(int64)
		숏포지션_수량 := 원래값[p숏포지션_수량].(int64)
		종목별_포트폴리오 := NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
		
		종목별_포트폴리오_모음 := make([]C종목별_포트폴리오, 0)
		종목별_포트폴리오_모음 = append(종목별_포트폴리오_모음, 종목별_포트폴리오)
		
		전략 := 원래값[p전략].(C전략)
		전략별_포트폴리오_원래값 := NC전략별_포트폴리오(전략, 종목별_포트폴리오_모음)
	
		전략별_포트폴리오 := 포트폴리오.G전략별_포트폴리오(전략)
		
		F같은값_확인(테스트, 
					전략별_포트폴리오.G롱포지션_금액(), 
					전략별_포트폴리오_원래값.G롱포지션_금액())
					
		F같은값_확인(테스트, 
					전략별_포트폴리오.G숏포지션_금액(), 
					전략별_포트폴리오_원래값.G숏포지션_금액())
	}
	
	// G전략별_포트폴리오_모음() 테스트
	전략별_포트폴리오_모음 := 포트폴리오.G전략별_포트폴리오_모음()
	
	F같은값_확인(테스트, len(전략별_포트폴리오_모음), len(원래값_모음))
	
	전략별_포트폴리오_모음_원래값 := make([]C전략별_포트폴리오, 0)
	for _, 원래값 := range 원래값_모음 {
		종목 := 원래값[p종목].(C종목)
		롱포지션_수량 := 원래값[p롱포지션_수량].(int64)
		숏포지션_수량 := 원래값[p숏포지션_수량].(int64)
		종목별_포트폴리오 := NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
		
		종목별_포트폴리오_모음 := make([]C종목별_포트폴리오, 0)
		종목별_포트폴리오_모음 = append(종목별_포트폴리오_모음, 종목별_포트폴리오)
		
		전략 := 원래값[p전략].(C전략)
		전략별_포트폴리오_원래값 := NC전략별_포트폴리오(전략, 종목별_포트폴리오_모음)
		
		전략별_포트폴리오_모음_원래값 = append(전략별_포트폴리오_모음_원래값, 전략별_포트폴리오_원래값)
	}
	
	for _, 전략별_포트폴리오_원래값 := range 전략별_포트폴리오_모음_원래값 {		
		찾았음 := false
		for _, 전략별_포트폴리오 := range 전략별_포트폴리오_모음 {
			if 전략별_포트폴리오.G전략().G코드() == 전략별_포트폴리오_원래값.G전략().G코드() {
				찾았음 = true
				
				F같은값_확인(테스트, 
					전략별_포트폴리오.G롱포지션_금액(), 
					전략별_포트폴리오_원래값.G롱포지션_금액())
					
				F같은값_확인(테스트, 
					전략별_포트폴리오.G숏포지션_금액(), 
					전략별_포트폴리오_원래값.G숏포지션_금액())
			}
		}
		
		F참인지_확인(테스트, 찾았음)
	}
	
	// G롱포지션_금액() 테스트
	롱포지션_금액_원래값 := NV통화(KRW, 0)
	숏포지션_금액_원래값 := NV통화(KRW, 0)
	순_금액_원래값 := NV통화(KRW, 0)
	총_금액_원래값 := NV통화(KRW, 0)
	
	for _, 원래값 := range 원래값_모음 {
		종목 := 원래값[p종목].(C종목)
		롱포지션_수량 := 원래값[p롱포지션_수량].(int64)
		숏포지션_수량 := 원래값[p숏포지션_수량].(int64)
		단가 := 단가_도우미(종목)
		
		롱포지션_금액_원래값.S더하기(단가.G변수형().S곱하기(롱포지션_수량))
		숏포지션_금액_원래값.S더하기(단가.G변수형().S곱하기(숏포지션_수량))
		순_금액_원래값.S더하기(단가.G변수형().S곱하기(롱포지션_수량 - 숏포지션_수량))
		총_금액_원래값.S더하기(단가.G변수형().S곱하기(롱포지션_수량 + 숏포지션_수량))
	}
	
	F같은값_확인(테스트, 포트폴리오.G롱포지션_금액(), 롱포지션_금액_원래값)
	F같은값_확인(테스트, 포트폴리오.G숏포지션_금액(), 숏포지션_금액_원래값)
	F같은값_확인(테스트, 포트폴리오.G순_금액(), 순_금액_원래값)
	F같은값_확인(테스트, 포트폴리오.G총_금액(), 총_금액_원래값)
}