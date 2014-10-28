// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"math/rand"
	"testing"
	"time"
)

func TestC종목(테스트 *testing.T) {
	종목 := NC종목("코드", "명칭")
	
	_, 상수형임 := 종목.(I상수형)
	F참인지_확인(테스트, 상수형임, "NC종목() 결과값이 상수형이 아님.")

	F같은값_확인(테스트, 종목.G코드(), "코드")
	F같은값_확인(테스트, 종목.G명칭(), "명칭")
	F같은값_확인(테스트, 종목.String(), "코드 명칭")
	
	임의값_생성기 := rand.New(rand.NewSource(time.Now().UnixNano()))
	_ = 종목.(I임의값_생성).Generate(임의값_생성기, 1).Interface().(C종목)
}