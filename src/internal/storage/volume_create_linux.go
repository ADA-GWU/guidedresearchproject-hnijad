//go:build linux

package storage

func CreateNewVolume(id int, dir string) (*Volume, error) {
	volumeName := dir + "/" + fmt.Sprintf("%d.data", id)

	if util.FileExists(volumeName) {
		log.Warningf("Volume %v exists", volumeName)
		return nil, errors.New(fmt.Sprintf("volume with %d id exists", id))
	}

	file, err := os.OpenFile(volumeName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Errorln("Error creating new volume:", err)
		return nil, err
	}

	syscall.Fallocate(int(file.Fd()), 1, 0, 5*1024*1024*1024)
	volume := &Volume{
		ID:        id,
		Dir:       dir,
		dataFile:  file,
		NeedleMap: make(map[int]*NeedleInfo),
		UsedSpace: 0,
	}
	return volume, nil
}
