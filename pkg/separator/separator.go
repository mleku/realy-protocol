package separator

func Add(dst []byte, custom ...byte) (r []byte) {
	sep := byte('\n')
	if len(custom) > 0 {
		sep = custom[0]
	}
	r = append(dst, sep)
	return
}
