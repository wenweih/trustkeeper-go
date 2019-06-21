package repository

// IBiz repository bussiness logic
type IBiz interface {
  Close() error
}

func (repo *repo) Close() error{
  return repo.close()
}
