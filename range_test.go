package tddbc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshihiro-shu/tddbc"
)

func TestClosedRange(t *testing.T) {
	type args struct {
		n1, n2 int
	}

	// 課題1-1 上端点、下端点
	lowerEndpointTests := []struct {
		name   string
		args   args
		expect int
	}{
		{
			name:   "(1, 10)の時、下端点は1を返す",
			args:   args{1, 10},
			expect: 1,
		},
		{
			name:   "(10, 1)の時、下端点は1を返す",
			args:   args{10, 1},
			expect: 1,
		},
		{
			name:   "(1, 1)の時、下端点は1を返す",
			args:   args{1, 1},
			expect: 1,
		},
	}
	for _, test := range lowerEndpointTests {
		t.Run(test.name, func(t *testing.T) {
			cr := tddbc.NewClosedRange(test.args.n1, test.args.n2)
			res := cr.LowerEndpoint()
			assert.Equal(t, res, test.expect)
		})
	}
	upperEndpointTests := []struct {
		name   string
		args   args
		expect int
	}{
		{
			name:   "(1, 10)の時、上端点は10を返す",
			args:   args{1, 10},
			expect: 10,
		},
		{
			name:   "(10, 1)の時、上端点は10を返す",
			args:   args{10, 1},
			expect: 10,
		},
		{
			name:   "(1, 1)の時、上端点は1を返す",
			args:   args{1, 1},
			expect: 1,
		},
	}
	for _, test := range upperEndpointTests {
		t.Run(test.name, func(t *testing.T) {
			cr := tddbc.NewClosedRange(test.args.n1, test.args.n2)
			res := cr.UpperEndpoint()
			assert.Equal(t, res, test.expect)
		})
	}

	// 課題1-2 文字列表現 "[a,b]"
	toStringTests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name:   `(1, 10)の時、"[1,10]"を返す`,
			args:   args{1, 10},
			expect: `"[1,10]"`,
		},
		{
			name:   `(10, 1)の時、"[1,10]"を返す`,
			args:   args{10, 1},
			expect: `"[1,10]"`,
		},
		{
			name:   `(1, 1)の時、"[1,1]"を返す`,
			args:   args{1, 1},
			expect: `"[1,1]"`,
		},
	}
	for _, test := range toStringTests {
		t.Run(test.name, func(t *testing.T) {
			cr := tddbc.NewClosedRange(test.args.n1, test.args.n2)
			res := cr.ToString()
			assert.Equal(t, res, test.expect)
		})
	}

	// 課題1-3 帰属
	containsTests := []struct {
		name         string
		args         args
		containsArgs int
		expect       bool
	}{
		{
			name:         "[1,10]は、5を含む",
			args:         args{1, 10},
			containsArgs: 5,
			expect:       true,
		},
		{
			name:         "[1,10]は、-1を含まない",
			args:         args{1, 10},
			containsArgs: -1,
			expect:       false,
		},
		{
			name:         "[1,10]は、0を含まない",
			args:         args{1, 10},
			containsArgs: 0,
			expect:       false,
		},
		{
			name:         "[1,10]は、11を含まない",
			args:         args{1, 10},
			containsArgs: 11,
			expect:       false,
		},
		{
			name:         "[1,10]は、1を含む",
			args:         args{1, 10},
			containsArgs: 1,
			expect:       true,
		},
		{
			name:         "[1,10]は、10を含む",
			args:         args{1, 10},
			containsArgs: 10,
			expect:       true,
		},
		{
			name:         "[1,1]は、1を含む",
			args:         args{1, 1},
			containsArgs: 1,
			expect:       true,
		},
	}
	for _, test := range containsTests {
		t.Run(test.name, func(t *testing.T) {
			cr := tddbc.NewClosedRange(test.args.n1, test.args.n2)
			res := cr.Contains(test.containsArgs)
			assert.Equal(t, res, test.expect)
		})
	}
}
