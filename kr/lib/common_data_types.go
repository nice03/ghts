// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	//"math/rand"
	//"reflect"
	"sync/atomic"
	"time"
)

// 종목
type sC종목 struct {	
	코드 string
	명칭 string

	출력_문자열 string
}

func (s *sC종목) G상수형임()       {}
func (s *sC종목) G코드() string { return s.코드 }
func (s *sC종목) G명칭() string { return s.명칭 }
func (s *sC종목) String() string { return s.코드 + " " + s.명칭 }

// 전략
type sC전략 struct { 코드 string }

func (s *sC전략) G상수형임()       {}
func (s *sC전략) G코드() string { return s.코드 }
func (s *sC전략) String() string { return s.코드 }

// 포트폴리오 변동내역
type sC포트폴리오_변동내역 struct {
	키 string
	종목 C종목
	전략 C전략	
	롱포지션_변동수량 int64
	숏포지션_변동수량 int64
	시점 time.Time
}

func (s *sC포트폴리오_변동내역) G키() string {
	if s.키 == "" {
		s.키 = s.종목.G코드() + "_" + s.전략.G코드()
	}

	return s.키
}
func (s *sC포트폴리오_변동내역) G전략() C전략 { return s.전략 }
func (s *sC포트폴리오_변동내역) G종목() C종목 { return s.종목 }
func (s *sC포트폴리오_변동내역) G롱포지션_변동수량() int64 { return s.롱포지션_변동수량 }
func (s *sC포트폴리오_변동내역) G숏포지션_변동수량() int64 { return s.숏포지션_변동수량 }
func (s *sC포트폴리오_변동내역) G시점() time.Time { return s.시점 }
func (s *sC포트폴리오_변동내역) String() string {
	문자열 := F포맷된_문자열(
		"전략 %v, 종목 %v, 롱포지션 변동수량 %v, 숏포지션 변동수량 %v, 시점 %v",
		s.G전략().G코드(), s.G종목().G코드(),
		s.G롱포지션_변동수량(), s.G숏포지션_변동수량(),
		s.G시점().Format(P시점_형식))

	return 문자열
}

// 포트폴리오 구성요소
type s포트폴리오_구성요소 struct {
	키 string	// 종목코드 + "_" + 전략코드
	종목 C종목
	전략 C전략
	롱포지션_수량 int64
	숏포지션_수량 int64
}

func (s *s포트폴리오_구성요소) G키() string { return s.키 }
func (s *s포트폴리오_구성요소) G종목() C종목 { return s.종목 }
func (s *s포트폴리오_구성요소) G종목코드() string { return s.종목.G코드() }
func (s *s포트폴리오_구성요소) G전략() C전략 { return s.전략 }
func (s *s포트폴리오_구성요소) G전략코드() string { return s.전략.G코드() }
func (s *s포트폴리오_구성요소) G롱포지션_수량() int64 {
	return atomic.LoadInt64(&s.롱포지션_수량)
}
func (s *s포트폴리오_구성요소) G숏포지션_수량() int64 {
	return atomic.LoadInt64(&s.숏포지션_수량)
}
func (s *s포트폴리오_구성요소) G순_수량() int64 {
	return s.G롱포지션_수량() - s.G숏포지션_수량()
}
func (s *s포트폴리오_구성요소) G총_수량() int64 {
	return s.G롱포지션_수량() + s.G숏포지션_수량()
}
func (s *s포트폴리오_구성요소) String() string {
	문자열 := F포맷된_문자열(
		"전략코드 %v, 종목코드 %v, 롱포지션_수량 %v, 숏포지션_수량 %v",
		s.G전략코드(), s.G종목코드(),
		s.G롱포지션_수량(), s.G숏포지션_수량())

	return 문자열
}

type sC포트폴리오_구성요소 struct { *s포트폴리오_구성요소 }

func (s *sC포트폴리오_구성요소) G상수형임() {}
func (s *sC포트폴리오_구성요소) G변수형() V포트폴리오_구성요소 {
	return NV포트폴리오_구성요소(
			s.G종목(), s.G전략(), s.G롱포지션_수량(), s.G숏포지션_수량())
}

type sV포트폴리오_구성요소 struct { *s포트폴리오_구성요소 }

func (s *sV포트폴리오_구성요소) G변수형임() {}
func (s *sV포트폴리오_구성요소) G상수형() C포트폴리오_구성요소 {
	return NC포트폴리오_구성요소(s.G종목(), s.G전략(), s.G롱포지션_수량(), s.G숏포지션_수량())
}
func (s *sV포트폴리오_구성요소) S변동(변동내역 C포트폴리오_변동내역) error {
	if s.G키() != 변동내역.G키() {
		에러 := F에러_생성("키 불일치. %s, %s", s.G키(), 변동내역.G키())

		F문자열_출력(에러.Error())
		return 에러
	}

	if 변동내역.G롱포지션_변동수량() != 0 {
		atomic.AddInt64(
				&(s.s포트폴리오_구성요소.롱포지션_수량),
				변동내역.G롱포지션_변동수량())
	}

	if 변동내역.G숏포지션_변동수량() != 0 {
		atomic.AddInt64(
				&(s.s포트폴리오_구성요소.숏포지션_수량),
				변동내역.G숏포지션_변동수량())
	}

	return nil
}

var 단가_도우미 = F최근_단가

type sC종목별_포트폴리오 struct {
	종목    C종목
	롱포지션_수량 int64
	숏포지션_수량 int64
}

func (s *sC종목별_포트폴리오) G상수형임() {}
func (s *sC종목별_포트폴리오) G종목() C종목 { return s.종목 }
func (s *sC종목별_포트폴리오) G롱포지션_수량() int64 { return s.롱포지션_수량 }
func (s *sC종목별_포트폴리오) G숏포지션_수량() int64 { return s.숏포지션_수량 }
func (s *sC종목별_포트폴리오) G순_수량() int64 { return s.롱포지션_수량 - s.숏포지션_수량 }
func (s *sC종목별_포트폴리오) G총_수량() int64 { return s.롱포지션_수량 + s.숏포지션_수량 }
func (s *sC종목별_포트폴리오) G롱포지션_금액() C통화 {
	v단가 := 단가_도우미(s.종목).G변수형()
	
	return v단가.S곱하기(s.G롱포지션_수량()).G상수형()
}
func (s *sC종목별_포트폴리오) G숏포지션_금액() C통화 {
	v단가 := 단가_도우미(s.종목).G변수형()

	return v단가.S곱하기(s.G숏포지션_수량()).G상수형()
}
func (s *sC종목별_포트폴리오) G순_금액() C통화 {
	v단가 := 단가_도우미(s.종목).G변수형()
	
	return v단가.S곱하기(s.G순_수량()).G상수형()
}
func (s *sC종목별_포트폴리오) G총_금액() C통화 {
	v단가 := 단가_도우미(s.종목).G변수형()
	
	return v단가.S곱하기(s.G총_수량()).G상수형()
}
func (s *sC종목별_포트폴리오) String() string {
	return F포맷된_문자열("종목 %s, " +
				"롱포지션 수량 %v, 숏포지션 수량 %v, 순수량 %v, 총수량 %v" +
				"롱포지션 금액 %v, 숏포지션 금액 %v, 순금액 %v, 총금액 %v",
				s.G종목().G코드(),
				s.G롱포지션_수량(), s.G숏포지션_수량(), s.G순_수량(), s.G총_수량(),
				s.G롱포지션_금액(), s.G숏포지션_금액(), s.G순_금액(), s.G총_금액())
}