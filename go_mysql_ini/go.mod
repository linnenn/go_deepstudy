module go_mysql_ini

go 1.15

require (
	github.com/go-ini/ini v1.57.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
)
replace (
	github.com/go_mysql_ini/conf => ./go_mysql_ini/conf
	github.com/go_mysql_ini/utill => ./go_mysql_ini/utill
)
