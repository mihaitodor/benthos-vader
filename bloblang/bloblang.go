package bloblang

import (
	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/jonreiter/govader"
)

func init() {
	sentimentSpec := bloblang.NewPluginSpec().
		Category("String Manipulation").
		Description("Computes sentiment parameters for the given string. Wraps the github.com/jonreiter/govader package. See its [docs](https://pkg.go.dev/github.com/jonreiter/govader) for more information.")

	if err := bloblang.RegisterMethodV2("sentiment", sentimentSpec,
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			analyzer := govader.NewSentimentIntensityAnalyzer()

			return bloblang.StringMethod(func(s string) (interface{}, error) {
				return analyzer.PolarityScores(s), nil
			}), nil
		},
	); err != nil {
		panic(err)
	}
}
