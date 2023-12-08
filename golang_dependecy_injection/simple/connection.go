package simple

type Connection struct {
	*File
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}

func (c *Connection) Close() string {
	return "Close " + c.Name
}