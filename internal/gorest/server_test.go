package gorest

// type MyWriter struct {
// 	Content []byte
// }

// func (w *MyWriter) Write(p []byte) (n int, err error) {
// 	w.Content = p
// 	return len(p), nil
// }

// func TestGinEngineAttached(t *testing.T) {
// 	log.SetOutput(io.Discard)
// 	repo := RepositoryMock{}
// 	s := NewServer(&repo, false)

// 	if fmt.Sprintf("%T", s.router) != "*gin.Engine" {
// 		t.Fatal("Gin engine not attached to server")
// 	}
// }

// func TestDummyAuthAttached(t *testing.T) {
// 	log.SetOutput(io.Discard)
// 	gin.SetMode(gin.TestMode)
// 	repo := RepositoryMock{}

// 	repo.FindFunc = func(s string) ([]byte, time.Time, error) {
// 		return []byte(""), time.Now(), nil
// 	}
// 	s := NewServer(&repo, false)

// 	req, _ := http.NewRequest("GET", "/", nil)
// 	req.Header.Add("Content-type", "application/json")
// 	req.Header.Add("Accept", "application/json")
// 	rec := httptest.NewRecorder()
// 	s.ServeHTTP(rec, req)

// 	if rec.Result().StatusCode != 401 {
// 		t.Fatalf("expected response code 401, received %d", rec.Result().StatusCode)
// 	}

// 	req.Header.Set("Authorization", "Bearer debug")
// 	rec = httptest.NewRecorder()
// 	s.ServeHTTP(rec, req)

// 	if rec.Result().StatusCode != 200 {
// 		t.Fatalf("expected response code 200, received %d", rec.Result().StatusCode)
// 	}

// }

// func TestGetPayloadSizeLimit(t *testing.T) {
// 	log.SetOutput(io.Discard)
// 	type test struct {
// 		inp string
// 		exp int64
// 	}
// 	var def int64 = 1000 // #default_payload_limit
// 	tests := []test{
// 		{"600", 600},
// 		{"10000", 10000},
// 		{"-5", def},
// 		{"1", 1},
// 	}

// 	for _, test := range tests {
// 		os.Setenv("PAYLOAD_BYTE_LIMIT", test.inp)
// 		limit := getPayloadSizeLimit()
// 		if limit != test.exp {
// 			t.Fatalf("Got %d, expected %d", limit, test.exp)
// 		}
// 	}

// 	os.Unsetenv("PAYLOAD_BYTE_LIMIT")
// 	limit := getPayloadSizeLimit()
// 	if limit != def {
// 		t.Fatalf("Got %d, expected %d", limit, def)
// 	}
// }
