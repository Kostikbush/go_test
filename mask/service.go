package mask

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}

type Service struct {
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) *Service {
	return &Service{prod: prod, pres: pres}
}


func (s *Service) MaskUrlInMessage(message string) string {
	byteMessage := []byte(message)
	var startIndex int;

  for i := range byteMessage {
    if  i+4 < len(byteMessage) && byteMessage[i] == 'h' &&
			 byteMessage[i+1] == 't' &&
			 byteMessage[i+2] == 't' &&
			 byteMessage[i+3] == 'p' && 
			 (byteMessage[i+4] == 's' || byteMessage[i+4] == ':'){
				if(byteMessage[i+4] == 's'){
					startIndex = i + 8
				}else {
					startIndex = i+7
				}

			continue
		}

		if(startIndex != 0 && byteMessage[i] != ' ' && startIndex <= i){
			byteMessage[i] = '*' 
			
			continue
		}

		if(byteMessage[i] == ' ') {
			startIndex = 0;

			
			continue
		}
  }

	return string(byteMessage)
}

func (s *Service) Run() error {
	lines, err := s.prod.produce()
	if err != nil {
		return err
	}

	masked := make([]string, 0, len(lines))
	for _, line := range lines {
		masked = append(masked, s.MaskUrlInMessage(line))
	}

	if err := s.pres.present(masked); err != nil {
		return err
	}
	return nil
}