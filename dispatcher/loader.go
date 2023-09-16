package dispatcher

type Loader func() (map[State]Process, error)
