package puzzle

import (
	"net"
	"testing"
	"time"
)

func TestDefineTwice(t *testing.T) {
	var x int
	config, _ := randomConfig()
	for _, e := range config.entries {
		if err := DefineVar(config, e.GetKey(), &x); err == nil {
			t.Errorf("expected error, got nil")
		}
	}
}

func TestGet(t *testing.T) {
	config, values := randomConfig()
	if v, err := Get[bool](config, "bool"); err != nil || v != values.b {
		t.Errorf("expected %v, got %v (error: %v)", values.b, v, err)
	}
	if v, err := Get[string](config, "string"); err != nil || v != values.s {
		t.Errorf("expected %v, got %v (error: %v)", values.b, v, err)
	}
	if v, err := Get[int](config, "int"); err != nil || v != values.i {
		t.Errorf("expected %v, got %v (error: %v)", values.i, v, err)
	}
	if v, err := Get[uint](config, "uint"); err != nil || v != values.u {
		t.Errorf("expected %v, got %v (error: %v)", values.u, v, err)
	}
	if v, err := Get[float32](config, "float32"); err != nil || v != values.f32 {
		t.Errorf("expected %v, got %v (error: %v)", values.f32, v, err)
	}
	if v, err := Get[float64](config, "float64"); err != nil || v != values.f64 {
		t.Errorf("expected %v, got %v (error: %v)", values.f64, v, err)
	}
	if v, err := Get[time.Duration](config, "duration"); err != nil || v != values.d {
		t.Errorf("expected %v, got %v (error: %v)", values.d, v, err)
	}
	if v, err := Get[net.IP](config, "ip"); err != nil || !v.Equal(values.ip) {
		t.Errorf("expected %v, got %v (error: %v)", values.ip, v, err)
	}
	if v, err := Get[[]byte](config, "bytes"); err != nil || string(v) != string(values.bytes) {
		t.Errorf("expected %v, got %v (error: %v)", values.bytes, v, err)
	}
	if v, err := Get[[]string](config, "string-slice"); err != nil || !sliceEqualFactory(v)(values.ss) {
		t.Errorf("expected %v, got %v (error: %v)", values.s, v, err)
	}
	if v, err := Get[uint8](config, "uint8"); err != nil || v != values.u8 {
		t.Errorf("expected %v, got %v (error: %v)", values.u8, v, err)
	}
	if v, err := Get[uint16](config, "uint16"); err != nil || v != values.u16 {
		t.Errorf("expected %v, got %v (error: %v)", values.u16, v, err)
	}
	if v, err := Get[uint32](config, "uint32"); err != nil || v != values.u32 {
		t.Errorf("expected %v, got %v (error: %v)", values.u32, v, err)
	}
	if v, err := Get[uint64](config, "uint64"); err != nil || v != values.u64 {
		t.Errorf("expected %v, got %v (error: %v)", values.u64, v, err)
	}
	if v, err := Get[int8](config, "int8"); err != nil || v != values.i8 {
		t.Errorf("expected %v, got %v (error: %v)", values.i8, v, err)
	}
	if v, err := Get[int16](config, "int16"); err != nil || v != values.i16 {
		t.Errorf("expected %v, got %v (error: %v)", values.i16, v, err)
	}
	if v, err := Get[int32](config, "int32"); err != nil || v != values.i32 {
		t.Errorf("expected %v, got %v (error: %v)", values.i32, v, err)
	}
	if v, err := Get[int64](config, "int64"); err != nil || v != values.i64 {
		t.Errorf("expected %v, got %v (error: %v)", values.i64, v, err)
	}
}
