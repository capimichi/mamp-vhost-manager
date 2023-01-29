package model

type Vhost struct {
	//<VirtualHost *:80>
	//ServerAdmin webmaster@dummy-host2.example.com
	//DocumentRoot "/Applications/MAMP/Library/docs/dummy-host2.example.com"
	//ServerName dummy-host2.example.com
	//ErrorLog "logs/dummy-host2.example.com-error_log"
	//CustomLog "logs/dummy-host2.example.com-access_log" common
	//</VirtualHost>

	Name string
	ServerAdmin string
	ServerName string
	DocumentRoot string

}