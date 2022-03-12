package main

import (
	"os"
	"time"
	"fmt"
	// "log"
	"testing"
	"github.com/ory/dockertest/v3"
	// "github.com/ory/dockertest/docker"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func TestMain(m *testing.M) {
	resource, pool := createContainer()
	defer closeContainer(resource, pool)
	m.Run()
}

// テストDBセットアップ関数

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	pwd, _ := os.Getwd()

	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		panic(err)
	}

	// Container起動時の設定
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:		"5.7",
		Env: []string{
			"MYSQL_ROOT_PASSWORD="+os.Getenv("MYSQL_ROOT_PASSWORD"),
		},
		// ExposedPorts: []string{"5432"},
		// PortBindings: map[docker.Port][]docker.PortBinding{
		// 	"5432": {
		// 		{HostIP: "0.0.0.0", HostPort: port},
		// 	},
		// },
		Mounts: []string{
			pwd + "/mysql.conf/my.cnf:/etc/mysql/my.cnf",
			pwd + "/mysql.conf/todo.sql:/docker-entrypoint-initdb.d/todo.sql",
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_cli",
		},
	}

	// コンテナの起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		// log.Fatalf("Could not start resource: %s", err)
		panic(err)
	}

	return resource, pool
}

// コンテナの削除
func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	if err := pool.Purge(resource); err != nil {
		panic(err)
	}
}

func ConnectTestDB(resource *dockertest.Resource, pool *dockertest.Pool) *gorm.DB {
	var db *gorm.DB
	if err := pool.Retry(func() error {
		// Containerが完全に立ち上がるまでのダウンタイム
		time.Sleep(time.Second * 10)

		var err error

		mysqlInfo := fmt.Sprintf("root:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True",
			os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
		
		db, err = gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})

		return err

	}); err != nil {
		panic(err)
	}
	return db
}
