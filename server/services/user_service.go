package services

func GetUserById() {
	a8m, err := client.User.
		Create().
		SetName("a8m").
		SetNillableAge(age).
		AddGroups(g1, g2).
		SetSpouse(nati).
		Save(ctx)
}
