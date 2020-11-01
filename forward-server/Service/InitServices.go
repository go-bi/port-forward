package Service


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)


var (
	OrmerS     orm.Ormer
	//ForWardServ   = new(ForWardServer)
	ForWardServ   = NewForWardServer()
	MagicServ   = NewMagicServiceV1()
	//MagicServ   = new(MagicServer)
	ConsoleServ = new(ConsoleServer)
	SysDataS   = new(SysDataService)
)

func init() {
	//开启DEBUG模式，输出SQL信息
	orm.Debug = true

	//_ "github.com/mattn/go-sqlite3"
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:data/data.db?cache=shared&loc=auto")

	OrmerS = orm.NewOrm()
	OrmerS.Using("default")


}