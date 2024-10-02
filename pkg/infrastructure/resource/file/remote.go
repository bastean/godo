package file

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bastean/godo/pkg/domain/errors"
)

type Remote struct{}

func (*Remote) Read(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Read",
			What:  fmt.Sprintf("Cannot obtain the remote resource from [%s]", url),
			Who:   err,
		})
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Read",
			What:  fmt.Sprintf("Cannot read the body of the response from [%s]", url),
			Who:   err,
		})
	}

	if response.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Read",
			What:  fmt.Sprintf("Failure response of [%s] with status code [%d] and body [%s].", url, response.StatusCode, content),
		})
	}

	return content, nil
}
