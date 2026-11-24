package cfg

func LoadGameData() error {
	funcs := []func() error{
		loadSkillId,
		loadInstData,
		loadGachaPool,
		loadSkinPool,
		loadChapterData,
	}

	for _, action := range funcs {
		if err := action(); err != nil {
			return err
		}
	}
	return nil
}
