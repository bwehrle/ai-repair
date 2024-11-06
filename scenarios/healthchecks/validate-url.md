In order to validate that the process supports a URL, a test should be written.  For the Go language, the following type of test should be written.  Replace `PLACEHOLDER` with the url to be validated and `Service` with the service to be invoked.

```go
func Test_ValidateUrl_Example(t *testing.T) {
    url := "PLACEHOLDER"
    t.Run("validate returns", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, url, nil)
		response := httptest.NewRecorder()
		Service(response, request)
		if response.StatusCode != 200 {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
```