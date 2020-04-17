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
    "instance_types":[
        {"name":"基础型","type":"s1","class":101,"desc":"最垃圾的配置, 一般不用"},
        {"name":"企业型e1","type":"e1","class":201,"desc":"虽然叫企业型, 但是并不保证高可用。"},
        {"name":"企业型e2","type":"e2","class":202,"desc":"(首选)企业型e2, 保证高可用"},
        {"name":"专业增强型","type":"p1","class":301,"desc":"比企业e2好一点"}
    ],
    "volume_types":[
        {"name":"性能型","type":0,"desc":""},
        {"name":"基础型","type":100,"desc":"基础型硬盘是 (只能被基础型主机挂载)"},
        {"name":"容量型","type":2,"desc":""},
        {"name":"SSD企业级硬盘","type":200,"desc":"SSD 企业级硬盘是，一般用于 企业e2机器"},
        {"name":"超高性能型","type":3,"desc":"超高性能型是 (只能被超高性能主机挂载)"},
        {"name":"NeonSAN(企业级分布式SAN)","type":5,"desc":"最好的硬盘，超级贵。(企业级分布式SAN)"}
    ],
    "image_types":[
        {"name":"ubuntu16.04","image":"xeu1843","desc":""},
        {"name":"centos7.6","image":"xeu1843","desc":""},
        {"name":"debian9","image":"xeu1843","desc":""},
        {"name":"自定义centos7-kernel5.8","image":"xeu1843","desc":""}
    ],
    "zones":[
        {"name":"pek3d","desc":"北京三区D区"},
        {"name":"pek3c","desc":"北京三区C区"}
    ],
    "vxnets":[
        {"name":"vxnet-sn2rnad","desc":"online"},
        {"name":"vxnet-xxxxx","desc":"dev"},
        {"name":"vxnet-yyyyy","desc":"qingxu"}
    ],
    "keypairs":[
        {"name":"key1","desc":"ops公共"}
    ],
    "relation":{
        "s1":[0,100],
        "e1":[2,3,5,200],
        "e2":[2,3,5,200],
        "p1":[2,3,5,200]
    }
}
`
