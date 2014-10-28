// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"math/rand"
	"reflect"
	//"sync/atomic"
	//"time"
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
func (s *sC종목) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	c := NC문자열("").(I임의값_생성)
	코드 := c.Generate(임의값_생성기, 크기).Interface().(C문자열).G값()
	명칭 := c.Generate(임의값_생성기, 크기).Interface().(C문자열).G값()

	return reflect.ValueOf(&sC종목{코드: 코드, 명칭: 명칭})
}


