package dao

import (
	"fmt"
	"fuying-web/config"
	"fuying-web/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
var db *gorm.DB

// 包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	fmt.Println(config.Config.Root)
	fmt.Println(config.Config.Port)

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Config.Root, config.Config.Pwd, config.Config.Host, config.Config.Port, config.Config.Name)
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
		return
	}

	sqlDB, _ := db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	if db == nil {
		return nil
	}
	return db

}
func GetOwnerIdList() (minersList []model.JmUserStudyLog, err error) {
	db.Raw(`select detail_id,pid from jm_user_study_log `).Scan(&minersList)
	return
}
