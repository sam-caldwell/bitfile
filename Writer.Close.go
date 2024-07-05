package bitfile

// Close - Close method to close the file behind the writer object.
func (o *Writer) Close() (err error) {

	if o.file != nil {
		err = o.file.Close()
		if err == nil {
			o.file = nil
		}
	}

	return err

}
