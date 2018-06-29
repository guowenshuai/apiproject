package models

import (
	"github.com/astaxie/beego/logs"

	//"github.com/astaxie/beego"
	beeOrm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/guowenshuai/apiproject/db"
	"github.com/pkg/errors"
	"strconv"
)

//var w io.Writer

/*
func init() {
	mysqluser, mysqlpwd, mysqlurl, mysqlport, mysqldb :=
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb")

	dburl := mysqluser + ":" + mysqlpwd + "@tcp(" + mysqlurl + ":" + mysqlport + ")/" + mysqldb
	beego.Info("db url: ", dburl)

	//beeOrm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test_work?charset=utf8, 30, 30")

	// 注册数据库
	if err := beeOrm.RegisterDataBase("default", "mysql", dburl+"?charset=utf8, 30, 30"); err != nil {
		beego.Error("register db err: ", err.Error())
	}
	beego.Info("connect db success")

	// 注册驱动
	beeOrm.RegisterDriver("mysql", beeOrm.DRMySQL)
	// 注册模型
	beeOrm.RegisterModel(new(Job))
	// 打开调试模式
	beeOrm.Debug = true
	 //自动创建数据表, 	defalut数据库, 是否重建表, 是否打印详细
	beeOrm.RunSyncdb("default", false, true)
}
*/

func init() {
	// 注册模型
	beeOrm.RegisterModel(new(Job))
	//自动创建数据表, 	defalut数据库, 是否重建表, 是否打印详细
	beeOrm.RunSyncdb("default", false, true)
}

type Job struct {
	Id      int
	Company string
	Boss    string
	City    string
	Address string
}

// 自定义表名
func (j *Job) TableName() string {
	return "jobs"
}

// 插入数据
func (j *Job) Insert() (int64, error) {
	o := beeOrm.NewOrm()
	return o.Insert(j)
}

func InsertJobs(jobs []Job) (int64, error) {
	o := beeOrm.NewOrm()
	return o.InsertMulti(len(jobs), jobs)
}

// 查询所有
func GetAllJobs() (total int64, jobs []Job, err error) {
	o := beeOrm.NewOrm()
	qs := o.QueryTable("jobs")
	if total, err = qs.All(&jobs); err != nil {
		total, jobs = 0, []Job{}
	}
	logs.Info("get all %d jobs", total)
	return

}

// 根据id查询
func JobById(id string) (job Job, err error) {
	logs.Info("job by id: ", id)
	o := beeOrm.NewOrm()
	job.Id, _ = strconv.Atoi(id)
	err = o.Read(&job)
	return
}

// 根据boss字段查询
func JobByBoss(boss string) (job Job, err error) {
	logs.Info("job by boss ", boss)
	o := beeOrm.NewOrm()
	job.Boss = boss
	err = o.Read(&job, "Boss")
	return
}

// 根据id删除
func JobDelete(id string) (int64, error) {
	o := beeOrm.NewOrm()
	job := Job{}
	job.Id, _ = strconv.Atoi(id)
	return o.Delete(&job)
}

// 根据id更新job
func JobUpdate(id string, job Job) (int64, error) {
	o := beeOrm.NewOrm()
	jobId, _ := strconv.Atoi(id)
	if o.Read(&Job{Id: jobId}) == nil {
		job.Id = jobId
		return o.Update(&job)
	} else {
		return 0, errors.New("can not find job via id: " + id)
	}
}
