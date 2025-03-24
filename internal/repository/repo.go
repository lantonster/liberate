package repository

type Repo struct {
	UserRepo UserRepo
}

func NewRepo(
	UserRepo UserRepo,
) *Repo {
	return &Repo{
		UserRepo: UserRepo,
	}
}
