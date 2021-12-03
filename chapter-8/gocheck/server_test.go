package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "gopkg.in/check.v1" // エクスポートされた識別子はパッケージ名を省略してアクセスできる
)

// 1. テストスイートに保存されるテストフィクスチャのデータ
type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	// テストスイートの登録
	Suite(&PostTestSuite{})
}

// パッケージtestingとの統合
func Test(t *testing.T) { TestingT(t) }

// 2. テストフィクスチャを作成
// テストケース関数の最初の文として実行
func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

// テストケース関数の最後の文として実行
func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

// テストケース全体が実行される前に実行
func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

// テストケース全体が実行される後に実行
func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finishing Post Test Suite")
}

func (s *PostTestSuite) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestPutPost(c *C) {
	json := strings.NewReader(
		`{"content":"Updated post", "author":"Sau Sheong"}`,
	)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 1)
	c.Check(s.post.Content, Equals, "Updated post")
}
