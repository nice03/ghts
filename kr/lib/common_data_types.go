// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	//"math/rand"
	//"reflect"
	"sync"
	"time"
)

// 종목
type sC종목 struct {	
	코드 string
	명칭 string
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
func (s *s포트폴리오_구성요소) G롱포지션_수량() int64 { return s.롱포지션_수량 }
func (s *s포트폴리오_구성요소) G숏포지션_수량() int64 { return s.숏포지션_수량 }
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

type sV포트폴리오_구성요소 struct {
	잠금 sync.RWMutex
	*s포트폴리오_구성요소
}

func (s *sV포트폴리오_구성요소) G변수형임() {}
func (s *sV포트폴리오_구성요소) G상수형() C포트폴리오_구성요소 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return NC포트폴리오_구성요소(
				s.G종목(), 
				s.G전략(), 
				s.G롱포지션_수량(), 
				s.G숏포지션_수량())
}
func (s *sV포트폴리오_구성요소) G롱포지션_수량() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return s.s포트폴리오_구성요소.G롱포지션_수량()
}
func (s *sV포트폴리오_구성요소) G숏포지션_수량() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return s.s포트폴리오_구성요소.G숏포지션_수량()
}
func (s *sV포트폴리오_구성요소) G순_수량() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return s.s포트폴리오_구성요소.G순_수량()
}
func (s *sV포트폴리오_구성요소) G총_수량() int64 {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return s.s포트폴리오_구성요소.G총_수량()
}
func (s *sV포트폴리오_구성요소) String() string {
	s.잠금.RLock()
	defer s.잠금.RUnlock()
	
	return s.s포트폴리오_구성요소.String()
}

func (s *sV포트폴리오_구성요소) S변동(변동내역 C포트폴리오_변동내역) error {
	s.잠금.Lock()
	defer s.잠금.Unlock()
	
	if s.G키() != 변동내역.G키() {
		에러 := F에러_생성("키 불일치. %s, %s", s.G키(), 변동내역.G키())

		F문자열_출력(에러.Error())
		return 에러
	}

	if 변동내역.G롱포지션_변동수량() != 0 {
		s.s포트폴리오_구성요소.롱포지션_수량 += 변동내역.G롱포지션_변동수량()
	}

	if 변동내역.G숏포지션_변동수량() != 0 {
		s.s포트폴리오_구성요소.숏포지션_수량 += 변동내역.G숏포지션_변동수량()
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

type sC전략별_포트폴리오 struct {
	전략 C전략
	저장소 map[string]C종목별_포트폴리오
}

func (s *sC전략별_포트폴리오) G상수형임() {}
func (s *sC전략별_포트폴리오) G전략() C전략 { return s.전략 }
func (s *sC전략별_포트폴리오) G종목_모음() []C종목 {
	종목_모음 := make([]C종목, len(s.저장소))

	인덱스 := 0
	for _, 값 := range s.저장소 {
		종목_모음[인덱스] = 값.(C종목별_포트폴리오).G종목()
		인덱스++
	}

	return 종목_모음
}
func (s *sC전략별_포트폴리오) G종목별_포트폴리오(종목 C종목) C종목별_포트폴리오 {
	return s.저장소[종목.G코드()]
}
func (s *sC전략별_포트폴리오) G종목별_포트폴리오_모음() []C종목별_포트폴리오 {
	종목별_포트폴리오_모음 := make([]C종목별_포트폴리오, len(s.저장소))

	인덱스 := 0
	for _, 값 := range s.저장소 {
		종목별_포트폴리오_모음[인덱스] = 값.(C종목별_포트폴리오)
		인덱스++
	}

	return 종목별_포트폴리오_모음
}
func (s *sC전략별_포트폴리오) G롱포지션_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	금액 := NV통화(종목별_포트폴리오_모음[0].G롱포지션_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		금액.S더하기(종목별_포트폴리오.G롱포지션_금액())
	}

	return 금액.G상수형()
}
func (s *sC전략별_포트폴리오) G숏포지션_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	금액 := NV통화(종목별_포트폴리오_모음[0].G숏포지션_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		금액.S더하기(종목별_포트폴리오.G숏포지션_금액())
	}

	return 금액.G상수형()
}
func (s *sC전략별_포트폴리오) G순_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	금액 := NV통화(종목별_포트폴리오_모음[0].G순_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		금액.S더하기(종목별_포트폴리오.G순_금액())
	}

	return 금액.G상수형()
}
func (s *sC전략별_포트폴리오) G총_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	금액 := NV통화(종목별_포트폴리오_모음[0].G총_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		금액.S더하기(종목별_포트폴리오.G총_금액())
	}

	return 금액.G상수형()
}
func(s *sC전략별_포트폴리오) String() string {
	return F포맷된_문자열("전략 %s, 롱포지션 %v, 숏포지션 %v, 순금액 %v, 총금액 %v",
			s.전략.G코드(), s.G롱포지션_금액(), s.G숏포지션_금액(),
			s.G순_금액(), s.G총_금액())
}

// 포트폴리오
type sV포트폴리오 struct {
	자본 V통화	
	전체_저장소 V문자열키_맵
	종목별_저장소 V문자열키_맵
	전략별_저장소 V문자열키_맵
}

func (s *sV포트폴리오) G자본() C통화 { return s.자본.G상수형() }

func (s *sV포트폴리오) G종목_모음() []C종목 {
	종목코드_모음 := s.종목별_저장소.G키_모음()
	종목_모음 := make([]C종목, len(종목코드_모음))
	
	인덱스 := 0
	for _, 종목코드 := range 종목코드_모음 {
		종목별_맵, _ := s.종목별_저장소.G값(종목코드)	
		구성요소 := 종목별_맵.(V문자열키_맵).G임의값()
		종목_모음[인덱스] = 구성요소.(V포트폴리오_구성요소).G종목()
		
		인덱스++
	}
	
	return 종목_모음
}

func (s *sV포트폴리오) G종목별_포트폴리오(종목 C종목) C종목별_포트폴리오 {
	종목별_맵, 존재함 := s.종목별_저장소.G값(종목.G코드())
	
	if !존재함 {
		return NC종목별_포트폴리오(종목, 0, 0)	
	}
	
	롱포지션_수량, 숏포지션_수량 := int64(0), int64(0)
	
	구성요소_모음 := 종목별_맵.(V문자열키_맵).G값_모음()
	
	for _, v := range 구성요소_모음 {
		롱포지션_수량 = 롱포지션_수량 + v.(V포트폴리오_구성요소).G롱포지션_수량()
		숏포지션_수량 = 숏포지션_수량 + v.(V포트폴리오_구성요소).G숏포지션_수량()
	}

	return NC종목별_포트폴리오(종목, 롱포지션_수량, 숏포지션_수량)
}

func (s *sV포트폴리오) G종목별_포트폴리오_모음() []C종목별_포트폴리오 {
	종목_모음 := s.G종목_모음()
	종목별_포트폴리오_모음 := make([]C종목별_포트폴리오, len(종목_모음))
	
	인덱스 := 0	
	for _, 종목 := range 종목_모음 {
		종목별_포트폴리오_모음[인덱스] = s.G종목별_포트폴리오(종목)
		인덱스++
	}
	
	return 종목별_포트폴리오_모음
}

func (s *sV포트폴리오) G전략_모음() []C전략 {
	전략코드_모음 := s.전략별_저장소.G키_모음()
	전략_모음 := make([]C전략, len(전략코드_모음))
	
	인덱스 := 0
	for _, 전략코드 := range 전략코드_모음 {
		전략별_맵, _ := s.전략별_저장소.G값(전략코드)
		구성요소 := 전략별_맵.(V문자열키_맵).G임의값()
		전략_모음[인덱스] = 구성요소.(V포트폴리오_구성요소).G전략()
		
		인덱스++
	}
	
	return 전략_모음
}

func (s *sV포트폴리오) G전략별_포트폴리오(전략 C전략) C전략별_포트폴리오 {
	전략별_맵, 존재함 := s.전략별_저장소.G값(전략.G코드())
	
	if !존재함 {
		return NC전략별_포트폴리오(전략, make([]C종목별_포트폴리오, 0, 0))
	}

	종목별_포트폴리오_모음 := make([]C종목별_포트폴리오, 전략별_맵.(V문자열키_맵).G수량())
	
	포트폴리오_구성요소_모음 := 전략별_맵.(V문자열키_맵).G값_모음()

	인덱스 := 0
	for _, v := range 포트폴리오_구성요소_모음 {
		구성요소 := v.(V포트폴리오_구성요소)
		
		종목별_포트폴리오 :=
			NC종목별_포트폴리오(
				구성요소.G종목(),
				구성요소.G롱포지션_수량(),
				구성요소.G숏포지션_수량())

		종목별_포트폴리오_모음[인덱스] = 종목별_포트폴리오

		인덱스++
	}

	return NC전략별_포트폴리오(전략, 종목별_포트폴리오_모음)
}

func (s *sV포트폴리오) G전략별_포트폴리오_모음() []C전략별_포트폴리오 {
	전략_모음 := s.G전략_모음()
	전략별_포트폴리오_모음 := make([]C전략별_포트폴리오, len(전략_모음))
	
	인덱스 := 0
	for _, 전략 := range 전략_모음 {
		전략별_포트폴리오_모음[인덱스] = s.G전략별_포트폴리오(전략)
		인덱스++
	}
	
	return 전략별_포트폴리오_모음
}

func (s *sV포트폴리오) G롱포지션_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	롱포지션_금액 := NV통화(종목별_포트폴리오_모음[0].G롱포지션_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		롱포지션_금액.S더하기(종목별_포트폴리오.G롱포지션_금액())
	}

	return 롱포지션_금액.G상수형()
}

func (s *sV포트폴리오) G숏포지션_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	숏포지션_금액 := NV통화(종목별_포트폴리오_모음[0].G숏포지션_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		숏포지션_금액.S더하기(종목별_포트폴리오.G숏포지션_금액())
	}

	return 숏포지션_금액.G상수형()
}

func (s *sV포트폴리오) G순_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	순_금액 := NV통화(종목별_포트폴리오_모음[0].G순_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		순_금액.S더하기(종목별_포트폴리오.G순_금액())
	}

	return 순_금액.G상수형()
}

func (s *sV포트폴리오) G총_금액() C통화 {
	종목별_포트폴리오_모음 := s.G종목별_포트폴리오_모음()

	총_금액 := NV통화(종목별_포트폴리오_모음[0].G총_금액().G종류(), 0)

	for _, 종목별_포트폴리오 := range 종목별_포트폴리오_모음 {
		총_금액.S더하기(종목별_포트폴리오.G총_금액())
	}

	return 총_금액.G상수형()
}

func (s *sV포트폴리오) S변동(변동내역 C포트폴리오_변동내역) {
	F메모("포트폴리오 변동내역 DB에 저장하는 기능 추가할 것.")
	
	// 전체 저장소
	s.전체_저장소.S없으면_추가(
		변동내역.G키(), 
		NV포트폴리오_구성요소(변동내역.G종목(), 변동내역.G전략(), 0, 0))
	구성요소, _ := s.g포트폴리오_구성요소(변동내역.G키())	

	// 종목별 저장소
	s.종목별_저장소.S없으면_추가(구성요소.G종목코드(), NV기본_문자열키_맵())
	종목별_맵, _ := s.종목별_저장소.G값(구성요소.G종목코드())
	종목별_맵.(V문자열키_맵).S없으면_추가(구성요소.G전략코드(), 구성요소)

	// 전략별 저장소
	s.전략별_저장소.S없으면_추가(구성요소.G전략코드(), NV기본_문자열키_맵())
	전략별_맵, _ := s.전략별_저장소.G값(구성요소.G전략코드())
	전략별_맵.(V문자열키_맵).S없으면_추가(구성요소.G종목코드(), 구성요소)
	
	// 변동내역 기록
	구성요소.S변동(변동내역)
}

func (s *sV포트폴리오) g포트폴리오_구성요소(키 string) (V포트폴리오_구성요소, bool) {
	v, 존재함 := s.전체_저장소.G값(키)

	if !존재함 {
		return nil, false
	}

	return v.(V포트폴리오_구성요소), true
}