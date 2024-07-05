package bitfile

// Close - Close the current Reader
func (o *Reader) Close() (err error) {

	if o.file != nil {

		err = o.file.Close()

		if err == nil {

			o.file = nil

		}

	}

	return err

}
