package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"log"
	"testing"
	"time"
)

func newConn() *zk.Conn {
	zkList := []string{"localhost:2181"}
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		panic(err)
	}
	return conn
}

// 查询路径下的所有节点
func Test_Children(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	children, _, err := conn.Children("/")
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("children:", children)
}

// 创建单层节点
func Test_CreateZNode(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	nodePath := "/config"
	//flags参数用于设置结点类型
	//0=永久节点，节点不会被删除，除非手动删除
	//1=短暂节点（zk.FlagEphemeral），session断开时，节点会被删除
	//2=顺序节点（zk.FlagSequence），会在节点后面自动添加序号
	//3=短暂节点+顺序节点，短暂且自动添加序号
	flags := int32(0)
	//acl参数用于设置权限控制
	acl := zk.WorldACL(zk.PermAll)
	resPath, err := conn.Create(nodePath, nil, flags, acl)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("resPath:", resPath)
}

// 判断节点是否存在
func Test_Exists(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	path := "/config"
	exist, stat, err := conn.Exists(path)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("exist:", exist)
	log.Println("stat:", stat)
}

// 创建多层节点
func Test_CreateMultilayerZNode(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	path := "/config/project01/mysql/host"
	err := createMultilayerZNode(conn, path)
	if err != nil {
		panic(err)
	}
}

// 逐层创建节点（递归）
func createMultilayerZNode(conn *zk.Conn, path string) error {
	log.Println("multilayer path", path)

	//判断节点是否存在
	exist, _, err := conn.Exists(path)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	//截取父节点的路径
	parentPath := path[:len(path)-1]
	i := len(parentPath) - 1
	for i >= 0 {
		if parentPath[i] == '/' {
			break
		}
		i--
	}
	parentPath = path[:i]
	if i > 1 {
		err = createMultilayerZNode(conn, parentPath[:i])
		if err != nil {
			return err
		}
	}

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	resPath, err := conn.Create(path, nil, flags, acl)
	if err != nil {
		log.Println("err", err)
		return err
	}
	log.Println("resPath", resPath)
	return nil
}

// 删除单层结点
func Test_Delete(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	path := "/config/project01/mysql/host"
	err := conn.Delete(path, -1)
	if err != nil {
		log.Println("err:", err)
	}
}

// 查询节点的数据
func Test_Get(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	path := "/config/project01/mysql/host"
	data, stat, err := conn.Get(path)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("data:", string(data))
	log.Println("stat:", stat)
}

// 修改节点的数据
func Test_Set(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	path := "/config/project01/mysql/host"
	data := []byte("localhost:3306")
	stat, err := conn.Set(path, data, -1)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("stat:", stat)
}

// 修改节点的数据（带版本）
func Test_SetVersion(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	path := "/config/project01/mysql/host"
	data, stat, err := conn.Get(path)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("stat.version:", stat.Version)

	_, err = conn.Set(path, data, stat.Version-1)
	if err != nil {
		log.Println("set with stat.version-1")
		log.Println("err:", err)
	}

	_, err = conn.Set(path, data, stat.Version)
	if err != nil {
		log.Println("err:", err)
	}
	log.Println("set with stat.version")
}
