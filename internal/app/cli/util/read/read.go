package read

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/bastean/godo/internal/pkg/service/errors"
)

func Remote(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Remote",
			What:  fmt.Sprintf("Cannot obtain the remote resource from [%s]", url),
			Who:   err,
		})
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Remote",
			What:  fmt.Sprintf("Cannot read the body of the response from [%s]", url),
			Who:   err,
		})
	}

	if response.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Remote",
			What:  fmt.Sprintf("Failure response of [%s] with status code [%d] and body [%s].", url, response.StatusCode, content),
		})
	}

	return content, nil
}

func Local(path string) ([]byte, error) {
	content, err := os.ReadFile(path)

	switch {
	case errors.Is(err, os.ErrNotExist):
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Local",
			What:  fmt.Sprintf("File does not exist [%s]", path),
		})
	case err != nil:
		return nil, errors.NewFailure(&errors.Bubble{
			Where: "Local",
			What:  fmt.Sprintf("Failure to read file [%s]", path),
			Who:   err,
		})
	}

	return content, nil
}

func File(route string) ([]byte, error) {
	if strings.HasPrefix(route, "http") {
		return Remote(route)
	}

	return Local(route)
}

func JSON(data []byte, v any) error {
	var errJsonUnmarshalType *json.UnmarshalTypeError

	err := json.Unmarshal(data, &v)

	switch {
	case errors.As(err, &errJsonUnmarshalType):
		return errors.NewFailure(&errors.Bubble{
			Where: "JSON",
			What:  fmt.Sprintf("Invalid type field [%s] required type is [%s] and [%s] was obtained", errJsonUnmarshalType.Field, errJsonUnmarshalType.Type, errJsonUnmarshalType.Value),
		})
	case err != nil:
		return errors.NewFailure(&errors.Bubble{
			Where: "JSON",
			What:  "Failure to read JSON",
			Who:   err,
		})
	}

	return nil
}
