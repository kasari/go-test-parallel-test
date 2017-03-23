## go test の並列実行のメモ

### -cpuと-parallel
http://qiita.com/tenntenn/items/2832149c6fca15b66397

### まず
go testは明示的にコード内で指定しない限り、勝手にテストを並列に実行することはない。  
ひとつひとつ`TestXXX`を実行していく流れとなる。  
では、どのように指定するかというと、以下のように`t.Parallel()`を実行する。  
```go
func TestHoge(t *testing.T) {
	t.Parallel()
}
```

### 実際のテスト結果を交えて、以下の内容を見る
- パッケージをまたいだ時の`t.Parallel()`の挙動
- 並列実行時の共有オブジェクトの挙動

```
go test -v ./...
```
```
=== RUN   Test1
=== RUN   Test1_2
--- PASS: Test1_2 (0.91s)
	1_test.go:26: 0 Test1_2
	1_test.go:26: 1 Test1_2
	1_test.go:26: 2 Test1_2
=== RUN   Test2
--- PASS: Test2 (0.60s)
	2_test.go:17: 0 Test2
	2_test.go:17: 1 Test2
	2_test.go:17: 2 Test2
--- PASS: Test1 (0.91s)
	1_test.go:17: 0 Test2
	1_test.go:17: 1 Test2
	1_test.go:17: 2 Test2
PASS
ok  	github.com/kasari/test-parallel-test	1.831s
=== RUN   TestOther
--- PASS: TestOther (0.91s)
	other_test.go:17: 0 TestOther
	other_test.go:17: 1 TestOther
	other_test.go:17: 2 TestOther
PASS
ok  	github.com/kasari/test-parallel-test/other	0.918s
?   	github.com/kasari/test-parallel-test/share	[no test files]
```

### 結果
- TestOtherに`t.Parallel()`を書いたにも関わらず、Test1と同時に動いていない。よってパッケージをまたいだテストの並列実行は行われない
- Test1とTest2でどちらも共有オブジェクトを触っているため、欲しい値にならない。こういう場合は、`t．Parallel()`を呼ばないほうが良い
