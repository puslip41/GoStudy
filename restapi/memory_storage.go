package restapi

import "errors"

type MemoryStorage struct {
	m  map[string]StorageItem
}

func (s *MemoryStorage)Register(id, password, name, email string) error {
	if s.m == nil {
		s.m = map[string]StorageItem{}
	}

	if _, exist := s.m[id]; exist == false {
		s.m[id] = StorageItem{ID:id, Password:password, Name:name, EMail:email}
	} else {
		return errors.New("exist id: " + id)
	}
	return nil
}

func (s *MemoryStorage)Modify(id, password, name, email string) error {
	if err := isCreateMap(s.m); err != nil {
		return err
	} else {
		if item, err := getItem(&s.m, id); err != nil {
			return err
		} else {
			if item.Password != password {
				return errors.New("not equals password")
			} else {
				s.m[id] = StorageItem{ID:id, Password:password, Name:name, EMail:email}

				return nil
			}
		}
	}

	return errors.New("not declare returns")
}

func (s *MemoryStorage)Delete(id, password string) error {
	if err := isCreateMap(s.m); err != nil {
		return err
	} else {
		if item, err := getItem(&s.m, id); err != nil {
			return err
		} else {
			if item.Password != password {
				return errors.New("not equals password")
			} else {
				delete(s.m, id)

				return nil
			}
		}
	}

	return errors.New("not declare returns")
}

func (s *MemoryStorage)Query(id string) (string, string, error) {
	if err := isCreateMap(s.m); err != nil {
		return "", "", err
	} else {
		if item, err := getItem(&s.m, id); err != nil {
			return "", "", err
		} else {
			return item.Name, item.EMail, nil
		}
	}

	return "", "", errors.New("not declare returns")
}

func (s *MemoryStorage)Count() int {
	return len(s.m)
}

func isCreateMap(m map[string]StorageItem) error {
	if m == nil {
		return errors.New("m is not created map")
	}

	return nil
}

func getItem(m *map[string]StorageItem, id string) (StorageItem, error) {
	item, exist := (*m)[id]

	if exist == false {
		return StorageItem{}, errors.New("not exist id: " + id)
	}

	return item, nil
}
