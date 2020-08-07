package gettext

func init() {
	b, err := ReadFile("po.json")
	if err != nil {
		panic(err)
	}

	err = SetupTranslationMapping(b)
	if err != nil {
		panic(err)
	}
}
