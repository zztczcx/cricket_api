package batting

import (
	"os"
	"strings"
	"testing"
)

func Test_Proudce(t *testing.T) {
        f, err := os.CreateTemp("", "test.*.csv")
        if err != nil {
                t.Fatal(err)
        }
        defer os.Remove(f.Name())

        Content := "header\ncontent2\ncontent3"
        if _, err := f.Write([]byte(Content)); err != nil {
		t.Fatal(err)
	}

        n := f.Name()
	l := &loader{
                dataFile: &n, 
        }

        dataSource := l.produce()
        content := []string{}
        for d := range dataSource {
                content = append(content, d)
        }
        
        //header was skipped
        expectedContent := "content2\ncontent3"
        resultContent := strings.Join(content[:], "\n")

        if expectedContent != resultContent {
                t.Errorf("expected %s\n, got %s", expectedContent, resultContent)
        }
}
