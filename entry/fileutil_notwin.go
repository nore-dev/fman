package entry

func FileHidden(file string) (bool, error) {
	return file[0:1] == ".", nil
}
