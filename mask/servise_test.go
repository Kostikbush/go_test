package mask

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProducerMock struct{ mock.Mock }

var linesMock = []string{
		"See http://example.com now",
		"x https://a.b/c y",
		"hello world",
		"go to http://ex.com",
		"a/b c",
}

var linesMaskedMock = []string{
	"See http://*********** now",
	"x https://***** y",
	"hello world",
	"go to http://******",
	"a/b c",
}

func (p *ProducerMock) produce() ([]string, error) {
	args := p.Called()
	return args.Get(0).([]string), args.Error(1)
}

type PresenterMock struct{ mock.Mock }

func (p *PresenterMock) present(lines []string) error {
	args := p.Called(lines)
	return args.Error(0)
}

func TestRun_Success(t *testing.T) {
	prod := new(ProducerMock)
	pres := new(PresenterMock)

	prod.On("produce").Return(linesMock, nil).Once()
	pres.On("present", linesMaskedMock).Return(nil).Once()

	maskService := NewService(prod, pres)

	err := maskService.Run()

	assert.NoError(t, err)
	prod.AssertExpectations(t)
	pres.AssertExpectations(t)
}

func TestRun_ProducerError(t *testing.T) {
	prod := new(ProducerMock)
	pres := new(PresenterMock)

	prod.On("produce").Return([]string(nil), errors.New("error")).Once()

	maskService := NewService(prod, pres)

	err := maskService.Run()

	assert.Error(t, err)
	pres.AssertNotCalled(t, "Present", mock.Anything)
}

func TestRun_PresenterError(t *testing.T) {
	prod := new(ProducerMock)
	pres := new(PresenterMock)

	prod.On("produce").Return(linesMock, nil).Once()
	pres.On("present", linesMaskedMock).Return(errors.New("error")).Once()

	maskService := NewService(prod, pres)

	err := maskService.Run()

	assert.Error(t, err)

}