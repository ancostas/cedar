package model

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/evergreen-ci/pail"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestTestResultsIterator(t *testing.T) {
	const numResults = 10
	for testName, testCase := range map[string]func(ctx context.Context, t *testing.T, iter TestResultsIterator, results map[string]TestResult){
		"VerifyInitialState": func(ctx context.Context, t *testing.T, iter TestResultsIterator, results map[string]TestResult) {
			assert.Zero(t, iter.Item())
			assert.Zero(t, iter.Err())
			assert.False(t, iter.Exhausted())
			assert.NoError(t, iter.Close())
		},
		"NextItemMatches": func(ctx context.Context, t *testing.T, iter TestResultsIterator, results map[string]TestResult) {
			var numItems int
			for iter.Next(ctx) {
				expectedItem, ok := results[iter.Item().TestName]
				require.True(t, ok)
				assert.Equal(t, expectedItem, iter.Item())
				assert.False(t, iter.Exhausted())
				numItems++
			}
			require.Zero(t, iter.Err())
			assert.True(t, iter.Exhausted())
			assert.Equal(t, len(results), numItems)

			// Verify that it's safe to call next after it's exhausted.
			assert.False(t, iter.Next(ctx))
			assert.Zero(t, iter.Err())
		},
		"ClosePreventsNext": func(ctx context.Context, t *testing.T, iter TestResultsIterator, results map[string]TestResult) {
			require.NoError(t, iter.Close())
			assert.False(t, iter.Next(ctx))
		},
	} {
		t.Run(testName, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			tmpDir, err := ioutil.TempDir(".", "test-results-iterator-test")
			require.NoError(t, err)
			defer func() {
				assert.NoError(t, os.RemoveAll(tmpDir))
			}()

			tr := getTestResults()
			testBucket, err := pail.NewLocalBucket(pail.LocalOptions{
				Path:   tmpDir,
				Prefix: filepath.Join(testResultsCollection, tr.ID),
			})
			require.NoError(t, err)

			results := map[string]TestResult{}
			for i := 0; i < numResults; i++ {
				result := getTestResult()
				results[result.TestName] = result
				data, err := bson.Marshal(result)
				require.NoError(t, err)
				require.NoError(t, testBucket.Put(ctx, result.TestName, bytes.NewReader(data)))
			}

			testCase(ctx, t, NewTestResultsIterator(testBucket), results)
		})
	}
}
