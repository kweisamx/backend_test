# Backend Test

## 使用方法
```
git clone https://github.com/kweisamx/backend_test.git
cd backend_test
go run server.go
```
接著可以用http get 去 localhost:8080 或是 127.0.0.1:8080 去做request, 一分鐘超過六十下就會400 Bad Reqeuest了

## 測試程式

### 單元測試
測試程式測試handler對70個request的回傳值做檢查, 正常來說前面60個都要回傳當前的次數, 最後10個則是都要回傳error並檢驗 request status 是否為400 Bad Reqeust
```
go test server_test.go server.go
```
執行完看到下方commment表示通過測試

```
ok  	command-line-arguments	0.017s
``` 

### 實際測試

由於我們還要檢查60秒後是否開啟, 故我寫了一隻test.go方便查驗
主要就是一樣一開始會送出70個request, 正常後面10個的直應該要是error

再來我們過六十秒再送10個, 這常來說這10個應該都已經可以得到正確的回傳, 若上述有地方錯誤代表撰寫有問題

```
go run server.go // 先開啟web server
// 開另一個termial

go run test.go
```

termial 可以看到是否有測試成功, 會出現下列顯示

```
First Request is OK, wait 60 seconds
Second Reqeust is OK, test finish
```

## 心得以及作法
寫一點心得哈哈
作為我當完兵後的第一個題目, 這個真的有趣很多了(比起刷題...)

在golang的web端上, 有太多太多現成的web framework, 不過要求不能用到, 這邊我們就用最原生的net/http 應戰！

雖然一直很喜歡這語言, 但都沒有好好深究, 上一次好像也是快一年多前再寫一個side project時用到了, 那個side project已經消失了...
剛收到信的當晚就來寫, 號稱只要一個晚上就能寫完的, 不過我歷經一年都用python, 以及刷題用c++, 加上當兵的智商降低..., 疑 寫code突然變得好辛苦！？

首先我在會建立一個struct, 這裡面存三個值 第一個是IP, 第二個是在一分鐘內敲了多少下(判斷是否超過六十), 第三個就是倒數一分鐘
然後在開始寫, 首先我遇到的第一個問題是, 這些額外的值要怎麼寫進web 裡面？ 這邊我一開始用很髒的作法, 不管三七二十一先用global的變數再說(好孩子別學.. 除非是singleton), 在歷經一番坡折後, 終於寫好雛形了, 不過過程中也注意到lock的使用, 否則會有race condiction. 資料不同步的問題, 再來我們在每秒檢查這個連線端的list,如果過了60秒我們就要解除限制, 這邊也沒有用thread, 我選擇用goroutine.
再來花了一些功夫, 再想辦法把這global 的這麼髒的寫法給改掉, 遇到第一個難題, slice竟然也沒辦法 pass by reference!?, 這跟在C語言系列的array有點些許的不一樣, 終於找到用pointer的作法了, 這真的蠻特別的哈哈有空來看一下這個存取記憶體的方式

最後最後, 就是最重要的測試, 我知道golang 有內建一些測試的方法, 不過我也從來沒有用過所以研究一下, 畢竟好的測試會是一個產品重要的品質保證
