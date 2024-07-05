package bitfile

// Name - Return the file name of any currently open bitfile.
func (o *Reader) Name() string {

	if o.file != nil {

		return o.file.Name()

	}

	return ""

}
