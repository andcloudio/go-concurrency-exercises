package absdata

//TODO: complete the below program to make data structure concurrent safe.

// Data structure that can be used by multiple goroutines.
type Data struct {
	// add lock field

	val interface{}
}

// New - create new data
func New(v interface{}) *Data {
	// modify to initialize the lock field

	return &Data{v}
}

// Set value
func (d *Data) Set(v interface{}) {
	// protect assignment with lock

	d.val = v
}

// Get value
func (d *Data) Get() interface{} {
	// allow concurrent read, with write getting exclusie access.

	return d.val
}

// Int value
func (d *Data) Int() int {
	return d.Get().(int)
}

// Bool value
func (d *Data) Bool() bool {
	return d.Get().(bool)
}

// String value
func (d *Data) String() string {
	return d.Get().(string)
}
