# blog_system
後端部落格系統

# 產品業務功能
##  會員系統
- 會員註冊
- 登入
- 登出
- 忘記密碼

## 文章管理
- 新增文章
- 查詢文章清單
- 取得指定文章
- 更新文章
- 刪除文章
  
## 標籤管理
- 新增標籤
- 查詢標籤清單
- 更新標籤
- 刪除標籤

## 其他功能
- 多圖片上傳

## 內部公共元件
- convert: 統一的字元轉換處理工具 ex: str to uint32、str to int
- email: 處理email與SMTP Server相關互動邏輯
- errcode: 這裡定義error的處理方式與業務邏輯錯誤的定義
- logger: 處理log的格式化與分級輸出相關邏輯
- pagination: 定義分頁相關邏輯、計算 PageOffset、PageSize



## 基礎建設
- Mysql
- Redis
- SMTP Server

configs/config.yaml
```
# 伺服器設定檔
Server: 
  # Run mode: "debug" or "release"
  RunMode: release
  # 監聽地址
  ListenAddr: 127.0.0.1
  # 監聽的HTTP Port
  HttpPort: 8080
  # Read timeout for incoming requests
  ReadTimeout: 60
  # Write timeout for outgoing responses
  WriteTimeout: 60

# 應用程式設定檔
App:
  # 設定預設 PageSize
  DefaultPageSize: 10
  # 設定最大 PageSize
  MaxPageSize: 100
  # 存放 Log 的位置
  LogSavePath: storage/logs
  # Log 的名稱
  LogFileName: app
  # Log 的副檔名
  LogFileExt: .log
    # 檔案上傳位置
  UploadSavePath: storage/uploads
  # 上傳檔案伺服器URL
  UploadServerURL: http://127.0.0.1:8080/static
  # 上傳圖片單檔大小限制: MB
  UploadImageMaxSize: 5
  # 多檔上傳圖片總檔案大小限制: Byte
  UploadMultiImageTotalMaxSize: 10485760
  # 上傳圖片允許的副檔名
  UploadImageAllowExts:
    - .jpg
    - .jpeeg
    - .png

# Configuration for the database
Database:
  # 資料庫類型 (e.g., "mysql", "postgres")
  DBType: mysql
  # 資料庫使用者名稱
  Username: 
  # 資料庫密碼
  Password: 
  # 資料庫的 Host
  Host: 127.0.0.1:3306
  # 資料庫名稱
  DBName: membership_system
  # 資料表的 Prefix 
  TablePrefix: blog_
  # 資料庫 Character set 設定
  Charset: utf8
  # 是否啟動資料庫的ParseTime選項
  ParseTime: true
  # 最大閒置連線數
  MaxIdleConns: 10
  # 最大連線數
  MaxOpenConns: 30

#SMTP資訊
Email:
  Host: 
  Port: 
  UserName: 
  Password: 
  IsSSL: 
  From: 
  To: 

#Redis資訊
Redis:
  Host: 
  Port: 
```
