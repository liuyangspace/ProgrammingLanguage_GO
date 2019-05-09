package data


/*
命名空间:
包:
	声明:
		package packageName			// 声明当前程序的包名
	使用:
		import [aliasName] "packName"		// 导入其他包, aliasName 为 . 时，可忽略包名调用包内函数e.g:
												import "./pack1" 				//导入同一路径下的包
												import "container/list" 		//导入标准库包, 通过 packName.Item 调用成员
												import . "./pack1"				//使用.来做为包的别名时，你可以不通过包名来使用其中的项目。例如：test := ReturnStr()
												import _ "./pack1/pack1"		//只导入其副作用，只执行它的init函数并初始化其中的全局变量。
		import {							// 导入多个包
			"packName1",
			"packName2"
		}
		import ("packName1"; "packName2")	// 导入多个包
	安装: #多用于外部包，安装在 $GOROOT/src/ 目录
		托管安装: go install 命令,e.g: go install codesite.ext/author/goExample/goex //安装自http://codesite.ext/author/goExample/goex托管的包
		编译安装:
			# 以Linux/OS X 为例
			1) 编写 Makefile 文件
			2) chmod 777 ./Makefile 确保可执行性
			3) 终端执行 make 或 gomake 工具：他们都会生成一个包含静态库 pack1.a 的 _obj 目录。
			4) 同go install命令同样复制 pack1.a 到本地的 $GOROOT/pkg 的目录中一个以操作系统为名的子目录下。
		git安装:
	API文档工具: godoc -http=:6060 -goroot="."
可见性规则:
	#使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用。
	函数名首字母小写即为 private : func getId() {}
	函数名首字母大写即为 public : func Printf() {}
# 网站和版本控制系统的其他的选择(括号中为网站所使用的版本控制系统)：
	BitBucket(hg/Git)
	GitHub(Git)
	Google Code(hg/Git/svn)
	Launchpad(bzr)
# 外部库，如：
	MySQL(GoMySQL), PostgreSQL(go-pgsql), MongoDB (mgo, gomongo), CouchDB (couch-go), ODBC (godbcl), Redis (redis.go) and SQLite3 (gosqlite) database drivers
	SDL bindings
	Google's Protocal Buffers(goprotobuf)
	XML-RPC(go-xmlrpc)
	Twitter(twitterstream)
	OAuth libraries(GoAuth)
 */
