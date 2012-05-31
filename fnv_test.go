package fnv

import (
	"testing"
)


func Test_FNV1_32(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1_32(fnv_test_str[i]) != fnv1_32_results[i] {
				t.FailNow()
		}
	}
}

func Test_FNV1a_32(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1a_32(fnv_test_str[i]) != fnv1a_32_results[i] {
				t.FailNow()
		}
	}
}

func Test_FNV1(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1(fnv_test_str[i]) != fnv1_64_results[i] {
				t.FailNow()
		}
	}
}

func Test_FNV1a(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1a(fnv_test_str[i]) != fnv1a_64_results[i] {
				t.FailNow()
		}
	}
}

func Test_FNV1aMod(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1aMod(fnv_test_str[i], 100) > 100 {
			t.FailNow()
		}
	}
}

func Test_FNV1aMod_32(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1aMod_32(fnv_test_str[i], 100) > 100 {
			t.FailNow()
		}
	}
}

func Test_FNV1Mod(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1Mod(fnv_test_str[i], 100) > 100 {
			t.FailNow()
		}
	}
}

func Test_FNV1Mod_32(t *testing.T) {
	for i := range fnv_test_str {
		if Hash1Mod_32(fnv_test_str[i], 100) > 100 {
			t.FailNow()
		}
	}
}

func Benchmark_FNV1(t *testing.B) {
	for a := range fnv_test_str {
		Hash1(fnv_test_str[a])
	}
}

func Benchmark_FNV1_32(t *testing.B) {
	for a := range fnv_test_str {
		Hash1_32(fnv_test_str[a])
	}
}

func Benchmark_FNV1a(t *testing.B) {

	for a := range fnv_test_str {
		Hash1a(fnv_test_str[a])
	}
}

func Benchmark_FNV1a_32(t *testing.B) {

	for a := range fnv_test_str {
		Hash1a_32(fnv_test_str[a])
	}
}