在 golang 中，實作 interface 為隱性

如果實作沒有完整，在編譯的時候並不會被診斷出問題

常在執行時才發現問題

因此建議使用宣告的方式去顯性驗證

```
var _ MyInterface = (*MyStructPoint)(nil)
```
