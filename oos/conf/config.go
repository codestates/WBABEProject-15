package conf

import (
	"os"

	"github.com/naoina/toml"
)

type Log struct {
	Level   string
	Fpath   string
	Msize   int
	Mage    int
	Mbackup int
}

type Config struct {
	Server struct {
		Mode string
		Port string
	}
	Log
	DB map[string]map[string]interface{}
}

func GetConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		//toml 파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		}
	}
	return c
}
	/* [코드리뷰]
	 * 해당 GetConfig 코드에는 하나의 function에서 간결한 이중 조건문이 발생하고 있습니다.
	 * 15 line으로 간결하게 구성되어, 보기에 어렵지 않은 좋은 코드입니다.
	 * 그러나 코드 라인수가 많아지고, 비즈니스 로직이 풍부해 지는 것을 고려한다면
	 * return으로 나갈 수 빠지는 case를 최소화하고, if문을 줄이는 방향으로 개발을 진행해보시는 연습을 추천드립니다.
	 * 점진적으로 코드의 가독성이 조금 더 향상하게 될 것입니다.
	 * as-is: 
	 if file, err := os.Open(fpath); err != nil {
			panic(err)
		} else {
			defer file.Close()
			//toml 파일 디코딩
			if err := toml.NewDecoder(file).Decode(c); err != nil {
				panic(err)
			}
	}
	 * to-be:
	 if file, err := os.Open(fpath); err == nil {
			defer file.Close()
			//toml 파일 디코딩
			if err := toml.NewDecoder(file).Decode(c); err == nil {
				return c
			} 
		}
		return nil


	 * 또한 function에서 정상일 경우에는 config를 return하지만, 
	 * 나머지 경우에서는 프로그램이 즉시 종료될 것으로 예상됩니다.
	 * panic 보다는 return err를 통한 err 관리로 변경하면 더 견고한 코드가 될 것으로 보입니다.
	 */