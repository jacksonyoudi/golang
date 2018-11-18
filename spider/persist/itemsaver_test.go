package persist

import (
	"golang/spider/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{

	}

	//  测试不要依赖 外界

	Save(profile)
}
