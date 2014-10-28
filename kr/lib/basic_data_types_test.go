// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestS반환값(테스트 *testing.T) {
	테스트.Parallel()

	값_모음 := []I가변형{nil, uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1), float64(1), true, false, "test", time.Now(),
		NC부호없는_정수(1), NC정수(1), NC실수(1), NC정밀수(1), NC통화(KRW, 1),
		NC시점(time.Now()), NC문자열("테스트")}

	에러 := F에러_생성("에러")
	for _, 값 := range 값_모음 {
		정상값 := N반환값(값, nil)
		에러값 := N반환값(nil, 에러)

		F같은값_확인(테스트, 정상값.G값(), 값)
		F같은값_확인(테스트, 정상값.G에러(), nil)

		F같은값_확인(테스트, 에러값.G값(), nil)
		F같은값_확인(테스트, 에러값.G에러(), 에러)
	}
}

func testI정수(테스트 *testing.T, 생성자 I가변형) {	
	입력값 := int64(100)
	입력값_백업 := int64(100)

	var i정수 I정수
	var i기본_문자열 I기본_문자열
	var i임의값_생성 I임의값_생성

	// 생성자 테스트
	switch 생성자.(type) {
	case func(int64) C정수:
		생성자_ := 생성자.(func(int64) C정수)
		i정수 = 생성자_(입력값)
	case func(int64) V정수:
		생성자_ := 생성자.(func(int64) V정수)
		i정수 = 생성자_(입력값)
	case func(*sC정수64) V정수:
		생성자_ := 생성자.(func(*sC정수64) V정수)
		i정수 = 생성자_(&sC정수64{&s정수64{입력값}})
	case func(*sV정수64) C정수:
		생성자_ := 생성자.(func(*sV정수64) C정수)
		i정수 = 생성자_(&sV정수64{s정수64: &s정수64{입력값}})
	default:
		F문자열_출력("예상치 못한 생성자 형식. " + reflect.TypeOf(생성자).String())
		테스트.Fail()
	}

	i기본_문자열 = i정수.(I기본_문자열)
	i임의값_생성 = i정수.(I임의값_생성)

	// G값() 테스트
	F같은값_확인(테스트, i정수.G값(), 입력값)
	F같은값_확인(테스트, i정수.G정수(), 입력값)
	F같은값_확인(테스트, i정수.G실수(), 입력값)
	F같은값_확인(테스트, i정수.G정밀수(), 입력값)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatInt(입력값, 10))

	// 입력값 변경 후 독립성 확인.
	입력값 = 입력값 + 1

	F다른값_확인(테스트, 입력값, 입력값_백업)
	F다른값_확인(테스트, i정수.G값(), 입력값)
	F같은값_확인(테스트, i정수.G값(), 입력값_백업)
	F같은값_확인(테스트, i정수.G정수(), 입력값_백업)
	F같은값_확인(테스트, i정수.G실수(), 입력값_백업)
	F같은값_확인(테스트, i정수.G정밀수(), 입력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatInt(입력값_백업, 10))

	// 출력값 변경 후 독립성 확인.
	출력값 := i정수.G값()
	출력값_백업 := 출력값
	F같은값_확인(테스트, 출력값, 출력값_백업)

	출력값 = 출력값 + 1
	F다른값_확인(테스트, 출력값, 출력값_백업)
	F다른값_확인(테스트, i정수.G값(), 출력값)
	F같은값_확인(테스트, i정수.G값(), 출력값_백업)
	F같은값_확인(테스트, i정수.G정수(), 출력값_백업)
	F같은값_확인(테스트, i정수.G실수(), 출력값_백업)
	F같은값_확인(테스트, i정수.G정밀수(), 출력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatInt(출력값_백업, 10))

	// Generate() 테스트
	맵 := make(map[int64]bool)
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))

	양수_수량 := 0
	음수_수량 := 0

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		reflect값 := i임의값_생성.Generate(임의값_생성기, 1)
		값 := reflect값.Interface().(I정수).G값()
		맵[값] = true

		if 값 > 0 {
			양수_수량++
		} else if 값 < 0 {
			음수_수량++
		}
	}

	F참인지_확인(테스트, len(맵) > 80)
	F참인지_확인(테스트, 양수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)
	F참인지_확인(테스트, 음수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)

	// 상수형, 변수형 혼용되지 않는 지 확인
	switch i정수.(type) {
	case C정수:
		c := i정수.(C정수)

		_, ok := c.(*sC정수64)
		F참인지_확인(테스트, ok)

		_, ok = c.(I상수형)
		F참인지_확인(테스트, ok)

		_, ok = c.(I변수형)
		F거짓인지_확인(테스트, ok)
	case V정수:
		v := i정수.(V정수)

		_, ok := v.(*sV정수64)
		F참인지_확인(테스트, ok)

		_, ok = v.(I상수형)
		F거짓인지_확인(테스트, ok)

		_, ok = v.(I변수형)
		F참인지_확인(테스트, ok)
	default:
		F문자열_출력("예상치 못한 경우. %v", reflect.TypeOf(i정수))
	}
}

func TestC정수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI정수(테스트, NC정수)
	testI정수(테스트, (*sC정수64).G변수형)

	c := NC정수(100)
	v := c.G변수형()
	F같은값_확인(테스트, c, v)
}

func TestV정수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI정수(테스트, NV정수)
	testI정수(테스트, (*sV정수64).G상수형)

	v := NV정수(100)
	c := v.G상수형()
	F같은값_확인(테스트, c, v)

	F같은값_확인(테스트, NV정수(100).S값(300), 300)
	F같은값_확인(테스트, NV정수(100).S절대값(), 100)
	F같은값_확인(테스트, NV정수(-100).S절대값(), 100)
	F같은값_확인(테스트, NV정수(100).S더하기(100), 200)
	F같은값_확인(테스트, NV정수(100).S빼기(100), 0)
	F같은값_확인(테스트, NV정수(100).S곱하기(100), 10000)
	F같은값_확인(테스트, NV정수(100).S나누기(100), 1)

	F패닉발생_확인(테스트, NV정수(100).S나누기, 0)
}

func testI부호없는_정수(테스트 *testing.T, 생성자 I가변형) {	
	입력값 := uint64(100)
	입력값_백업 := uint64(100)

	var i부호없는_정수 I부호없는_정수
	var i기본_문자열 I기본_문자열
	var i임의값_생성 I임의값_생성

	// 생성자 테스트
	switch 생성자.(type) {
	case func(uint64) C부호없는_정수:
		생성자_ := 생성자.(func(uint64) C부호없는_정수)
		i부호없는_정수 = 생성자_(입력값)
	case func(uint64) V부호없는_정수:
		생성자_ := 생성자.(func(uint64) V부호없는_정수)
		i부호없는_정수 = 생성자_(입력값)
	case func(*sC부호없는_정수64) V부호없는_정수:
		생성자_ := 생성자.(func(*sC부호없는_정수64) V부호없는_정수)
		i부호없는_정수 = 생성자_(&sC부호없는_정수64{&s부호없는_정수64{입력값}})
	case func(*sV부호없는_정수64) C부호없는_정수:
		생성자_ := 생성자.(func(*sV부호없는_정수64) C부호없는_정수)
		i부호없는_정수 = 생성자_(&sV부호없는_정수64{s부호없는_정수64: &s부호없는_정수64{입력값}})
	default:
		F문자열_출력("예상치 못한 생성자 형식. " + reflect.TypeOf(생성자).String())
		테스트.Fail()
	}

	i기본_문자열 = i부호없는_정수.(I기본_문자열)
	i임의값_생성 = i부호없는_정수.(I임의값_생성)

	// G값() 테스트
	F같은값_확인(테스트, i부호없는_정수.G값(), 입력값)
	F같은값_확인(테스트, i부호없는_정수.G정수(), 입력값)
	F같은값_확인(테스트, i부호없는_정수.G실수(), 입력값)
	F같은값_확인(테스트, i부호없는_정수.G정밀수(), 입력값)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatUint(입력값, 10))

	// 입력값 변경 후 독립성 확인.
	입력값 = 입력값 + 1

	F다른값_확인(테스트, 입력값, 입력값_백업)
	F다른값_확인(테스트, i부호없는_정수.G값(), 입력값)
	F같은값_확인(테스트, i부호없는_정수.G값(), 입력값_백업)
	F같은값_확인(테스트, i부호없는_정수.G정수(), 입력값_백업)
	F같은값_확인(테스트, i부호없는_정수.G실수(), 입력값_백업)
	F같은값_확인(테스트, i부호없는_정수.G정밀수(), 입력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatUint(입력값_백업, 10))

	// 출력값 변경 후 독립성 확인.
	출력값 := i부호없는_정수.G값()
	출력값_백업 := 출력값
	F같은값_확인(테스트, 출력값, 출력값_백업)

	출력값 = 출력값 + 1

	F다른값_확인(테스트, 출력값, 출력값_백업)
	F다른값_확인(테스트, i부호없는_정수.G값(), 출력값)
	F같은값_확인(테스트, i부호없는_정수.G값(), 출력값_백업)
	F같은값_확인(테스트, i부호없는_정수.G정수(), 출력값_백업)
	F같은값_확인(테스트, i부호없는_정수.G실수(), 출력값_백업)
	F같은값_확인(테스트, i부호없는_정수.G정밀수(), 출력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatUint(출력값_백업, 10))

	// Generate() 테스트
	맵 := make(map[uint64]bool)
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		reflect값 := i임의값_생성.Generate(임의값_생성기, 1)
		값 := reflect값.Interface().(I부호없는_정수).G값()
		맵[값] = true
	}

	F참인지_확인(테스트, len(맵) > 80)

	switch i부호없는_정수.(type) {
	case C부호없는_정수:
		c := i부호없는_정수.(C부호없는_정수)

		_, ok := c.(*sC부호없는_정수64)
		F참인지_확인(테스트, ok)

		_, ok = c.(I상수형)
		F참인지_확인(테스트, ok)

		_, ok = c.(I변수형)
		F거짓인지_확인(테스트, ok)
	case V부호없는_정수:
		v := i부호없는_정수.(V부호없는_정수)

		_, ok := v.(*sV부호없는_정수64)
		F참인지_확인(테스트, ok)

		_, ok = v.(I상수형)
		F거짓인지_확인(테스트, ok)

		_, ok = v.(I변수형)
		F참인지_확인(테스트, ok)
	default:
		F문자열_출력("예상치 못한 경우. %v", reflect.TypeOf(i부호없는_정수))
	}
}

func TestC부호없는_정수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI부호없는_정수(테스트, NC부호없는_정수)
	testI부호없는_정수(테스트, (*sC부호없는_정수64).G변수형)

	c := NC부호없는_정수(100)
	v := c.G변수형()
	F같은값_확인(테스트, c, v)
}

func TestV부호없는_정수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI부호없는_정수(테스트, NV부호없는_정수)
	testI부호없는_정수(테스트, (*sV부호없는_정수64).G상수형)

	v := NV부호없는_정수(100)
	c := v.G상수형()
	F같은값_확인(테스트, c, v)

	F같은값_확인(테스트, NV부호없는_정수(100).S값(300), 300)
	F같은값_확인(테스트, NV부호없는_정수(100).S더하기(100), 200)
	F같은값_확인(테스트, NV부호없는_정수(100).S빼기(100), 0)
	F같은값_확인(테스트, NV부호없는_정수(100).S곱하기(100), 10000)
	F같은값_확인(테스트, NV부호없는_정수(100).S나누기(100), 1)

	F문자열_출력_일시정지_시작()
	F패닉발생_확인(테스트, NV부호없는_정수(100).S나누기, 0)
	F문자열_출력_일시정지_종료()
}

func testI실수(테스트 *testing.T, 생성자 I가변형) {
	입력값 := float64(100.1)
	입력값_백업 := float64(100.1)

	var i실수 I실수
	var i기본_문자열 I기본_문자열
	var i임의값_생성 I임의값_생성

	// 생성자 테스트
	switch 생성자.(type) {
	case func(float64) C실수:
		생성자_ := 생성자.(func(float64) C실수)
		i실수 = 생성자_(입력값)
	case func(float64) V실수:
		생성자_ := 생성자.(func(float64) V실수)
		i실수 = 생성자_(입력값)
	case func(*sC실수64) V실수:
		생성자_ := 생성자.(func(*sC실수64) V실수)
		i실수 = 생성자_(&sC실수64{&s실수64{입력값}})
	case func(*sV실수64) C실수:
		생성자_ := 생성자.(func(*sV실수64) C실수)
		i실수 = 생성자_(&sV실수64{s실수64: &s실수64{입력값}})
	default:
		F문자열_출력("예상치 못한 생성자 형식. " + reflect.TypeOf(생성자).String())
		테스트.Fail()
	}

	i기본_문자열 = i실수.(I기본_문자열)
	i임의값_생성 = i실수.(I임의값_생성)

	// G값() 테스트
	F같은값_확인(테스트, i실수.G값(), 입력값)
	F같은값_확인(테스트, i실수.G실수(), 입력값)
	F같은값_확인(테스트, i실수.G정밀수(), 입력값)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatFloat(입력값, 'f', -1, 64))

	// 입력값 변경 후 독립성 확인.
	입력값 = 입력값 + 1

	F다른값_확인(테스트, 입력값, 입력값_백업)
	F다른값_확인(테스트, i실수.G값(), 입력값)
	F같은값_확인(테스트, i실수.G값(), 입력값_백업)
	F같은값_확인(테스트, i실수.G실수(), 입력값_백업)
	F같은값_확인(테스트, i실수.G정밀수(), 입력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatFloat(입력값_백업, 'f', -1, 64))

	// 출력값 변경 후 독립성 확인.
	출력값 := i실수.G값()
	출력값_백업 := 출력값
	F같은값_확인(테스트, 출력값, 출력값_백업)

	출력값 = 출력값 + 1.1
	F다른값_확인(테스트, 출력값, 출력값_백업)
	F다른값_확인(테스트, i실수.G값(), 출력값)
	F같은값_확인(테스트, i실수.G값(), 출력값_백업)
	F같은값_확인(테스트, i실수.G실수(), 출력값_백업)
	F같은값_확인(테스트, i실수.G정밀수(), 출력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatFloat(출력값_백업, 'f', -1, 64))

	// Generate() 테스트
	맵 := make(map[float64]bool)
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))

	양수_수량 := 0
	음수_수량 := 0

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		reflect값 := i임의값_생성.Generate(임의값_생성기, 1)
		값 := reflect값.Interface().(I실수).G값()
		맵[값] = true

		if 값 > 0 {
			양수_수량++
		} else if 값 < 0 {
			음수_수량++
		}
	}

	F참인지_확인(테스트, len(맵) > 80)
	F참인지_확인(테스트, 양수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)
	F참인지_확인(테스트, 음수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)

	// 상수형, 변수형 혼용되지 않는 지 확인
	switch i실수.(type) {
	case C실수:
		c := i실수.(C실수)

		_, ok := c.(*sC실수64)
		F참인지_확인(테스트, ok)

		_, ok = c.(I상수형)
		F참인지_확인(테스트, ok)

		_, ok = c.(I변수형)
		F거짓인지_확인(테스트, ok)
	case V실수:
		v := i실수.(V실수)

		_, ok := v.(*sV실수64)
		F참인지_확인(테스트, ok)

		_, ok = v.(I상수형)
		F거짓인지_확인(테스트, ok)

		_, ok = v.(I변수형)
		F참인지_확인(테스트, ok)
	default:
		F문자열_출력("예상치 못한 경우. %v", reflect.TypeOf(i실수))
	}
}

func TestC실수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI실수(테스트, NC실수)
	testI실수(테스트, (*sC실수64).G변수형)

	c := NC실수(100.1)
	v := c.G변수형()
	F같은값_확인(테스트, c, v)
}

func TestV실수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI실수(테스트, NV실수)
	testI실수(테스트, (*sV실수64).G상수형)

	v := NV실수(100.1)
	c := v.G상수형()
	F같은값_확인(테스트, c, v)

	F같은값_확인(테스트, NV실수(100.0).S값(300.0), 300)
	F같은값_확인(테스트, NV실수(100.0).S절대값(), 100)
	F같은값_확인(테스트, NV실수(-100.0).S절대값(), 100)
	F같은값_확인(테스트, NV실수(100.0).S더하기(100.0), 200)
	F같은값_확인(테스트, NV실수(100.0).S빼기(100.0), 0)
	F같은값_확인(테스트, NV실수(100.0).S곱하기(100.0), 10000)
	F같은값_확인(테스트, NV실수(100.0).S나누기(100.0), 1)
	F같은값_확인(테스트, NV실수(100.0).S역수(), 0.01)

	F문자열_출력_일시정지_시작()
	F참인지_확인(테스트, math.IsInf(NV실수(100.0).S나누기(0.0).G값(), 0))
	F참인지_확인(테스트, math.IsInf(NV실수(0.0).S역수().G값(), 0))
	F문자열_출력_일시정지_종료()
}

func testI시점(테스트 *testing.T, 생성자 I가변형) {
	입력값 := time.Now()
	입력값_백업 := F시점_복사(입력값)
	F같은값_확인(테스트, 입력값, 입력값_백업)

	var i시점, i시점2 I시점
	var i기본_문자열 I기본_문자열
	var i임의값_생성 I임의값_생성

	// 생성자 테스트
	switch 생성자.(type) {
	case func(I가변형) C시점:
		생성자_ := 생성자.(func(I가변형) C시점)
		i시점 = 생성자_(입력값)
		i시점2 = 생성자_(입력값.Format(P시점_형식))
	case func(I가변형) V시점:
		생성자_ := 생성자.(func(I가변형) V시점)
		i시점 = 생성자_(입력값)
		i시점2 = 생성자_(입력값.Format(P시점_형식))
	case func(*sC시점) V시점:
		생성자_ := 생성자.(func(*sC시점) V시점)
		i시점 = 생성자_(&sC시점{s시점: &s시점{입력값}})
		i시점2 = nil
	case func(*sV시점) C시점:
		생성자_ := 생성자.(func(*sV시점) C시점)
		i시점 = 생성자_(&sV시점{s시점: &s시점{입력값}})
		i시점2 = nil
	default:
		F문자열_출력("예상치 못한 생성자 형식. " + reflect.TypeOf(생성자).String())
		테스트.Fail()
	}

	i기본_문자열 = i시점.(I기본_문자열)
	i임의값_생성 = i시점.(I임의값_생성)

	// G값() 테스트
	F같은값_확인(테스트, i시점.G값(), 입력값)
	F같은값_확인(테스트, i기본_문자열.String(), 입력값.Format(P시점_형식))

	// 문자열로 생성된 값도 확인.
	if i시점2 != nil {
		F같은값_확인(테스트, i시점2.G값(), 입력값)
	}

	// 입력값 변경 후 독립성 확인.
	입력값 = 입력값.AddDate(0, 0, 1)

	F다른값_확인(테스트, 입력값, 입력값_백업)
	F다른값_확인(테스트, i시점.G값(), 입력값)
	F같은값_확인(테스트, i시점.G값(), 입력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), 입력값_백업.Format(P시점_형식))

	// 출력값 변경 후 독립성 확인.
	출력값 := i시점.G값()
	출력값_백업 := F시점_복사(출력값)
	F같은값_확인(테스트, 출력값, 출력값_백업)

	출력값 = 출력값.AddDate(0, 0, 1)

	F다른값_확인(테스트, 출력값, 출력값_백업)
	F다른값_확인(테스트, i시점.G값(), 출력값)
	F같은값_확인(테스트, i시점.G값(), 출력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), 출력값_백업.Format(P시점_형식))

	// Generate() 테스트
	맵 := make(map[string]bool)
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))

	미래시점_수량 := 0
	과거시점_수량 := 0

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		reflect값 := i임의값_생성.Generate(임의값_생성기, 1)
		시점 := reflect값.Interface().(I시점).G값()

		맵[시점.Format(P시점_형식)] = true

		현재시점 := time.Now()

		if 시점.After(현재시점) {
			미래시점_수량++
		} else if 시점.Before(현재시점) {
			과거시점_수량++
		}
	}

	F참인지_확인(테스트, len(맵) > 80)
	F참인지_확인(테스트, 미래시점_수량 > 25, "미래시점_수량 %v개, 과거시점_수량 %v개",
		미래시점_수량, 과거시점_수량)
	F참인지_확인(테스트, 과거시점_수량 > 25, "미래시점_수량 %v개, 과거시점_수량 %v개",
		미래시점_수량, 과거시점_수량)

	// 상수형, 변수형 혼용되지 않는 지 확인
	switch i시점.(type) {
	case C시점:
		c := i시점.(C시점)

		_, ok := c.(*sC시점)
		F참인지_확인(테스트, ok)

		_, ok = c.(I상수형)
		F참인지_확인(테스트, ok)

		_, ok = c.(I변수형)
		F거짓인지_확인(테스트, ok)
	case V시점:
		v := i시점.(V시점)

		_, ok := v.(*sV시점)
		F참인지_확인(테스트, ok)

		_, ok = v.(I상수형)
		F거짓인지_확인(테스트, ok)

		_, ok = v.(I변수형)
		F참인지_확인(테스트, ok)
	default:
		F문자열_출력("예상치 못한 경우. %v", reflect.TypeOf(i시점))
	}
}

func TestC시점(테스트 *testing.T) {
	테스트.Parallel()
	
	testI시점(테스트, NC시점)
	testI시점(테스트, (*sC시점).G변수형)

	c := NC시점(time.Now())
	v := c.G변수형()
	F같은값_확인(테스트, c, v)
}

func TestV시점(테스트 *testing.T) {
	테스트.Parallel()
	
	testI시점(테스트, NV시점)
	testI시점(테스트, (*sV시점).G상수형)

	v := NV시점(time.Now())
	c := v.G상수형()
	F같은값_확인(테스트, c, v)

	F같은값_확인(테스트,
		NV시점("2000-01-01").S일자_더하기(1, 1, 1),
		NC시점("2001-02-02"))
}

func TestN정밀수_생성자(테스트 *testing.T) {
	테스트.Parallel()
	
	입력값_모음 := []I가변형{
		uint(100), uint8(100), uint16(100), uint32(100), uint64(100),
		int(100), int8(100), int16(100), int32(100), int64(100),
		NC부호없는_정수(100), NC정수(100)}

	for _, 입력값 := range 입력값_모음 {
		F같은값_확인(테스트, NC정밀수(입력값), 100)
		F같은값_확인(테스트, NV정밀수(입력값), 100)
	}

	입력값_모음 = []I가변형{
		float32(100.025), float64(100.025),
		NC실수(100.025), NC정밀수(100.025), "100.025"}

	for _, 입력값 := range 입력값_모음 {
		F같은값_확인(테스트, NC정밀수(입력값), 100.025)
		F같은값_확인(테스트, NV정밀수(입력값), 100.025)
	}

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_종료()

	입력값_모음 = []I가변형{
		nil, "변환 불가능한 문자열", true, time.Now(), NC통화(KRW, 100)}

	for _, 입력값 := range 입력값_모음 {
		F참인지_확인(테스트, F_nil값임(NC정밀수(입력값)))
		F참인지_확인(테스트, F_nil값임(NV정밀수(입력값)))
	}
}

func testI정밀수(테스트 *testing.T, 생성자 I가변형) {
	입력값 := 100.1
	입력값_백업 := 100.1

	var i정밀수 I정밀수
	var i기본_문자열 I기본_문자열
	var i임의값_생성 I임의값_생성

	// 생성자 테스트
	switch 생성자.(type) {
	case func(I가변형) C정밀수:
		생성자_ := 생성자.(func(I가변형) C정밀수)
		i정밀수 = 생성자_(입력값)
	case func(I가변형) V정밀수:
		생성자_ := 생성자.(func(I가변형) V정밀수)
		i정밀수 = 생성자_(입력값)
	case func(*sC정밀수) V정밀수:
		생성자_ := 생성자.(func(*sC정밀수) V정밀수)
		bigRat, _ := new(big.Rat).SetString(F문자열(입력값))
		i정밀수 = 생성자_(&sC정밀수{&s정밀수{bigRat}})
	case func(*sV정밀수) C정밀수:
		생성자_ := 생성자.(func(*sV정밀수) C정밀수)
		bigRat, _ := new(big.Rat).SetString(F문자열(입력값))
		i정밀수 = 생성자_(&sV정밀수{s정밀수: &s정밀수{bigRat}})
	default:
		F문자열_출력("예상치 못한 생성자 형식. " + reflect.TypeOf(생성자).String())
		테스트.Fail()
	}

	i기본_문자열 = i정밀수.(I기본_문자열)
	i임의값_생성 = i정밀수.(I임의값_생성)

	// G값() 테스트
	F같은값_확인(테스트, i정밀수.G값(), 입력값)
	F같은값_확인(테스트, i정밀수.GRat(), 입력값)
	F같은값_확인(테스트, i정밀수.G실수(), 입력값)
	F같은값_확인(테스트, i정밀수.G정밀수(), 입력값)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatFloat(입력값, 'f', -1, 64))

	// 입력값 변경 후 독립성 확인.
	입력값 = 입력값 + 1

	F다른값_확인(테스트, 입력값, 입력값_백업)
	F다른값_확인(테스트, i정밀수.G값(), 입력값)
	F같은값_확인(테스트, i정밀수.G값(), 입력값_백업)
	F같은값_확인(테스트, i정밀수.GRat(), 입력값_백업)
	F같은값_확인(테스트, i정밀수.G실수(), 입력값_백업)
	F같은값_확인(테스트, i정밀수.G정밀수(), 입력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), strconv.FormatFloat(입력값_백업, 'f', -1, 64))

	// 출력값 변경 후 독립성 확인.
	출력값 := i정밀수.GRat()
	출력값_백업 := new(big.Rat).Set(출력값)
	F같은값_확인(테스트, 출력값, 출력값_백업)

	출력값.Add(출력값, 출력값)
	F다른값_확인(테스트, 출력값, 출력값_백업)
	F다른값_확인(테스트, i정밀수.G값(), 출력값)
	F같은값_확인(테스트, i정밀수.G값(), 출력값_백업)
	F같은값_확인(테스트, i정밀수.GRat(), 출력값_백업)
	F같은값_확인(테스트, i정밀수.G실수(), 출력값_백업)
	F같은값_확인(테스트, i정밀수.G정밀수(), 출력값_백업)
	F같은값_확인(테스트, i기본_문자열.String(), F마지막_0_제거(출력값_백업.FloatString(100)))

	// G같음()
	F참인지_확인(테스트, i정밀수.G같음(입력값_백업))
	F거짓인지_확인(테스트, i정밀수.G같음(입력값_백업+1))
	F거짓인지_확인(테스트, i정밀수.G같음(입력값_백업-1))

	무효한_비교값_모음 := []I가변형{
		nil, true, "숫자 아님", NC통화(KRW, 100), time.Now()}

	for _, 무효한_비교값 := range 무효한_비교값_모음 {
		F거짓인지_확인(테스트, i정밀수.G같음(무효한_비교값))
	}

	// G비교()
	F참인지_확인(테스트, i정밀수.G비교(입력값_백업-1) == 1, "에러 %v %v, %v %v", reflect.TypeOf(i정밀수), i정밀수, reflect.TypeOf(입력값_백업), 입력값_백업 - 1)
	F같은값_확인(테스트, i정밀수.G비교(입력값_백업-1), 1)
	F같은값_확인(테스트, i정밀수.G비교(입력값_백업), 0)
	F같은값_확인(테스트, i정밀수.G비교(입력값_백업+1), -1)

	무효한_비교값_모음 = []I가변형{
		nil, true, "숫자 아님", NC통화(KRW, 100), time.Now()}

	for _, 무효한_비교값 := range 무효한_비교값_모음 {
		F같은값_확인(테스트, i정밀수.G비교(무효한_비교값), -2)
	}

	// Generate() 테스트
	맵 := make(map[string]bool)
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))

	양수_수량 := 0
	음수_수량 := 0

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		reflect값 := i임의값_생성.Generate(임의값_생성기, 1)
		값 := reflect값.Interface().(I정밀수)

		맵[값.(I기본_문자열).String()] = true

		if 값.G비교(0.0) == 1 {
			양수_수량++
		} else if 값.G비교(0.0) == -1 {
			음수_수량++
		}
	}

	F참인지_확인(테스트, len(맵) > 80)
	F참인지_확인(테스트, 양수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)
	F참인지_확인(테스트, 음수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)

	// 상수형, 변수형 혼용되지 않는 지 확인
	switch i정밀수.(type) {
	case C정밀수:
		c := i정밀수.(C정밀수)

		_, ok := c.(*sC정밀수)
		F참인지_확인(테스트, ok)

		_, ok = c.(I상수형)
		F참인지_확인(테스트, ok)

		_, ok = c.(I변수형)
		F거짓인지_확인(테스트, ok)
	case V정밀수:
		v := i정밀수.(V정밀수)

		_, ok := v.(*sV정밀수)
		F참인지_확인(테스트, ok)

		_, ok = v.(I상수형)
		F거짓인지_확인(테스트, ok)

		_, ok = v.(I변수형)
		F참인지_확인(테스트, ok)
	default:
		F문자열_출력("예상치 못한 경우. %v", reflect.TypeOf(i정밀수))
	}
}

func TestC정밀수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI정밀수(테스트, NC정밀수)
	testI정밀수(테스트, (*sC정밀수).G변수형)

	c := NC정밀수(100.1)
	v := c.G변수형()
	F같은값_확인(테스트, c, v)
}

func TestV정밀수(테스트 *testing.T) {
	테스트.Parallel()
	
	testI정밀수(테스트, NV정밀수)
	testI정밀수(테스트, (*sV정밀수).G상수형)

	v := NV정밀수(100.1)
	c := v.G상수형()
	F같은값_확인(테스트, c, v)

	F같은값_확인(테스트, NV정밀수(10.0045).S반올림(2), 10.0)
	F같은값_확인(테스트, NV정밀수(10.0045).S반올림(3), 10.005)
	F같은값_확인(테스트, NV정밀수(100.1).S값(300), 300)
	F같은값_확인(테스트, NV정밀수(100.1).S절대값(), 100.1)
	F같은값_확인(테스트, NV정밀수(-100.1).S절대값(), 100.1)
	F같은값_확인(테스트, NV정밀수(100.1).S더하기(100.1), 200.2)
	F같은값_확인(테스트, NV정밀수(100.1).S빼기(100.1), 0.0)
	F같은값_확인(테스트, NV정밀수(100.1).S곱하기(100.1), 10020.01)
	F같은값_확인(테스트, NV정밀수(100.1).S나누기(100.1), 1.0)

	F문자열_출력_일시정지_시작()
	F같은값_확인(테스트, NV정밀수(100.1).S나누기(0), nil)
	F문자열_출력_일시정지_종료()

	F같은값_확인(테스트, NV정밀수(100).S역수(), 0.01)

	F문자열_출력_일시정지_시작()
	F같은값_확인(테스트, NV정밀수(0.0).S역수(), nil)
	F문자열_출력_일시정지_종료()

	F같은값_확인(테스트, NV정밀수(100).S반대부호값(), -100)
	F같은값_확인(테스트, NV정밀수(-100).S반대부호값(), 100)
}

func TestP통화종류(테스트 *testing.T) {
	테스트.Parallel()
	
	F같은값_확인(테스트, KRW.String(), "KRW")
	F같은값_확인(테스트, USD.String(), "USD")
	F같은값_확인(테스트, CNY.String(), "CNY")
	F같은값_확인(테스트, EUR.String(), "EUR")
	F같은값_확인(테스트, INVALID_CURRENCY_TYPE.String(), "INVALID_CURRENCY_TYPE")
}

func TestN통화_생성자(테스트 *testing.T) {
	테스트.Parallel()
	
	입력값_모음 := []I가변형{
		uint(100), uint8(100), uint16(100), uint32(100), uint64(100),
		int(100), int8(100), int16(100), int32(100), int64(100),
		NC부호없는_정수(100), NC정수(100), float32(100), float64(100),
		NC실수(100), NC정밀수(100), "100.0"}

	for _, 입력값 := range 입력값_모음 {
		F같은값_확인(테스트, NC통화(KRW, 입력값).G종류(), KRW)
		F같은값_확인(테스트, NC통화(KRW, 입력값).G값(), 100)
		F같은값_확인(테스트, NV통화(KRW, 입력값).G종류(), KRW)
		F같은값_확인(테스트, NV통화(KRW, 입력값).G값(), 100)

		F같은값_확인(테스트, NC원화(입력값).G종류(), KRW)
		F같은값_확인(테스트, NC원화(입력값).G값(), 100)
		F같은값_확인(테스트, NV원화(입력값).G종류(), KRW)
		F같은값_확인(테스트, NV원화(입력값).G값(), 100)

		F같은값_확인(테스트, NC달러(입력값).G종류(), USD)
		F같은값_확인(테스트, NC달러(입력값).G값(), 100)
		F같은값_확인(테스트, NV달러(입력값).G종류(), USD)
		F같은값_확인(테스트, NV달러(입력값).G값(), 100)

		F같은값_확인(테스트, NC위안화(입력값).G종류(), CNY)
		F같은값_확인(테스트, NC위안화(입력값).G값(), 100)
		F같은값_확인(테스트, NV위안화(입력값).G종류(), CNY)
		F같은값_확인(테스트, NV위안화(입력값).G값(), 100)

		F같은값_확인(테스트, NC유로화(입력값).G종류(), EUR)
		F같은값_확인(테스트, NC유로화(입력값).G값(), 100)
		F같은값_확인(테스트, NV유로화(입력값).G종류(), EUR)
		F같은값_확인(테스트, NV유로화(입력값).G값(), 100)
	}

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_종료()

	입력값_모음 = []I가변형{nil, "변환 불가능한 문자열", true,
		time.Now(), NC통화(KRW, 100)}

	for _, 입력값 := range 입력값_모음 {
		F참인지_확인(테스트, F_nil값임(NC통화(KRW, 입력값)))
		F참인지_확인(테스트, F_nil값임(NV통화(KRW, 입력값)))

		F참인지_확인(테스트, F_nil값임(NC원화(입력값)))
		F참인지_확인(테스트, F_nil값임(NV원화(입력값)))

		F참인지_확인(테스트, F_nil값임(NC달러(입력값)))
		F참인지_확인(테스트, F_nil값임(NV달러(입력값)))

		F참인지_확인(테스트, F_nil값임(NC위안화(입력값)))
		F참인지_확인(테스트, F_nil값임(NV위안화(입력값)))

		F참인지_확인(테스트, F_nil값임(NC유로화(입력값)))
		F참인지_확인(테스트, F_nil값임(NV유로화(입력값)))
	}
}

func testI통화(테스트 *testing.T, 생성자 I가변형) {
	// 통화종류를 임의로 선택하기.
	통화종류 := F임의_통화종류()
	초기값 := 11111.1111

	var i통화 I통화
	var i기본_문자열 I기본_문자열
	var i임의값_생성 I임의값_생성

	// 생성자 테스트
	switch 생성자.(type) {
	case func(P통화종류, I가변형) C통화:
		생성자_ := 생성자.(func(P통화종류, I가변형) C통화)
		i통화 = 생성자_(통화종류, 초기값)
	case func(P통화종류, I가변형) V통화:
		생성자_ := 생성자.(func(P통화종류, I가변형) V통화)
		i통화 = 생성자_(통화종류, 초기값)
	case func(I가변형) C통화:
		생성자_ := 생성자.(func(I가변형) C통화)
		i통화 = 생성자_(초기값)
		if i통화.G종류() != 통화종류 {
			통화종류 = i통화.G종류()
		}
	case func(I가변형) V통화:
		생성자_ := 생성자.(func(I가변형) V통화)
		i통화 = 생성자_(초기값)
		if i통화.G종류() != 통화종류 {
			통화종류 = i통화.G종류()
		}
	case func(*sC통화) V통화:
		생성자_ := 생성자.(func(*sC통화) V통화)
		i통화 = 생성자_(&sC통화{종류: 통화종류, 금액: NC정밀수(초기값)})
	case func(*sV통화) C통화:
		생성자_ := 생성자.(func(*sV통화) C통화)
		i통화 = 생성자_(&sV통화{종류: 통화종류, 금액: NV정밀수(초기값)})
	default:
		테스트.Errorf("%stestI통화() : 알려지지 않은  생성자 타입 %v.",
			F소스코드_위치(2), reflect.TypeOf(생성자))
	}

	입력값 := F반올림(초기값, F통화종류별_정밀도(통화종류))
	//입력값_백업 := F반올림(초기값, F통화종류별_정밀도(통화종류))

	문자열_예상값 := 통화종류.String() + " 11,111"

	if F통화종류별_정밀도(통화종류) > 0 {
		문자열_예상값 += "."
	}

	for 반복횟수 := 0; 반복횟수 < F통화종류별_정밀도(통화종류); 반복횟수++ {
		문자열_예상값 += "1"
	}

	i기본_문자열 = i통화.(I기본_문자열)
	i임의값_생성 = i통화.(I임의값_생성)

	F같은값_확인(테스트, i통화.G종류(), 통화종류)
	F같은값_확인(테스트, i통화.G값(), 입력값)
	F같은값_확인(테스트, i기본_문자열.String(), 문자열_예상값)

	// 통화 그 자체로 비교되는 지 확인
	F같은값_확인(테스트, i통화, NC통화(통화종류, 입력값))

	// G같음()
	F참인지_확인(테스트, i통화.G같음(NC통화(통화종류, 입력값)))
	F거짓인지_확인(테스트, i통화.G같음(NC통화(INVALID_CURRENCY_TYPE, 입력값)))
	F거짓인지_확인(테스트, i통화.G같음(NC통화(통화종류, 입력값.G실수()+1)))
	F거짓인지_확인(테스트, i통화.G같음(NC통화(통화종류, nil)))
	F거짓인지_확인(테스트, i통화.G같음(nil))

	// G비교()
	F같은값_확인(테스트, i통화.G비교(NC통화(통화종류, 입력값.G실수()-1)), 1)
	F같은값_확인(테스트, i통화.G비교(NC통화(통화종류, 입력값)), 0)
	F같은값_확인(테스트, i통화.G비교(NC통화(통화종류, 입력값.G실수()+1)), -1)
	F같은값_확인(테스트, i통화.G비교(NC통화(INVALID_CURRENCY_TYPE, 입력값)), -2)
	F같은값_확인(테스트, i통화.G비교(nil), -2)
	F같은값_확인(테스트, i통화.G비교(NC통화(통화종류, nil)), -2)

	// 입력값 변경 후 독립성 확인. => 입력값이 상수형이라서 변경 불가임. 생략.
	// 출력값 변경 후 독립성 확인. => 출력값이 상수형이라서 변경 불가임. 생략.

	// Generate() 테스트
	맵_종류 := make(map[P통화종류]bool)
	맵_값 := make(map[string]bool)
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))

	양수_수량 := 0
	음수_수량 := 0

	for 반복횟수 := 0; 반복횟수 < 100; 반복횟수++ {
		reflect값 := i임의값_생성.Generate(임의값_생성기, 1)
		값 := reflect값.Interface().(I통화)
		맵_종류[값.G종류()] = true
		맵_값[값.G값().String()] = true

		if 값.G값().G비교(0.0) == 1 {
			양수_수량++
		} else if 값.G값().G비교(0.0) == -1 {
			음수_수량++
		}
	}

	F참인지_확인(테스트, len(맵_종류) == 4)
	F참인지_확인(테스트, len(맵_값) > 80, len(맵_값))
	F참인지_확인(테스트, 양수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)
	F참인지_확인(테스트, 음수_수량 > 30, "양수 %v개, 음수 %v개", 양수_수량, 음수_수량)

	// 상수형, 변수형 혼용되지 않는 지 확인
	switch i통화.(type) {
	case C통화:
		c := i통화.(C통화)

		_, ok := c.(*sC통화)
		F참인지_확인(테스트, ok)

		_, ok = c.(I상수형)
		F참인지_확인(테스트, ok)

		_, ok = c.(I변수형)
		F거짓인지_확인(테스트, ok)
	case V통화:
		v := i통화.(V통화)

		_, ok := v.(*sV통화)
		F참인지_확인(테스트, ok)

		_, ok = v.(I상수형)
		F거짓인지_확인(테스트, ok)

		_, ok = v.(I변수형)
		F참인지_확인(테스트, ok)
	default:
		F문자열_출력("예상치 못한 경우. %v", reflect.TypeOf(i통화))
	}
}

func TestC통화(테스트 *testing.T) {
	테스트.Parallel()
	
	testI통화(테스트, NC통화)
	testI통화(테스트, (*sC통화).G변수형)
	testI통화(테스트, NC원화)
	testI통화(테스트, NC달러)
	testI통화(테스트, NC위안화)
	testI통화(테스트, NC유로화)

	c := NC통화(KRW, 100)
	v := c.G변수형()
	F같은값_확인(테스트, c, v)
}

func TestV통화(테스트 *testing.T) {
	테스트.Parallel()
	
	testI통화(테스트, NV통화)
	testI통화(테스트, (*sV통화).G상수형)
	testI통화(테스트, NV원화)
	testI통화(테스트, NV달러)
	testI통화(테스트, NV위안화)
	testI통화(테스트, NV유로화)

	v := NV통화(KRW, 100)
	c := v.G상수형()
	F같은값_확인(테스트, c, v)

	// 통화종류를 매번 다르게 선택하기.
	통화종류 := F임의_통화종류()

	F같은값_확인(테스트, NV통화(통화종류, 100).S값(300), NC통화(통화종류, 300))
	F같은값_확인(테스트, NV통화(통화종류, 100).S절대값(), NC통화(통화종류, 100))
	F같은값_확인(테스트, NV통화(통화종류, -100).S절대값(), NC통화(통화종류, 100))
	F같은값_확인(테스트, NV통화(통화종류, 100).S더하기(100), NC통화(통화종류, 200))
	F같은값_확인(테스트, NV통화(통화종류, 100).S빼기(100), NC통화(통화종류, 0))
	F같은값_확인(테스트, NV통화(통화종류, 100).S곱하기(100), NC통화(통화종류, 10000))
	F같은값_확인(테스트, NV통화(통화종류, 100).S나누기(100), NC통화(통화종류, 1))
	F같은값_확인(테스트, NV통화(통화종류, 100).S반대부호값(), NC통화(통화종류, -100))
	F같은값_확인(테스트, NV통화(통화종류, -100).S반대부호값(), NC통화(통화종류, 100))

	F문자열_출력_일시정지_시작()
	F참인지_확인(테스트, F_nil값임(NV통화(통화종류, 100).S나누기(0)))
	F문자열_출력_일시정지_종료()
}

func TestC매개변수(테스트 *testing.T) {
	테스트.Parallel()
	
	c := NC매개변수("이름", 1.0)

	_, ok := c.(*sC매개변수)
	F참인지_확인(테스트, ok)

	_, ok = c.(I상수형)
	F참인지_확인(테스트, ok)

	F같은값_확인(테스트, c.G이름(), "이름")
	F같은값_확인(테스트, c.G값(), 1.0)
	F참인지_확인(테스트, c.G숫자형식임())
	F거짓인지_확인(테스트, c.G문자열형식임())
	F거짓인지_확인(테스트, c.G시점형식임())
	F거짓인지_확인(테스트, c.G참거짓형식임())
}

func TestC안전한_가변형(테스트 *testing.T) {
	테스트.Parallel()
	
	채널 := make(chan int)
	함수 := func () {}
	
	값_모음 := []I가변형{
		nil, uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1), float64(1), true, false, "test", time.Now(),
		NC부호없는_정수(1), NC정수(1), NC실수(1), NC정밀수(1), NC통화(KRW, 1),
		NC시점(time.Now()), NC문자열("테스트"), 채널, 함수}

	for _, 값 := range 값_모음 {
		c := NC안전한_가변형(값)
		
		if 값 == nil { 
			F참인지_확인(테스트, F_nil값임(c))
			return
		} else {
			F거짓인지_확인(테스트, F_nil값임(c))
		}
		
		if reflect.TypeOf(값).Kind() == reflect.Func { continue }
		
		F같은값_확인(테스트, c.G값(), 값)
	}
}

func TestI안전한_배열(테스트 *testing.T) {
	테스트.Parallel()
	
	배열0 := N안전한_배열()
	배열1 := 배열0.S추가("첫번째")

	// 아래 배열 2개는 달라야 함.
	배열2_1 := 배열1.S추가("두번째-1")
	배열2_2 := 배열1.S추가("두번째-2")

	// G비어있음() 테스트
	F참인지_확인(테스트, 배열0.G비어있음())
	F거짓인지_확인(테스트, 배열1.G비어있음())
	F거짓인지_확인(테스트, 배열2_1.G비어있음())
	F거짓인지_확인(테스트, 배열2_2.G비어있음())

	// G길이() 테스트
	F같은값_확인(테스트, 배열0.G길이(), 0)
	F같은값_확인(테스트, 배열1.G길이(), 1)
	F같은값_확인(테스트, 배열2_1.G길이(), 2)
	F같은값_확인(테스트, 배열2_2.G길이(), 2) // 2_1과 2_2는 독립적이어야 함.

	// G슬라이스() 테스트
	F같은값_확인(테스트, 배열1.G슬라이스()[0], "첫번째")
	F같은값_확인(테스트, 배열2_1.G슬라이스()[0], "첫번째")
	F같은값_확인(테스트, 배열2_1.G슬라이스()[1], "두번째-1")
	F같은값_확인(테스트, 배열2_2.G슬라이스()[0], "첫번째")
	F같은값_확인(테스트, 배열2_2.G슬라이스()[1], "두번째-2")

	// 유효한 값 받아들이는 지 확인.
	입력값_모음 := []I가변형{uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		int(1), int8(1), int16(1), int32(1), int64(1),
		float32(1), float64(1), true, false, "test",
		time.Now(), NC부호없는_정수(1), NC정수(1), NC실수(1), NC정밀수(1),
		NC시점(time.Now()), NC문자열("테스트"), NC통화(KRW, 1), NC매개변수("테스트", 1.1),
		NV부호없는_정수(1), NV정수(1), NV실수(1), NV정밀수(1.1),
		NV시점(time.Now()), NV통화(KRW, 100), make(chan int)}

	for _, 입력값 := range 입력값_모음 {
		배열 := N안전한_배열().S추가(입력값)
		F같은값_확인(테스트, 배열.G슬라이스()[0], 입력값)
	}
}

func TestI안전한_맵(테스트 *testing.T) {
	테스트.Parallel()
	
	맵 := N안전한_맵()

	F참인지_확인(테스트, 맵.G비어있음())
	F같은값_확인(테스트, len(맵.G키_모음()), 0)

	// 원소가 추가될 때마다 새로운 맵이 추가 되는 지 확인.
	맵1 := N안전한_맵().S값("첫번째", "1")
	맵2_1 := 맵1.S값("두번째", "2_1")
	맵2_2 := 맵1.S값("두번째", "2_2")

	F같은값_확인(테스트, 맵1.G길이(), 1)
	F같은값_확인(테스트, 맵1.G존재함("첫번째"), true)
	F같은값_확인(테스트, 맵1.G값("첫번째"), "1")
	F같은값_확인(테스트, 맵1.G존재함("두번째"), false)
	F같은값_확인(테스트, 맵1.G값("두번째"), nil)

	F같은값_확인(테스트, 맵2_1.G길이(), 2)
	F같은값_확인(테스트, 맵2_1.G존재함("첫번째"), true)
	F같은값_확인(테스트, 맵2_1.G값("첫번째"), "1")
	F같은값_확인(테스트, 맵2_1.G존재함("두번째"), true)
	F같은값_확인(테스트, 맵2_1.G값("두번째"), "2_1")

	F같은값_확인(테스트, 맵2_2.G길이(), 2)
	F같은값_확인(테스트, 맵2_2.G값("첫번째"), "1")
	F같은값_확인(테스트, 맵2_2.G존재함("첫번째"), true)
	F같은값_확인(테스트, 맵2_2.G값("두번째"), "2_2")
	F같은값_확인(테스트, 맵2_2.G존재함("두번째"), true)

	// immutable 확인.
	맵1 = 맵1.S값("첫번째", "100")

	F같은값_확인(테스트, 맵1.G길이(), 1)
	F같은값_확인(테스트, 맵1.G값("첫번째"), "100")

	F같은값_확인(테스트, 맵2_1.G길이(), 2)
	F같은값_확인(테스트, 맵2_1.G값("첫번째"), "1")

	F같은값_확인(테스트, 맵2_2.G길이(), 2)
	F같은값_확인(테스트, 맵2_2.G값("첫번째"), "1")

	맵2_1 = 맵2_1.S값("첫번째", "1000_001")

	F같은값_확인(테스트, 맵1.G값("첫번째"), "100")
	F같은값_확인(테스트, 맵2_1.G값("첫번째"), "1000_001")
	F같은값_확인(테스트, 맵2_2.G값("첫번째"), "1")

	맵2_2 = 맵2_1.S값("두번째", "200_200")
	F같은값_확인(테스트, 맵2_2.G값("두번째"), "200_200")

	// G키_모음() 테스트
	키_모음_1 := 맵1.G키_모음()
	F같은값_확인(테스트, len(키_모음_1), 1)
	F같은값_확인(테스트, 키_모음_1[0], "첫번째")

	키_모음_2_1 := 맵2_1.G키_모음()
	F같은값_확인(테스트, len(키_모음_2_1), 2)
	F같은값_확인(테스트, 키_모음_2_1[0], "첫번째")
	F같은값_확인(테스트, 키_모음_2_1[1], "두번째")

	키_모음_2_2 := 맵2_2.G키_모음()
	F같은값_확인(테스트, len(키_모음_2_2), 2)
	F같은값_확인(테스트, 키_모음_2_2[0], "첫번째")
	F같은값_확인(테스트, 키_모음_2_2[1], "두번째")

	// S삭제() 테스트
	맵1 = 맵1.S삭제("첫번째")
	F같은값_확인(테스트, 맵1.G비어있음(), true)
	F같은값_확인(테스트, 맵1.G길이(), 0)

	F같은값_확인(테스트, 맵2_1.G길이(), 2)
	F같은값_확인(테스트, 맵2_2.G길이(), 2)

	맵2_1 = 맵2_1.S삭제("두번째")

	F같은값_확인(테스트, 맵2_1.G길이(), 1)
	F같은값_확인(테스트, 맵2_2.G길이(), 2)

	맵 = N안전한_맵()
	길이 := 100

	for 인덱스 := 0; 인덱스 < 길이; 인덱스++ {
		맵 = 맵.S값(F문자열(인덱스), 인덱스*10)
	}

	F같은값_확인(테스트, 맵.G길이(), 길이)

	맵 = 맵.S삭제("42").S삭제("7").S삭제("19").S삭제("99")

	F같은값_확인(테스트, 맵.G길이(), 길이-4)

	for 인덱스 := 0; 인덱스 < 길이; 인덱스++ {
		switch 인덱스 {
		case 7, 19, 42, 99:
			F같은값_확인(테스트, 맵.G존재함(F문자열(인덱스)), false)
			F같은값_확인(테스트, 맵.G값(F문자열(인덱스)), nil)

		default:
			F같은값_확인(테스트, 맵.G존재함(F문자열(인덱스)), true)
			F같은값_확인(테스트, 맵.G값(F문자열(인덱스)), 인덱스*10)
		}
	}
}

func TestV문자열키_맵(테스트 *testing.T) {
	테스트.Parallel()
	
	맵 := NV문자열키_맵()
	
	입력값_모음 := []I가변형{1, 1.1, "테스트1", "테스트2", true}
	
	// G값(), S값() 테스트
	for _, 입력값 := range 입력값_모음 {
		맵.S값(F문자열(입력값), 입력값)
		
		값, 존재함 := 맵.G값(F문자열(입력값))
		
		F참인지_확인(테스트, 존재함)
		F같은값_확인(테스트, 값, 입력값)
	}
	
	값, 존재함 := 맵.G값("2")	// 존재하지 않는 값.
	F거짓인지_확인(테스트, 존재함)
	F참인지_확인(테스트, 값 == nil)
	
	
	// G키_모음() 테스트
	F같은값_확인(테스트, len(맵.G키_모음()), len(입력값_모음))
	
	for _, 키 := range 맵.G키_모음() {
		일치 := false
		
		for _, 입력값 := range 입력값_모음 {
			if 키 ==  F문자열(입력값) {
				일치 = true
				
				break
			}
		}
		
		F참인지_확인(테스트, 일치, "%v", 키)
	}
	
	// G키_값_모음() 테스트	
	F같은값_확인(테스트, len(맵.G키_값_모음()), len(입력값_모음))
	
	for _, 키_값 := range 맵.G키_값_모음() {	
		일치 := false
		
		for _, 입력값 := range 입력값_모음 {
			if 키_값.G키() == F문자열(입력값) &&
			   F값_같음(키_값.G값(), 입력값) {
				일치 = true
				
				break
			}
		}
		
		F참인지_확인(테스트, 일치, "%v, %v", 키_값.G키(), 키_값.G값())
	}
	
	// S없으면_추가() 테스트
	맵.S없으면_추가("1", 1)		// 이미 존재하는 값
	F같은값_확인(테스트, len(맵.G키_모음()), len(입력값_모음))
	
	맵.S없으면_추가("2", 2)		// 존재하지 않던 값
	F같은값_확인(테스트, len(맵.G키_모음()), len(입력값_모음) + 1)
	
	// 임의값 삽입 테스트
	총_수량 := 20000
	
	// 데드락등 동시처리로 인한 문제가 발생하기 쉽도록 일부러 동시처리 수량을 크게 잡음.
	동시처리_수량 := runtime.NumCPU() * 500
	
	맵 = NV문자열키_맵()	
	임의값_모음 := F임의값_생성(총_수량)
	
	입력_채널 := make(chan I가변형, 총_수량)	
	종료_통보_채널 := make(chan bool)
	
	// 순차처리하는 일반 맵에 먼저 채워넣음. (총 수량을 나중에 비교함.)
	보통_맵 := make(map[string]I가변형)
	for _, 임의값 := range 임의값_모음 {
		보통_맵[F문자열(임의값)] = 임의값
	}
	
	// 모든 데이터를 미리 채널 버퍼에 채워놓음.
	go func() {
		for _, 임의값 := range 임의값_모음 {
			입력_채널 <- 임의값
		}
	
		// 이게 중요하더라는.... 이거 안 하면 데드락이 발생함.
		close(입력_채널)
	}()
	
	// 동시처리 테스트
	for 반복횟수 := 0; 반복횟수 < 동시처리_수량; 반복횟수++ {
		go func() {
			for 임의값 := range 입력_채널 {
				맵.S값(F문자열(임의값), 임의값)
			}
			
			종료_통보_채널 <- true
		}()
	}
	
	// 처리가 끝날 때까지 대기할 것.
	// 이것도 엄청 중요하더라는.. 
	// 이거 안 하면 데이터가 맵에 채워지기도 전에 수량 비교 테스트가 진행되고, 테스트 실패함.
	for 반복횟수 := 0 ; 반복횟수 < 동시처리_수량 ; 반복횟수++ {
		<-종료_통보_채널
	}
	
	평균_수량 := len(보통_맵) / len(맵.(*sV문자열키_맵).중앙_저장소)
	최소_수량 := int(float64(평균_수량) * 0.5)
	
	for 키, 맵_조각 := range 맵.(*sV문자열키_맵).중앙_저장소 {		
		맵_조각별_수량 := len(맵_조각.저장소)
		
		F참인지_확인(테스트, 맵_조각별_수량 > 최소_수량, 
					"균일하게 퍼지지 않음. 키 %v, 최소_수량 %v, 수량 %v", 
					키, 최소_수량, 맵_조각별_수량)
	}
	
	// G값() 테스트.
	for 키, 원래_값 := range 보통_맵 {
		값, 존재함 := 맵.G값(키)	
		
		F참인지_확인(테스트, 존재함, "G값() 테스트 실패. 키 %v, 값 %v", 키, 원래_값)
		F같은값_확인(테스트, 원래_값, 값)
	}
	
	// G키_모음() 테스트.
	키_모음 := 맵.G키_모음()
	F같은값_확인(테스트, len(보통_맵), len(키_모음))
	
	for _, 키 := range 키_모음 {
		_, 존재함 := 보통_맵[키]
		
		F참인지_확인(테스트, 존재함, "G키_모음() 테스트 실패. 키 %v", 키)
	}
		
	// G키_값_모음() 테스트.
	키_값_모음 := 맵.G키_값_모음()
	F같은값_확인(테스트, len(보통_맵), len(키_값_모음))
	
	for _, 키_값 := range 키_값_모음 {
		키 := 키_값.G키()
		값, 존재함 := 보통_맵[키]
		
		F참인지_확인(테스트, 존재함, "G키_값_모음() 테스트 실패. 키 %v", 키)
		F같은값_확인(테스트, 키_값.G값(), 값)
	}
	
	// S없으면_추가() 테스트
	for _, 임의값 := range 임의값_모음 {
		맵.S값(F문자열(임의값), 임의값)
	}
	
	F같은값_확인(테스트, len(보통_맵), len(맵.G키_모음()))
}