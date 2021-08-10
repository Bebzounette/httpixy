package httpixy

var (
	Headers = []string{"Client-IP","Connection","Contact","Forwarded","From","Host","Origin","Referer","True-Client-IP","X-Client-IP","X-Custom-IP-Authorization","X-Forward-For","X-Forwarded-For","X-Forwarded-Host","X-Forwarded-Server","X-Host","X-HTTP-Host-Override","X-Original-URL","X-Originating-IP","X-Real-IP","X-Remote-Addr","X-Remote-IP","X-Rewrite-URL","X-Wap-Profile"}
	Inject = []string{"127.0.0.1","localhost","0.0.0.0","0","127.1","127.0.1","2130706433"}
	Url string
	File string
	To int
)
