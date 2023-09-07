# membership_system
後端會員系統

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
