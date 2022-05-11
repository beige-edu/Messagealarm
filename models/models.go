package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic("请配置数据库引擎")
	}
	err = orm.RegisterDataBase("default", "mysql",
		beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_password")+"@tcp("+beego.AppConfig.String("db_host")+":"+beego.AppConfig.String("db_port")+")/"+beego.AppConfig.String("db_name")+"?charset=utf8mb4&loc=Local")
	if err != nil {
		panic("数据库连接失败")
	}
	// 注册模型
	orm.RegisterModel(new(PrometheusAlertDB), new(AlertRecord), new(User), new(MessagePushLogs))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		panic("自动创建更新表失败")
	}
}

// PrometheusAlertDB 分类
type PrometheusAlertDB struct {
	Id      int
	Tpltype string
	Tpluse  string
	Tplname string `orm:"index"`
	Tpl     string `orm:"type(text)"`
	Created time.Time
}

func GetAllTpl() ([]*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	Tpl_all := make([]*PrometheusAlertDB, 0)
	qs := o.QueryTable("PrometheusAlertDB")
	_, err := qs.All(&Tpl_all)
	return Tpl_all, err
}

func GetTpl(id int) (*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	tpl_one := new(PrometheusAlertDB)
	qs := o.QueryTable("PrometheusAlertDB")
	err := qs.Filter("id", id).One(tpl_one)
	if err != nil {
		return nil, err
	}
	return tpl_one, err
}

func GetTplOne(name string) (*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	tpl_one := new(PrometheusAlertDB)
	qs := o.QueryTable("PrometheusAlertDB")
	err := qs.Filter("Tplname", name).One(tpl_one)
	if err != nil {
		return tpl_one, err
	}
	return tpl_one, err
}

func DelTpl(id int) error {
	o := orm.NewOrm()
	tpl_one := &PrometheusAlertDB{Id: id}
	_, err := o.Delete(tpl_one)
	return err
}

func AddTpl(id int, tplname, t_type, t_use, tpl string) error {
	o := orm.NewOrm()
	qs := o.QueryTable("PrometheusAlertDB")
	bExist := qs.Filter("Tplname", tplname).Exist()
	var err error
	if bExist {
		err = errors.New("模版名称已经存在！")
		return err
	}
	Template_table := &PrometheusAlertDB{
		Id:      id,
		Tplname: tplname,
		Tpltype: t_type,
		Tpluse:  t_use,
		Tpl:     tpl,
		Created: time.Now(),
	}
	// 插入数据
	_, err = o.Insert(Template_table)
	return err
}

func UpdateTpl(id int, tplname, t_type, t_use, tpl string) error {
	o := orm.NewOrm()
	tpl_update := &PrometheusAlertDB{Id: id}
	err := o.Read(tpl_update)
	if err == nil {
		tpl_update.Id = id
		tpl_update.Tplname = tplname
		tpl_update.Tpltype = t_type
		tpl_update.Tpluse = t_use
		tpl_update.Tpl = tpl
		tpl_update.Created = time.Now()
		_, err := o.Update(tpl_update)
		return err
	}
	return err
}

type AlertRecord struct {
	Id           int64
	SendType     string
	Alertname    string
	AlertLevel   string
	BusinessType string
	Instance     string
	StartsAt     string
	EndsAt       string
	Summary      string
	Description  string
	HandleStatus string
	AlertStatus  string
	AlertJson    string
	Remark       string
	Revision     int
	CreatedBy    string
	CreatedTime  time.Time
	UpdatedBy    string
	UpdatedTime  time.Time
}

func (alertRecord *AlertRecord) TableName() string {
	return "alert_record"
}

func AddAlertRecord(sendType string, alertname string, alertLevel string, businessType string, instance string,
	startsAt string, endsAt string, summary string, description string, alertStatus string, alertJson string, remark string) error {
	o := orm.NewOrm()
	var err error

	alertRecord := &AlertRecord{
		//Id: id,
		SendType:     sendType,
		Alertname:    alertname,
		AlertLevel:   alertLevel,
		BusinessType: businessType,
		Instance:     instance,
		StartsAt:     startsAt,
		EndsAt:       endsAt,
		Summary:      summary,
		Description:  description,
		HandleStatus: "0",
		AlertStatus:  alertStatus,
		AlertJson:    alertJson,
		Remark:       remark,
		Revision:     1,
		CreatedBy:    "system",
		CreatedTime:  time.Now(),
		UpdatedBy:    "system",
		UpdatedTime:  time.Now(),
	}
	// 插入数据
	_, err = o.Insert(alertRecord)
	return err
}

//User 第三方授权表
type User struct {
	Id        int
	AppKey    string    `json:"app_key"`
	AppSecret string    `json:"app_secret"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetUserByAppKey 通过appke appSecret获取第三方用户
func GetUserByAppKey(appkey, appSecret string) (*User, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("User")
	err := qs.Filter("app_key", appkey).Filter("app_secret", appSecret).One(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// MessagePushLogs 消息推送日志
type MessagePushLogs struct {
	Id        int64     `orm:"auto"`                            //主键
	Appkey    string    `orm:"size(100);index" json:"appkey"`   //用户AppKey
	SendType  string    `orm:"size(50);index" json:"send_type"` //发送类型
	Source    string    `orm:"size(100)" json:"source"`         //发起方
	Content   string    `orm:"type(text)" json:"content"`       //发送内容
	Status    int8      `json:"status"`                         //发送状态
	CreatedAt time.Time `json:"created_at"`                     //创建时间
	UpdatedAt time.Time `json:"updated_at"`                     //更新时间
}

// InsertMessagePushLog 插入推送消息
func InsertMessagePushLog(logs *MessagePushLogs) error {
	var err error
	o := orm.NewOrm()
	_, err = o.Insert(logs)
	return err
}
