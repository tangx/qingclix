package global

import "github.com/tangx/qingclix/utils"

var (
	HomeDir    string = utils.HomeDir()
	AuthFile   string = HomeDir + "/.qingcloud/config.yaml"
	ConfigFile string = HomeDir + "/.qingclix/config.json"
)

// Global Flags Vars
var (
	// logrus 日志等级
	// 0: Panic (最高), 6: Trace (最低)
	Verbose int

	// 是否跳过 Contract 购买
	SkipContract bool

	// Count 购买数量
	Count int

	// Dryrun 购买机器直接返回。
	Dryrun bool
)

var QingTypes string = `
{
    "instance_type":{
        
        "基础型":{"name":"基础型","type":"s1","class":101,"desc":"最垃圾的配置, 一般不用"},
        "企业型e1":{"name":"企业型e1","type":"e1","class":201,"desc":"虽然叫企业型, 但是并不保证高可用。"},
        "企业型e2":{"name":"企业型e2","type":"e2","class":202,"desc":"(首选)企业型e2, 保证高可用"},
        "专业增强型":{"name":"专业增强型","type":"p1","class":301,"desc":"比企业e2好一点"}
    },
    "volume_type":{
        "性能型":{"name":"性能型","type":0,"desc":""},
        "容量型":{"name":"容量型","type":2,"desc":""},
        "基础型":{"name":"基础型","type":100,"desc":"基础型硬盘是 100 (只能被基础型主机挂载)"},
        "SSD企业级硬盘": {"name":"SSD企业级硬盘","type":200,"desc":"SSD 企业级硬盘是，一般用于 企业e2机器"},
        "超高性能型":{"name":"超高性能型","type":3,"desc":"超高性能型是 3 (只能被超高性能主机挂载)"},
        "NeonSAN(企业级分布式SAN)":{"name":"NeonSAN(企业级分布式SAN)","type":5,"desc":"最好的硬盘，超级贵。(企业级分布式SAN)"}
    },
    "image_type":{
        "ubuntu16.04":{"name":"ubuntu16.04","image":"xeu1843","desc":""},
        "centos7.6":{"name":"centos7.6","image":"xeu1843","desc":""},
        "debian9":{"name":"debian9","image":"xeu1843","desc":""},
        "自定义centos7-kernel5.8":{"name":"自定义centos7-kernel5.8","image":"xeu1843","desc":""}
    }
}
`
