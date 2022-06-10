module github.com/hexya-addons/mail

go 1.13

require (
	github.com/hexya-addons/base v0.1.6
	github.com/hexya-addons/baseSetup v0.0.0-20200620173052-9ef2a7b4c7fe
	github.com/hexya-addons/bus v0.1.7
	github.com/hexya-addons/web v0.1.7
	github.com/hexya-erp/hexya v0.1.7
	github.com/hexya-erp/pool v1.0.2
	github.com/jaytaylor/html2text v0.0.0-20200412013138-3577fbdbcff7
	github.com/jordan-wright/email v0.0.0-20200602115436-fd8a7622303e
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
)

replace github.com/hexya-erp/pool v1.0.2 => /home/npiganeau/Progs/hexya/hexya-addons/mail/pool

replace github.com/hexya-erp/hexya v0.1.7 => /home/npiganeau/Progs/hexya/hexya

replace github.com/hexya-addons/base v0.1.6 => /home/npiganeau/Progs/hexya/hexya-addons/base
