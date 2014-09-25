package common

import (
	//"math/big"
	//"time"
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