package puzzle

import "testing"

func TestEnvName(t *testing.T) {
	truth := map[string]string{
		"test":              "TEST",
		"testN00b":          "TESTN00B",
		"a.b.c":             "A_B_C",
		"ok-go":             "OK_GO",
		"this--is_Awesome!": "THIS__IS_AWESOME_",
		"@user":             "_USER",
		"tomtom&nana":       "TOMTOM_NANA",
		"Oh my  gOd":        "OH_MY__GOD",
	}

	for k, v := range truth {
		e := NewEntry[bool](k)

		if e.Metadata.EnvName != v {
			t.Errorf("Expected %s, got %s", v, e.Metadata.EnvName)
		}
	}
}

func TestFlagName(t *testing.T) {
	truth := map[string]string{
		"test":              "test",
		"testN00b":          "testn00b",
		"a.b.c":             "a-b-c",
		"ok-go":             "ok-go",
		"this--is_Awesome!": "this--is-awesome",
		"@user":             "user",
		"tomtom&nana":       "tomtom-nana",
		"Oh my  gOd":        "oh-my--god",
	}

	for k, v := range truth {
		e := NewEntry[bool](k)
		if e.Metadata.FlagName != v {
			t.Errorf("Expected %s, got %s", v, e.Metadata.FlagName)
		}
	}
}

func TestString(t *testing.T) {
	c, _ := randomConfig()
	for _, e := range c.entries {
		s0 := e.String()
		e.Set(s0)
		s1 := e.String()
		if s0 != s1 {
			t.Errorf("Expected %s, got %s", s0, s1)
		}
	}
}

func TestSetWithBadInput(t *testing.T) {
	bad := "b4d"
	c, _ := randomConfig()
	for _, e := range c.entries {
		switch any(e.GetValue()).(type) {
		case string, []string:
			continue
		default:
			if err := e.Set(bad); err == nil {
				t.Errorf("Expected error, got nil")
			}
		}

	}
}
