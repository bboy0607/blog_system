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
# Configuration for the server
Server: 
  # Run mode: "debug" or "release"
  RunMode: release
  # Address to listen on
  ListenAddr: 127.0.0.1
  # HTTP port to listen on
  HttpPort: 8080
  # Read timeout for incoming requests
  ReadTimeout: 60
  # Write timeout for outgoing responses
  WriteTimeout: 60

# Configuration for the application
App:
  # Default page size for pagination
  DefaultPageSize: 10
  # Maximum page size for pagination
  MaxPageSize: 100
  # Path to save log files
  LogSavePath: storage/logs
  # Log file name
  LogFileName: app
  # Log file extension
  LogFileExt: .log

# Configuration for the database
Database:
  # Type of database (e.g., "mysql", "postgres")
  DBType: mysql
  # Username for database connection
  Username:
  # Password for database connection
  Password:
  # Host of the database server
  Host:
  # Name of the database
  DBName:
  # Prefix for table names
  TablePrefix:
  # Character set for the database
  Charset: utf8
  # Parse time option for database connection
  ParseTime: true
  # Maximum idle connections
  MaxIdleConns: 10
  # Maximum open connections
  MaxOpenConns: 30

#SMTP資訊
Email:
  Host: 
  Port:
  UserName: 
  Password: 
  IsSSL: true
  From: 
  To: 

#Redis資訊
Redis:
  Host: 127.0.0.1
  Port: 6379
```
