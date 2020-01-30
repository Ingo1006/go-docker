package subsystem

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/sirupsen/logrus"
)

type MemorySubSystem struct {
}

func (*MemorySubSystem) Name() string {
	return "cpu"
}

func (m *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subsystemCgroupPath, err := GetCgroupPath(m.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	if res.MemoryLimit != "" {
		// 设置cgroup内存限制，
		// 将这个限制写入到cgroup对应目录的 memory.limit_in_bytes文件中即可
		err := ioutil.WriteFile(path.Join(subsystemCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MemorySubSystem) Remove(cgroupPath string) error {
	subsystemCgroupPath, err := GetCgroupPath(m.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	return os.Remove(subsystemCgroupPath)
}

func (m *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	subsystemCgroupPath, err := GetCgroupPath(m.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(subsystemCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err !=nil {
		logrus.Errorf("write pid to tasks, err: %v", err)
		return err
	}
	return nil
}